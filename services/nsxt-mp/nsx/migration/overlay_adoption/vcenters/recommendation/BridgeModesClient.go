// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: BridgeModes
// Used by client-side stubs.

package recommendation

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type BridgeModesClient interface {

	// Get bridge modes recommended for an EXTEND_DVPG workflow.
	//
	// @param vcenterIdParam (required)
	// @param dvpgIdParam ID of the DVPG that will be bridged to an overlay segment (required)
	// @param projectIdParam NSX multi-tenancy project id (optional)
	// @param segmentNameForAutoCreationParam Segment name for auto creation (optional)
	// @param segmentPathParam Path of an existing segment to which the DVPG will be bridged (optional)
	// @param vlanTransportZonePathParam Path of the VLAN transport-zone chosen to bridge the DVPG. (optional)
	// @return com.vmware.nsx.model.RecommendationListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(vcenterIdParam string, dvpgIdParam string, projectIdParam *string, segmentNameForAutoCreationParam *string, segmentPathParam *string, vlanTransportZonePathParam *string) (nsxModel.RecommendationListResult, error)
}

type bridgeModesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewBridgeModesClient(connector vapiProtocolClient_.Connector) *bridgeModesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.overlay_adoption.vcenters.recommendation.bridge_modes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	bIface := bridgeModesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *bridgeModesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *bridgeModesClient) Get(vcenterIdParam string, dvpgIdParam string, projectIdParam *string, segmentNameForAutoCreationParam *string, segmentPathParam *string, vlanTransportZonePathParam *string) (nsxModel.RecommendationListResult, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bridgeModesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bridgeModesGetInputType(), typeConverter)
	sv.AddStructField("VcenterId", vcenterIdParam)
	sv.AddStructField("DvpgId", dvpgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("SegmentNameForAutoCreation", segmentNameForAutoCreationParam)
	sv.AddStructField("SegmentPath", segmentPathParam)
	sv.AddStructField("VlanTransportZonePath", vlanTransportZonePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.RecommendationListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.overlay_adoption.vcenters.recommendation.bridge_modes", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.RecommendationListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BridgeModesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.RecommendationListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
