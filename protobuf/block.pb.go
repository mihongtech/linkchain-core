// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/block.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BlockHeader struct {
	Version              *uint32    `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Height               *uint32    `protobuf:"varint,2,req,name=height" json:"height,omitempty"`
	Time                 *int64     `protobuf:"varint,3,req,name=time" json:"time,omitempty"`
	Nounce               *uint32    `protobuf:"varint,4,req,name=nounce" json:"nounce,omitempty"`
	Difficulty           *uint32    `protobuf:"varint,5,req,name=difficulty" json:"difficulty,omitempty"`
	Prev                 *Hash      `protobuf:"bytes,6,req,name=prev" json:"prev,omitempty"`
	TxRoot               *Hash      `protobuf:"bytes,7,req,name=txRoot" json:"txRoot,omitempty"`
	Status               *Hash      `protobuf:"bytes,8,req,name=status" json:"status,omitempty"`
	Sign                 *Signature `protobuf:"bytes,9,opt,name=sign" json:"sign,omitempty"`
	Data                 []byte     `protobuf:"bytes,10,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a48bcf14e684fd, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *BlockHeader) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *BlockHeader) GetNounce() uint32 {
	if m != nil && m.Nounce != nil {
		return *m.Nounce
	}
	return 0
}

func (m *BlockHeader) GetDifficulty() uint32 {
	if m != nil && m.Difficulty != nil {
		return *m.Difficulty
	}
	return 0
}

func (m *BlockHeader) GetPrev() *Hash {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *BlockHeader) GetTxRoot() *Hash {
	if m != nil {
		return m.TxRoot
	}
	return nil
}

func (m *BlockHeader) GetStatus() *Hash {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BlockHeader) GetSign() *Signature {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Block struct {
	Header               *BlockHeader  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	TxList               *Transactions `protobuf:"bytes,2,req,name=txList" json:"txList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a48bcf14e684fd, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTxList() *Transactions {
	if m != nil {
		return m.TxList
	}
	return nil
}

type Blocks struct {
	Block                []*Block `protobuf:"bytes,1,rep,name=block" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blocks) Reset()         { *m = Blocks{} }
func (m *Blocks) String() string { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()    {}
func (*Blocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a48bcf14e684fd, []int{2}
}

func (m *Blocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blocks.Unmarshal(m, b)
}
func (m *Blocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blocks.Marshal(b, m, deterministic)
}
func (m *Blocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blocks.Merge(m, src)
}
func (m *Blocks) XXX_Size() int {
	return xxx_messageInfo_Blocks.Size(m)
}
func (m *Blocks) XXX_DiscardUnknown() {
	xxx_messageInfo_Blocks.DiscardUnknown(m)
}

var xxx_messageInfo_Blocks proto.InternalMessageInfo

func (m *Blocks) GetBlock() []*Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "protobuf.BlockHeader")
	proto.RegisterType((*Block)(nil), "protobuf.Block")
	proto.RegisterType((*Blocks)(nil), "protobuf.Blocks")
}

func init() { proto.RegisterFile("protobuf/block.proto", fileDescriptor_65a48bcf14e684fd) }

var fileDescriptor_65a48bcf14e684fd = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdf, 0x4e, 0x83, 0x30,
	0x14, 0xc6, 0x03, 0x63, 0x6c, 0x1e, 0xfc, 0x93, 0x54, 0xb7, 0x34, 0xbb, 0x30, 0x84, 0x44, 0xe5,
	0x46, 0x96, 0xf0, 0x08, 0x5e, 0xed, 0xc2, 0xab, 0xea, 0x0b, 0x74, 0xac, 0x6c, 0x8d, 0xa3, 0x5d,
	0xe8, 0x61, 0xd1, 0x77, 0xf3, 0xe1, 0x4c, 0x0b, 0x8c, 0xc5, 0xec, 0xae, 0xdf, 0xf9, 0xfd, 0x72,
	0xa0, 0x5f, 0xe1, 0xe1, 0x50, 0x6b, 0xd4, 0xeb, 0xa6, 0x5c, 0xae, 0xf7, 0xba, 0xf8, 0xca, 0x5c,
	0x24, 0xd3, 0x7e, 0xba, 0x98, 0x9d, 0x78, 0xa1, 0xab, 0x4a, 0xab, 0x56, 0x58, 0x2c, 0x4e, 0x63,
	0xac, 0xb9, 0x32, 0xbc, 0x40, 0xd9, 0xb3, 0xe4, 0xd7, 0x87, 0xe8, 0xcd, 0x2e, 0x5b, 0x09, 0xbe,
	0x11, 0x35, 0xa1, 0x30, 0x39, 0x8a, 0xda, 0x48, 0xad, 0xa8, 0x17, 0xfb, 0xe9, 0x0d, 0xeb, 0x23,
	0x99, 0x43, 0xb8, 0x13, 0x72, 0xbb, 0x43, 0xea, 0x3b, 0xd0, 0x25, 0x42, 0x20, 0x40, 0x59, 0x09,
	0x3a, 0x8a, 0xfd, 0x74, 0xc4, 0xdc, 0xd9, 0xba, 0x4a, 0x37, 0xaa, 0x10, 0x34, 0x68, 0xdd, 0x36,
	0x91, 0x47, 0x80, 0x8d, 0x2c, 0x4b, 0x59, 0x34, 0x7b, 0xfc, 0xa1, 0x63, 0xc7, 0xce, 0x26, 0x24,
	0x81, 0xe0, 0x50, 0x8b, 0x23, 0x0d, 0x63, 0x3f, 0x8d, 0xf2, 0xdb, 0xac, 0xff, 0xf1, 0x6c, 0xc5,
	0xcd, 0x8e, 0x39, 0x46, 0x9e, 0x21, 0xc4, 0x6f, 0xa6, 0x35, 0xd2, 0xc9, 0x45, 0xab, 0xa3, 0xd6,
	0x33, 0xc8, 0xb1, 0x31, 0x74, 0x7a, 0xd9, 0x6b, 0x29, 0x79, 0x81, 0xc0, 0xc8, 0xad, 0xa2, 0x57,
	0xb1, 0x97, 0x46, 0xf9, 0xfd, 0x60, 0x7d, 0xc8, 0xad, 0xe2, 0xd8, 0xd4, 0x82, 0x39, 0xc1, 0x5e,
	0x74, 0xc3, 0x91, 0x53, 0x88, 0xbd, 0xf4, 0x9a, 0xb9, 0x73, 0x52, 0xc2, 0xd8, 0xb5, 0x47, 0x5e,
	0x6d, 0x3b, 0xb6, 0x41, 0x57, 0x5b, 0x94, 0xcf, 0x86, 0x3d, 0x67, 0xf5, 0xb2, 0x4e, 0x22, 0x99,
	0xbd, 0xc4, 0xbb, 0x34, 0x6d, 0x99, 0x51, 0x3e, 0x1f, 0xf4, 0xcf, 0xe1, 0x8d, 0x0c, 0xeb, 0xac,
	0x64, 0x09, 0xa1, 0x5b, 0x63, 0xc8, 0x13, 0x8c, 0xdd, 0xe3, 0x53, 0x2f, 0x1e, 0xa5, 0x51, 0x7e,
	0xf7, 0xef, 0x3b, 0xac, 0xa5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x1c, 0x98, 0x42, 0x2b,
	0x02, 0x00, 0x00,
}
