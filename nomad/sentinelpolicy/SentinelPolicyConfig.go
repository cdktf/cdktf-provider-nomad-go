// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelPolicyConfig struct {
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
	// Specifies the enforcement level of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/sentinel_policy#enforcement_level SentinelPolicy#enforcement_level}
	EnforcementLevel *string `field:"required" json:"enforcementLevel" yaml:"enforcementLevel"`
	// Unique name for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/sentinel_policy#name SentinelPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Sentinel policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/sentinel_policy#policy SentinelPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Specifies the scope for this policy. Only 'submit-job' is currently supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/sentinel_policy#scope SentinelPolicy#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Description for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/sentinel_policy#description SentinelPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/sentinel_policy#id SentinelPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

