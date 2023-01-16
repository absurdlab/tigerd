package proto

import (
	"absurdlab.io/tigerd/oidc"
	clientv1alpha1 "buf.build/gen/go/absurdlab/tigerdapis/protocolbuffers/go/client/v1alpha1"
)

func ToClientRegistrationV1Alpha1(r *oidc.ClientRegistration) *clientv1alpha1.Registration {
	if r == nil {
		return nil
	}

	return &clientv1alpha1.Registration{
		ClientId:                         r.Id,
		ClientIdIssuedAt:                 protoTimestampOfEpochSeconds(r.CreatedAt),
		ClientSecret:                     r.Secret,
		ClientSecretExpireAt:             protoTimestampOfEpochSeconds(r.SecretExpireAt),
		RegistrationAccessToken:          r.RegistrationAccessToken,
		RegistrationClientUri:            r.RegistrationClientUri,
		RegistrationClientRenewSecretUri: r.RegistrationClientRenewSecretUri,
		Metadata:                         ToClientMetadataV1Alpha1(r.ClientMetadata),
	}
}

func ToClientPublicDataV1Alpha1(r *oidc.ClientRegistration) *clientv1alpha1.PublicData {
	return &clientv1alpha1.PublicData{
		ClientId:   r.Id,
		ClientName: r.ClientMetadata.Name,
		ClientUri:  r.ClientMetadata.ClientUri,
		LogoUri:    r.ClientMetadata.LogoUri,
		PolicyUri:  r.ClientMetadata.PolicyUri,
		TosUri:     r.ClientMetadata.TosUri,
		Contacts:   r.ClientMetadata.Contacts,
	}
}

func ToClientMetadataV1Alpha1(m *oidc.ClientMetadata) *clientv1alpha1.Metadata {
	if m == nil {
		return nil
	}

	return &clientv1alpha1.Metadata{
		ClientName:                   m.Name,
		Contacts:                     m.Contacts,
		LogoUri:                      m.LogoUri,
		ClientUri:                    m.ClientUri,
		PolicyUri:                    m.PolicyUri,
		TosUri:                       m.TosUri,
		RedirectUris:                 m.RedirectUris,
		PostLogoutRedirectUris:       m.PostLogoutRedirectUris,
		ResponseTypes:                m.ResponseTypes,
		GrantTypes:                   m.GrantTypes,
		Scopes:                       m.Scopes,
		ApplicationType:              m.ApplicationType,
		JwksUri:                      m.JwksUri,
		Jwks:                         m.Jwks.String(),
		SubjectType:                  m.SubjectType,
		SectorIdentifierUri:          m.SectorIdentifierUri,
		IdTokenSignedResponseAlg:     m.IdTokenSignedResponseAlg,
		IdTokenEncryptedResponseAlg:  m.IdTokenEncryptedResponseAlg,
		IdTokenEncryptedResponseEnc:  m.IdTokenEncryptedResponseEnc,
		UserinfoSignedResponseAlg:    m.UserinfoSignedResponseAlg,
		UserinfoEncryptedResponseAlg: m.UserinfoEncryptedResponseAlg,
		UserinfoEncryptedResponseEnc: m.UserinfoEncryptedResponseEnc,
		RequestObjectSigningAlg:      m.RequestObjectSigningAlg,
		RequestObjectEncryptionAlg:   m.RequestObjectEncryptionAlg,
		RequestObjectEncryptionEnc:   m.RequestObjectEncryptionEnc,
		TokenEndpointAuthMethod:      m.TokenEndpointAuthMethod,
		TokenEndpointAuthSigningAlg:  m.TokenEndpointAuthSigningAlg,
		DefaultMaxAge:                protoDurationOfSeconds(m.DefaultMaxAge),
		RequireAuthTime:              m.RequireAuthTime,
		DefaultAcrValues:             m.DefaultAcrValues,
		RequestUris:                  m.RequestUris,
		InteractionRedirectUri:       m.InteractionRedirectUri,
		PreferJwtAccessToken:         m.PreferJwtAccessToken,
		AccessTokenAudiences:         m.AccessTokenAudiences,
		AuthorizationCodeGrantAccessTokenLifespanOverride:  protoDurationOfSeconds(m.AuthorizationCodeGrantAccessTokenLifespanOverride),
		AuthorizationCodeGrantRefreshTokenLifespanOverride: protoDurationOfSeconds(m.AuthorizationCodeGrantRefreshTokenLifespanOverride),
		AuthorizationCodeGrantIdTokenLifespanOverride:      protoDurationOfSeconds(m.AuthorizationCodeGrantIdTokenLifespanOverride),
		ImplicitGrantAccessTokenLifespanOverride:           protoDurationOfSeconds(m.ImplicitGrantAccessTokenLifespanOverride),
		ImplicitGrantIdTokenLifespanOverride:               protoDurationOfSeconds(m.ImplicitGrantIdTokenLifespanOverride),
		ClientCredentialsGrantAccessTokenLifespanOverride:  protoDurationOfSeconds(m.ClientCredentialsGrantAccessTokenLifespanOverride),
		RefreshTokenGrantAccessTokenLifespanOverride:       protoDurationOfSeconds(m.RefreshTokenGrantAccessTokenLifespanOverride),
		RefreshTokenGrantRefreshTokenLifespanOverride:      protoDurationOfSeconds(m.RefreshTokenGrantRefreshTokenLifespanOverride),
		RefreshTokenGrantIdTokenLifespanOverride:           protoDurationOfSeconds(m.RefreshTokenGrantIdTokenLifespanOverride),
	}
}
