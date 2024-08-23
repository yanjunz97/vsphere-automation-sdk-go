// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PortProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package ports

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func portProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["port_profile_binding_map_id"] = "portProfileBindingMapId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/port-profile-binding-maps/{portProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
}

func portProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["port_profile_binding_map_id"] = "portProfileBindingMapId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/port-profile-binding-maps/{portProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_port_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	fieldNameMap["vpc_subnet_port_profile_binding_map"] = "VpcSubnetPortProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_port_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	fieldNameMap["vpc_subnet_port_profile_binding_map"] = "VpcSubnetPortProfileBindingMap"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_subnet_port_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["port_profile_binding_map_id"] = "portProfileBindingMapId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"vpc_subnet_port_profile_binding_map",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/port-profile-binding-maps/{portProfileBindingMapId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_port_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	fieldNameMap["vpc_subnet_port_profile_binding_map"] = "VpcSubnetPortProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
}

func portProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_port_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_profile_binding_map_id"] = "PortProfileBindingMapId"
	fieldNameMap["vpc_subnet_port_profile_binding_map"] = "VpcSubnetPortProfileBindingMap"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_subnet_port_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetPortProfileBindingMapBindingType)
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["port_profile_binding_map_id"] = "portProfileBindingMapId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"vpc_subnet_port_profile_binding_map",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/port-profile-binding-maps/{portProfileBindingMapId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
