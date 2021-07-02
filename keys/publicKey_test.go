package keys

import (
	"github.com/joeqian10/neo3-gogogo-legacy/crypto"
	"github.com/joeqian10/neo3-gogogo-legacy/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicKeyToScriptHash(t *testing.T) {
	for _, testCase := range KeyCases {
		pubKey, err := crypto.NewECPointFromString(testCase.PublicKey)
		assert.Nil(t, err)
		scriptHash := PublicKeyToScriptHash(pubKey)
		s := scriptHash.String()
		assert.Equal(t, testCase.ScriptHash, s)
	}
}

func TestPublicKeyToAddress(t *testing.T) {
	for _, testCase := range KeyCases {
		pubKey, err := crypto.NewECPointFromString(testCase.PublicKey)
		assert.Nil(t, err)
		address := PublicKeyToAddress(pubKey, helper.DefaultAddressVersion)
		assert.Equal(t, testCase.Address, address)
	}
}
