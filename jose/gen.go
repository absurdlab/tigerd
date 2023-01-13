package jose

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"github.com/go-jose/go-jose/v3"
)

// GenerateSignatureKey generates a signature Key with the given kid and algorithm.
func GenerateSignatureKey(kid string, alg string, bits int) *JSONWebKey {
	var key any

	switch alg {
	case RS256, RS384, RS512,
		ES256, ES384, ES512,
		PS256, PS384, PS512:
		key = mustGenSig(alg, bits)
	default:
		panic("unsupported algorithm")
	}

	return &JSONWebKey{
		Key:       key,
		KeyID:     kid,
		Algorithm: alg,
		Use:       UseSig,
	}
}

// GenerateEncryptionKey generates an encryption Key with the given kid and algorithm.
func GenerateEncryptionKey(kid string, alg string, bits int) *JSONWebKey {
	var key any

	switch alg {
	case RSA1_5, RSA_OAEP, RSA_OAEP_256,
		ECDH_ES, ECDH_ES_A128KW, ECDH_ES_A192KW, ECDH_ES_A256KW:
		key = mustGenEnc(alg, bits)
	default:
		panic("unsupported algorithm")
	}

	return &JSONWebKey{
		Key:       key,
		KeyID:     kid,
		Algorithm: alg,
		Use:       UseEnc,
	}
}

func mustGenSig(alg string, bits int) any {
	_, pk, err := keygenSig(jose.SignatureAlgorithm(alg), bits)
	if err != nil {
		panic(err)
	}
	return pk
}

func mustGenEnc(alg string, bits int) any {
	_, pk, err := keygenEnc(jose.KeyAlgorithm(alg), bits)
	if err != nil {
		panic(err)
	}
	return pk
}

// This method is copied directly from gopkg.in/square/go-jose.v2/jwk-keygen package.
func keygenSig(alg jose.SignatureAlgorithm, bits int) (crypto.PublicKey, crypto.PrivateKey, error) {
	switch alg {
	case jose.ES256, jose.ES384, jose.ES512, jose.EdDSA:
		keylen := map[jose.SignatureAlgorithm]int{
			jose.ES256: 256,
			jose.ES384: 384,
			jose.ES512: 521, // sic!
			jose.EdDSA: 256,
		}
		if bits != 0 && bits != keylen[alg] {
			return nil, nil, errors.New("this alg does not support arbitrary key length")
		}
	case jose.RS256, jose.RS384, jose.RS512, jose.PS256, jose.PS384, jose.PS512:
		if bits == 0 {
			bits = 2048
		}
		if bits < 2048 {
			return nil, nil, errors.New("too short key for RSA `alg`, 2048+ should required")
		}
	}
	switch alg {
	case jose.ES256:
		// The cryptographic operations are implemented using constant-time algorithms.
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		return key.Public(), key, err
	case jose.ES384:
		// NB: The cryptographic operations do not use constant-time algorithms.
		key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		return key.Public(), key, err
	case jose.ES512:
		// NB: The cryptographic operations do not use constant-time algorithms.
		key, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		return key.Public(), key, err
	case jose.EdDSA:
		pub, key, err := ed25519.GenerateKey(rand.Reader)
		return pub, key, err
	case jose.RS256, jose.RS384, jose.RS512, jose.PS256, jose.PS384, jose.PS512:
		key, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			return nil, nil, err
		}
		return key.Public(), key, err
	default:
		return nil, nil, errors.New("unknown `alg` for `use` = `sig`")
	}
}

// This method is copied directly from gopkg.in/square/go-jose.v2/jwk-keygen package.
func keygenEnc(alg jose.KeyAlgorithm, bits int) (crypto.PublicKey, crypto.PrivateKey, error) {
	switch alg {
	case jose.RSA1_5, jose.RSA_OAEP, jose.RSA_OAEP_256:
		if bits == 0 {
			bits = 2048
		}
		if bits < 2048 {
			return nil, nil, errors.New("too short key for RSA `alg`, 2048+ should required")
		}
		key, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			return nil, nil, err
		}
		return key.Public(), key, err
	case jose.ECDH_ES, jose.ECDH_ES_A128KW, jose.ECDH_ES_A192KW, jose.ECDH_ES_A256KW:
		var crv elliptic.Curve
		switch bits {
		case 0, 256:
			crv = elliptic.P256()
		case 384:
			crv = elliptic.P384()
		case 521:
			crv = elliptic.P521()
		default:
			return nil, nil, errors.New("unknown elliptic curve bit length, use one of 256, 384, 521")
		}
		key, err := ecdsa.GenerateKey(crv, rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		return key.Public(), key, err
	default:
		return nil, nil, errors.New("unknown `alg` for `use` = `enc`")
	}
}
