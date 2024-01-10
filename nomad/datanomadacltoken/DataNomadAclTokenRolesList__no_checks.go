// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadacltoken

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadAclTokenRolesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataNomadAclTokenRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadAclTokenRolesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadAclTokenRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadAclTokenRolesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadAclTokenRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadAclTokenRolesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

