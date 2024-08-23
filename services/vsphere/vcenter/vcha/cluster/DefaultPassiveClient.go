
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Passive
 * Functions that implement the generated PassiveClient interface
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

type DefaultPassiveClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultPassiveClient(connector client.Connector) *DefaultPassiveClient {
	interfaceName := "com.vmware.vcenter.vcha.cluster.passive"
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
	pIface := DefaultPassiveClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	pIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	pIface.methodNameToDefMap["check"] = pIface.checkMethodDefinition()
	pIface.methodNameToDefMap["redeploy"] = pIface.redeployMethodDefinition()
	return &pIface
}

func (pIface *DefaultPassiveClient) Check(specParam PassiveCheckSpec) (PassiveCheckResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "check")
	sv := bindings.NewStructValueBuilder(passiveCheckInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput PassiveCheckResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := passiveCheckRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput PassiveCheckResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), passiveCheckOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(PassiveCheckResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *DefaultPassiveClient) Redeploy(specParam PassiveRedeploySpec) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "redeploy")
	sv := bindings.NewStructValueBuilder(passiveRedeployInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := passiveRedeployRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}


func (pIface *DefaultPassiveClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := pIface.connector.GetApiProvider().Invoke(pIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (pIface *DefaultPassiveClient) checkMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(passiveCheckInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(passiveCheckOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "check")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnverifiedPeerBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.UnverifiedPeer error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InvalidElementConfigurationBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.InvalidElementConfiguration error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.check method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (pIface *DefaultPassiveClient) redeployMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
	typeConverter := pIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(passiveRedeployInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(passiveRedeployOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.redeploy method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.redeploy method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "redeploy")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	pIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.redeploy method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.redeploy method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnverifiedPeerBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.redeploy method's errors.UnverifiedPeer error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultPassiveClient.redeploy method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
