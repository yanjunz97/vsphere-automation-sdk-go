
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Job
 * Functions that implement the generated JobClient interface
 */


package reconciliation

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultJobClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultJobClient(connector client.Connector) *DefaultJobClient {
	interfaceName := "com.vmware.appliance.recovery.reconciliation.job"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
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
	jIface := DefaultJobClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	jIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	jIface.methodNameToDefMap["create"] = jIface.createMethodDefinition()
	jIface.methodNameToDefMap["get"] = jIface.getMethodDefinition()
	return &jIface
}

func (jIface *DefaultJobClient) Create(specParam JobCreateSpec) (JobInfo, error) {
	typeConverter := jIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(jIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(jobCreateInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput JobInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := jobCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	jIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := jIface.Invoke(jIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput JobInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), jobCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(JobInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), jIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (jIface *DefaultJobClient) Get() (JobInfo, error) {
	typeConverter := jIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(jIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(jobGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput JobInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := jobGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	jIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := jIface.Invoke(jIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput JobInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), jobGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(JobInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), jIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (jIface *DefaultJobClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := jIface.connector.GetApiProvider().Invoke(jIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (jIface *DefaultJobClient) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(jIface.interfaceName)
	typeConverter := jIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(jobCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(jobCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	jIface.errorBindingMap[errors.FeatureInUse{}.Error()] = errors.FeatureInUseBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.FeatureInUseBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.create method's errors.FeatureInUse error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	jIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.create method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	jIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.create method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	jIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.create method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (jIface *DefaultJobClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(jIface.interfaceName)
	typeConverter := jIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(jobGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(jobGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	jIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	jIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultJobClient.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
