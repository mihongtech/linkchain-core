package meta

import (
	"encoding/hex"
	"errors"

	"github.com/mihongtech/linkchain-core/common/btcec"
	"github.com/mihongtech/linkchain-core/common/serialize"
	"github.com/mihongtech/linkchain-core/protobuf"

	"github.com/golang/protobuf/proto"
)

type Signature struct {
	Code []byte `json:"code"`
}

func NewSignature(code []byte) *Signature {
	return &Signature{Code: code}
}

//Serialize/Deserialize
func (sign *Signature) Serialize() serialize.SerializeStream {
	peer := protobuf.Signature{
		Code: proto.NewBuffer(sign.Code).Bytes(),
	}
	return &peer
}

func (sign *Signature) Deserialize(s serialize.SerializeStream) error {
	data := *s.(*protobuf.Signature)
	sign.Code = data.Code
	return nil
}

func (sign *Signature) String() string {
	return hex.EncodeToString(sign.Code)
}

func (sign *Signature) EncodeToBytes() ([]byte, error) {
	return proto.Marshal(sign.Serialize())
}

func (sign *Signature) DecodeFromBytes(buff []byte) error {
	var protoSignature protobuf.Signature
	if err := proto.Unmarshal(buff, &protoSignature); err != nil {
		return err
	}
	return sign.Deserialize(&protoSignature)
}

func (sign *Signature) Verify(hash []byte, address []byte) error {
	signer, err := btcec.GetSigner(hash, sign.Code)
	if err != nil {
		return err
	}
	id := NewAddress(signer)

	if id.IsEqual(BytesToAddress(address)) {
		return nil
	}
	return errors.New("Verify sign failed")
}
