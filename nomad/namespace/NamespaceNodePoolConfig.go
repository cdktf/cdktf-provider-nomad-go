// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespace


type NamespaceNodePoolConfig struct {
	// The list of node pools allowed to be used in this namespace. Cannot be used with denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/namespace#allowed Namespace#allowed}
	Allowed *[]*string `field:"optional" json:"allowed" yaml:"allowed"`
	// The node pool to use when none are specified in the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/namespace#default Namespace#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// The list of node pools not allowed to be used in this namespace. Cannot be used with allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/namespace#denied Namespace#denied}
	Denied *[]*string `field:"optional" json:"denied" yaml:"denied"`
}

