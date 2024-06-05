// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quotaspecification


type QuotaSpecificationLimits struct {
	// Region in which this limit has affect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/quota_specification#region QuotaSpecification#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// region_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/quota_specification#region_limit QuotaSpecification#region_limit}
	RegionLimit *QuotaSpecificationLimitsRegionLimit `field:"required" json:"regionLimit" yaml:"regionLimit"`
}

