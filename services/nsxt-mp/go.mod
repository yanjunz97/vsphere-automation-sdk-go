module github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp

go 1.17

replace (
	github.com/vmware/vsphere-automation-sdk-go/lib => github.com/yanjunz97/vsphere-automation-sdk-go/lib v0.0.0-20240823062859-191e51e646c1
	github.com/vmware/vsphere-automation-sdk-go/runtime => github.com/yanjunz97/vsphere-automation-sdk-go/runtime v0.0.0-20240823062859-191e51e646c1
)

require (
	github.com/vmware/vsphere-automation-sdk-go/lib v0.7.0
	github.com/vmware/vsphere-automation-sdk-go/runtime v0.7.0
)

require (
	github.com/beevik/etree v1.1.0 // indirect
	github.com/gibson042/canonicaljson-go v1.0.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	golang.org/x/text v0.3.8 // indirect
)
