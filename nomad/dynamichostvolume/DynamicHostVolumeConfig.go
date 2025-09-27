// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamichostvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamicHostVolumeConfig struct {
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
	// capability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#capability DynamicHostVolume#capability}
	Capability interface{} `field:"required" json:"capability" yaml:"capability"`
	// Volume name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#name DynamicHostVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Plugin ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#plugin_id DynamicHostVolume#plugin_id}
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
	// Requested maximum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#capacity_max DynamicHostVolume#capacity_max}
	CapacityMax *string `field:"optional" json:"capacityMax" yaml:"capacityMax"`
	// Requested minimum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#capacity_min DynamicHostVolume#capacity_min}
	CapacityMin *string `field:"optional" json:"capacityMin" yaml:"capacityMin"`
	// constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#constraint DynamicHostVolume#constraint}
	Constraint interface{} `field:"optional" json:"constraint" yaml:"constraint"`
	// Volume namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#namespace DynamicHostVolume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Node ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#node_id DynamicHostVolume#node_id}
	NodeId *string `field:"optional" json:"nodeId" yaml:"nodeId"`
	// Node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#node_pool DynamicHostVolume#node_pool}
	NodePool *string `field:"optional" json:"nodePool" yaml:"nodePool"`
	// Parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume#parameters DynamicHostVolume#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

