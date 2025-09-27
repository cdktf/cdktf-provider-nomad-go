// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamichostvolumeregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamicHostVolumeRegistrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#capability DynamicHostVolumeRegistration#capability}
	Capability interface{} `field:"required" json:"capability" yaml:"capability"`
	// Host path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#host_path DynamicHostVolumeRegistration#host_path}
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// Volume name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#name DynamicHostVolumeRegistration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Node ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#node_id DynamicHostVolumeRegistration#node_id}
	NodeId *string `field:"required" json:"nodeId" yaml:"nodeId"`
	// Provisioned capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#capacity DynamicHostVolumeRegistration#capacity}
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
	// Volume namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#namespace DynamicHostVolumeRegistration#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/dynamic_host_volume_registration#parameters DynamicHostVolumeRegistration#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

