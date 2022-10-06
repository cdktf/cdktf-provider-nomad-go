package job


type JobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#create Job#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/job#update Job#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

