// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadjwks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadJwksKeysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataNomadJwksKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadJwksKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadJwksKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadJwksKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadJwksKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadJwksKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

