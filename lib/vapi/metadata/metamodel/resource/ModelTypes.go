// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Model.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package resource

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

func modelListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	fieldNameMap["resource_id"] = "ResourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ModelListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewSetType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.structure"}, ""), reflect.TypeOf(map[string]bool{}))
}

func modelListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	fieldNameMap["resource_id"] = "ResourceId"
	paramsTypeMap["resource_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	paramsTypeMap["resourceId"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	pathParams["resource_id"] = "resourceId"
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
		"/vapi/metadata/metamodel/resource/{resourceId}/model",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}
