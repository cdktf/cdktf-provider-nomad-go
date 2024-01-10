// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datanomadjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNomadJobTaskGroupsTaskList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataNomadJobTaskGroupsTaskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNomadJobTaskGroupsTaskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobTaskGroupsTaskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobTaskGroupsTaskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNomadJobTaskGroupsTaskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNomadJobTaskGroupsTaskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

