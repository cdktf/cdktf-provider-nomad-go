// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespace


type NamespaceCapabilities struct {
	// Disabled task drivers for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/namespace#disabled_task_drivers Namespace#disabled_task_drivers}
	DisabledTaskDrivers *[]*string `field:"optional" json:"disabledTaskDrivers" yaml:"disabledTaskDrivers"`
	// Enabled task drivers for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/namespace#enabled_task_drivers Namespace#enabled_task_drivers}
	EnabledTaskDrivers *[]*string `field:"optional" json:"enabledTaskDrivers" yaml:"enabledTaskDrivers"`
}

