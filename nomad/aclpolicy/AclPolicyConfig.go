// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclPolicyConfig struct {
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
	// Unique name for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_policy#name AclPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// HCL or JSON representation of the rules to enforce on this policy.
	//
	// Use file() to specify a file as input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_policy#rules_hcl AclPolicy#rules_hcl}
	RulesHcl *string `field:"required" json:"rulesHcl" yaml:"rulesHcl"`
	// Description for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_policy#description AclPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_policy#id AclPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// job_acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/acl_policy#job_acl AclPolicy#job_acl}
	JobAcl *AclPolicyJobAcl `field:"optional" json:"jobAcl" yaml:"jobAcl"`
}

