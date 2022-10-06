package schedulerconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchedulerConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/scheduler_config#id SchedulerConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, tasks may exceed their reserved memory limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/scheduler_config#memory_oversubscription_enabled SchedulerConfig#memory_oversubscription_enabled}
	MemoryOversubscriptionEnabled interface{} `field:"optional" json:"memoryOversubscriptionEnabled" yaml:"memoryOversubscriptionEnabled"`
	// Options to enable preemption for various schedulers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/scheduler_config#preemption_config SchedulerConfig#preemption_config}
	PreemptionConfig *map[string]interface{} `field:"optional" json:"preemptionConfig" yaml:"preemptionConfig"`
	// Specifies whether scheduler binpacks or spreads allocations on available nodes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/scheduler_config#scheduler_algorithm SchedulerConfig#scheduler_algorithm}
	SchedulerAlgorithm *string `field:"optional" json:"schedulerAlgorithm" yaml:"schedulerAlgorithm"`
}

