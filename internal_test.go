// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
This test file is part of the mvcutil package rather than than the
mvcutil_test package so it can bridge access to the internals to properly test
cases which are either not possible or can't reliably be tested via the public
interface. The functions are only exported while the tests are being run.
*/

package mvcutil

import (
	"github.com/metasv/mvcd/bsvec"
	"github.com/metasv/mvcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// SetBlockBytes sets the internal serialized block byte buffer to the passed
// buffer.  It is used to inject errors and is only available to the test
// package.
func (b *Block) SetBlockBytes(buf []byte) {
	b.serializedBlock = buf
}

// TstAppDataDir makes the internal appDataDir function available to the test
// package.
func TstAppDataDir(goos, appName string, roaming bool) string {
	return appDataDir(goos, appName, roaming)
}

// TstLegacyAddressPubKeyHash makes a LegacyAddressPubKeyHash, setting the
// unexported fields with the parameters hash and netID.
func TstLegacyAddressPubKeyHash(hash [ripemd160.Size]byte,
	netID byte) *AddressPubKeyHash {

	return &AddressPubKeyHash{
		hash:  hash,
		netID: netID,
	}
}

// TstLegacyAddressScriptHash makes a LegacyAddressScriptHash, setting the
// unexported fields with the parameters hash and netID.
func TstLegacyAddressScriptHash(hash [ripemd160.Size]byte,
	netID byte) *AddressScriptHash {

	return &AddressScriptHash{
		hash:  hash,
		netID: netID,
	}
}

// TstAddressPubKey makes an AddressPubKey, setting the unexported fields with
// the parameters.
func TstAddressPubKey(serializedPubKey []byte, pubKeyFormat PubKeyFormat,
	netID byte) *AddressPubKey {

	pubKey, _ := bsvec.ParsePubKey(serializedPubKey, bsvec.S256())
	return &AddressPubKey{
		pubKeyFormat: pubKeyFormat,
		pubKey:       (*bsvec.PublicKey)(pubKey),
		pubKeyHashID: netID,
	}
}

// TstLegacyAddressSAddr returns the expected script address bytes for
// P2PKH and P2SH legacy addresses.
func TstLegacyAddressSAddr(addr string) []byte {
	decoded := base58.Decode(addr)
	return decoded[1 : 1+ripemd160.Size]
}
