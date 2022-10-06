package job

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobConfig struct {
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
	// Job specification. If you want to point to a file use the file() function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#jobspec Job#jobspec}
	Jobspec *string `field:"required" json:"jobspec" yaml:"jobspec"`
	// The Consul token used to submit this job.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#consul_token Job#consul_token}
	ConsulToken *string `field:"optional" json:"consulToken" yaml:"consulToken"`
	// If true, the job will be deregistered on destroy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#deregister_on_destroy Job#deregister_on_destroy}
	DeregisterOnDestroy interface{} `field:"optional" json:"deregisterOnDestroy" yaml:"deregisterOnDestroy"`
	// If true, the job will be deregistered when the job ID changes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#deregister_on_id_change Job#deregister_on_id_change}
	DeregisterOnIdChange interface{} `field:"optional" json:"deregisterOnIdChange" yaml:"deregisterOnIdChange"`
	// If true, the provider will return immediately after creating or updating, instead of monitoring.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#detach Job#detach}
	Detach interface{} `field:"optional" json:"detach" yaml:"detach"`
	// hcl2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#hcl2 Job#hcl2}
	Hcl2 *JobHcl2 `field:"optional" json:"hcl2" yaml:"hcl2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#id Job#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, the `jobspec` will be parsed as json instead of HCL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#json Job#json}
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Override any soft-mandatory Sentinel policies that fail.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#policy_override Job#policy_override}
	PolicyOverride interface{} `field:"optional" json:"policyOverride" yaml:"policyOverride"`
	// Whether to purge the job when the resource is destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#purge_on_destroy Job#purge_on_destroy}
	PurgeOnDestroy interface{} `field:"optional" json:"purgeOnDestroy" yaml:"purgeOnDestroy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#timeouts Job#timeouts}
	Timeouts *JobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The Vault token used to submit this job.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#vault_token Job#vault_token}
	VaultToken *string `field:"optional" json:"vaultToken" yaml:"vaultToken"`
}

