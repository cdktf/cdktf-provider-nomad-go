package datanomaddatacenters

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNomadDatacentersConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/datacenters#id DataNomadDatacenters#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If enabled, this flag will ignore nodes that are down when listing datacenters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/datacenters#ignore_down_nodes DataNomadDatacenters#ignore_down_nodes}
	IgnoreDownNodes interface{} `field:"optional" json:"ignoreDownNodes" yaml:"ignoreDownNodes"`
	// Prefix value used for filtering results.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/datacenters#prefix DataNomadDatacenters#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

