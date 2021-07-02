package crypto

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo-legacy/helper"
)

func BytesToScriptHash(script []byte) *helper.UInt160 {
	return helper.UInt160FromBytes(Hash160(script))
}

func ScriptHashToAddress(scriptHash *helper.UInt160, version byte) string {
	data := append([]byte{version}, scriptHash.ToByteArray()...)
	return Base58CheckEncode(data)
}

func AddressToScriptHash(address string, version byte) (*helper.UInt160, error) {
	data, err := Base58CheckDecode(address)
	if err != nil {
		return nil, err
	}
	if data == nil || len(data) != 21 || data[0] != version {
		return nil, fmt.Errorf("invalid address string")
	}
	return helper.UInt160FromBytes(data[1:]), nil
}
