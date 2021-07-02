package keys

import (
	"github.com/joeqian10/neo3-gogogo-legacy/crypto"
	"github.com/joeqian10/neo3-gogogo-legacy/helper"
	"github.com/joeqian10/neo3-gogogo-legacy/sc"
)

// returns a NEO-specific hash of the public key.
func PublicKeyToScriptHash(p *crypto.ECPoint) *helper.UInt160 {
	b, _ := sc.CreateSignatureRedeemScript(p)
	hash := crypto.Hash160(b)
	return helper.UInt160FromBytes(hash)
}

// returns a base58-encoded NEO-specific address based on the key hash.
func PublicKeyToAddress(p *crypto.ECPoint, version byte) string {
	return crypto.ScriptHashToAddress(PublicKeyToScriptHash(p), version)
}
