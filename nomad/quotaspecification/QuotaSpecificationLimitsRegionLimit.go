// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quotaspecification


type QuotaSpecificationLimitsRegionLimit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/quota_specification#cpu QuotaSpecification#cpu}.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/quota_specification#memory_mb QuotaSpecification#memory_mb}.
	MemoryMb *float64 `field:"optional" json:"memoryMb" yaml:"memoryMb"`
}

