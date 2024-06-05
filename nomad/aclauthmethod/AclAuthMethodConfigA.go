// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod


type AclAuthMethodConfigA struct {
	// A list of allowed values that can be used for the redirect URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#allowed_redirect_uris AclAuthMethod#allowed_redirect_uris}
	AllowedRedirectUris *[]*string `field:"optional" json:"allowedRedirectUris" yaml:"allowedRedirectUris"`
	// List of auth claims that are valid for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#bound_audiences AclAuthMethod#bound_audiences}
	BoundAudiences *[]*string `field:"optional" json:"boundAudiences" yaml:"boundAudiences"`
	// The value against which to match the iss claim in a JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#bound_issuer AclAuthMethod#bound_issuer}
	BoundIssuer *[]*string `field:"optional" json:"boundIssuer" yaml:"boundIssuer"`
	// Mappings of claims (key) that will be copied to a metadata field (value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#claim_mappings AclAuthMethod#claim_mappings}
	ClaimMappings *map[string]*string `field:"optional" json:"claimMappings" yaml:"claimMappings"`
	// Duration of leeway when validating all claims in the form of a time duration such as "5m" or "1h".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#clock_skew_leeway AclAuthMethod#clock_skew_leeway}
	ClockSkewLeeway *string `field:"optional" json:"clockSkewLeeway" yaml:"clockSkewLeeway"`
	// PEM encoded CA certs for use by the TLS client used to talk with the OIDC Discovery URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#discovery_ca_pem AclAuthMethod#discovery_ca_pem}
	DiscoveryCaPem *[]*string `field:"optional" json:"discoveryCaPem" yaml:"discoveryCaPem"`
	// Duration of leeway when validating expiration of a JWT in the form of a time duration such as "5m" or "1h".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#expiration_leeway AclAuthMethod#expiration_leeway}
	ExpirationLeeway *string `field:"optional" json:"expirationLeeway" yaml:"expirationLeeway"`
	// PEM encoded CA cert for use by the TLS client used to talk with the JWKS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#jwks_ca_cert AclAuthMethod#jwks_ca_cert}
	JwksCaCert *string `field:"optional" json:"jwksCaCert" yaml:"jwksCaCert"`
	// JSON Web Key Sets url for authenticating signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#jwks_url AclAuthMethod#jwks_url}
	JwksUrl *string `field:"optional" json:"jwksUrl" yaml:"jwksUrl"`
	// List of PEM-encoded public keys to use to authenticate signatures locally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#jwt_validation_pub_keys AclAuthMethod#jwt_validation_pub_keys}
	JwtValidationPubKeys *[]*string `field:"optional" json:"jwtValidationPubKeys" yaml:"jwtValidationPubKeys"`
	// Mappings of list claims (key) that will be copied to a metadata field (value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#list_claim_mappings AclAuthMethod#list_claim_mappings}
	ListClaimMappings *map[string]*string `field:"optional" json:"listClaimMappings" yaml:"listClaimMappings"`
	// Duration of leeway when validating not before values of a token in the form of a time duration such as "5m" or "1h".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#not_before_leeway AclAuthMethod#not_before_leeway}
	NotBeforeLeeway *string `field:"optional" json:"notBeforeLeeway" yaml:"notBeforeLeeway"`
	// The OAuth Client ID configured with the OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#oidc_client_id AclAuthMethod#oidc_client_id}
	OidcClientId *string `field:"optional" json:"oidcClientId" yaml:"oidcClientId"`
	// The OAuth Client Secret configured with the OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#oidc_client_secret AclAuthMethod#oidc_client_secret}
	OidcClientSecret *string `field:"optional" json:"oidcClientSecret" yaml:"oidcClientSecret"`
	// Nomad will not make a request to the identity provider to get OIDC UserInfo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#oidc_disable_userinfo AclAuthMethod#oidc_disable_userinfo}
	OidcDisableUserinfo interface{} `field:"optional" json:"oidcDisableUserinfo" yaml:"oidcDisableUserinfo"`
	// The OIDC Discovery URL, without any .well-known component (base path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#oidc_discovery_url AclAuthMethod#oidc_discovery_url}
	OidcDiscoveryUrl *string `field:"optional" json:"oidcDiscoveryUrl" yaml:"oidcDiscoveryUrl"`
	// List of OIDC scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#oidc_scopes AclAuthMethod#oidc_scopes}
	OidcScopes *[]*string `field:"optional" json:"oidcScopes" yaml:"oidcScopes"`
	// A list of supported signing algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#signing_algs AclAuthMethod#signing_algs}
	SigningAlgs *[]*string `field:"optional" json:"signingAlgs" yaml:"signingAlgs"`
}

