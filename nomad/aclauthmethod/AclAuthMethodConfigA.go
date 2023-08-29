// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod


type AclAuthMethodConfigA struct {
	// A list of allowed values that can be used for the redirect URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#allowed_redirect_uris AclAuthMethod#allowed_redirect_uris}
	AllowedRedirectUris *[]*string `field:"required" json:"allowedRedirectUris" yaml:"allowedRedirectUris"`
	// The OAuth Client ID configured with the OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#oidc_client_id AclAuthMethod#oidc_client_id}
	OidcClientId *string `field:"required" json:"oidcClientId" yaml:"oidcClientId"`
	// The OAuth Client Secret configured with the OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#oidc_client_secret AclAuthMethod#oidc_client_secret}
	OidcClientSecret *string `field:"required" json:"oidcClientSecret" yaml:"oidcClientSecret"`
	// The OIDC Discovery URL, without any .well-known component (base path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#oidc_discovery_url AclAuthMethod#oidc_discovery_url}
	OidcDiscoveryUrl *string `field:"required" json:"oidcDiscoveryUrl" yaml:"oidcDiscoveryUrl"`
	// List of auth claims that are valid for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#bound_audiences AclAuthMethod#bound_audiences}
	BoundAudiences *[]*string `field:"optional" json:"boundAudiences" yaml:"boundAudiences"`
	// Mappings of claims (key) that will be copied to a metadata field (value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#claim_mappings AclAuthMethod#claim_mappings}
	ClaimMappings *map[string]*string `field:"optional" json:"claimMappings" yaml:"claimMappings"`
	// PEM encoded CA certs for use by the TLS client used to talk with the OIDC Discovery URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#discovery_ca_pem AclAuthMethod#discovery_ca_pem}
	DiscoveryCaPem *[]*string `field:"optional" json:"discoveryCaPem" yaml:"discoveryCaPem"`
	// Mappings of list claims (key) that will be copied to a metadata field (value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#list_claim_mappings AclAuthMethod#list_claim_mappings}
	ListClaimMappings *map[string]*string `field:"optional" json:"listClaimMappings" yaml:"listClaimMappings"`
	// List of OIDC scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#oidc_scopes AclAuthMethod#oidc_scopes}
	OidcScopes *[]*string `field:"optional" json:"oidcScopes" yaml:"oidcScopes"`
	// A list of supported signing algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/acl_auth_method#signing_algs AclAuthMethod#signing_algs}
	SigningAlgs *[]*string `field:"optional" json:"signingAlgs" yaml:"signingAlgs"`
}

