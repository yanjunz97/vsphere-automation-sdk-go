
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Witness
 * Functions that implement the generated WitnessClient interface
 */


package cluster

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultWitnessClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultWitnessClient(connector client.Connector) *DefaultWitnessClient {
	interfaceName := "com.vmware.vcenter.vcha.cluster.witness"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "check"),
		core.NewMethodIdentifier(interfaceIdentifier, "redeploy"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	wIface := DefaultWitnessClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	wIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	wIface.methodNameToDefMap["check"] = wIface.checkMethodDefinition()
	wIface.methodNameToDefMap["redeploy"] = wIface.redeployMethodDefinition()
	return &wIface
}

func (wIface *DefaultWitnessClient) Check(specParam WitnessCheckSpec) (WitnessCheckResult, error) {
	typeConverter := wIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(wIface.interfaceIdentifier, "check")
	sv := bindings.NewStructValueBuilder(witnessCheckInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput WitnessCheckResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := witnessCheckRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.Invoke(wIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput WitnessCheckResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), witnessCheckOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(WitnessCheckResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (wIface *DefaultWitnessClient) Redeploy(specParam WitnessRedeploySpec) error {
	typeConverter := wIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(wIface.interfaceIdentifier, "redeploy")
	sv := bindings.NewStructValueBuilder(witnessRedeployInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := witnessRedeployRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.Invoke(wIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}


func (wIface *DefaultWitnessClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := wIface.connector.GetApiProvider().Invoke(wIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (wIface *DefaultWitnessClient) checkMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(wIface.interfaceName)
	typeConverter := wIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(witnessCheckInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(witnessCheckOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "check")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	wIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnverifiedPeerBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.UnverifiedPeer error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InvalidElementConfigurationBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.InvalidElementConfiguration error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.check method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (wIface *DefaultWitnessClient) redeployMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(wIface.interfaceName)
	typeConverter := wIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(witnessRedeployInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(witnessRedeployOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.redeploy method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.redeploy method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "redeploy")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	wIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.redeploy method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.redeploy method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnverifiedPeerBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.redeploy method's errors.UnverifiedPeer error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	wIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultWitnessClient.redeploy method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
