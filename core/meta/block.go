package meta

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/mihongtech/linkchain-core/common/math"
	"github.com/mihongtech/linkchain-core/common/serialize"
	"github.com/mihongtech/linkchain-core/common/trie"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/protobuf"

	"github.com/golang/protobuf/proto"
)

type Block struct {
	Header BlockHeader  `json:"header"`
	TXs    Transactions `json:"txs"`
}

func NewBlock(header BlockHeader, txs []Transaction) *Block {

	return &Block{
		Header: header,
		TXs:    *NewTransactions(txs...),
	}
}

func (b *Block) SetTx(newTXs ...Transaction) error {
	if err := b.TXs.SetTx(newTXs...); err != nil {
		return err
	}
	b.Header.SetMerkleRoot(b.CalculateTxTreeRoot()) //calculate merkle root

	return nil
}

func (b *Block) SetSign(signature math.ISignature) {
	b.Header.Sign = *signature.(*Signature)
}

func (b *Block) GetHeight() uint32 {
	return b.Header.Height
}

func (b *Block) GetBlockID() *BlockID {
	return b.Header.GetBlockID()
}

func (b *Block) GetTime() time.Time {
	return b.Header.Time
}

func (b *Block) GetStatus() *TreeID {
	return &b.Header.Status
}

func (b *Block) GetPrevBlockID() *BlockID {
	return &b.Header.Prev
}

func (b *Block) GetMerkleRoot() *TreeID {
	return b.Header.GetMerkleRoot()
}

//Serialize/Deserialize
func (b *Block) Serialize() serialize.SerializeStream {
	header := b.Header.Serialize().(*protobuf.BlockHeader)

	block := protobuf.Block{
		Header: header,
		TxList: b.TXs.Serialize().(*protobuf.Transactions),
	}

	return &block
}

func (b *Block) Deserialize(s serialize.SerializeStream) error {
	data := *s.(*protobuf.Block)
	err := b.Header.Deserialize(data.Header)
	if err != nil {
		return err
	}

	return b.TXs.Deserialize(data.TxList)
}

func (b *Block) String() string {
	data, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (b *Block) EncodeToBytes() ([]byte, error) {
	return proto.Marshal(b.Serialize())
}

func (b *Block) DecodeFromBytes(buff []byte) error {
	var protoBlock protobuf.Block
	if err := proto.Unmarshal(buff, &protoBlock); err != nil {
		return err
	}
	return b.Deserialize(&protoBlock)
}

func (b *Block) GetTxs() []Transaction {
	return b.TXs.Txs
}

func (b *Block) GetTx(id TxID) (*Transaction, error) {
	return b.TXs.GetTx(id)
}

func (b *Block) CalculateTxTreeRoot() TreeID {
	transactions := make(map[math.Hash][]byte)
	for index, t := range b.TXs.Txs {
		transactions[*b.TXs.Txs[index].GetTxID()] = t.Data
	}
	hash, _ := GetMakeTreeID(transactions)
	return hash
}

func (b *Block) IsGensis() bool {
	return b.Header.IsGensis()
}

type BlockHeader struct {
	// Version of the block.  This is not the same as the protocol version.
	Version uint32 `json:"version,int"`

	//the height of block
	Height uint32 `json:"height,int"`

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	Time time.Time `json:"Time"`

	// Nonce used to generate the block.
	Nonce uint32 `json:"nonce"`

	// Difficulty target for the block.
	Difficulty uint32 `json:"difficulty"`

	// Hash of the previous block header in the block chain.
	Prev BlockID `json:"prev"`

	// Merkle tree reference to hash of all transactions for the block.
	TxRoot TreeID `json:"txroot"`

	// The status of the whole system
	Status TreeID `json:"status"`

	// The sign of miner
	Sign Signature `json:"sign"`

	// Data used to extenion the block.
	Data []byte `json:"data"`

	//The Hash of this block
	hash BlockID
}

func NewBlockHeader(version uint32, height uint32, time time.Time, nounce uint32, difficulty uint32, prev BlockID, root TreeID, status TreeID, sign Signature, extra []byte) *BlockHeader {

	return &BlockHeader{
		Version: version,
		Height:  height,
		Time:    time,
		Nonce:   nounce,

		Difficulty: difficulty,
		Prev:       prev,
		TxRoot:     root,
		Status:     status,
		Sign:       sign,
		Data:       extra,
	}
}

func (bh *BlockHeader) GetBlockID() *BlockID {
	if bh.hash.IsEmpty() {
		if err := bh.Deserialize(bh.Serialize()); err != nil {
			log.Error("BlockHeader", "GetBlockID() error", err)
			return nil
		}
	}
	return &bh.hash
}

func (bh *BlockHeader) GetMerkleRoot() *TreeID {
	return &bh.TxRoot
}

func (bh *BlockHeader) SetMerkleRoot(root TreeID) {
	bh.TxRoot = root
}

//Serialize/Deserialize
func (bh *BlockHeader) Serialize() serialize.SerializeStream {
	prevHash := bh.Prev.Serialize().(*protobuf.Hash)
	merkleRoot := bh.TxRoot.Serialize().(*protobuf.Hash)
	status := bh.Status.Serialize().(*protobuf.Hash)
	sign := bh.Sign.Serialize().(*protobuf.Signature)
	header := protobuf.BlockHeader{
		Version:    proto.Uint32(bh.Version),
		Height:     proto.Uint32(bh.Height),
		Time:       proto.Int64(bh.Time.Unix()),
		Nounce:     proto.Uint32(bh.Nonce),
		Difficulty: proto.Uint32(bh.Difficulty),
		Prev:       prevHash,
		TxRoot:     merkleRoot,
		Status:     status,
		Sign:       sign,
		Data:       proto.NewBuffer(bh.Data).Bytes(),
	}
	return &header
}

func (bh *BlockHeader) Deserialize(s serialize.SerializeStream) error {
	data := s.(*protobuf.BlockHeader)
	bh.Version = *data.Version
	bh.Height = *data.Height
	bh.Time = time.Unix(*data.Time, 0)
	bh.Nonce = *data.Nounce
	bh.Difficulty = *data.Difficulty
	if err := bh.Prev.Deserialize(data.Prev); err != nil {
		return err
	}

	if err := bh.TxRoot.Deserialize(data.TxRoot); err != nil {
		return err
	}

	if err := bh.Status.Deserialize(data.Status); err != nil {
		return err
	}

	if err := bh.Sign.Deserialize(data.Sign); err != nil {
		return err
	}

	bh.Data = data.Data

	t := protobuf.BlockHeader{
		Version:    data.Version,
		Height:     data.Height,
		Time:       data.Time,
		Nounce:     data.Nounce,
		Difficulty: data.Difficulty,
		Prev:       data.Prev,
		TxRoot:     data.TxRoot,
		Status:     data.Status,
		Data:       data.Data,
	}

	buffer, err := proto.Marshal(&t)
	if err != nil {
		return err
	}

	bh.hash = *MakeBlockId(buffer)
	return nil
}

func (bh *BlockHeader) String() string {
	data, err := json.Marshal(bh)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (bh *BlockHeader) EncodeToBytes() ([]byte, error) {
	return proto.Marshal(bh.Serialize())
}

func (bh *BlockHeader) DecodeFromBytes(buff []byte) error {
	var protoBlockHeader protobuf.BlockHeader
	if err := proto.Unmarshal(buff, &protoBlockHeader); err != nil {
		return err
	}
	return bh.Deserialize(&protoBlockHeader)
}

func (bh *BlockHeader) IsGensis() bool {
	return bh.Height == 0 && bh.Prev.IsEmpty()
}

func GetMakeTreeID(txs map[math.Hash][]byte) (math.Hash, error) {
	trie := new(trie.Trie)
	for k, v := range txs {
		trie.Update(k.Bytes(), v)
	}
	return trie.Hash(), nil
}

type Blocks []*Block

type BlockBy func(b1, b2 *Block) bool

func (self BlockBy) Sort(blocks Blocks) {
	bs := blockSorter{
		blocks: blocks,
		by:     self,
	}
	sort.Sort(bs)
}

type blockSorter struct {
	blocks Blocks
	by     func(b1, b2 *Block) bool
}

func (self blockSorter) Len() int { return len(self.blocks) }
func (self blockSorter) Swap(i, j int) {
	self.blocks[i], self.blocks[j] = self.blocks[j], self.blocks[i]
}
func (self blockSorter) Less(i, j int) bool { return self.by(self.blocks[i], self.blocks[j]) }

func Number(b1, b2 *Block) bool { return (*b1).Header.Height < ((*b2).Header.Height) }
