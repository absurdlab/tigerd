package oidc

// AuthenticationRequest is the OpenID Connect authentication request.
//
// This object dubs as both container for request parameters and the payload of the request parameter (or request_uri,
// if passed by reference). Parameters specified from both source will be merged to form a single AuthenticationRequest,
// which is used for downstream processing.
type AuthenticationRequest struct {
	ClientId             string           `json:"client_id,omitempty" query:"client_id"`
	ResponseType         string           `json:"response_type,omitempty" query:"response_type"`
	RedirectUri          string           `json:"redirect_uri,omitempty" query:"redirect_uri"`
	Scope                string           `json:"scope,omitempty" query:"scope"`
	State                string           `json:"state,omitempty" query:"state"`
	CodeChallenge        string           `json:"code_challenge,omitempty" query:"code_challenge"`
	CodeChallengeMethod  string           `json:"code_challenge_method,omitempty" query:"code_challenge_method"`
	ResponseMode         string           `json:"response_mode,omitempty" query:"response_mode"`
	Nonce                string           `json:"nonce,omitempty" query:"nonce"`
	Display              string           `json:"display,omitempty" query:"display"`
	Prompt               string           `json:"prompt,omitempty" query:"prompt"`
	LoginHint            string           `json:"login_hint,omitempty" query:"login_hint"`
	MaxAge               uint64           `json:"max_age,omitempty" query:"max_age"`
	UiLocales            string           `json:"ui_locales,omitempty" query:"ui_locales"`
	IdTokenHint          string           `json:"id_token_hint,omitempty" query:"id_token_hint"`
	AcrValues            string           `json:"acr_values,omitempty" query:"acr_values"`
	Claims               *ClaimsParameter `json:"claims,omitempty" query:"claims"`
	ClaimsLocales        string           `json:"claims_locales,omitempty" query:"claims_locales"`
	AccessTokenAudiences string           `json:"access_token_audiences,omitempty" query:"access_token_audiences"`
	RequestUri           string           `json:"-" query:"request_uri"`
	RequestObject        string           `json:"-" query:"request"`
}

// TokenRequest is the OpenID Connect token request.
//
// This object omits the token endpoint authentication parameters, such as client_secret, client_assertion and
// client_assertion_type, or client_id and client_secret value passed in the Authorization header.
type TokenRequest struct {
	ClientId     string `form:"client_id"`
	GrantType    string `form:"grant_type"`
	Scope        string `form:"scope"`
	RedirectUri  string `form:"redirect_uri"`
	Code         string `form:"code"`
	CodeVerifier string `form:"code_verifier"`
	RefreshToken string `form:"refresh_token"`
}
