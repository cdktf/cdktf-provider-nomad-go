package datanomadvolumes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNomadVolumesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/volumes#id DataNomadVolumes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Volume namespace filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/volumes#namespace DataNomadVolumes#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Volume node filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/volumes#node_id DataNomadVolumes#node_id}
	NodeId *string `field:"optional" json:"nodeId" yaml:"nodeId"`
	// Plugin ID filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/volumes#plugin_id DataNomadVolumes#plugin_id}
	PluginId *string `field:"optional" json:"pluginId" yaml:"pluginId"`
	// Volume Type (currently only 'csi').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/volumes#type DataNomadVolumes#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

