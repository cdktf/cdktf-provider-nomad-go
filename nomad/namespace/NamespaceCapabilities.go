// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespace


type NamespaceCapabilities struct {
	// Disabled network modes for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/namespace#disabled_network_modes Namespace#disabled_network_modes}
	DisabledNetworkModes *[]*string `field:"optional" json:"disabledNetworkModes" yaml:"disabledNetworkModes"`
	// Disabled task drivers for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/namespace#disabled_task_drivers Namespace#disabled_task_drivers}
	DisabledTaskDrivers *[]*string `field:"optional" json:"disabledTaskDrivers" yaml:"disabledTaskDrivers"`
	// Enabled network modes for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/namespace#enabled_network_modes Namespace#enabled_network_modes}
	EnabledNetworkModes *[]*string `field:"optional" json:"enabledNetworkModes" yaml:"enabledNetworkModes"`
	// Enabled task drivers for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/namespace#enabled_task_drivers Namespace#enabled_task_drivers}
	EnabledTaskDrivers *[]*string `field:"optional" json:"enabledTaskDrivers" yaml:"enabledTaskDrivers"`
}

