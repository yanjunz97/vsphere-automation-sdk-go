// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ClusterConfigs
// Used by client-side stubs.

package security

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ClusterConfigsClient interface {

	// Read security cluster firewall configuration
	//
	// @param clusterExternalIdParam Cluster ID (required)
	// @return com.vmware.nsx_policy.model.SecurityClusterConfiguration
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(clusterExternalIdParam string) (nsx_policyModel.SecurityClusterConfiguration, error)

	// API will list all security cluster wise configuration
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.SecurityClusterConfigurationListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.SecurityClusterConfigurationListResult, error)

	// Patch security cluster configuration for DFW.
	//
	// @param clusterExternalIdParam Cluster ID (required)
	// @param securityClusterConfigurationParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(clusterExternalIdParam string, securityClusterConfigurationParam nsx_policyModel.SecurityClusterConfiguration) error

	// Update security cluster configuration for DFW.
	//
	// @param clusterExternalIdParam Cluster ID (required)
	// @param securityClusterConfigurationParam (required)
	// @return com.vmware.nsx_policy.model.SecurityClusterConfiguration
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(clusterExternalIdParam string, securityClusterConfigurationParam nsx_policyModel.SecurityClusterConfiguration) (nsx_policyModel.SecurityClusterConfiguration, error)
}

type clusterConfigsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewClusterConfigsClient(connector vapiProtocolClient_.Connector) *clusterConfigsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.settings.security.cluster_configs")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := clusterConfigsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *clusterConfigsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *clusterConfigsClient) Get(clusterExternalIdParam string) (nsx_policyModel.SecurityClusterConfiguration, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clusterConfigsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clusterConfigsGetInputType(), typeConverter)
	sv.AddStructField("ClusterExternalId", clusterExternalIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SecurityClusterConfiguration
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.security.cluster_configs", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SecurityClusterConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClusterConfigsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SecurityClusterConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clusterConfigsClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.SecurityClusterConfigurationListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clusterConfigsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clusterConfigsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SecurityClusterConfigurationListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.security.cluster_configs", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SecurityClusterConfigurationListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClusterConfigsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SecurityClusterConfigurationListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clusterConfigsClient) Patch(clusterExternalIdParam string, securityClusterConfigurationParam nsx_policyModel.SecurityClusterConfiguration) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clusterConfigsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clusterConfigsPatchInputType(), typeConverter)
	sv.AddStructField("ClusterExternalId", clusterExternalIdParam)
	sv.AddStructField("SecurityClusterConfiguration", securityClusterConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.security.cluster_configs", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *clusterConfigsClient) Update(clusterExternalIdParam string, securityClusterConfigurationParam nsx_policyModel.SecurityClusterConfiguration) (nsx_policyModel.SecurityClusterConfiguration, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clusterConfigsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clusterConfigsUpdateInputType(), typeConverter)
	sv.AddStructField("ClusterExternalId", clusterExternalIdParam)
	sv.AddStructField("SecurityClusterConfiguration", securityClusterConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SecurityClusterConfiguration
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.security.cluster_configs", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SecurityClusterConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClusterConfigsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SecurityClusterConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
