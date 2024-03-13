// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclpolicy


type AclPolicyJobAcl struct {
	// Job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_policy#job_id AclPolicy#job_id}
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_policy#group AclPolicy#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_policy#namespace AclPolicy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/acl_policy#task AclPolicy#task}
	Task *string `field:"optional" json:"task" yaml:"task"`
}

