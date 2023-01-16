package proto

import (
	"absurdlab.io/tigerd/oidc"
	flowv1alpha1 "buf.build/gen/go/absurdlab/tigerdapis/protocolbuffers/go/flow/v1alpha1"
	"github.com/samber/lo"
)

func ToClaimsRequestV1Alpha1(p *oidc.ClaimsParameter) *flowv1alpha1.ClaimsRequest {
	if p == nil {
		return nil
	}

	return &flowv1alpha1.ClaimsRequest{
		IdToken:  lo.MapValues(p.IdToken, claimOptionV1Alpha1),
		Userinfo: lo.MapValues(p.Userinfo, claimOptionV1Alpha1),
	}
}

func claimOptionV1Alpha1(o *oidc.ClaimOption, _ string) *flowv1alpha1.ClaimsRequest_Option {
	if o == nil {
		return nil
	}

	p := &flowv1alpha1.ClaimsRequest_Option{Essential: o.Essential}
	switch {
	case len(o.Value) > 0:
		p.Values = append([]string{}, o.Value)
	case len(o.Values) > 0:
		p.Values = append([]string{}, o.Values...)
	}

	return p
}
