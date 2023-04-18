//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobTaskGroupsTaskVolumeMountsList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobTaskGroupsTaskVolumeMountsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobTaskGroupsTaskVolumeMountsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskGroupsTaskVolumeMountsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobTaskGroupsTaskVolumeMountsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobTaskGroupsTaskVolumeMountsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

