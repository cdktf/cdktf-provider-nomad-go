package volume


type VolumeCapability struct {
	// Defines whether a volume should be available concurrently.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#access_mode Volume#access_mode}
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// The storage API that will be used by the volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/volume#attachment_mode Volume#attachment_mode}
	AttachmentMode *string `field:"required" json:"attachmentMode" yaml:"attachmentMode"`
}

