// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamichostvolume


type DynamicHostVolumeConstraint struct {
	// An attribute to check to constrain volume placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#attribute DynamicHostVolume#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The operator to use for comparison.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#operator DynamicHostVolume#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The requested value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#value DynamicHostVolume#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

