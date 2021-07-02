package blockchain

import (
	"github.com/joeqian10/neo3-gogogo-legacy/io"
)

//StorageKey used to store StorageItem on blockchain
type StorageKey struct {
	Id  int
	Key []byte
}

//Deserialize deserialize from byte array
func (sk *StorageKey) Deserialize(br *io.BinaryReader) {
	var id int32
	br.ReadLE(&id)
	sk.Id = int(id)
	sk.Key = br.ReadAllBytes()
}

//Serialize serialize to byte array
func (sk *StorageKey) Serialize(writer *io.BinaryWriter) {
	writer.WriteLE(int32(sk.Id))
	writer.WriteLE(sk.Key)
}
