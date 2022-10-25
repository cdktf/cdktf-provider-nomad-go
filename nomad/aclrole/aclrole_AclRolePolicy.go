package aclrole


type AclRolePolicy struct {
	// The name of the ACL policy to link.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/acl_role#name AclRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

