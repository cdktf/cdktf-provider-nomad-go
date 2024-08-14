// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclrole


type AclRolePolicy struct {
	// The name of the ACL policy to link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/acl_role#name AclRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

