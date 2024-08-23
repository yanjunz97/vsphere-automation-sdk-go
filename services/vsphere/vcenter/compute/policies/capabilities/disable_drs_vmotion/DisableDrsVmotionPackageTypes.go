/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.compute.policies.capabilities.disable_drs_vmotion.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package disable_drs_vmotion

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``CreateSpec`` class contains information used to create a new DisableDrsVmotion policy, see Policies#create, in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CreateSpec struct {
    // Identifier of a tag that can be associated with a virtual machine. Virtual machines with this tag will not be migrated away from the host on which they are powered on, except when the host is put into maintenance mode or failed over. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	VmTag string
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
    // Name of the policy. The name needs to be unique within this vCenter server. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
}

// The ``Info`` class contains information about a DisableDrsVmotion policy, see Policies#get in VMware Cloud on AWS. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type Info struct {
    // Identifier of a tag that can be associated with a virtual machine. Virtual machines with this tag will not be migrated away from the host on which they are powered on, except when the host is put into maintenance mode or failed over. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	VmTag string
    // Name of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Name string
    // Description of the policy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Description string
    // Identifier of the capability this policy is based on. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Capability string
}




func CreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
	fieldNameMap["vm_tag"] = "VmTag"
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.disable_drs_vmotion.create_spec", fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
	fieldNameMap["vm_tag"] = "VmTag"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["capability"] = bindings.NewIdType([]string{"com.vmware.vcenter.compute.policies.Capability"}, "")
	fieldNameMap["capability"] = "Capability"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.compute.policies.capabilities.disable_drs_vmotion.info", fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


