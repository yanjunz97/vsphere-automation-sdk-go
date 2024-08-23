
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: LibraryItems
 * Functions that implement the generated LibraryItemsClient interface
 */


package vm_template

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultLibraryItemsClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultLibraryItemsClient(connector client.Connector) *DefaultLibraryItemsClient {
	interfaceName := "com.vmware.vcenter.vm_template.library_items"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
		core.NewMethodIdentifier(interfaceIdentifier, "deploy"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	lIface := DefaultLibraryItemsClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	lIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	lIface.methodNameToDefMap["create"] = lIface.createMethodDefinition()
	lIface.methodNameToDefMap["deploy"] = lIface.deployMethodDefinition()
	lIface.methodNameToDefMap["get"] = lIface.getMethodDefinition()
	return &lIface
}

func (lIface *DefaultLibraryItemsClient) Create(specParam LibraryItemsCreateSpec) (string, error) {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(libraryItemsCreateInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := libraryItemsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), libraryItemsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *DefaultLibraryItemsClient) Deploy(templateLibraryItemParam string, specParam LibraryItemsDeploySpec) (string, error) {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "deploy")
	sv := bindings.NewStructValueBuilder(libraryItemsDeployInputType(), typeConverter)
	sv.AddStructField("TemplateLibraryItem", templateLibraryItemParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := libraryItemsDeployRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), libraryItemsDeployOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *DefaultLibraryItemsClient) Get(templateLibraryItemParam string) (*LibraryItemsInfo, error) {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(libraryItemsGetInputType(), typeConverter)
	sv.AddStructField("TemplateLibraryItem", templateLibraryItemParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *LibraryItemsInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := libraryItemsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput *LibraryItemsInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), libraryItemsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*LibraryItemsInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (lIface *DefaultLibraryItemsClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := lIface.connector.GetApiProvider().Invoke(lIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (lIface *DefaultLibraryItemsClient) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
	typeConverter := lIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(libraryItemsCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(libraryItemsCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	lIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.AlreadyExists error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ResourceInUse{}.Error()] = errors.ResourceInUseBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ResourceInUseBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.ResourceInUse error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnableToAllocateResourceBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.UnableToAllocateResource error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef10, errError10 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError10 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError10).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef10.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef11, errError11 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError11 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.create method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError11).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef11.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (lIface *DefaultLibraryItemsClient) deployMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
	typeConverter := lIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(libraryItemsDeployInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(libraryItemsDeployOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "deploy")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	lIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.AlreadyExists error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnableToAllocateResourceBindingType())
	if errError8 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.UnableToAllocateResource error - %s",
			bindings.VAPIerrorsToError(errError8).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef9, errError9 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError9 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError9).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef9.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef10, errError10 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError10 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.deploy method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError10).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef10.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (lIface *DefaultLibraryItemsClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
	typeConverter := lIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(libraryItemsGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(libraryItemsGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	lIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's errors.ResourceInaccessible error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultLibraryItemsClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
