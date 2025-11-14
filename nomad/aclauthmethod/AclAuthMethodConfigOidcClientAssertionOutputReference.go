// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v10/jsii"

	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v10/aclauthmethod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclAuthMethodConfigOidcClientAssertionOutputReference interface {
	cdktf.ComplexObject
	Audience() *[]*string
	SetAudience(val *[]*string)
	AudienceInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExtraHeaders() *map[string]*string
	SetExtraHeaders(val *map[string]*string)
	ExtraHeadersInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *AclAuthMethodConfigOidcClientAssertion
	SetInternalValue(val *AclAuthMethodConfigOidcClientAssertion)
	KeyAlgorithm() *string
	SetKeyAlgorithm(val *string)
	KeyAlgorithmInput() *string
	KeySource() *string
	SetKeySource(val *string)
	KeySourceInput() *string
	PrivateKey() AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference
	PrivateKeyInput() *AclAuthMethodConfigOidcClientAssertionPrivateKey
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutPrivateKey(value *AclAuthMethodConfigOidcClientAssertionPrivateKey)
	ResetAudience()
	ResetExtraHeaders()
	ResetKeyAlgorithm()
	ResetPrivateKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AclAuthMethodConfigOidcClientAssertionOutputReference
type jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) Audience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) AudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ExtraHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ExtraHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) InternalValue() *AclAuthMethodConfigOidcClientAssertion {
	var returns *AclAuthMethodConfigOidcClientAssertion
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) KeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) KeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) KeySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) KeySourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) PrivateKey() AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference {
	var returns AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) PrivateKeyInput() *AclAuthMethodConfigOidcClientAssertionPrivateKey {
	var returns *AclAuthMethodConfigOidcClientAssertionPrivateKey
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAclAuthMethodConfigOidcClientAssertionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AclAuthMethodConfigOidcClientAssertionOutputReference {
	_init_.Initialize()

	if err := validateNewAclAuthMethodConfigOidcClientAssertionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigOidcClientAssertionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAclAuthMethodConfigOidcClientAssertionOutputReference_Override(a AclAuthMethodConfigOidcClientAssertionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigOidcClientAssertionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetAudience(val *[]*string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetExtraHeaders(val *map[string]*string) {
	if err := j.validateSetExtraHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraHeaders",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetInternalValue(val *AclAuthMethodConfigOidcClientAssertion) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetKeyAlgorithm(val *string) {
	if err := j.validateSetKeyAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetKeySource(val *string) {
	if err := j.validateSetKeySourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySource",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) PutPrivateKey(value *AclAuthMethodConfigOidcClientAssertionPrivateKey) {
	if err := a.validatePutPrivateKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrivateKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ResetAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ResetExtraHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetExtraHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ResetKeyAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

