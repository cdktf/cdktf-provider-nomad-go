package aclauthmethod

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethod",
		reflect.TypeOf((*AclAuthMethod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "configInput", GoGetter: "ConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "default", GoGetter: "Default"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInput", GoGetter: "DefaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxTokenTtl", GoGetter: "MaxTokenTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTokenTtlInput", GoGetter: "MaxTokenTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConfig", GoMethod: "PutConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefault", GoMethod: "ResetDefault"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tokenLocality", GoGetter: "TokenLocality"},
			_jsii_.MemberProperty{JsiiProperty: "tokenLocalityInput", GoGetter: "TokenLocalityInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AclAuthMethod{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfig",
		reflect.TypeOf((*AclAuthMethodConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigA",
		reflect.TypeOf((*AclAuthMethodConfigA)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-nomad.aclAuthMethod.AclAuthMethodConfigAOutputReference",
		reflect.TypeOf((*AclAuthMethodConfigAOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedRedirectUris", GoGetter: "AllowedRedirectUris"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRedirectUrisInput", GoGetter: "AllowedRedirectUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "boundAudiences", GoGetter: "BoundAudiences"},
			_jsii_.MemberProperty{JsiiProperty: "boundAudiencesInput", GoGetter: "BoundAudiencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "claimMappings", GoGetter: "ClaimMappings"},
			_jsii_.MemberProperty{JsiiProperty: "claimMappingsInput", GoGetter: "ClaimMappingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "discoveryCaPem", GoGetter: "DiscoveryCaPem"},
			_jsii_.MemberProperty{JsiiProperty: "discoveryCaPemInput", GoGetter: "DiscoveryCaPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "listClaimMappings", GoGetter: "ListClaimMappings"},
			_jsii_.MemberProperty{JsiiProperty: "listClaimMappingsInput", GoGetter: "ListClaimMappingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcClientId", GoGetter: "OidcClientId"},
			_jsii_.MemberProperty{JsiiProperty: "oidcClientIdInput", GoGetter: "OidcClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcClientSecret", GoGetter: "OidcClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "oidcClientSecretInput", GoGetter: "OidcClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcDiscoveryUrl", GoGetter: "OidcDiscoveryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "oidcDiscoveryUrlInput", GoGetter: "OidcDiscoveryUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcScopes", GoGetter: "OidcScopes"},
			_jsii_.MemberProperty{JsiiProperty: "oidcScopesInput", GoGetter: "OidcScopesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBoundAudiences", GoMethod: "ResetBoundAudiences"},
			_jsii_.MemberMethod{JsiiMethod: "resetClaimMappings", GoMethod: "ResetClaimMappings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiscoveryCaPem", GoMethod: "ResetDiscoveryCaPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetListClaimMappings", GoMethod: "ResetListClaimMappings"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcScopes", GoMethod: "ResetOidcScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetSigningAlgs", GoMethod: "ResetSigningAlgs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgs", GoGetter: "SigningAlgs"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgsInput", GoGetter: "SigningAlgsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AclAuthMethodConfigAOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
