package quotaspecification


type QuotaSpecificationLimitsRegionLimit struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/quota_specification#cpu QuotaSpecification#cpu}.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/quota_specification#memory_mb QuotaSpecification#memory_mb}.
	MemoryMb *float64 `field:"optional" json:"memoryMb" yaml:"memoryMb"`
}

