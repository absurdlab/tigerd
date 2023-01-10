package oidc

const (
	space = " "
)

// response_type.
const (
	ResponseTypeCode    = "code"
	ResponseTypeToken   = "token"
	ResponseTypeIdToken = "id_token"
)

// grant_type. The password grant_type is explicitly unsupported due to security vulnerability.
const (
	GrantTypeAuthorizationCode = "authorization_code"
	GrantTypeImplicit          = "implicit"
	GrantTypeClientCredentials = "client_credentials"
	GrantTypeRefreshToken      = "refresh_token"
)

// response_mode.
const (
	ResponseModeQuery    = "query"
	ResponseModeFragment = "fragment"
)

// subject_type
const (
	SubjectTypePublic   = "public"
	SubjectTypePairwise = "pairwise"
)

// canonical scopes
const (
	ScopeOpenId        = "openid"
	ScopeOfflineAccess = "offline_access"
	ScopeProfile       = "profile"
	ScopeEmail         = "email"
	ScopeAddress       = "address"
)

const (
	ClaimSub      = "sub"
	ClaimAuthTime = "auth_time"
	ClaimAmr      = "amr"
	ClaimAcr      = "acr"
	ClaimNonce    = "nonce"
	ClaimCodeHash = "c_hash"
)

// token_endpoint_auth_method. The client_secret_jwt method is explicitly unsupported due to security vulnerability.
const (
	ClientSecretBasic = "client_secret_basic"
	ClientSecretPost  = "client_secret_post"
	PrivateKeyJwt     = "private_key_jwt"
)

// IsNoneOrEmpty checks if value is empty or equals to "none".
func IsNoneOrEmpty(value string) bool {
	switch value {
	case "", "none":
		return true
	default:
		return false
	}
}
