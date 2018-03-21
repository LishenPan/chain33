// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	account.proto
	blockchain.proto
	common.proto
	config.proto
	db.proto
	executor.proto
	p2p.proto
	rpc.proto
	transaction.proto
	wallet.proto

It has these top-level messages:
	Account
	ReceiptExecAccountTransfer
	ReceiptAccountTransfer
	ReqBalance
	Accounts
	ReqTokenBalance
	Header
	Block
	Blocks
	BlockDetails
	Headers
	BlockOverview
	BlockDetail
	Receipts
	ReceiptCheckTxList
	ChainStatus
	ReqBlocks
	MempoolSize
	ReplyBlockHeight
	BlockBody
	IsCaughtUp
	Reply
	ReqString
	ReplyString
	ReplyStrings
	ReqInt
	Int64
	ReqHash
	ReplyHash
	ReqNil
	ReqHashes
	ReplyHashes
	KeyValue
	Config
	MemPool
	Consensus
	Wallet
	Store
	LocalStore
	BlockChain
	P2P
	Rpc
	LeafNode
	InnerNode
	MAVLProof
	StoreNode
	LocalDBSet
	LocalDBList
	LocalDBGet
	LocalReplyValue
	StoreSet
	StoreGet
	StoreReplyValue
	Genesis
	CoinsAction
	CoinsGenesis
	CoinsTransfer
	CoinsWithdraw
	Hashlock
	HashlockAction
	HashlockLock
	HashlockUnlock
	HashlockSend
	HashRecv
	Hashlockquery
	Ticket
	TicketAction
	TicketMiner
	TicketBind
	TicketOpen
	TicketGenesis
	TicketClose
	TicketList
	TicketInfos
	ReplyTicketList
	ReplyWalletTickets
	ReceiptTicket
	ReceiptTicketBind
	ExecTxList
	Query
	RetrievePara
	Retrieve
	RetrieveAction
	BackupRetrieve
	PreRetrieve
	PerformRetrieve
	CancelRetrieve
	TokenAction
	TokenPreCreate
	TokenFinishCreate
	TokenRevokeCreate
	Token
	ReqTokens
	ReplyTokens
	ReceiptToken
	Trade
	TradeForSell
	SellOrder
	SellOrderReceipt
	ReqAddrTokens
	ReplySellOrders
	TokenRecv
	ReplyAddrRecvForTokens
	TradeForBuy
	TradeForRevokeSell
	ReceiptTradeBase
	ReceiptTradeSell
	ReceiptTradeBuy
	ReceiptTradeRevoke
	TradeBuyDone
	ReplyTradeBuyOrders
	P2PGetPeerInfo
	P2PPeerInfo
	P2PVersion
	P2PVerAck
	P2PPing
	P2PPong
	P2PGetAddr
	P2PAddr
	P2PExternalInfo
	P2PGetBlocks
	P2PGetMempool
	P2PInv
	Inventory
	P2PGetData
	P2PTx
	P2PBlock
	BroadCastData
	P2PGetHeaders
	P2PHeaders
	InvData
	InvDatas
	Peer
	PeerList
	CreateTx
	UnsignTx
	SignedTx
	Transaction
	Signature
	AddrOverview
	ReqAddr
	HexTx
	ReplyTxInfo
	ReqTxList
	ReplyTxList
	TxHashList
	ReplyTxInfos
	ReceiptLog
	Receipt
	ReceiptData
	TxResult
	TransactionDetail
	TransactionDetails
	ReqAddrs
	WalletTxDetail
	WalletTxDetails
	WalletAccountStore
	WalletPwHash
	WalletStatus
	WalletAccounts
	WalletAccount
	WalletUnLock
	GenSeedLang
	GetSeedByPw
	SaveSeedByPw
	ReplySeed
	ReqWalletSetPasswd
	ReqNewAccount
	MinerFlag
	ReqWalletTransactionList
	ReqWalletImportPrivKey
	ReqWalletSendToAddress
	ReqWalletSetFee
	ReqWalletSetLabel
	ReqWalletMergeBalance
	ReqStr
	ReplyStr
	ReqTokenPreCreate
	ReqTokenFinishCreate
	ReqTokenRevokeCreate
	ReqSellToken
	ReqBuyToken
	ReqRevokeSell
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account 的信息
type Account struct {
	Currency int32  `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	Balance  int64  `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	Frozen   int64  `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	Addr     string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 账户余额改变的一个交易回报（合约内）
type ReceiptExecAccountTransfer struct {
	ExecAddr string   `protobuf:"bytes,1,opt,name=execAddr" json:"execAddr,omitempty"`
	Prev     *Account `protobuf:"bytes,2,opt,name=prev" json:"prev,omitempty"`
	Current  *Account `protobuf:"bytes,3,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptExecAccountTransfer) Reset()                    { *m = ReceiptExecAccountTransfer{} }
func (m *ReceiptExecAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptExecAccountTransfer) ProtoMessage()               {}
func (*ReceiptExecAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReceiptExecAccountTransfer) GetExecAddr() string {
	if m != nil {
		return m.ExecAddr
	}
	return ""
}

func (m *ReceiptExecAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptExecAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 账户余额改变的一个交易回报（coins内）
type ReceiptAccountTransfer struct {
	Prev    *Account `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *Account `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptAccountTransfer) Reset()                    { *m = ReceiptAccountTransfer{} }
func (m *ReceiptAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptAccountTransfer) ProtoMessage()               {}
func (*ReceiptAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReceiptAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 查询一个地址列表在某个执行器中余额
type ReqBalance struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Execer    string   `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqBalance) Reset()                    { *m = ReqBalance{} }
func (m *ReqBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqBalance) ProtoMessage()               {}
func (*ReqBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReqBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// Account 的列表
type Accounts struct {
	Acc []*Account `protobuf:"bytes,1,rep,name=acc" json:"acc,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Accounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ReqTokenBalance struct {
	Addresses   []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	TokenSymbol string   `protobuf:"bytes,2,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	Execer      string   `protobuf:"bytes,3,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqTokenBalance) Reset()                    { *m = ReqTokenBalance{} }
func (m *ReqTokenBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenBalance) ProtoMessage()               {}
func (*ReqTokenBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReqTokenBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqTokenBalance) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *ReqTokenBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*ReceiptExecAccountTransfer)(nil), "types.ReceiptExecAccountTransfer")
	proto.RegisterType((*ReceiptAccountTransfer)(nil), "types.ReceiptAccountTransfer")
	proto.RegisterType((*ReqBalance)(nil), "types.ReqBalance")
	proto.RegisterType((*Accounts)(nil), "types.Accounts")
	proto.RegisterType((*ReqTokenBalance)(nil), "types.ReqTokenBalance")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6a, 0xf3, 0x30,
	0x10, 0xc5, 0x51, 0x9c, 0x7f, 0x9e, 0xf0, 0x7d, 0x05, 0x2d, 0x82, 0x08, 0x5d, 0x08, 0xad, 0xbc,
	0x28, 0x59, 0xb4, 0x27, 0x68, 0xa0, 0x17, 0x50, 0x73, 0x01, 0x45, 0x19, 0x43, 0x48, 0x2a, 0x39,
	0x92, 0x52, 0xea, 0x1e, 0xa0, 0xe7, 0x2e, 0x52, 0xe4, 0xd4, 0x34, 0x50, 0xb2, 0xf3, 0x9b, 0x37,
	0x33, 0xef, 0xe7, 0x41, 0xf0, 0x4f, 0x69, 0x6d, 0x4f, 0x26, 0x2c, 0x1b, 0x67, 0x83, 0xa5, 0xa3,
	0xd0, 0x36, 0xe8, 0xc5, 0x1e, 0x26, 0xcf, 0xe7, 0x3a, 0x5d, 0xc0, 0x54, 0x9f, 0x9c, 0x43, 0xa3,
	0x5b, 0x46, 0x38, 0xa9, 0x46, 0xf2, 0xa2, 0x29, 0x83, 0xc9, 0x46, 0x1d, 0x94, 0xd1, 0xc8, 0x06,
	0x9c, 0x54, 0x85, 0xec, 0x24, 0x9d, 0xc3, 0xb8, 0x76, 0xf6, 0x13, 0x0d, 0x2b, 0x92, 0x91, 0x15,
	0xa5, 0x30, 0x54, 0xdb, 0xad, 0x63, 0x43, 0x4e, 0xaa, 0x52, 0xa6, 0x6f, 0xf1, 0x45, 0x60, 0x21,
	0x51, 0xe3, 0xae, 0x09, 0x2f, 0x1f, 0xa8, 0x73, 0xf0, 0xda, 0x29, 0xe3, 0x6b, 0x74, 0x11, 0x00,
	0x63, 0x39, 0x8e, 0x91, 0x34, 0x76, 0xd1, 0x54, 0xc0, 0xb0, 0x71, 0xf8, 0x9e, 0xd2, 0x67, 0x8f,
	0xff, 0x97, 0x89, 0x7e, 0x99, 0x37, 0xc8, 0xe4, 0xd1, 0x0a, 0x26, 0x67, 0xe0, 0x90, 0x58, 0xae,
	0xdb, 0x3a, 0x5b, 0xd4, 0x30, 0xcf, 0x1c, 0xbf, 0x19, 0xba, 0x1c, 0x72, 0x5b, 0xce, 0xe0, 0xef,
	0x9c, 0x15, 0x80, 0xc4, 0xe3, 0x2a, 0x9f, 0xea, 0x1e, 0xca, 0x78, 0x06, 0xf4, 0x1e, 0x3d, 0x23,
	0xbc, 0xa8, 0x4a, 0xf9, 0x53, 0x88, 0x87, 0x8c, 0x7f, 0x8b, 0x2e, 0x2d, 0x2d, 0x65, 0x56, 0xe2,
	0x01, 0xa6, 0x79, 0xaf, 0xa7, 0x1c, 0x0a, 0xa5, 0x75, 0x9a, 0xbd, 0x4e, 0x8d, 0x96, 0xd8, 0xc1,
	0x9d, 0xc4, 0xe3, 0xda, 0xee, 0xd1, 0xdc, 0x16, 0xcb, 0x61, 0x16, 0x62, 0xf7, 0x6b, 0xfb, 0xb6,
	0xb1, 0x87, 0x9c, 0xdd, 0x2f, 0xf5, 0xc0, 0x8a, 0x3e, 0xd8, 0x66, 0x9c, 0x1e, 0xd2, 0xd3, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x36, 0x04, 0xec, 0x59, 0x02, 0x00, 0x00,
}
