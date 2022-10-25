//go:build no_runtime_type_checking

package acltoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclTokenRoleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclTokenRoleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclTokenRoleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclTokenRoleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclTokenRoleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclTokenRoleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclTokenRoleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

