package oidc

import (
	"encoding/json"
	"time"
)

// Session is the overall state during the entire OpenID Connect lifecycle.
type Session struct {
	// Id is the identifier of the Session, issued at creation. This id will be embedded in the grants issued so
	// to relate the grant to the Session.
	Id string `json:"id"`
	// CreatedAt is the creation timestamp of the Session.
	CreatedAt uint64 `json:"created_at,omitempty"`
	// LastUpdatedAt is the last update timestamp of the Session.
	LastUpdatedAt uint64 `json:"last_updated_at,omitempty"`
	// ExpireAt is the expiry timestamp of the Session. A value of zero means this Session does not expire.
	ExpireAt uint64 `json:"expire_at,omitempty"`
	// Client is the snapshot of the requesting ClientRegistration. This might be nil depending on the whether the
	// client is configured to keep_configuration_for_existing_sessions. When nil, read the latest ClientRegistration
	// from database using the client_id from AuthenticationRequest, and keep this property nil.
	Client *ClientRegistration `json:"client,omitempty"`
	// AuthenticationRequest is the effective OpenID Connect authentication request. It has undergone server processing
	// and might not be entirely equivalent to the parameters submitted at the authorization endpoint.
	AuthenticationRequest *AuthenticationRequest `json:"auth_request,omitempty"`
	// Authentication is the authenticated user information
	Authentication *Authentication `json:"auth,omitempty"`
	// Consent is the status of all requested and projected scopes.
	Consent map[string]*ScopeState `json:"consent,omitempty"`
	// Context is the contextual data that gets passed to the next interaction flow. This property is only updated
	// by flow responses. The system does not touch it.
	Context map[string]string `json:"context,omitempty"`
	// Code is authorization code state.
	Code json.RawMessage `json:"code,omitempty"`
	// AccessToken is access token state.
	AccessToken json.RawMessage `json:"access_token,omitempty"`
	// RefreshToken is refresh token state.
	RefreshToken json.RawMessage `json:"refresh_token,omitempty"`
}

func (s *Session) HasExpired() bool {
	return s.ExpireAt > 0 && time.Unix(int64(s.ExpireAt), 0).Before(time.Now().UTC())
}

type Authentication struct {
	Subject  string   `json:"subject,omitempty"`
	AuthTime uint64   `json:"auth_time,omitempty"`
	Amr      []string `json:"amr,omitempty"`
	Acr      string   `json:"acr,omitempty"`
	Remember uint64   `json:"remember,omitempty"`
}

type ScopeState struct {
	// Decision is nil to start with, true if granted and false if rejected.
	Decision *bool `json:"decision,omitempty"`
	// System can be set to true if the grant is made by the system.
	System bool `json:"system,omitempty"`
	// Once can be set to true if the grant is for current request only.
	Once bool `json:"once,omitempty"`
}
