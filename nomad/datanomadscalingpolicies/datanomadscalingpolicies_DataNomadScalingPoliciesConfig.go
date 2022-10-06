package datanomadscalingpolicies

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNomadScalingPoliciesConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/scaling_policies#id DataNomadScalingPolicies#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Job ID to use to filter scaling policies.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/scaling_policies#job_id DataNomadScalingPolicies#job_id}
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Scaling policy type used to filter scaling policies.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/d/scaling_policies#type DataNomadScalingPolicies#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

