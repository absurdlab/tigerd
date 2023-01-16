package oidc

import (
	"absurdlab.io/tigerd/jose"
	"context"
	"time"
)

// ClientRegistration is the full OpenID Connect Dynamic Registration client.
type ClientRegistration struct {
	*ClientMetadata

	// Id is the generated client identifier, issued at time of creation. This value is readonly after creation.
	Id string `json:"client_id,omitempty"`
	// CreatedAt is the creation unix timestamp value. This value is readonly after creation.
	CreatedAt uint64 `json:"client_id_issued_at,omitempty"`
	// LastUpdatedAt is the last update unix timestamp value. This value is updated on every client modification.
	LastUpdatedAt uint64 `json:"client_last_updated_at,omitempty"`
	// Secret contains the client_secret value. It is normally presented in a hash encoded form, so that its original
	// value cannot be recovered. However, temporarily after client registration or client secret renewal, its value
	// may contain the plain-text form of the secret, in order to be presented to the API client. The rule of thumb
	// is that, if this client is read from the database, the Secret value is encoded.
	Secret string `json:"client_secret,omitempty"`
	// SecretExpireAt is the unix time stamp of secret expiry. If value is zero, the secret will not expire.
	SecretExpireAt uint64 `json:"client_secret_expire_at,omitempty"`
	// RegistrationAccessToken is the JWT access_token to access client management endpoints. This value is generated on
	// client registration and regenerated upon every client modification. This access token does NOT expire. It is
	// normally presented in a hash encoded form, so that is original value cannot be recovered. However, after each
	// client modification, its value may contain the plain-text form of the access token, in order to be presented to
	// the API client. The rule of thumb is that, if this client is read from the database, the RegistrationAccessToken
	// value is encoded.
	RegistrationAccessToken string `json:"registration_access_token,omitempty"`
	// RegistrationClientUri is the client registration URL where client data can be managed. This endpoint can be
	// accessed with different HTTP verbs in order to perform different operations.
	//
	//	POST: create client registration
	//	PUT: replace client registration metadata
	//	PATCH: partially modify client registration metadata
	// 	DELETE: remove client registration
	//	GET: retrieve client registration
	//
	// These endpoints are protected by access tokens. For PUT, PATCH, DELETE and GET operations, include RegistrationAccessToken
	// in the "Authorization" http header. For the POST operation, the required access token must be issued for a
	// client with "client:register" scope.
	RegistrationClientUri string `json:"registration_client_uri,omitempty"`
	// RegistrationClientRenewSecretUri is the URL at which the client's secret can be renewed. This endpoint can only
	// be used with the http POST verb. A RegistrationAccessToken must be included in the "Authorization" http header.
	RegistrationClientRenewSecretUri string `json:"registration_client_renew_secret_uri,omitempty"`
	// CachedResources contains all cached resources of this client registration. Cacheable resources include the
	// jwks_uri and each request_uris. A cached entry may have expired. Expired entries will not be consulted by
	// the server, but may still exist until next client update.
	CachedResources map[string]*ClientCacheEntry `json:"cached_resources,omitempty"`
}

// ClientMetadata is client's configurable metadata properties.
type ClientMetadata struct {
	Name                                               string              `json:"client_name,omitempty"`
	Contacts                                           []string            `json:"contacts,omitempty"`
	LogoUri                                            string              `json:"logo_uri,omitempty"`
	ClientUri                                          string              `json:"client_uri,omitempty"`
	PolicyUri                                          string              `json:"policy_uri,omitempty"`
	TosUri                                             string              `json:"tos_uri,omitempty"`
	RedirectUris                                       []string            `json:"redirect_uris,omitempty"`
	PostLogoutRedirectUris                             []string            `json:"post_logout_redirect_uris,omitempty"`
	ResponseTypes                                      []string            `json:"response_types,omitempty"`
	GrantTypes                                         []string            `json:"grant_types,omitempty"`
	Scopes                                             []string            `json:"scopes,omitempty"`
	ApplicationType                                    string              `json:"application_type,omitempty"`
	JwksUri                                            string              `json:"jwks_uri,omitempty"`
	Jwks                                               *jose.JSONWebKeySet `json:"jwks,omitempty"`
	SubjectType                                        string              `json:"subject_type,omitempty"`
	SectorIdentifierUri                                string              `json:"sector_identifier_uri,omitempty"`
	IdTokenSignedResponseAlg                           string              `json:"id_token_signed_response_alg,omitempty"`
	IdTokenEncryptedResponseAlg                        string              `json:"id_token_encrypted_response_alg,omitempty"`
	IdTokenEncryptedResponseEnc                        string              `json:"id_token_encrypted_response_enc,omitempty"`
	UserinfoSignedResponseAlg                          string              `json:"userinfo_signed_response_alg,omitempty"`
	UserinfoEncryptedResponseAlg                       string              `json:"userinfo_encrypted_response_alg,omitempty"`
	UserinfoEncryptedResponseEnc                       string              `json:"userinfo_encrypted_response_enc,omitempty"`
	RequestObjectSigningAlg                            string              `json:"request_object_signing_alg,omitempty"`
	RequestObjectEncryptionAlg                         string              `json:"request_object_encryption_alg,omitempty"`
	RequestObjectEncryptionEnc                         string              `json:"request_object_encryption_enc,omitempty"`
	TokenEndpointAuthMethod                            string              `json:"token_endpoint_auth_method,omitempty"`
	TokenEndpointAuthSigningAlg                        string              `json:"token_endpoint_auth_signing_alg,omitempty"`
	DefaultMaxAge                                      uint64              `json:"default_max_age,omitempty"`
	RequireAuthTime                                    bool                `json:"require_auth_time,omitempty"`
	DefaultAcrValues                                   []string            `json:"default_acr_values,omitempty"`
	RequestUris                                        []string            `json:"request_uris,omitempty"`
	InteractionRedirectUri                             string              `json:"interaction_redirect_uri,omitempty"`
	PreferJwtAccessToken                               bool                `json:"prefer_jwt_access_token,omitempty"`
	AccessTokenAudiences                               []string            `json:"access_token_audiences,omitempty"`
	AuthorizationCodeGrantAccessTokenLifespanOverride  uint64              `json:"authorization_code_grant_access_token_lifespan,omitempty"`
	AuthorizationCodeGrantRefreshTokenLifespanOverride uint64              `json:"authorization_code_grant_refresh_token_lifespan_override,omitempty"`
	AuthorizationCodeGrantIdTokenLifespanOverride      uint64              `json:"authorization_code_grant_id_token_lifespan_override,omitempty"`
	ImplicitGrantAccessTokenLifespanOverride           uint64              `json:"implicit_grant_access_token_lifespan_override,omitempty"`
	ImplicitGrantIdTokenLifespanOverride               uint64              `json:"implicit_grant_id_token_lifespan_override,omitempty"`
	ClientCredentialsGrantAccessTokenLifespanOverride  uint64              `json:"client_credentials_grant_access_token_lifespan_override,omitempty"`
	RefreshTokenGrantAccessTokenLifespanOverride       uint64              `json:"refresh_token_grant_access_token_lifespan_override,omitempty"`
	RefreshTokenGrantRefreshTokenLifespanOverride      uint64              `json:"refresh_token_grant_refresh_token_lifespan_override,omitempty"`
	RefreshTokenGrantIdTokenLifespanOverride           uint64              `json:"refresh_token_grant_id_token_lifespan_override,omitempty"`
}

// ClientCacheEntry is cached value with an optional expiry unix timestamp.
type ClientCacheEntry struct {
	Value    string `json:"value"`
	ExpireAt uint64 `json:"expire_at,omitempty"`
}

func (e *ClientCacheEntry) HasExpired() bool {
	return e.ExpireAt > 0 && time.Unix(int64(e.ExpireAt), 0).Before(time.Now().UTC())
}

// ClientRegistrationTokenIssuer abstracts the behavior for issuing a client registration token.
type ClientRegistrationTokenIssuer interface {
	// IssueClientRegistrationToken creates a JWT based access token with the client as the subject, the client
	// registration endpoint and the client secret renewal endpoint as the sole audience and "client:manage" scope.
	// The created token will be set on the client, and must not have an expiry time.
	IssueClientRegistrationToken(ctx context.Context, client *ClientRegistration) error
}
