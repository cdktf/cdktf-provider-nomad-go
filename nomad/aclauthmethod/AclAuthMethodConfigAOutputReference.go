// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v6/jsii"

	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v6/aclauthmethod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclAuthMethodConfigAOutputReference interface {
	cdktf.ComplexObject
	AllowedRedirectUris() *[]*string
	SetAllowedRedirectUris(val *[]*string)
	AllowedRedirectUrisInput() *[]*string
	BoundAudiences() *[]*string
	SetBoundAudiences(val *[]*string)
	BoundAudiencesInput() *[]*string
	ClaimMappings() *map[string]*string
	SetClaimMappings(val *map[string]*string)
	ClaimMappingsInput() *map[string]*string
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
	DiscoveryCaPem() *[]*string
	SetDiscoveryCaPem(val *[]*string)
	DiscoveryCaPemInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *AclAuthMethodConfigA
	SetInternalValue(val *AclAuthMethodConfigA)
	ListClaimMappings() *map[string]*string
	SetListClaimMappings(val *map[string]*string)
	ListClaimMappingsInput() *map[string]*string
	OidcClientId() *string
	SetOidcClientId(val *string)
	OidcClientIdInput() *string
	OidcClientSecret() *string
	SetOidcClientSecret(val *string)
	OidcClientSecretInput() *string
	OidcDiscoveryUrl() *string
	SetOidcDiscoveryUrl(val *string)
	OidcDiscoveryUrlInput() *string
	OidcScopes() *[]*string
	SetOidcScopes(val *[]*string)
	OidcScopesInput() *[]*string
	SigningAlgs() *[]*string
	SetSigningAlgs(val *[]*string)
	SigningAlgsInput() *[]*string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBoundAudiences()
	ResetClaimMappings()
	ResetDiscoveryCaPem()
	ResetListClaimMappings()
	ResetOidcScopes()
	ResetSigningAlgs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AclAuthMethodConfigAOutputReference
type jsiiProxy_AclAuthMethodConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) AllowedRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) AllowedRedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRedirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) BoundAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) BoundAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) ClaimMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) ClaimMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) DiscoveryCaPem() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"discoveryCaPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) DiscoveryCaPemInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"discoveryCaPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) InternalValue() *AclAuthMethodConfigA {
	var returns *AclAuthMethodConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) ListClaimMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listClaimMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) ListClaimMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listClaimMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcDiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcDiscoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcDiscoveryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcDiscoveryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oidcScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) OidcScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oidcScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) SigningAlgs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingAlgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) SigningAlgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingAlgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAclAuthMethodConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AclAuthMethodConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewAclAuthMethodConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AclAuthMethodConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAclAuthMethodConfigAOutputReference_Override(a AclAuthMethodConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetAllowedRedirectUris(val *[]*string) {
	if err := j.validateSetAllowedRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRedirectUris",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetBoundAudiences(val *[]*string) {
	if err := j.validateSetBoundAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundAudiences",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetClaimMappings(val *map[string]*string) {
	if err := j.validateSetClaimMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimMappings",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetDiscoveryCaPem(val *[]*string) {
	if err := j.validateSetDiscoveryCaPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryCaPem",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetInternalValue(val *AclAuthMethodConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetListClaimMappings(val *map[string]*string) {
	if err := j.validateSetListClaimMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listClaimMappings",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetOidcClientId(val *string) {
	if err := j.validateSetOidcClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcClientId",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetOidcClientSecret(val *string) {
	if err := j.validateSetOidcClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcClientSecret",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetOidcDiscoveryUrl(val *string) {
	if err := j.validateSetOidcDiscoveryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcDiscoveryUrl",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetOidcScopes(val *[]*string) {
	if err := j.validateSetOidcScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcScopes",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetSigningAlgs(val *[]*string) {
	if err := j.validateSetSigningAlgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgs",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AclAuthMethodConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ResetBoundAudiences() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundAudiences",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ResetClaimMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetClaimMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ResetDiscoveryCaPem() {
	_jsii_.InvokeVoid(
		a,
		"resetDiscoveryCaPem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ResetListClaimMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetListClaimMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ResetOidcScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ResetSigningAlgs() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningAlgs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclAuthMethodConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

