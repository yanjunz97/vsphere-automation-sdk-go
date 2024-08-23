/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Passive.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/vcha"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``CheckSpec`` class contains placement information for validation. This class was added in vSphere API 6.7.1.
type PassiveCheckSpec struct {
    // Contains the active node's management vCenter server credentials. This property was added in vSphere API 6.7.1.
	VcSpec *vcha.CredentialsSpec
    // Contains the node's placement information for validation. This property was added in vSphere API 6.7.1.
	Placement vcha.PlacementSpec
}

// The ``CheckResult`` class contains the warnings and errors that will occur during the clone operation. This class was added in vSphere API 6.7.1.
type PassiveCheckResult struct {
    // A list of problems which may require attention, but which are not fatal. This property was added in vSphere API 6.7.1.
	Warnings []std.LocalizableMessage
    // A list of problems which are fatal to the operation and the operation will fail. This property was added in vSphere API 6.7.1.
	Errors []std.LocalizableMessage
}

// The ``RedeploySpec`` class contains the redeploy specification. This class was added in vSphere API 6.7.1.
type PassiveRedeploySpec struct {
    // Contains the active node's management vCenter server credentials. This property was added in vSphere API 6.7.1.
	VcSpec *vcha.CredentialsSpec
    // Contains the node's placement information. This property was added in vSphere API 6.7.1.
	Placement vcha.PlacementSpec
    // Contains the VCHA HA network configuration of the node. All cluster communication (state replication, heartbeat, cluster messages) happens over this network. This property was added in vSphere API 6.7.1.
	HaIp *vcha.IpSpec
    // Failover IP address that this node must assume after the failover to serve client requests. This property was added in vSphere API 6.7.1.
	FailoverIp *vcha.IpSpec
}



func passiveCheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(PassiveCheckSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func passiveCheckOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PassiveCheckResultBindingType)
}

func passiveCheckRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(PassiveCheckSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"UnverifiedPeer": 400,"NotFound": 404,"InvalidElementConfiguration": 400,"NotAllowedInCurrentState": 400,"Unauthorized": 403,"Error": 500})
}

func passiveRedeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(PassiveRedeploySpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func passiveRedeployOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func passiveRedeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(PassiveRedeploySpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func PassiveCheckSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
	fieldNameMap["vc_spec"] = "VcSpec"
	fields["placement"] = bindings.NewReferenceType(vcha.PlacementSpecBindingType)
	fieldNameMap["placement"] = "Placement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.passive.check_spec", fields, reflect.TypeOf(PassiveCheckSpec{}), fieldNameMap, validators)
}

func PassiveCheckResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["warnings"] = "Warnings"
	fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["errors"] = "Errors"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.passive.check_result", fields, reflect.TypeOf(PassiveCheckResult{}), fieldNameMap, validators)
}

func PassiveRedeploySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
	fieldNameMap["vc_spec"] = "VcSpec"
	fields["placement"] = bindings.NewReferenceType(vcha.PlacementSpecBindingType)
	fieldNameMap["placement"] = "Placement"
	fields["ha_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.IpSpecBindingType))
	fieldNameMap["ha_ip"] = "HaIp"
	fields["failover_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.IpSpecBindingType))
	fieldNameMap["failover_ip"] = "FailoverIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.passive.redeploy_spec", fields, reflect.TypeOf(PassiveRedeploySpec{}), fieldNameMap, validators)
}

