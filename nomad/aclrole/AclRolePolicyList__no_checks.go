// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package aclrole

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclRolePolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclRolePolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclRolePolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclRolePolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclRolePolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclRolePolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclRolePolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

