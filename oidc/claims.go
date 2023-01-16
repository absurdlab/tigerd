package oidc

import (
	"encoding/json"
)

// ClaimsParameter represents the content of the "claim" parameter.
type ClaimsParameter struct {
	IdToken  map[string]*ClaimOption `json:"id_token,omitempty"`
	Userinfo map[string]*ClaimOption `json:"userinfo,omitempty"`
}

// UnmarshalParam implicitly implements echo.BindMarshaler method to allow easy request parameter binding.
func (p *ClaimsParameter) UnmarshalParam(param string) error {
	return json.Unmarshal([]byte(param), &p)
}

// ClaimOption is the option for a single claim request.
type ClaimOption struct {
	Essential bool     `json:"essential"`
	Value     string   `json:"value,omitempty"`
	Values    []string `json:"values,omitempty"`
}
