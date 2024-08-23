// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: BridgeProfiles.
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

func bridgeProfilesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vcenter_id"] = vapiBindings_.NewStringType()
	fields["dvpg_id"] = vapiBindings_.NewStringType()
	fields["bridge_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["project_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["segment_name_for_auto_creation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["segment_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vlan_transport_zone_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vcenter_id"] = "VcenterId"
	fieldNameMap["dvpg_id"] = "DvpgId"
	fieldNameMap["bridge_mode"] = "BridgeMode"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_name_for_auto_creation"] = "SegmentNameForAutoCreation"
	fieldNameMap["segment_path"] = "SegmentPath"
	fieldNameMap["vlan_transport_zone_path"] = "VlanTransportZonePath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgeProfilesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.RecommendationListResultBindingType)
}

func bridgeProfilesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["bridge_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["project_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["segment_name_for_auto_creation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["segment_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vlan_transport_zone_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["vcenter_id"] = "VcenterId"
	fieldNameMap["dvpg_id"] = "DvpgId"
	fieldNameMap["bridge_mode"] = "BridgeMode"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_name_for_auto_creation"] = "SegmentNameForAutoCreation"
	fieldNameMap["segment_path"] = "SegmentPath"
	fieldNameMap["vlan_transport_zone_path"] = "VlanTransportZonePath"
	paramsTypeMap["vcenter_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["dvpg_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_name_for_auto_creation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["project_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["bridge_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["segment_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vlan_transport_zone_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vcenterId"] = vapiBindings_.NewStringType()
	pathParams["vcenter_id"] = "vcenterId"
	queryParams["dvpg_id"] = "dvpg_id"
	queryParams["segment_name_for_auto_creation"] = "segment_name_for_auto_creation"
	queryParams["project_id"] = "project_id"
	queryParams["bridge_mode"] = "bridge_mode"
	queryParams["segment_path"] = "segment_path"
	queryParams["vlan_transport_zone_path"] = "vlan_transport_zone_path"
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
		"/api/v1/migration/overlay-adoption/vcenters/{vcenterId}/recommendation/bridge-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
