package provider


type NomadProviderHeaders struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad#name NomadProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad#value NomadProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

