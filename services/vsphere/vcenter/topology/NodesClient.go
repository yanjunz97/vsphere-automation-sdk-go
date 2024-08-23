/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Nodes
 * Used by client-side stubs.
 */

package topology


// The ``Nodes`` interface provides methods to retrieve vCenter and Platform Services Controller nodes information in the topology. This interface was added in vSphere API 6.7.2.
type NodesClient interface {

    // Returns information about all vCenter and Platform Services Controller nodes matching the NodesFilterSpec. This method was added in vSphere API 6.7.2.
    //
    // @param filterParam  Specification of matching vCenter and Platform Services Controller nodes for which information should be returned.
    // If null, the behavior is equivalent to a NodesFilterSpec with all properties null which means all nodes match the filter.
    // @return commonly used information for all vCenter and Platform Services Controller nodes matching the NodesFilterSpec.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Unauthorized  if the user doesn't have the required privileges.
    // @throws InvalidArgument  if the NodesFilterSpec#types property contains a value that is not supported.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	List(filterParam *NodesFilterSpec) ([]NodesSummary, error)

    // Retrieve details for a given identifier of the vCenter or Platform Services Controller node. This method was added in vSphere API 6.7.2.
    //
    // @param nodeParam  Identifier of the vCenter or Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the node.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.VCenter.name``.
    // @return vCenter or Platform Services Controller node details with replication partners and client affinity information as applicable. See NodesInfo.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Unauthorized  if the user doesn't have the required privileges.
    // @throws NotFound  if a node doesn't exist for given node identifier.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	Get(nodeParam string) (NodesInfo, error)
}
