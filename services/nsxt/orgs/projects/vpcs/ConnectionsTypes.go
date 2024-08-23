// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Connections.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package vpcs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func connectionsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["vpc_connection_id"] = "VpcConnectionId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectionsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
}

func connectionsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_connection_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["vpc_connection_id"] = "VpcConnectionId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_connection_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcConnectionId"] = vapiBindings_.NewStringType()
	pathParams["vpc_connection_id"] = "vpcConnectionId"
	pathParams["vpc_id"] = "vpcId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/connections/{vpcConnectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func connectionsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["vpc_connection_id"] = "VpcConnectionId"
	fieldNameMap["vpc_connection"] = "VpcConnection"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectionsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func connectionsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_connection_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["vpc_connection_id"] = "VpcConnectionId"
	fieldNameMap["vpc_connection"] = "VpcConnection"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_connection"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
	paramsTypeMap["vpc_connection_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcConnectionId"] = vapiBindings_.NewStringType()
	pathParams["vpc_connection_id"] = "vpcConnectionId"
	pathParams["vpc_id"] = "vpcId"
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
		"vpc_connection",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/connections/{vpcConnectionId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func connectionsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["vpc_connection_id"] = "VpcConnectionId"
	fieldNameMap["vpc_connection"] = "VpcConnection"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectionsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
}

func connectionsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_connection_id"] = vapiBindings_.NewStringType()
	fields["vpc_connection"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["vpc_connection_id"] = "VpcConnectionId"
	fieldNameMap["vpc_connection"] = "VpcConnection"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_connection"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcConnectionBindingType)
	paramsTypeMap["vpc_connection_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcConnectionId"] = vapiBindings_.NewStringType()
	pathParams["vpc_connection_id"] = "vpcConnectionId"
	pathParams["vpc_id"] = "vpcId"
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
		"vpc_connection",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/connections/{vpcConnectionId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
