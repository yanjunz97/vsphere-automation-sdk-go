/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Scsi
 * Used by client-side stubs.
 */

package adapter


// The ``Scsi`` interface provides methods for configuring the virtual SCSI adapters of a virtual machine.
type ScsiClient interface {

    // Returns commonly used information about the virtual SCSI adapters belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual SCSI adapters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(vmParam string) ([]ScsiSummary, error)

    // Returns information about a virtual SCSI adapter.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual SCSI adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.ScsiAdapter``.
    // @return Information about the specified virtual SCSI adapter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual SCSI adapter is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string, adapterParam string) (ScsiInfo, error)

    // Adds a virtual SCSI adapter to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual SCSI adapter.
    // @return Virtual SCSI adapter identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.ScsiAdapter``.
    // @throws Error if the system reported that the SCSI adapter was created but was unable to confirm the creation because the identifier of the new adapter could not be determined.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended
    // @throws NotFound if the virtual machine is not found.
    // @throws UnableToAllocateResource if there are no more available SCSI buses on the virtual machine.
    // @throws ResourceInUse if the specified SCSI bus is in use.
    // @throws InvalidArgument if the specified SATA bus or PCI address is out of bounds.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the guest operating system of the virtual machine is not supported and spec includes null properties that default to guest-specific values.
	Create(vmParam string, specParam ScsiCreateSpec) (string, error)

    // Updates the configuration of a virtual SCSI adapter.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual SCSI adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.ScsiAdapter``.
    // @param specParam Specification for updating the virtual SCSI adapter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual SCSI adapter is not found.
    // @throws NotAllowedInCurrentState if one or more of the properties specified in the ``spec`` parameter cannot be modified due to the current power state of the virtual machine or the connection state of the virtual SCSI adapter.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Update(vmParam string, adapterParam string, specParam ScsiUpdateSpec) error

    // Removes a virtual SCSI adapter from the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual SCSI adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.ScsiAdapter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended
    // @throws NotFound if the virtual machine or virtual SCSI adapter is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Delete(vmParam string, adapterParam string) error
}
