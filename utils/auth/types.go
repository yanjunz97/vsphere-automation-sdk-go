/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

// Scheme represents the Authentication Scheme.
type Scheme string

// Info represents the details about Authentication Scheme.
type Info interface {
	// AuthScheme returns Authentication Scheme.
	AuthScheme() Scheme
	// AuthInterface gets the interface of the used Authentication Scheme.
	AuthInterface() interface{}
	// SecurityContext creates and return new Security Context from the Authentication.
	SecurityContext() (core.SecurityContext, error)
}
