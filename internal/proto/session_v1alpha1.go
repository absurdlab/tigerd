package proto

import (
	"absurdlab.io/tigerd/oidc"
	flowv1alpha1 "buf.build/gen/go/absurdlab/tigerdapis/protocolbuffers/go/flow/v1alpha1"
	"github.com/samber/lo"
	"strings"
)

func ToOpenIDConnectStateV1Alpha1(s *oidc.Session, c *oidc.ClientRegistration) *flowv1alpha1.OpenIDConnectState {
	if s == nil {
		return nil
	}

	return &flowv1alpha1.OpenIDConnectState{
		Sid: s.Id,
		Oidc: &flowv1alpha1.OpenIDConnectParameters{
			Client:               ToClientPublicDataV1Alpha1(c),
			Display:              s.AuthenticationRequest.Display,
			LoginHint:            s.AuthenticationRequest.LoginHint,
			AcrValues:            spaceSplit(s.AuthenticationRequest.AcrValues),
			UiLocales:            spaceSplit(s.AuthenticationRequest.UiLocales),
			ProjectedClaims:      ToClaimsRequestV1Alpha1(s.AuthenticationRequest.Claims),
			ClaimsLocales:        spaceSplit(s.AuthenticationRequest.ClaimsLocales),
			AccessTokenAudiences: spaceSplit(s.AuthenticationRequest.AccessTokenAudiences),
			Scopes:               spaceSplit(s.AuthenticationRequest.Scope),
			MaxAge:               protoDurationOfSeconds(s.AuthenticationRequest.MaxAge),
		},
		Auth:    ToAuthenticationV1Alpha1(s.Authentication),
		Consent: ToConsentV1Alpha1(s.Consent).GetScopes(),
		Context: s.Context,
	}
}

func ToAuthenticationV1Alpha1(u *oidc.Authentication) *flowv1alpha1.Authentication {
	if u == nil {
		return nil
	}

	return &flowv1alpha1.Authentication{
		Subject:     u.Subject,
		AuthTime:    protoTimestampOfEpochSeconds(u.AuthTime),
		Amr:         u.Amr,
		Acr:         u.Acr,
		RememberFor: protoDurationOfSeconds(u.Remember),
	}
}

func ToConsentV1Alpha1(c map[string]*oidc.ScopeState) *flowv1alpha1.Consent {
	if len(c) == 0 {
		return nil
	}

	return &flowv1alpha1.Consent{
		Scopes: lo.MapValues(c, func(value *oidc.ScopeState, key string) *flowv1alpha1.Scope {
			s := &flowv1alpha1.Scope{
				Value:     key,
				Decision:  flowv1alpha1.Scope_DECISION_UNSPECIFIED,
				System:    value.System,
				GrantOnce: value.Once,
			}

			if value.Decision != nil {
				if *value.Decision {
					s.Decision = flowv1alpha1.Scope_DECISION_GRANT
				} else {
					s.Decision = flowv1alpha1.Scope_DECISION_DENY
				}
			}

			return s
		}),
	}
}

func spaceSplit(value string) []string {
	return strings.Split(value, " ")
}
