package db

import (
	"bytes"
	"context"
	"strings"
	"time"

	"github.com/XiaoMi/pegasus-go-client/pegasus"
	log "github.com/inconshreveable/log15"
	"github.com/syndtr/goleveldb/leveldb/util"
	"gitlab.33.cn/chain33/chain33/types"
)

var slog = log.New("module", "db.pegasus")
var pdbBench = &SsdbBench{}
var HashKeyLen = 24

func init() {
	dbCreator := func(name string, dir string, cache int) (DB, error) {
		return NewPegasusDB(name, dir, cache)
	}
	registerDBCreator(goPegasusDbBackendStr, dbCreator, false)
}

type PegasusDB struct {
	TransactionDB

	cfg    *pegasus.Config
	name   string
	client pegasus.Client
	table  pegasus.TableConnector
}

func printPegasusBenchmark() {
	tick := time.Tick(time.Minute * 5)
	for {
		<-tick
		slog.Info(pdbBench.String())
	}
}

func NewPegasusDB(name string, dir string, cache int) (*PegasusDB, error) {
	database := &PegasusDB{name: name}
	database.cfg = parsePegasusNodes(dir)

	if database.cfg == nil {
		slog.Error("no valid instance exists, exit!")
		return nil, types.ErrDataBaseDamage
	}
	var err error
	database.client = pegasus.NewClient(*database.cfg)
	tb, err := database.client.OpenTable(context.Background(), database.name)
	if err != nil {
		slog.Error("connect to pegasus error!", "pegasus", database.cfg, "error", err)
		database.client.Close()
		return nil, types.ErrDataBaseDamage
	}
	database.table = tb

	go printPegasusBenchmark()
	return database, nil
}

// url pattern: ip:port,ip:port
func parsePegasusNodes(url string) *pegasus.Config {
	hosts := strings.Split(url, ",")
	if hosts == nil {
		slog.Error("invalid url")
		return nil
	}

	cfg := &pegasus.Config{hosts}
	return cfg
}

func (db *PegasusDB) Get(key []byte) ([]byte, error) {
	start := time.Now()
	hashKey := getHashKey(key)
	value, err := db.table.Get(context.Background(), hashKey, key)
	if err != nil {
		//slog.Error("Get value error", "error", err, "key", key, "keyhex", hex.EncodeToString(key), "keystr", string(key))
		return nil, err
	}
	if value == nil {
		return nil, ErrNotFoundInDb
	}

	pdbBench.read(1, time.Since(start))
	return value, nil
}

func (db *PegasusDB) BatchGet(keys [][]byte) (values [][]byte, err error) {
	start := time.Now()
	defer pdbBench.read(len(keys), time.Since(start))

	var (
		keyMap  map[int][]byte
		hashMap map[string][][]byte
		valMap  map[string][]byte
		hashKey []byte
	)
	keyMap = make(map[int][]byte)
	hashMap = make(map[string][][]byte)
	valMap = make(map[string][]byte)

	// 这里其实也需要对hashKey进行分别计算，然后分组查询，最后汇总结果

	// 首先，记录查询key的顺序，并对keys进行哈希分组
	for i, v := range keys {
		keyMap[i] = v
		hashKey = getHashKey(v)
		if value, ok := hashMap[string(hashKey)]; ok {
			hashMap[string(hashKey)] = append(value, v)
		} else {
			hashMap[string(hashKey)] = [][]byte{v}
		}
	}

	// 然后，使用hashKey进行分组查询
	for k, v := range hashMap {
		vals, err := db.batchGet([]byte(k), v)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(v); i++ {
			valMap[string(v[i])] = vals[i]
		}
	}

	// 最后，按照查询顺序，从新组装结果
	for i := 0; i < len(keys); i++ {
		if v, ok := valMap[string(keyMap[i])]; ok {
			values = append(values, v)
		} else {
			return nil, ErrNotFoundInDb
		}
	}

	return values, nil
}

func (db *PegasusDB) batchGet(hashKey []byte, keys [][]byte) (values [][]byte, err error) {
	vals, _, err := db.table.MultiGet(context.Background(), hashKey, keys)
	if err != nil {
		//slog.Error("Get value error", "error", err, "key", key, "keyhex", hex.EncodeToString(key), "keystr", string(key))
		return nil, err
	}
	if vals == nil {
		return nil, ErrNotFoundInDb
	}
	for _, v := range vals {
		values = append(values, v.Value)
	}
	return values, nil
}

func (db *PegasusDB) Set(key []byte, value []byte) error {
	start := time.Now()
	hashKey := getHashKey(key)
	err := db.table.Set(context.Background(), hashKey, key, value)
	if err != nil {
		slog.Error("Set", "error", err)
		return err
	}
	pdbBench.write(1, time.Since(start))
	return nil
}

func (db *PegasusDB) SetSync(key []byte, value []byte) error {
	return db.Set(key, value)
}

func (db *PegasusDB) Delete(key []byte) error {
	start := time.Now()
	defer pdbBench.write(1, time.Since(start))
	hashKey := getHashKey(key)
	err := db.table.Del(context.Background(), hashKey, key)
	if err != nil {
		slog.Error("Delete", "error", err)
		return err
	}
	return nil
}

func (db *PegasusDB) DeleteSync(key []byte) error {
	return db.Delete(key)
}

func (db *PegasusDB) Close() {
	db.table.Close()
	db.client.Close()
}

func (db *PegasusDB) Print() {
}

func (db *PegasusDB) Stats() map[string]string {
	return nil
}

func (db *PegasusDB) Iterator(begin []byte, end []byte, reverse bool) Iterator {
	var (
		err   error
		vals  []*pegasus.KeyValue
		start []byte
		over  []byte
	)
	if end == nil {
		end = bytesPrefix(begin)
	}
	if bytes.Equal(end, types.EmptyValue) {
		end = nil
	}
	limit := util.Range{begin, end}
	hashKey := getHashKey(begin)

	if reverse {
		start = begin
		over = limit.Limit
	} else {
		start = limit.Limit
		over = begin
	}
	dbit := &PegasusIt{itBase: itBase{begin, end, reverse}, index: -1, table: db.table, itbegin: start, itend: over}
	opts := &pegasus.MultiGetOptions{StartInclusive: false, StopInclusive: false, MaxFetchCount: IteratorPageSize, Reverse: dbit.reverse}
	vals, _, err = db.table.MultiGetRangeOpt(context.Background(), hashKey, begin, limit.Limit, opts)
	if err != nil {
		slog.Error("create iterator error!")
		return nil
	}
	if len(vals) > 0 {
		dbit.vals = vals
		// 如果返回的数据大小刚好满足分页，则假设下一页还有数据
		if len(dbit.vals) == IteratorPageSize {
			dbit.nextPage = true
			// 下一页数据的开始，等于本页数据的结束，不过在下次查询时需要设置StartInclusiv=false，因为本条数据已经包含
			dbit.tmpEnd = dbit.vals[IteratorPageSize-1].SortKey
		}
	}
	return dbit
}

type PegasusIt struct {
	itBase
	table    pegasus.TableConnector
	vals     []*pegasus.KeyValue
	index    int
	nextPage bool
	tmpEnd   []byte

	// 迭代开始位置
	itbegin []byte
	// 迭代结束位置
	itend []byte
	// 当前所属的页数（从0开始）
	pageNo int
}

func (dbit *PegasusIt) Close() {
	dbit.index = -1
}

func (dbit *PegasusIt) Next() bool {
	if len(dbit.vals) > dbit.index+1 {
		dbit.index++
		return true
	} else {
		// 如果有下一页数据，则自动抓取
		if dbit.nextPage {
			return dbit.cacheNextPage(dbit.tmpEnd)
		}
		return false
	}
}

func (dbit *PegasusIt) initPage(begin, end []byte) bool {
	var (
		vals []*pegasus.KeyValue
		err  error
	)
	opts := &pegasus.MultiGetOptions{StartInclusive: false, StopInclusive: false, MaxFetchCount: IteratorPageSize, Reverse: dbit.reverse}
	hashKey := getHashKey(begin)
	vals, _, err = dbit.table.MultiGetRangeOpt(context.Background(), hashKey, begin, end, opts)

	if err != nil {
		slog.Error("get iterator next page error", "error", err, "begin", begin, "end", dbit.itend, "reverse", dbit.reverse)
		return false
	}

	if len(vals) > 0 {
		// 这里只改变keys，不改变index
		dbit.vals = vals

		// 如果返回的数据大小刚好满足分页，则假设下一页还有数据
		if len(vals) == IteratorPageSize {
			dbit.nextPage = true
			dbit.tmpEnd = dbit.vals[IteratorPageSize-1].SortKey
		} else {
			dbit.nextPage = false
		}
		return true
	} else {
		return false
	}
}

// 获取下一页的数据
func (dbit *PegasusIt) cacheNextPage(flag []byte) bool {
	var (
		over []byte
	)
	// 如果是逆序，则取从开始到flag的数据
	if dbit.reverse {
		over = dbit.itbegin
	} else {
		over = dbit.itend
	}
	// 如果是正序，则取从flag到结束的数据
	if dbit.initPage(flag, over) {
		dbit.index = 0
		dbit.pageNo++
		return true
	} else {
		return false
	}
}

func (dbit *PegasusIt) checkKeyCmp(key1, key2 []byte, reverse bool) bool {
	if reverse {
		return bytes.Compare(key1, key2) < 0
	}
	return bytes.Compare(key1, key2) > 0
}

func (dbit *PegasusIt) findInPage(key []byte) int {
	pos := -1
	for i, v := range dbit.vals {
		if i < dbit.index {
			continue
		}
		if dbit.checkKeyCmp(key, v.SortKey, dbit.reverse) {
			continue
		} else {
			pos = i
			break
		}
	}
	return pos
}

func (dbit *PegasusIt) Seek(key []byte) bool {
	pos := -1
	pos = dbit.findInPage(key)

	// 如果第一页已经找到，不会走入此逻辑
	for pos == -1 && dbit.nextPage {
		if dbit.cacheNextPage(dbit.tmpEnd) {
			pos = dbit.findInPage(key)
		} else {
			break
		}
	}

	dbit.index = pos
	return dbit.Valid()
}

func (dbit *PegasusIt) Rewind() bool {
	// 目前代码的Rewind调用都是在第一页，正常情况下走不到else分支；
	// 但为了代码健壮性考虑，这里增加对else分支的处理
	if dbit.pageNo == 0 {
		dbit.index = 0
		return true
	}

	// 当数据取到第N页的情况时，Rewind需要返回到第一页第一条
	if dbit.initPage(dbit.itbegin, dbit.itend) {
		dbit.index = 0
		dbit.pageNo = 0
		return true
	} else {
		return false
	}
}

func (dbit *PegasusIt) Key() []byte {
	if dbit.index >= 0 && dbit.index < len(dbit.vals) {
		return dbit.vals[dbit.index].SortKey
	} else {
		return nil
	}
}
func (dbit *PegasusIt) Value() []byte {
	if dbit.index >= len(dbit.vals) {
		slog.Error("get iterator value error: index out of bounds")
		return nil
	}

	return dbit.vals[dbit.index].Value
}

func (dbit *PegasusIt) Error() error {
	return nil
}

func (dbit *PegasusIt) ValueCopy() []byte {
	v := dbit.Value()
	value := make([]byte, len(v))
	copy(value, v)
	return value
}

func (dbit *PegasusIt) Valid() bool {
	start := time.Now()
	if dbit.index < 0 {
		return false
	}
	if len(dbit.vals) <= dbit.index {
		return false
	}
	key := dbit.vals[dbit.index].SortKey
	pdbBench.read(1, time.Since(start))
	return dbit.checkKey(key)
}

type PegasusBatch struct {
	table    pegasus.TableConnector
	batchset map[string][]byte
	batchdel map[string][]byte
}

func (db *PegasusDB) NewBatch(sync bool) Batch {
	return &PegasusBatch{table: db.table, batchset: make(map[string][]byte), batchdel: make(map[string][]byte)}
}

func (db *PegasusBatch) Set(key, value []byte) {
	db.batchset[string(key)] = value
	delete(db.batchdel, string(key))
}

func (db *PegasusBatch) Delete(key []byte) {
	db.batchset[string(key)] = []byte("")
	delete(db.batchset, string(key))
	db.batchdel[string(key)] = key
}

// 注意本方法的实现逻辑，因为ssdb没有提供删除和更新同时进行的批量操作；
// 所以这里先执行更新操作（删除的KEY在这里会将VALUE设置为空）；
// 然后再执行删除操作；
// 这样即使中间执行出错，也不会导致删除结果未写入的情况（值已经被置空）；
func (db *PegasusBatch) Write() error {
	start := time.Now()

	// 这里其实也需要对hashKey进行分别计算，然后分组查询，最后汇总结果
	if len(db.batchset) > 0 {
		var (
			keysMap map[string][][]byte
			valsMap map[string][][]byte
			hashKey []byte
			byteKey []byte
			keys    [][]byte
			values  [][]byte
		)
		keysMap = make(map[string][][]byte)
		valsMap = make(map[string][][]byte)

		// 首先，使用hashKey进行数据分组
		for k, v := range db.batchset {
			byteKey = []byte(k)
			hashKey = getHashKey(byteKey)
			if value, ok := keysMap[string(hashKey)]; ok {
				keysMap[string(hashKey)] = append(value, byteKey)
				valsMap[string(hashKey)] = append(valsMap[string(hashKey)], v)
			} else {
				keysMap[string(hashKey)] = [][]byte{byteKey}
				valsMap[string(hashKey)] = [][]byte{v}
			}
		}
		// 然后，再分别提交修改
		for k, v := range keysMap {
			keys = v
			values = valsMap[k]

			err := db.table.MultiSet(context.Background(), []byte(k), keys, values)
			if err != nil {
				slog.Error("Write (multi_set)", "error", err)
				return err
			}
		}
	}

	if len(db.batchdel) > 0 {
		var (
			keysMap map[string][][]byte
			hashKey []byte
			byteKey []byte
		)
		keysMap = make(map[string][][]byte)

		// 首先，使用hashKey进行数据分组
		for k := range db.batchdel {
			byteKey = []byte(k)
			hashKey = getHashKey(byteKey)
			if value, ok := keysMap[string(hashKey)]; ok {
				keysMap[string(hashKey)] = append(value, byteKey)
			} else {
				keysMap[string(hashKey)] = [][]byte{byteKey}
			}
		}

		// 然后，再分别提交删除
		for k, v := range keysMap {
			err := db.table.MultiDel(context.Background(), []byte(k), v)
			if err != nil {
				slog.Error("Write (multi_del)", "error", err)
				return err
			}
		}
	}

	pdbBench.write(len(db.batchset)+len(db.batchdel), time.Since(start))
	return nil
}

func (db *PegasusBatch) ValueSize() int {
	return len(db.batchset)
}

func (db *PegasusBatch) Reset() {
	db.batchset = make(map[string][]byte)
	db.batchdel = make(map[string][]byte)
}

func getHashKey(key []byte) []byte {
	if len(key) <= HashKeyLen {
		return key
	}
	return key[:HashKeyLen]
}
