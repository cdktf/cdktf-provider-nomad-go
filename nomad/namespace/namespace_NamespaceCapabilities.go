package namespace


type NamespaceCapabilities struct {
	// Disabled task drivers for the namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/namespace#disabled_task_drivers Namespace#disabled_task_drivers}
	DisabledTaskDrivers *[]*string `field:"optional" json:"disabledTaskDrivers" yaml:"disabledTaskDrivers"`
	// Enabled task drivers for the namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/namespace#enabled_task_drivers Namespace#enabled_task_drivers}
	EnabledTaskDrivers *[]*string `field:"optional" json:"enabledTaskDrivers" yaml:"enabledTaskDrivers"`
}

