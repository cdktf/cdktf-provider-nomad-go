// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-nomad-go/nomad/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-nomad-go/nomad/v8/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs nomad}.
type NomadProvider interface {
	cdktf.TerraformProvider
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	CaFile() *string
	SetCaFile(val *string)
	CaFileInput() *string
	CaPem() *string
	SetCaPem(val *string)
	CaPemInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertFile() *string
	SetCertFile(val *string)
	CertFileInput() *string
	CertPem() *string
	SetCertPem(val *string)
	CertPemInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsulToken() *string
	SetConsulToken(val *string)
	ConsulTokenInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Headers() interface{}
	SetHeaders(val interface{})
	HeadersInput() interface{}
	HttpAuth() *string
	SetHttpAuth(val *string)
	HttpAuthInput() *string
	IgnoreEnvVars() *map[string]interface{}
	SetIgnoreEnvVars(val *map[string]interface{})
	IgnoreEnvVarsInput() *map[string]interface{}
	KeyFile() *string
	SetKeyFile(val *string)
	KeyFileInput() *string
	KeyPem() *string
	SetKeyPem(val *string)
	KeyPemInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecretId() *string
	SetSecretId(val *string)
	SecretIdInput() *string
	SkipVerify() interface{}
	SetSkipVerify(val interface{})
	SkipVerifyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	VaultToken() *string
	SetVaultToken(val *string)
	VaultTokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetCaFile()
	ResetCaPem()
	ResetCertFile()
	ResetCertPem()
	ResetConsulToken()
	ResetHeaders()
	ResetHttpAuth()
	ResetIgnoreEnvVars()
	ResetKeyFile()
	ResetKeyPem()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSecretId()
	ResetSkipVerify()
	ResetVaultToken()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NomadProvider
type jsiiProxy_NomadProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_NomadProvider) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CaPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CaPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CertPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) CertPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) ConsulToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) ConsulTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) Headers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) HttpAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) HttpAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) IgnoreEnvVars() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"ignoreEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) IgnoreEnvVarsInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"ignoreEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) KeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) KeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) KeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) KeyPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) SecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) SkipVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) SkipVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) VaultToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadProvider) VaultTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultTokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs nomad} Resource.
func NewNomadProvider(scope constructs.Construct, id *string, config *NomadProviderConfig) NomadProvider {
	_init_.Initialize()

	if err := validateNewNomadProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NomadProvider{}

	_jsii_.Create(
		"@cdktf/provider-nomad.provider.NomadProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs nomad} Resource.
func NewNomadProvider_Override(n NomadProvider, scope constructs.Construct, id *string, config *NomadProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nomad.provider.NomadProvider",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NomadProvider)SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetCaFile(val *string) {
	_jsii_.Set(
		j,
		"caFile",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetCaPem(val *string) {
	_jsii_.Set(
		j,
		"caPem",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetCertFile(val *string) {
	_jsii_.Set(
		j,
		"certFile",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetCertPem(val *string) {
	_jsii_.Set(
		j,
		"certPem",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetConsulToken(val *string) {
	_jsii_.Set(
		j,
		"consulToken",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetHeaders(val interface{}) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetHttpAuth(val *string) {
	_jsii_.Set(
		j,
		"httpAuth",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetIgnoreEnvVars(val *map[string]interface{}) {
	if err := j.validateSetIgnoreEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreEnvVars",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetKeyFile(val *string) {
	_jsii_.Set(
		j,
		"keyFile",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetKeyPem(val *string) {
	_jsii_.Set(
		j,
		"keyPem",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetSecretId(val *string) {
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetSkipVerify(val interface{}) {
	if err := j.validateSetSkipVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipVerify",
		val,
	)
}

func (j *jsiiProxy_NomadProvider)SetVaultToken(val *string) {
	_jsii_.Set(
		j,
		"vaultToken",
		val,
	)
}

// Generates CDKTF code for importing a NomadProvider resource upon running "cdktf plan <stack-name>".
func NomadProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNomadProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.provider.NomadProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NomadProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNomadProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.provider.NomadProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NomadProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNomadProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.provider.NomadProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NomadProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNomadProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nomad.provider.NomadProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NomadProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nomad.provider.NomadProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NomadProvider) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NomadProvider) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NomadProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		n,
		"resetAlias",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetCaFile() {
	_jsii_.InvokeVoid(
		n,
		"resetCaFile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetCaPem() {
	_jsii_.InvokeVoid(
		n,
		"resetCaPem",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetCertFile() {
	_jsii_.InvokeVoid(
		n,
		"resetCertFile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetCertPem() {
	_jsii_.InvokeVoid(
		n,
		"resetCertPem",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetConsulToken() {
	_jsii_.InvokeVoid(
		n,
		"resetConsulToken",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetHeaders() {
	_jsii_.InvokeVoid(
		n,
		"resetHeaders",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetHttpAuth() {
	_jsii_.InvokeVoid(
		n,
		"resetHttpAuth",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetIgnoreEnvVars() {
	_jsii_.InvokeVoid(
		n,
		"resetIgnoreEnvVars",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetKeyFile() {
	_jsii_.InvokeVoid(
		n,
		"resetKeyFile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetKeyPem() {
	_jsii_.InvokeVoid(
		n,
		"resetKeyPem",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		n,
		"resetRegion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetSecretId() {
	_jsii_.InvokeVoid(
		n,
		"resetSecretId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetSkipVerify() {
	_jsii_.InvokeVoid(
		n,
		"resetSkipVerify",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) ResetVaultToken() {
	_jsii_.InvokeVoid(
		n,
		"resetVaultToken",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

