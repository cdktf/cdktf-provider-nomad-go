// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclAuthMethodConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#config AclAuthMethod#config}
	Config *AclAuthMethodConfigA `field:"required" json:"config" yaml:"config"`
	// Defines the maximum life of a token created by this method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#max_token_ttl AclAuthMethod#max_token_ttl}
	MaxTokenTtl *string `field:"required" json:"maxTokenTtl" yaml:"maxTokenTtl"`
	// The identifier of the ACL Auth Method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#name AclAuthMethod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Defines whether the ACL Auth Method creates a local or global token when performing SSO login.
	//
	// This field must be set to either "local" or "global".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#token_locality AclAuthMethod#token_locality}
	TokenLocality *string `field:"required" json:"tokenLocality" yaml:"tokenLocality"`
	// ACL Auth Method SSO workflow type. Currently, the only supported type is "OIDC.".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#type AclAuthMethod#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines whether this ACL Auth Method is to be set as default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#default AclAuthMethod#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#id AclAuthMethod#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Defines the token format for the authenticated users. This can be lightly templated using HIL '${foo}' syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_auth_method#token_name_format AclAuthMethod#token_name_format}
	TokenNameFormat *string `field:"optional" json:"tokenNameFormat" yaml:"tokenNameFormat"`
}

