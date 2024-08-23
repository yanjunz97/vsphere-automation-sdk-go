// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Transportzones
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

type TransportzonesClient interface {

	// Get VLAN transport-zones recommended for an EXTEND_DVPG workflow.
	//
	// @param vcenterIdParam (required)
	// @param dvpgIdParam ID of the DVPG that will be bridged to an overlay segment (required)
	// @param projectIdParam NSX multi-tenancy project id (optional)
	// @param segmentNameForAutoCreationParam Segment name for auto creation (optional)
	// @param segmentPathParam Path of an existing segment to which the DVPG will be bridged (optional)
	// @return com.vmware.nsx.model.RecommendationListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(vcenterIdParam string, dvpgIdParam string, projectIdParam *string, segmentNameForAutoCreationParam *string, segmentPathParam *string) (nsxModel.RecommendationListResult, error)
}

type transportzonesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTransportzonesClient(connector vapiProtocolClient_.Connector) *transportzonesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.overlay_adoption.vcenters.recommendation.transportzones")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := transportzonesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportzonesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportzonesClient) Get(vcenterIdParam string, dvpgIdParam string, projectIdParam *string, segmentNameForAutoCreationParam *string, segmentPathParam *string) (nsxModel.RecommendationListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportzonesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportzonesGetInputType(), typeConverter)
	sv.AddStructField("VcenterId", vcenterIdParam)
	sv.AddStructField("DvpgId", dvpgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("SegmentNameForAutoCreation", segmentNameForAutoCreationParam)
	sv.AddStructField("SegmentPath", segmentPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.RecommendationListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.overlay_adoption.vcenters.recommendation.transportzones", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.RecommendationListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportzonesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.RecommendationListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
