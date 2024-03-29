package math

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/mihongtech/linkchain-core/common"
	"github.com/mihongtech/linkchain-core/common/serialize"
	"github.com/mihongtech/linkchain-core/protobuf"
)

// HashSize of array used to store hashes.  See Hash.
const HashSize = 32

// MaxHashStringSize is the maximum length of a Hash hash string.
const MaxHashStringSize = HashSize * 2

// ErrHashStrSize describes an error that indicates the caller specified a hash
// string that has too many characters.
var ErrHashStrSize = fmt.Errorf("max hash string length is %v bytes", MaxHashStringSize)

// Hash is used in several of the bitcoin messages and common structures.  It
// typically represents the double sha256 of data.
type Hash [HashSize]byte

func (hash Hash) GetString() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

func (hash Hash) IsEmpty() bool {
	isEmpty := true
	for i := 0; i < HashSize; i++ {
		if hash[i] != 0 {
			isEmpty = false
			break
		}
	}
	return isEmpty
}

// CloneBytes returns a copy of the bytes which represent the hash as a byte
// slice.
//
// NOTE: It is generally cheaper to just slice the hash directly thereby reusing
// the same bytes rather than calling this method.
func (hash *Hash) CloneBytes() []byte {
	newHash := make([]byte, HashSize)
	copy(newHash, hash[:])

	return newHash
}

// SetBytes sets the bytes which represent the hash.  An error is returned if
// the number of bytes passed in is not HashSize.
func (hash *Hash) SetBytes(newHash []byte) error {
	if len(newHash) > len(hash) {
		newHash = newHash[len(newHash)-HashSize:]
	}

	copy(hash[HashSize-len(newHash):], newHash)
	return nil
}

// IsEqual returns true if target is the same as hash.
func (hash *Hash) IsEqual(h *Hash) bool {
	if hash == nil && h == nil {
		return true
	}
	if hash == nil || h == nil {
		return false
	}

	return *hash == *h
}

// NewHash returns a new Hash from a byte slice.  An error is returned if
// the number of bytes passed in is not HashSize.
func NewHash(newHash []byte) (*Hash, error) {
	var sh Hash
	err := sh.SetBytes(newHash)
	if err != nil {
		return nil, err
	}
	return &sh, err
}

// BytesToHash returns a new Hash from a byte slice.
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

// NewHashFromStr creates a Hash from a hash string.  The string should be
// the hexadecimal string of a byte-reversed hash, but any missing characters
// result in zero padding at the end of the Hash.
func NewHashFromStr(hash string) (*Hash, error) {
	ret := new(Hash)
	err := Decode(ret, hash)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Decode decodes the byte-reversed hexadecimal string encoding of a Hash to a
// destination.
func Decode(dst *Hash, src string) error {
	// Return error if hash string is too long.
	if len(src) > MaxHashStringSize {
		return ErrHashStrSize
	}

	// Hex decoder expects the hash to be a multiple of two.  When not, pad
	// with a leading zero.
	var srcBytes []byte
	if len(src)%2 == 0 {
		srcBytes = []byte(src)
	} else {
		srcBytes = make([]byte, 1+len(src))
		srcBytes[0] = '0'
		copy(srcBytes[1:], src)
	}

	// Hex decode the source bytes to a temporary destination.
	var reversedHash Hash
	_, err := hex.Decode(reversedHash[HashSize-hex.DecodedLen(len(srcBytes)):], srcBytes)
	if err != nil {
		return err
	}

	// Reverse copy from the temporary hash to destination.  Because the
	// temporary was zeroed, the written result will be correctly padded.
	for i, b := range reversedHash[:HashSize/2] {
		dst[i], dst[HashSize-1-i] = reversedHash[HashSize-1-i], b
	}

	return nil
}

//Serialize/Deserialize
func (hash *Hash) Serialize() serialize.SerializeStream {
	h := protobuf.Hash{
		Data: proto.NewBuffer(hash.CloneBytes()).Bytes(),
	}
	return &h
}

func (hash *Hash) Deserialize(s serialize.SerializeStream) error {
	h := *s.(*protobuf.Hash)
	hash.SetBytes(h.Data)
	return nil
}

// String returns the Hash as the hexadecimal string of the byte-reversed
// hash.
func (hash Hash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

func (hash *Hash) EncodeToBytes() ([]byte, error) {
	return proto.Marshal(hash.Serialize())
}

func (hash *Hash) DecodeFromBytes(buff []byte) error {
	var protoHash protobuf.Hash
	if err := proto.Unmarshal(buff, &protoHash); err != nil {
		return err
	}
	return hash.Deserialize(&protoHash)
}

func (hash *Hash) ToString() string {
	return hash.String()
}

// Big converts a hash to a big integer.
func (h Hash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }

func StringToHash(s string) Hash { return BytesToHash([]byte(s)) }
func BigToHash(b *big.Int) Hash  { return BytesToHash(b.Bytes()) }
func HexToHash(s string) Hash    { return BytesToHash(common.FromHex(s)) }

func (h Hash) Bytes() []byte { return h[:] }

//Json Hash convert to Hex
func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *Hash) UnmarshalJSON(data []byte) error {
	str := ""
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	hash, err := NewHashFromStr(str)
	if err != nil {
		return err
	}
	return h.SetBytes(hash.CloneBytes())
}
