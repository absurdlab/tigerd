package jose

import "github.com/go-jose/go-jose/v3"

// supported signature algorithms
const (
	RS256 = string(jose.RS256)
	RS384 = string(jose.RS384)
	RS512 = string(jose.RS512)
	PS256 = string(jose.PS256)
	PS384 = string(jose.PS384)
	PS512 = string(jose.PS512)
	ES256 = string(jose.ES256)
	ES384 = string(jose.ES384)
	ES512 = string(jose.ES512)
)

// supported encryption algorithms
//goland:noinspection GoSnakeCaseUsage
const (
	RSA1_5         = string(jose.RSA1_5)
	RSA_OAEP       = string(jose.RSA_OAEP)
	RSA_OAEP_256   = string(jose.RSA_OAEP_256)
	DIRECT         = string(jose.DIRECT)
	ECDH_ES        = string(jose.ECDH_ES)
	ECDH_ES_A128KW = string(jose.ECDH_ES_A128KW)
	ECDH_ES_A192KW = string(jose.ECDH_ES_A192KW)
	ECDH_ES_A256KW = string(jose.ECDH_ES_A256KW)
)

// supported encryption encodings
//goland:noinspection GoSnakeCaseUsage
const (
	A128CBC_HS256 = string(jose.A128CBC_HS256)
	A192CBC_HS384 = string(jose.A192CBC_HS384)
	A256CBC_HS512 = string(jose.A256CBC_HS512)
	A128GCM       = string(jose.A128GCM)
	A192GCM       = string(jose.A192GCM)
	A256GCM       = string(jose.A256GCM)
)

const (
	NONE = "none"
)

// IsNoneOrEmpty checks if value is empty or equals to "none".
func IsNoneOrEmpty(value string) bool {
	switch value {
	case "", NONE:
		return true
	default:
		return false
	}
}
