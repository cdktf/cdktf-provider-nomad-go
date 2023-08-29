// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-nomad.provider.NomadProvider",
		reflect.TypeOf((*NomadProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "address", GoGetter: "Address"},
			_jsii_.MemberProperty{JsiiProperty: "addressInput", GoGetter: "AddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "caFile", GoGetter: "CaFile"},
			_jsii_.MemberProperty{JsiiProperty: "caFileInput", GoGetter: "CaFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "caPem", GoGetter: "CaPem"},
			_jsii_.MemberProperty{JsiiProperty: "caPemInput", GoGetter: "CaPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certFile", GoGetter: "CertFile"},
			_jsii_.MemberProperty{JsiiProperty: "certFileInput", GoGetter: "CertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "certPem", GoGetter: "CertPem"},
			_jsii_.MemberProperty{JsiiProperty: "certPemInput", GoGetter: "CertPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "consulToken", GoGetter: "ConsulToken"},
			_jsii_.MemberProperty{JsiiProperty: "consulTokenInput", GoGetter: "ConsulTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpAuth", GoGetter: "HttpAuth"},
			_jsii_.MemberProperty{JsiiProperty: "httpAuthInput", GoGetter: "HttpAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreEnvVars", GoGetter: "IgnoreEnvVars"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreEnvVarsInput", GoGetter: "IgnoreEnvVarsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyFile", GoGetter: "KeyFile"},
			_jsii_.MemberProperty{JsiiProperty: "keyFileInput", GoGetter: "KeyFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyPem", GoGetter: "KeyPem"},
			_jsii_.MemberProperty{JsiiProperty: "keyPemInput", GoGetter: "KeyPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaFile", GoMethod: "ResetCaFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaPem", GoMethod: "ResetCaPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertFile", GoMethod: "ResetCertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertPem", GoMethod: "ResetCertPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsulToken", GoMethod: "ResetConsulToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpAuth", GoMethod: "ResetHttpAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreEnvVars", GoMethod: "ResetIgnoreEnvVars"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyFile", GoMethod: "ResetKeyFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyPem", GoMethod: "ResetKeyPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretId", GoMethod: "ResetSecretId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVaultToken", GoMethod: "ResetVaultToken"},
			_jsii_.MemberProperty{JsiiProperty: "secretId", GoGetter: "SecretId"},
			_jsii_.MemberProperty{JsiiProperty: "secretIdInput", GoGetter: "SecretIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vaultToken", GoGetter: "VaultToken"},
			_jsii_.MemberProperty{JsiiProperty: "vaultTokenInput", GoGetter: "VaultTokenInput"},
		},
		func() interface{} {
			j := jsiiProxy_NomadProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-nomad.provider.NomadProviderConfig",
		reflect.TypeOf((*NomadProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-nomad.provider.NomadProviderHeaders",
		reflect.TypeOf((*NomadProviderHeaders)(nil)).Elem(),
	)
}
