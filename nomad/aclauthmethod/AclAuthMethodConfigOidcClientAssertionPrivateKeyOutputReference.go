// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v10/jsii"

	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v10/aclauthmethod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *AclAuthMethodConfigOidcClientAssertionPrivateKey
	SetInternalValue(val *AclAuthMethodConfigOidcClientAssertionPrivateKey)
	KeyId() *string
	SetKeyId(val *string)
	KeyIdHeader() *string
	SetKeyIdHeader(val *string)
	KeyIdHeaderInput() *string
	KeyIdInput() *string
	PemCert() *string
	SetPemCert(val *string)
	PemCertFile() *string
	SetPemCertFile(val *string)
	PemCertFileInput() *string
	PemCertInput() *string
	PemKey() *string
	SetPemKey(val *string)
	PemKeyFile() *string
	SetPemKeyFile(val *string)
	PemKeyFileInput() *string
	PemKeyInput() *string
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
	ResetKeyId()
	ResetKeyIdHeader()
	ResetPemCert()
	ResetPemCertFile()
	ResetPemKey()
	ResetPemKeyFile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference
type jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) InternalValue() *AclAuthMethodConfigOidcClientAssertionPrivateKey {
	var returns *AclAuthMethodConfigOidcClientAssertionPrivateKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) KeyIdHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) KeyIdHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) PemKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference {
	_init_.Initialize()

	if err := validateNewAclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference_Override(a AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetInternalValue(val *AclAuthMethodConfigOidcClientAssertionPrivateKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetKeyId(val *string) {
	if err := j.validateSetKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetKeyIdHeader(val *string) {
	if err := j.validateSetKeyIdHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyIdHeader",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetPemCert(val *string) {
	if err := j.validateSetPemCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemCert",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetPemCertFile(val *string) {
	if err := j.validateSetPemCertFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemCertFile",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetPemKey(val *string) {
	if err := j.validateSetPemKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemKey",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetPemKeyFile(val *string) {
	if err := j.validateSetPemKeyFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemKeyFile",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ResetKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ResetKeyIdHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyIdHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ResetPemCert() {
	_jsii_.InvokeVoid(
		a,
		"resetPemCert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ResetPemCertFile() {
	_jsii_.InvokeVoid(
		a,
		"resetPemCertFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ResetPemKey() {
	_jsii_.InvokeVoid(
		a,
		"resetPemKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ResetPemKeyFile() {
	_jsii_.InvokeVoid(
		a,
		"resetPemKeyFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AclAuthMethodConfigOidcClientAssertionPrivateKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

