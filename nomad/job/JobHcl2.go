// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobHcl2 struct {
	// If true, HCL2 file system functions will be enabled when parsing the `jobspec`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/job#allow_fs Job#allow_fs}
	AllowFs interface{} `field:"optional" json:"allowFs" yaml:"allowFs"`
	// Additional variables to use when templating the job with HCL2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/job#vars Job#vars}
	Vars *map[string]*string `field:"optional" json:"vars" yaml:"vars"`
}

