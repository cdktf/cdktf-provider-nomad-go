// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package quotaspecification

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QuotaSpecificationLimitsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QuotaSpecificationLimitsList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QuotaSpecificationLimitsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QuotaSpecificationLimitsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QuotaSpecificationLimitsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QuotaSpecificationLimitsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QuotaSpecificationLimitsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQuotaSpecificationLimitsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

