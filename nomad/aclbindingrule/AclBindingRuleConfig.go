// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclbindingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclBindingRuleConfig struct {
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
	// Name of the auth method for which this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_binding_rule#auth_method AclBindingRule#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
	// Adjusts how this binding rule is applied at login time. Valid values are "role" and "policy".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_binding_rule#bind_type AclBindingRule#bind_type}
	BindType *string `field:"required" json:"bindType" yaml:"bindType"`
	// Target of the binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_binding_rule#bind_name AclBindingRule#bind_name}
	BindName *string `field:"optional" json:"bindName" yaml:"bindName"`
	// Description for this ACL binding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_binding_rule#description AclBindingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_binding_rule#id AclBindingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A boolean expression that matches against verified identity attributes returned from the auth method during login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_binding_rule#selector AclBindingRule#selector}
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
}

