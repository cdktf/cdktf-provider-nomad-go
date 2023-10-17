// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schedulerconfig

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchedulerConfig) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SchedulerConfig) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSchedulerConfig_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSchedulerConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSchedulerConfig_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSchedulerConfig_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetMemoryOversubscriptionEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetPreemptionConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SchedulerConfig) validateSetSchedulerAlgorithmParameters(val *string) error {
	return nil
}

func validateNewSchedulerConfigParameters(scope constructs.Construct, id *string, config *SchedulerConfigConfig) error {
	return nil
}

