package acltoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclTokenConfig struct {
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
	// The type of token to create, 'client' or 'management'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#type AclToken#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Provides a TTL for the token in the form of a time duration such as "5m" or "1h".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#expiration_ttl AclToken#expiration_ttl}
	ExpirationTtl *string `field:"optional" json:"expirationTtl" yaml:"expirationTtl"`
	// Whether the token should be replicated to all regions or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#global AclToken#global}
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#id AclToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Human-readable name for this token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#name AclToken#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ACL policies to associate with the token, if it's a 'client' type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#policies AclToken#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.19/docs/resources/acl_token#role AclToken#role}
	Role interface{} `field:"optional" json:"role" yaml:"role"`
}

