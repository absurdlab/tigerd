package oidc

import (
	"absurdlab.io/tigerd/jose"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

// NewProvider return a new Provider configuration that reflects this implementation. The Provider metadata depends on
// the baseUrl value which may change for different deployments.
func NewProvider(baseUrl string) *Provider {
	return &Provider{
		Issuer:                                     baseUrl,
		AuthorizationEndpoint:                      fmt.Sprintf("%s/oauth/authorize", baseUrl),
		ResumeAuthorizationEndpoint:                fmt.Sprintf("%s/oauth/authorize/resume", baseUrl),
		TokenEndpoint:                              fmt.Sprintf("%s/oauth/token", baseUrl),
		UserInfoEndpoint:                           fmt.Sprintf("%s/userinfo", baseUrl),
		JSONWebKeySetURI:                           fmt.Sprintf("%s/jwks.json", baseUrl),
		RegistrationEndpoint:                       fmt.Sprintf("%s/register", baseUrl),
		RegistrationRenewSecretEndpoint:            fmt.Sprintf("%s/register/secret", baseUrl),
		ServiceDocumentation:                       fmt.Sprintf("%s/docs", baseUrl),
		OPPolicyURI:                                fmt.Sprintf("%s/policy", baseUrl),
		OPTermsOfServiceURI:                        fmt.Sprintf("%s/tos", baseUrl),
		ScopesSupported:                            []string{ScopeOpenId, ScopeOfflineAccess, ScopeProfile, ScopeEmail, ScopeAddress},
		ResponseTypesSupported:                     []string{ResponseTypeCode, ResponseTypeIdToken, strings.Join([]string{ResponseTypeCode, ResponseTypeIdToken}, space)},
		ResponseModesSupported:                     []string{ResponseModeQuery, ResponseModeFragment},
		GrantTypesSupported:                        []string{GrantTypeAuthorizationCode, GrantTypeImplicit, GrantTypeClientCredentials, GrantTypeRefreshToken},
		SubjectTypesSupported:                      []string{SubjectTypePublic, SubjectTypePairwise},
		IdTokenSigningAlgValuesSupported:           []string{jose.RS256, jose.PS256, jose.ES256, jose.NONE},
		IdTokenEncryptionAlgValuesSupported:        []string{jose.RSA_OAEP_256, jose.ECDH_ES, jose.ECDH_ES_A256KW, jose.NONE},
		IdTokenEncryptionEncValuesSupported:        []string{jose.A256CBC_HS512, jose.A256GCM, jose.NONE},
		UserInfoSigningAlgValuesSupported:          []string{jose.RS256, jose.PS256, jose.ES256, jose.NONE},
		UserInfoEncryptionAlgValuesSupported:       []string{jose.RSA_OAEP_256, jose.ECDH_ES, jose.ECDH_ES_A256KW, jose.NONE},
		UserInfoEncryptionEncValuesSupported:       []string{jose.A256CBC_HS512, jose.A256GCM, jose.NONE},
		RequestObjectSigningAlgValuesSupported:     []string{jose.RS256, jose.PS256, jose.ES256, jose.NONE},
		RequestObjectEncryptionAlgValuesSupported:  []string{jose.RSA_OAEP_256, jose.ECDH_ES, jose.ECDH_ES_A256KW, jose.NONE},
		RequestObjectEncryptionEncValuesSupported:  []string{jose.A256CBC_HS512, jose.A256GCM, jose.NONE},
		TokenEndpointAuthMethodsSupported:          []string{ClientSecretBasic, ClientSecretPost, PrivateKeyJwt},
		TokenEndpointAuthSigningAlgValuesSupported: []string{jose.RS256, jose.ES256, jose.PS256},
		DisplayValuesSupported:                     []string{"page"},
		ClaimTypesSupported:                        []string{"normal"},
		ClaimsSupported:                            []string{ClaimSub, ClaimAuthTime, ClaimAmr, ClaimAcr, ClaimNonce, ClaimCodeHash},
		ClaimsParameterSupported:                   true,
		RequestParameterSupported:                  true,
		RequestURIParameterSupported:               true,
		RequireRequestURIRegistration:              true,
	}
}

// Provider is the OpenID Connect provider metadata defined in https://openid.net/specs/openid-connect-discovery-1_0.html.
type Provider struct {
	Issuer                                     string   `json:"issuer"`
	AuthorizationEndpoint                      string   `json:"authorization_endpoint"`
	ResumeAuthorizationEndpoint                string   `json:"resume_authorization_endpoint"`
	TokenEndpoint                              string   `json:"token_endpoint"`
	UserInfoEndpoint                           string   `json:"userinfo_endpoint"`
	JSONWebKeySetURI                           string   `json:"jwks_uri"`
	RegistrationEndpoint                       string   `json:"registration_endpoint"`
	RegistrationRenewSecretEndpoint            string   `json:"registration_renew_secret_endpoint"`
	ScopesSupported                            []string `json:"scopes_supported"`
	ResponseTypesSupported                     []string `json:"response_types_supported"`
	ResponseModesSupported                     []string `json:"response_modes_supported"`
	GrantTypesSupported                        []string `json:"grant_types_supported"`
	AcrValuesSupported                         []string `json:"acr_values_supported"`
	SubjectTypesSupported                      []string `json:"subject_types_supported"`
	IdTokenSigningAlgValuesSupported           []string `json:"id_token_signing_alg_values_supported"`
	IdTokenEncryptionAlgValuesSupported        []string `json:"id_token_encryption_alg_values_supported"`
	IdTokenEncryptionEncValuesSupported        []string `json:"id_token_encryption_enc_values_supported"`
	UserInfoSigningAlgValuesSupported          []string `json:"userinfo_signing_alg_values_supported"`
	UserInfoEncryptionAlgValuesSupported       []string `json:"userinfo_encryption_alg_values_supported"`
	UserInfoEncryptionEncValuesSupported       []string `json:"userinfo_encryption_enc_values_supported"`
	RequestObjectSigningAlgValuesSupported     []string `json:"request_object_signing_alg_values_supported"`
	RequestObjectEncryptionAlgValuesSupported  []string `json:"request_object_encryption_alg_values_supported"`
	RequestObjectEncryptionEncValuesSupported  []string `json:"request_object_encryption_enc_values_supported"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported"`
	DisplayValuesSupported                     []string `json:"display_values_supported"`
	ClaimTypesSupported                        []string `json:"claim_types_supported"`
	ClaimsSupported                            []string `json:"claims_supported"`
	ServiceDocumentation                       string   `json:"service_documentation"`
	UILocalesSupported                         []string `json:"ui_locales_supported"`
	ClaimsParameterSupported                   bool     `json:"claims_parameter_supported"`
	RequestParameterSupported                  bool     `json:"request_parameter_supported"`
	RequestURIParameterSupported               bool     `json:"request_uri_parameter_supported"`
	RequireRequestURIRegistration              bool     `json:"require_request_uri_registration"`
	OPPolicyURI                                string   `json:"op_policy_uri"`
	OPTermsOfServiceURI                        string   `json:"op_tos_uri"`
}

// IsSupportedResponseType checks if the given value is supported as response_type. Allowed input types are string and
// string slice. Any other types causes panic. Empty string does not pass check.
func (p *Provider) IsSupportedResponseType(value interface{}) error {
	switch v := value.(type) {
	case string:
		if !lo.Contains(p.ResponseTypesSupported, v) {
			return fmt.Errorf("[%v] is not supported as response_type", v)
		}
	case []string:
		if !lo.Every(p.ResponseTypesSupported, v) {
			return fmt.Errorf("[%v] contains unsupported response_type", strings.Join(v, "|"))
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for response_type", v))
	}
	return nil
}

// IsSupportedGrantType checks if the given value is supported as grant_type. Allowed input types are string and
// string slice. Any other types causes panic. Empty string does not pass check.
func (p *Provider) IsSupportedGrantType(value interface{}) error {
	switch v := value.(type) {
	case string:
		if !lo.Contains(p.GrantTypesSupported, v) {
			return fmt.Errorf("[%v] is not supported as grant_type", v)
		}
	case []string:
		if !lo.Every(p.GrantTypesSupported, v) {
			return fmt.Errorf("[%v] contains unsupported grant_type", strings.Join(v, "|"))
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for grant_type", v))
	}
	return nil
}

// IsSupportedResponseMode checks if the given value is supported as response_mode. Only string is allowed as input.
// Any other types causes panic. Empty string passes check.
func (p *Provider) IsSupportedResponseMode(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.ResponseModesSupported, v) {
			return fmt.Errorf("[%v] is not supported as response_mode", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for response_mode", v))
	}
	return nil
}

// IsSupportedSubjectType checks if the given value is supported as subject_type. Only string is allowed as input.
// Any other types causes panic. Empty string passes check.
func (p *Provider) IsSupportedSubjectType(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.SubjectTypesSupported, v) {
			return fmt.Errorf("[%v] is not supported as subject_type", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for subject_type", v))
	}
	return nil
}

// IsSupportedIdTokenSigningAlg checks if the given value is supported as id_token signing algorithm. Only string is
// allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedIdTokenSigningAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.IdTokenSigningAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as id_token signing algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for id_token signing algorithm", v))
	}
	return nil
}

// IsSupportedIdTokenEncryptionAlg checks if the given value is supported as id_token encryption algorithm. Only string
// is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedIdTokenEncryptionAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.IdTokenEncryptionAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as id_token encryption algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for id_token encryption algorithm", v))
	}
	return nil
}

// IsSupportedIdTokenEncryptionEnc checks if the given value is supported as id_token encryption encoding algorithm.
// Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedIdTokenEncryptionEnc(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.IdTokenEncryptionEncValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as id_token encryption encoding", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for id_token encryption encoding", v))
	}
	return nil
}

// IsSupportedUserInfoSigningAlg checks if the given value is supported as userinfo signing algorithm. Only string is
// allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedUserInfoSigningAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.UserInfoSigningAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as userinfo signing algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for userinfo signing algorithm", v))
	}
	return nil
}

// IsSupportedUserInfoEncryptionAlg checks if the given value is supported as userinfo encryption algorithm. Only string
// is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedUserInfoEncryptionAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.UserInfoEncryptionAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as userinfo encryption algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for userinfo encryption algorithm", v))
	}
	return nil
}

// IsSupportedUserInfoEncryptionEnc checks if the given value is supported as userinfo encryption encoding algorithm.
// Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedUserInfoEncryptionEnc(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.UserInfoEncryptionEncValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as userinfo encryption encoding", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for userinfo encryption encoding", v))
	}
	return nil
}

// IsSupportedRequestObjectSigningAlg checks if the given value is supported as request object signing algorithm.
// Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedRequestObjectSigningAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.RequestObjectSigningAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as request object signing algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for request object signing algorithm", v))
	}
	return nil
}

// IsSupportedRequestObjectEncryptionAlg checks if the given value is supported as request object encryption algorithm.
// Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedRequestObjectEncryptionAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.RequestObjectEncryptionAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as request object encryption algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for request object encryption algorithm", v))
	}
	return nil
}

// IsSupportedRequestObjectEncryptionEnc checks if the given value is supported as request object encryption encoding
// algorithm. Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedRequestObjectEncryptionEnc(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.RequestObjectEncryptionEncValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as request object encryption encoding", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for request object encryption encoding", v))
	}
	return nil
}

// IsSupportedTokenEndpointAuthMethod checks if the given value is supported as token endpoint authentication method.
// Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedTokenEndpointAuthMethod(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.TokenEndpointAuthMethodsSupported, v) {
			return fmt.Errorf("[%v] is not supported as token endpoint auth method", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for token endpoint auth method", v))
	}
	return nil
}

// IsSupportedTokenEndpointAuthSigningAlg checks if the given value is supported as token endpoint auth signing
// algorithm. Only string is allowed as input. Any other type causes panic. Empty value passes check.
func (p *Provider) IsSupportedTokenEndpointAuthSigningAlg(value interface{}) error {
	switch v := value.(type) {
	case string:
		if len(v) > 0 && !lo.Contains(p.TokenEndpointAuthSigningAlgValuesSupported, v) {
			return fmt.Errorf("[%v] is not supported as token endpoint auth signing algorithm", v)
		}
	default:
		panic(fmt.Errorf("unexpected type [%T] for token endpoint auth signing algorithm", v))
	}
	return nil
}
