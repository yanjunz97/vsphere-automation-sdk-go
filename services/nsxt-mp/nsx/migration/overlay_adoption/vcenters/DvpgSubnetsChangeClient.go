// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DvpgSubnetsChange
// Used by client-side stubs.

package vcenters

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DvpgSubnetsChangeClient interface {

	// The returned subnets change list is the cumulative result of all partial patchings done by PatchDvpgIpSubnetsChangeList and all whole updates done by UpdateDvpgIpSubnetsChangeList so far for the DVPG.
	//
	// @param vcenterIdParam (required)
	// @param dvpgIdParam (required)
	// @return com.vmware.nsx.model.DvpgIpSubnetsChangeList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(vcenterIdParam string, dvpgIdParam string) (nsxModel.DvpgIpSubnetsChangeList, error)

	// The request can contain some (not all) additional subnets changes; it will be rejected if any non-empty \"old_subnets\" is repeated or any \"old_subnets\" matches any \"new_subnets\" in the request. Each IpSubnetsChange in the requst will be added into the cumulative list of IP subnets changes for the DVPG in below five cases. 1. It is added into the list if its \"old_subnets\" is unset or empty; 2. else its \"new_subnets\" replaces the \"new_subnets\" in the existing IpSubnetsChange whose \"old_subnets\" matches its \"old_subnets\"; 3. else it is added into the list, and the existing IpSubnetsChange whose \"new_subnets\" and \"old_subnets\" match its \"old_subnets\" and \"new_subnets\" respectively is deleted; 4. else its \"new_subnets\" replaces the \"new_subnets\" in each existing IpSubnetsChange whose \"new_subnets\" matches its \"old_subnets\", and it is added into the list; 5. else it is added into the list.
	//
	// @param vcenterIdParam (required)
	// @param dvpgIpSubnetsChangeListParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(vcenterIdParam string, dvpgIpSubnetsChangeListParam nsxModel.DvpgIpSubnetsChangeList) error

	// The request must contain all subnets changes needed for the DVPG; it will be rejected if any non-empty \"old_subnets\" is repeated or any \"old_subnets\" matches any \"new_subnets\" in the request.
	//
	// @param vcenterIdParam (required)
	// @param dvpgIpSubnetsChangeListParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(vcenterIdParam string, dvpgIpSubnetsChangeListParam nsxModel.DvpgIpSubnetsChangeList) error
}

type dvpgSubnetsChangeClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDvpgSubnetsChangeClient(connector vapiProtocolClient_.Connector) *dvpgSubnetsChangeClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.overlay_adoption.vcenters.dvpg_subnets_change")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := dvpgSubnetsChangeClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *dvpgSubnetsChangeClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *dvpgSubnetsChangeClient) Get(vcenterIdParam string, dvpgIdParam string) (nsxModel.DvpgIpSubnetsChangeList, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dvpgSubnetsChangeGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dvpgSubnetsChangeGetInputType(), typeConverter)
	sv.AddStructField("VcenterId", vcenterIdParam)
	sv.AddStructField("DvpgId", dvpgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.DvpgIpSubnetsChangeList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.overlay_adoption.vcenters.dvpg_subnets_change", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.DvpgIpSubnetsChangeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DvpgSubnetsChangeGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.DvpgIpSubnetsChangeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dvpgSubnetsChangeClient) Patch(vcenterIdParam string, dvpgIpSubnetsChangeListParam nsxModel.DvpgIpSubnetsChangeList) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dvpgSubnetsChangePatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dvpgSubnetsChangePatchInputType(), typeConverter)
	sv.AddStructField("VcenterId", vcenterIdParam)
	sv.AddStructField("DvpgIpSubnetsChangeList", dvpgIpSubnetsChangeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.overlay_adoption.vcenters.dvpg_subnets_change", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *dvpgSubnetsChangeClient) Update(vcenterIdParam string, dvpgIpSubnetsChangeListParam nsxModel.DvpgIpSubnetsChangeList) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dvpgSubnetsChangeUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dvpgSubnetsChangeUpdateInputType(), typeConverter)
	sv.AddStructField("VcenterId", vcenterIdParam)
	sv.AddStructField("DvpgIpSubnetsChangeList", dvpgIpSubnetsChangeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.overlay_adoption.vcenters.dvpg_subnets_change", "update", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
