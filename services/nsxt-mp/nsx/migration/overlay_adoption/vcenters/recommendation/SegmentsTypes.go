// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Segments.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package recommendation

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func segmentsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vcenter_id"] = vapiBindings_.NewStringType()
	fields["dvpg_id"] = vapiBindings_.NewStringType()
	fields["ipv4_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["ipv6_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["project_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vcenter_id"] = "VcenterId"
	fieldNameMap["dvpg_id"] = "DvpgId"
	fieldNameMap["ipv4_subnet"] = "Ipv4Subnet"
	fieldNameMap["ipv6_subnet"] = "Ipv6Subnet"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.RecommendationListResultBindingType)
}

func segmentsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["vcenter_id"] = vapiBindings_.NewStringType()
	fields["dvpg_id"] = vapiBindings_.NewStringType()
	fields["ipv4_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["ipv6_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["project_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vcenter_id"] = "VcenterId"
	fieldNameMap["dvpg_id"] = "DvpgId"
	fieldNameMap["ipv4_subnet"] = "Ipv4Subnet"
	fieldNameMap["ipv6_subnet"] = "Ipv6Subnet"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["ipv6_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vcenter_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["dvpg_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["ipv4_subnet"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["project_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vcenterId"] = vapiBindings_.NewStringType()
	pathParams["vcenter_id"] = "vcenterId"
	queryParams["ipv6_subnet"] = "ipv6_subnet"
	queryParams["dvpg_id"] = "dvpg_id"
	queryParams["ipv4_subnet"] = "ipv4_subnet"
	queryParams["project_id"] = "project_id"
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
		"/api/v1/migration/overlay-adoption/vcenters/{vcenterId}/recommendation/segments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
