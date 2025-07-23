# \DefaultAPI

All URIs are relative to *http://IPAddress*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers**](DefaultAPI.md#DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers) | **Delete** /redfish/v1/SFSSApp/Alerts({uuid}) | Delete alert
[**DeleteRedfishV1SFSSAppCDCInstanceManagers**](DefaultAPI.md#DeleteRedfishV1SFSSAppCDCInstanceManagers) | **Delete** /redfish/v1/SFSSApp/CDCInstanceManagers({InstanceId}) | Delete CDC instance
[**DeleteRedfishV1SFSSAppIpAddressManagements**](DefaultAPI.md#DeleteRedfishV1SFSSAppIpAddressManagements) | **Delete** /redfish/v1/SFSSApp/IpAddressManagements({InterfaceId}) | Delete VLAN interface
[**DeleteRedfishV1SFSSAppNTPServerIP**](DefaultAPI.md#DeleteRedfishV1SFSSAppNTPServerIP) | **Delete** /redfish/v1/SFSSApp/NTP({ServerIP}) | Remove NTP configuration
[**DeleteRedfishV1SFSSAppSFSSImages**](DefaultAPI.md#DeleteRedfishV1SFSSAppSFSSImages) | **Delete** /redfish/v1/SFSSApp/SFSSImages | Delete image
[**GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums**](DefaultAPI.md#GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums) | **Get** /redfish/v1/SFSSApp/IpAddressManagements/Enums | Get interface enums
[**GetExpandRedfishV1SFSSAppBackups**](DefaultAPI.md#GetExpandRedfishV1SFSSAppBackups) | **Get** /redfish/v1/SFSSApp/Backups?$expand&#x3D;Backups | Get all backups
[**GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand**](DefaultAPI.md#GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand) | **Get** /redfish/v1/SFSSApp/FoundationalConfigs?$expand&#x3D;FoundationalConfigs | Get all foundational configuration
[**GetExpandRedfishV1SFSSAppIpAddressManagements**](DefaultAPI.md#GetExpandRedfishV1SFSSAppIpAddressManagements) | **Get** /redfish/v1/SFSSApp/IpAddressManagements?$expand&#x3D;IpAddressManagements | Get all interfaces
[**GetExpandRedfishV1SFSSAppRestores**](DefaultAPI.md#GetExpandRedfishV1SFSSAppRestores) | **Get** /redfish/v1/SFSSApp/Restores?$expand&#x3D;Restores | Get detailed restore information
[**GetIDRedfishV1SFSSApp**](DefaultAPI.md#GetIDRedfishV1SFSSApp) | **Get** /redfish/v1/SFSSApp/SFSSImages({version}) | Get specific image
[**GetIDRedfishV1SFSSAppRadiusServers**](DefaultAPI.md#GetIDRedfishV1SFSSAppRadiusServers) | **Get** /redfish/v1/SFSSApp/RadiusServers({IP}) | Get specific RADIUS server
[**GetIDRedfishV1SFSSAppTacacsServers**](DefaultAPI.md#GetIDRedfishV1SFSSAppTacacsServers) | **Get** /redfish/v1/SFSSApp/TacacsServers({IP}) | Get specific TACACS+ server
[**GetIDredfishV1SFSSAppSFSSInterfaceList**](DefaultAPI.md#GetIDredfishV1SFSSAppSFSSInterfaceList) | **Get** /redfish/v1/SFSSApp/SFSSInterfaceList | Get all interfaces
[**GetRedfishV1SFSSApp**](DefaultAPI.md#GetRedfishV1SFSSApp) | **Get** /redfish/v1/SFSSApp | Get SFSS application details
[**GetRedfishV1SFSSAppAlertsAlertId**](DefaultAPI.md#GetRedfishV1SFSSAppAlertsAlertId) | **Get** /redfish/v1/SFSSApp/Alerts({uuid}) | Get specific alert
[**GetRedfishV1SFSSAppAuthenticationSequence**](DefaultAPI.md#GetRedfishV1SFSSAppAuthenticationSequence) | **Get** /redfish/v1/SFSSApp/AuthenticationSequence | Get authentication sequence
[**GetRedfishV1SFSSAppAuthenticationSequenceEnums**](DefaultAPI.md#GetRedfishV1SFSSAppAuthenticationSequenceEnums) | **Get** /redfish/v1/SFSSApp/AuthenticationSequence/Enums | Get authentication sequence enums
[**GetRedfishV1SFSSAppBackups**](DefaultAPI.md#GetRedfishV1SFSSAppBackups) | **Get** /redfish/v1/SFSSApp/Backups | Get all backups
[**GetRedfishV1SFSSAppBackupsID**](DefaultAPI.md#GetRedfishV1SFSSAppBackupsID) | **Get** /redfish/v1/SFSSApp/Backups({ID}) | Get specific backup
[**GetRedfishV1SFSSAppCDCHealthStatus**](DefaultAPI.md#GetRedfishV1SFSSAppCDCHealthStatus) | **Get** /redfish/v1/SFSSApp/CDCHealthStatus | Get CDC health
[**GetRedfishV1SFSSAppCDCHealthStatusID**](DefaultAPI.md#GetRedfishV1SFSSAppCDCHealthStatusID) | **Get** /redfish/v1/SFSSApp/CDCHealthStatus({ID}) | Get specific CDC health
[**GetRedfishV1SFSSAppCDCInstanceManagers**](DefaultAPI.md#GetRedfishV1SFSSAppCDCInstanceManagers) | **Get** /redfish/v1/SFSSApp/CDCInstanceManagers | Get all CDCs
[**GetRedfishV1SFSSAppCDCInstanceManagersEnums**](DefaultAPI.md#GetRedfishV1SFSSAppCDCInstanceManagersEnums) | **Get** /redfish/v1/SFSSApp/CDCInstanceManagers/Enums | Get CDC enums
[**GetRedfishV1SFSSAppCDCInstanceManagersID**](DefaultAPI.md#GetRedfishV1SFSSAppCDCInstanceManagersID) | **Get** /redfish/v1/SFSSApp/CDCInstanceManagers({InstanceId}) | Get specific CDC information
[**GetRedfishV1SFSSAppDevice**](DefaultAPI.md#GetRedfishV1SFSSAppDevice) | **Get** /redfish/v1/SFSSApp/Device | Get device details
[**GetRedfishV1SFSSAppEnums**](DefaultAPI.md#GetRedfishV1SFSSAppEnums) | **Get** /redfish/v1/SFSSApp/SFSSImages/Enums | Get image enums
[**GetRedfishV1SFSSAppEvents**](DefaultAPI.md#GetRedfishV1SFSSAppEvents) | **Get** /redfish/v1/SFSSApp/Events | Get all events
[**GetRedfishV1SFSSAppEventsID**](DefaultAPI.md#GetRedfishV1SFSSAppEventsID) | **Get** /redfish/v1/SFSSApp/Events({ID}) | Get specific event
[**GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier**](DefaultAPI.md#GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier) | **Get** /redfish/v1/SFSSApp/FoundationalConfigs({InstanceIdentifer}) | Get specific foundational configuration
[**GetRedfishV1SFSSAppGlobalSettings**](DefaultAPI.md#GetRedfishV1SFSSAppGlobalSettings) | **Get** /redfish/v1/SFSSApp/GlobalSettings | Get global settings
[**GetRedfishV1SFSSAppIpAddressManagements**](DefaultAPI.md#GetRedfishV1SFSSAppIpAddressManagements) | **Get** /redfish/v1/SFSSApp/IpAddressManagements | Get all interfaces
[**GetRedfishV1SFSSAppIpAddressManagementsInterface**](DefaultAPI.md#GetRedfishV1SFSSAppIpAddressManagementsInterface) | **Get** /redfish/v1/SFSSApp/IpAddressManagements({interface}) | Get specific interface
[**GetRedfishV1SFSSAppLicenses**](DefaultAPI.md#GetRedfishV1SFSSAppLicenses) | **Get** /redfish/v1/SFSSApp/Licenses | Get license count
[**GetRedfishV1SFSSAppLicensesLicenseId**](DefaultAPI.md#GetRedfishV1SFSSAppLicensesLicenseId) | **Get** /redfish/v1/SFSSApp/Licenses({LicenseId}) | Get specific license
[**GetRedfishV1SFSSAppLicensesexpandLicenses**](DefaultAPI.md#GetRedfishV1SFSSAppLicensesexpandLicenses) | **Get** /redfish/v1/SFSSApp/Licenses?$expand&#x3D;Licenses | Get detailed license information
[**GetRedfishV1SFSSAppNTP**](DefaultAPI.md#GetRedfishV1SFSSAppNTP) | **Get** /redfish/v1/SFSSApp/NTP | Get NTP server information
[**GetRedfishV1SFSSAppRadiusServers**](DefaultAPI.md#GetRedfishV1SFSSAppRadiusServers) | **Get** /redfish/v1/SFSSApp/RadiusServers | Get all RADIUS servers
[**GetRedfishV1SFSSAppRadiusServersSequence**](DefaultAPI.md#GetRedfishV1SFSSAppRadiusServersSequence) | **Get** /redfish/v1/SFSSApp/RadiusServers/Sequence | Get all RADIUS servers
[**GetRedfishV1SFSSAppRestores**](DefaultAPI.md#GetRedfishV1SFSSAppRestores) | **Get** /redfish/v1/SFSSApp/Restores | Get all restores
[**GetRedfishV1SFSSAppRestoresID**](DefaultAPI.md#GetRedfishV1SFSSAppRestoresID) | **Get** /redfish/v1/SFSSApp/Restores({ID}) | Get specific restore
[**GetRedfishV1SFSSAppSFSSHealthStatus**](DefaultAPI.md#GetRedfishV1SFSSAppSFSSHealthStatus) | **Get** /redfish/v1/SFSSApp/SFSSHealthStatus | Get system health
[**GetRedfishV1SFSSAppSosReports**](DefaultAPI.md#GetRedfishV1SFSSAppSosReports) | **Get** /redfish/v1/SFSSApp/SosReports | Get SOS report
[**GetRedfishV1SFSSAppSosReportsID**](DefaultAPI.md#GetRedfishV1SFSSAppSosReportsID) | **Get** /redfish/v1/SFSSApp/SosReports({ID}) | Get specific SOS report
[**GetRedfishV1SFSSAppSosReportsexpandSosReports**](DefaultAPI.md#GetRedfishV1SFSSAppSosReportsexpandSosReports) | **Get** /redfish/v1/SFSSApp/SosReports?$expand&#x3D;SosReports | Get all SOS reports
[**GetRedfishV1SFSSAppTacacsServers**](DefaultAPI.md#GetRedfishV1SFSSAppTacacsServers) | **Get** /redfish/v1/SFSSApp/TacacsServers | Get all TACACS+ servers
[**GetRedfishV1SFSSAppTacacsServersSequence**](DefaultAPI.md#GetRedfishV1SFSSAppTacacsServersSequence) | **Get** /redfish/v1/SFSSApp/TacacsServers/Sequence | Get all TACACS+ servers
[**GetRedfishV1SFSSAppUserActivityAudit**](DefaultAPI.md#GetRedfishV1SFSSAppUserActivityAudit) | **Get** /redfish/v1/SFSSApp/UserActivityAudit | Get user activities
[**GetRedfishV1SFSSAppUserActivityAuditID**](DefaultAPI.md#GetRedfishV1SFSSAppUserActivityAuditID) | **Get** /redfish/v1/SFSSApp/UserActivityAudit({ID}) | Get specific user activity
[**GetRedfishV1SFSSApp_0**](DefaultAPI.md#GetRedfishV1SFSSApp_0) | **Get** /redfish/v1/SFSSApp/SFSSImages | Get all images
[**PostRedfishV1SFSSAppAlerts**](DefaultAPI.md#PostRedfishV1SFSSAppAlerts) | **Post** /redfish/v1/SFSSApp/Alerts | Add alert
[**PostRedfishV1SFSSAppAuthenticationSequence**](DefaultAPI.md#PostRedfishV1SFSSAppAuthenticationSequence) | **Post** /redfish/v1/SFSSApp/AuthenticationSequence | Add authentication sequence
[**PostRedfishV1SFSSAppBackups**](DefaultAPI.md#PostRedfishV1SFSSAppBackups) | **Post** /redfish/v1/SFSSApp/Backups | Perform backup
[**PostRedfishV1SFSSAppChangePassword**](DefaultAPI.md#PostRedfishV1SFSSAppChangePassword) | **Post** /redfish/v1/SFSSApp/ChangePassword | Change admin password
[**PostRedfishV1SFSSAppFabricManagerInfoPost**](DefaultAPI.md#PostRedfishV1SFSSAppFabricManagerInfoPost) | **Post** /redfish/v1/SFSSApp/CDCInstanceManagers | Configure CDC instance
[**PostRedfishV1SFSSAppGlobalSettings**](DefaultAPI.md#PostRedfishV1SFSSAppGlobalSettings) | **Post** /redfish/v1/SFSSApp/GlobalSettings | Configure global settings
[**PostRedfishV1SFSSAppIpAddressManagements**](DefaultAPI.md#PostRedfishV1SFSSAppIpAddressManagements) | **Post** /redfish/v1/SFSSApp/IpAddressManagements | Configure interface
[**PostRedfishV1SFSSAppLicenses**](DefaultAPI.md#PostRedfishV1SFSSAppLicenses) | **Post** /redfish/v1/SFSSApp/Licenses | Install a license
[**PostRedfishV1SFSSAppNTP1**](DefaultAPI.md#PostRedfishV1SFSSAppNTP1) | **Post** /redfish/v1/SFSSApp/NTP | Configure NTP server
[**PostRedfishV1SFSSAppRadiusServers**](DefaultAPI.md#PostRedfishV1SFSSAppRadiusServers) | **Post** /redfish/v1/SFSSApp/RadiusServers | Configure RADIUS server
[**PostRedfishV1SFSSAppRestores**](DefaultAPI.md#PostRedfishV1SFSSAppRestores) | **Post** /redfish/v1/SFSSApp/Restores | Restore a backup
[**PostRedfishV1SFSSAppSFSSImages**](DefaultAPI.md#PostRedfishV1SFSSAppSFSSImages) | **Post** /redfish/v1/SFSSApp/SFSSImages | Add image
[**PostRedfishV1SFSSAppTacacsServers**](DefaultAPI.md#PostRedfishV1SFSSAppTacacsServers) | **Post** /redfish/v1/SFSSApp/TacacsServers | Configure TACACS+ server
[**PutRedfishV1SFSSApp**](DefaultAPI.md#PutRedfishV1SFSSApp) | **Put** /redfish/v1/SFSSApp | Upgrade SFSS application
[**PutRedfishV1SFSSAppAlerts**](DefaultAPI.md#PutRedfishV1SFSSAppAlerts) | **Put** /redfish/v1/SFSSApp/Alerts({uuid}) | Update alert
[**PutRedfishV1SFSSAppCDCInstanceManagers**](DefaultAPI.md#PutRedfishV1SFSSAppCDCInstanceManagers) | **Put** /redfish/v1/SFSSApp/CDCInstanceManagers({InstanceId}) | Update CDC instance
[**PutRedfishV1SFSSAppIpAddressManagements**](DefaultAPI.md#PutRedfishV1SFSSAppIpAddressManagements) | **Put** /redfish/v1/SFSSApp/IpAddressManagements({InterfaceId}) | Update interface
[**PutRedfishV1SFSSAppLicenses**](DefaultAPI.md#PutRedfishV1SFSSAppLicenses) | **Put** /redfish/v1/SFSSApp/Licenses | Accept EULA
[**PutRedfishV1SFSSAppNTP**](DefaultAPI.md#PutRedfishV1SFSSAppNTP) | **Put** /redfish/v1/SFSSApp/NTP | Enable or disable NTP service
[**PutRedfishV1SFSSApp_0**](DefaultAPI.md#PutRedfishV1SFSSApp_0) | **Put** /redfish/v1/SFSSApp/SFSSImages | Upgrade application



## DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers

> DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers200Response DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers(ctx, uuid).Execute()

Delete alert



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "uuid_example" // string | Alert UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers`: DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Alert UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSAppAlertsCDCInstanceManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers200Response**](DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSAppCDCInstanceManagers

> PostRedfishV1SFSSAppFabricManagerInfoPost200Response DeleteRedfishV1SFSSAppCDCInstanceManagers(ctx, instanceId).Execute()

Delete CDC instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | InstanceIdentifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSAppCDCInstanceManagers(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSAppCDCInstanceManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSAppCDCInstanceManagers`: PostRedfishV1SFSSAppFabricManagerInfoPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSAppCDCInstanceManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | InstanceIdentifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSAppCDCInstanceManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostRedfishV1SFSSAppFabricManagerInfoPost200Response**](PostRedfishV1SFSSAppFabricManagerInfoPost200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSAppIpAddressManagements

> DeleteRedfishV1SFSSAppIpAddressManagements(ctx, interfaceId).Execute()

Delete VLAN interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	interfaceId := "interfaceId_example" // string | Vlan interface Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSAppIpAddressManagements(context.Background(), interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSAppIpAddressManagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | Vlan interface Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSAppIpAddressManagementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSAppNTPServerIP

> DeleteRedfishV1SFSSAppNTPServerIP200Response DeleteRedfishV1SFSSAppNTPServerIP(ctx, serverIP).Execute()

Remove NTP configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverIP := "serverIP_example" // string | IP address of the NTP server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSAppNTPServerIP(context.Background(), serverIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSAppNTPServerIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSAppNTPServerIP`: DeleteRedfishV1SFSSAppNTPServerIP200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSAppNTPServerIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverIP** | **string** | IP address of the NTP server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSAppNTPServerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRedfishV1SFSSAppNTPServerIP200Response**](DeleteRedfishV1SFSSAppNTPServerIP200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSAppSFSSImages

> PutRedfishV1SFSSApp200Response DeleteRedfishV1SFSSAppSFSSImages(ctx).Versiontbd(versiontbd).Execute()

Delete image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	versiontbd := "versiontbd_example" // string | ImageId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSAppSFSSImages(context.Background()).Versiontbd(versiontbd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSAppSFSSImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSAppSFSSImages`: PutRedfishV1SFSSApp200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSAppSFSSImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSAppSFSSImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versiontbd** | **string** | ImageId | 

### Return type

[**PutRedfishV1SFSSApp200Response**](PutRedfishV1SFSSApp200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums

> GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums200Response GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums(ctx).Execute()

Get interface enums



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums`: GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnumsRedfishV1SFSSAppIpAddressManagementsEnumsRequest struct via the builder pattern


### Return type

[**GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums200Response**](GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpandRedfishV1SFSSAppBackups

> GetExpandRedfishV1SFSSAppBackups200Response GetExpandRedfishV1SFSSAppBackups(ctx).Execute()

Get all backups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExpandRedfishV1SFSSAppBackups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExpandRedfishV1SFSSAppBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpandRedfishV1SFSSAppBackups`: GetExpandRedfishV1SFSSAppBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExpandRedfishV1SFSSAppBackups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpandRedfishV1SFSSAppBackupsRequest struct via the builder pattern


### Return type

[**GetExpandRedfishV1SFSSAppBackups200Response**](GetExpandRedfishV1SFSSAppBackups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand

> GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200Response GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand(ctx).Execute()

Get all foundational configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand`: GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpandRedfishV1SFSSAppFoundationalConfigsGETExpandRequest struct via the builder pattern


### Return type

[**GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200Response**](GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpandRedfishV1SFSSAppIpAddressManagements

> GetExpandRedfishV1SFSSAppIpAddressManagements200Response GetExpandRedfishV1SFSSAppIpAddressManagements(ctx).Execute()

Get all interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExpandRedfishV1SFSSAppIpAddressManagements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExpandRedfishV1SFSSAppIpAddressManagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpandRedfishV1SFSSAppIpAddressManagements`: GetExpandRedfishV1SFSSAppIpAddressManagements200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExpandRedfishV1SFSSAppIpAddressManagements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpandRedfishV1SFSSAppIpAddressManagementsRequest struct via the builder pattern


### Return type

[**GetExpandRedfishV1SFSSAppIpAddressManagements200Response**](GetExpandRedfishV1SFSSAppIpAddressManagements200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpandRedfishV1SFSSAppRestores

> GetExpandRedfishV1SFSSAppRestores200Response GetExpandRedfishV1SFSSAppRestores(ctx).Execute()

Get detailed restore information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExpandRedfishV1SFSSAppRestores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExpandRedfishV1SFSSAppRestores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpandRedfishV1SFSSAppRestores`: GetExpandRedfishV1SFSSAppRestores200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExpandRedfishV1SFSSAppRestores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpandRedfishV1SFSSAppRestoresRequest struct via the builder pattern


### Return type

[**GetExpandRedfishV1SFSSAppRestores200Response**](GetExpandRedfishV1SFSSAppRestores200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIDRedfishV1SFSSApp

> GetIDRedfishV1SFSSApp200Response GetIDRedfishV1SFSSApp(ctx, version).Execute()

Get specific image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	version := "version_example" // string | The version of the SFSS image to retrieve information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetIDRedfishV1SFSSApp(context.Background(), version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIDRedfishV1SFSSApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIDRedfishV1SFSSApp`: GetIDRedfishV1SFSSApp200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIDRedfishV1SFSSApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | The version of the SFSS image to retrieve information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIDRedfishV1SFSSAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIDRedfishV1SFSSApp200Response**](GetIDRedfishV1SFSSApp200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIDRedfishV1SFSSAppRadiusServers

> GetIDRedfishV1SFSSAppRadiusServers200Response GetIDRedfishV1SFSSAppRadiusServers(ctx, iP).Execute()

Get specific RADIUS server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iP := "iP_example" // string | IP address of the RADIUS server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetIDRedfishV1SFSSAppRadiusServers(context.Background(), iP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIDRedfishV1SFSSAppRadiusServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIDRedfishV1SFSSAppRadiusServers`: GetIDRedfishV1SFSSAppRadiusServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIDRedfishV1SFSSAppRadiusServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iP** | **string** | IP address of the RADIUS server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIDRedfishV1SFSSAppRadiusServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIDRedfishV1SFSSAppRadiusServers200Response**](GetIDRedfishV1SFSSAppRadiusServers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIDRedfishV1SFSSAppTacacsServers

> GetIDRedfishV1SFSSAppTacacsServers200Response GetIDRedfishV1SFSSAppTacacsServers(ctx, iP).Execute()

Get specific TACACS+ server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iP := "iP_example" // string | IP address of the TACACS+ server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetIDRedfishV1SFSSAppTacacsServers(context.Background(), iP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIDRedfishV1SFSSAppTacacsServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIDRedfishV1SFSSAppTacacsServers`: GetIDRedfishV1SFSSAppTacacsServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIDRedfishV1SFSSAppTacacsServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iP** | **string** | IP address of the TACACS+ server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIDRedfishV1SFSSAppTacacsServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIDRedfishV1SFSSAppTacacsServers200Response**](GetIDRedfishV1SFSSAppTacacsServers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIDredfishV1SFSSAppSFSSInterfaceList

> GetIDredfishV1SFSSAppSFSSInterfaceList200Response GetIDredfishV1SFSSAppSFSSInterfaceList(ctx).Execute()

Get all interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetIDredfishV1SFSSAppSFSSInterfaceList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIDredfishV1SFSSAppSFSSInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIDredfishV1SFSSAppSFSSInterfaceList`: GetIDredfishV1SFSSAppSFSSInterfaceList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIDredfishV1SFSSAppSFSSInterfaceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIDredfishV1SFSSAppSFSSInterfaceListRequest struct via the builder pattern


### Return type

[**GetIDredfishV1SFSSAppSFSSInterfaceList200Response**](GetIDredfishV1SFSSAppSFSSInterfaceList200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSApp

> GetRedfishV1SFSSApp200Response GetRedfishV1SFSSApp(ctx).Execute()

Get SFSS application details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSApp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSApp`: GetRedfishV1SFSSApp200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSApp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSApp200Response**](GetRedfishV1SFSSApp200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppAlertsAlertId

> GetRedfishV1SFSSAppAlertsAlertId(ctx, uuid).Execute()

Get specific alert



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "uuid_example" // string | Alert UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppAlertsAlertId(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppAlertsAlertId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Alert UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppAlertsAlertIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppAuthenticationSequence

> GetRedfishV1SFSSAppAuthenticationSequence200Response GetRedfishV1SFSSAppAuthenticationSequence(ctx).Execute()

Get authentication sequence



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppAuthenticationSequence(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppAuthenticationSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppAuthenticationSequence`: GetRedfishV1SFSSAppAuthenticationSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppAuthenticationSequence`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppAuthenticationSequenceRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppAuthenticationSequence200Response**](GetRedfishV1SFSSAppAuthenticationSequence200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppAuthenticationSequenceEnums

> GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response GetRedfishV1SFSSAppAuthenticationSequenceEnums(ctx).Execute()

Get authentication sequence enums



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppAuthenticationSequenceEnums(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppAuthenticationSequenceEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppAuthenticationSequenceEnums`: GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppAuthenticationSequenceEnums`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppAuthenticationSequenceEnumsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response**](GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppBackups

> GetRedfishV1SFSSAppBackups200Response GetRedfishV1SFSSAppBackups(ctx).Execute()

Get all backups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppBackups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppBackups`: GetRedfishV1SFSSAppBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppBackups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppBackupsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppBackups200Response**](GetRedfishV1SFSSAppBackups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppBackupsID

> GetRedfishV1SFSSAppBackupsID200Response GetRedfishV1SFSSAppBackupsID(ctx, iD).Execute()

Get specific backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iD := "iD_example" // string | Backup identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppBackupsID(context.Background(), iD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppBackupsID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppBackupsID`: GetRedfishV1SFSSAppBackupsID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppBackupsID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** | Backup identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppBackupsIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppBackupsID200Response**](GetRedfishV1SFSSAppBackupsID200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppCDCHealthStatus

> GetRedfishV1SFSSAppCDCHealthStatus200Response GetRedfishV1SFSSAppCDCHealthStatus(ctx).Execute()

Get CDC health



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppCDCHealthStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppCDCHealthStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppCDCHealthStatus`: GetRedfishV1SFSSAppCDCHealthStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppCDCHealthStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppCDCHealthStatusRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppCDCHealthStatus200Response**](GetRedfishV1SFSSAppCDCHealthStatus200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppCDCHealthStatusID

> GetRedfishV1SFSSAppCDCHealthStatusID200Response GetRedfishV1SFSSAppCDCHealthStatusID(ctx, iD).Execute()

Get specific CDC health



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iD := "iD_example" // string | CDC instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppCDCHealthStatusID(context.Background(), iD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppCDCHealthStatusID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppCDCHealthStatusID`: GetRedfishV1SFSSAppCDCHealthStatusID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppCDCHealthStatusID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** | CDC instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppCDCHealthStatusIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppCDCHealthStatusID200Response**](GetRedfishV1SFSSAppCDCHealthStatusID200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppCDCInstanceManagers

> GetRedfishV1SFSSAppCDCInstanceManagers200Response GetRedfishV1SFSSAppCDCInstanceManagers(ctx).Expand(expand).Execute()

Get all CDCs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	expand := "expand_example" // string | CDCInstanceManagers (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagers(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppCDCInstanceManagers`: GetRedfishV1SFSSAppCDCInstanceManagers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppCDCInstanceManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | CDCInstanceManagers | 

### Return type

[**GetRedfishV1SFSSAppCDCInstanceManagers200Response**](GetRedfishV1SFSSAppCDCInstanceManagers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppCDCInstanceManagersEnums

> GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response GetRedfishV1SFSSAppCDCInstanceManagersEnums(ctx).Execute()

Get CDC enums



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagersEnums(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagersEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppCDCInstanceManagersEnums`: GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagersEnums`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppCDCInstanceManagersEnumsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response**](GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppCDCInstanceManagersID

> GetRedfishV1SFSSAppCDCInstanceManagersID200Response GetRedfishV1SFSSAppCDCInstanceManagersID(ctx, instanceId).Execute()

Get specific CDC information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := int32(56) // int32 | The identifier of the CDC instance to retrieve information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagersID(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagersID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppCDCInstanceManagersID`: GetRedfishV1SFSSAppCDCInstanceManagersID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppCDCInstanceManagersID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int32** | The identifier of the CDC instance to retrieve information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppCDCInstanceManagersIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppCDCInstanceManagersID200Response**](GetRedfishV1SFSSAppCDCInstanceManagersID200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppDevice

> GetRedfishV1SFSSAppDevice200Response GetRedfishV1SFSSAppDevice(ctx).Body(body).Execute()

Get device details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppDevice(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppDevice`: GetRedfishV1SFSSAppDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetRedfishV1SFSSAppDevice200Response**](GetRedfishV1SFSSAppDevice200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppEnums

> GetRedfishV1SFSSAppEnums200Response GetRedfishV1SFSSAppEnums(ctx).Execute()

Get image enums



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppEnums(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppEnums`: GetRedfishV1SFSSAppEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppEnums`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppEnumsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppEnums200Response**](GetRedfishV1SFSSAppEnums200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppEvents

> GetRedfishV1SFSSAppEvents200Response GetRedfishV1SFSSAppEvents(ctx).Top2(top2).FilterCDCInstanceEq1(filterCDCInstanceEq1).Execute()

Get all events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	top2 := int32(56) // int32 | This parameter is used in pagination logic to fetch the top two records. (optional)
	filterCDCInstanceEq1 := "filterCDCInstanceEq1_example" // string | This parameter is used to fetch the specific record that matches the condition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppEvents(context.Background()).Top2(top2).FilterCDCInstanceEq1(filterCDCInstanceEq1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppEvents`: GetRedfishV1SFSSAppEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top2** | **int32** | This parameter is used in pagination logic to fetch the top two records. | 
 **filterCDCInstanceEq1** | **string** | This parameter is used to fetch the specific record that matches the condition. | 

### Return type

[**GetRedfishV1SFSSAppEvents200Response**](GetRedfishV1SFSSAppEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppEventsID

> GetRedfishV1SFSSAppEventsID200Response GetRedfishV1SFSSAppEventsID(ctx, iD).Execute()

Get specific event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iD := "iD_example" // string | Event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppEventsID(context.Background(), iD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppEventsID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppEventsID`: GetRedfishV1SFSSAppEventsID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppEventsID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** | Event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppEventsIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppEventsID200Response**](GetRedfishV1SFSSAppEventsID200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier

> GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier200Response GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier(ctx, instanceIdentifer).Execute()

Get specific foundational configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceIdentifer := "instanceIdentifer_example" // string | Instance identifier of the CDC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier(context.Background(), instanceIdentifer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier`: GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceIdentifer** | **string** | Instance identifier of the CDC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier200Response**](GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppGlobalSettings

> GetRedfishV1SFSSAppGlobalSettings200Response GetRedfishV1SFSSAppGlobalSettings(ctx).Execute()

Get global settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppGlobalSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppGlobalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppGlobalSettings`: GetRedfishV1SFSSAppGlobalSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppGlobalSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppGlobalSettingsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppGlobalSettings200Response**](GetRedfishV1SFSSAppGlobalSettings200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppIpAddressManagements

> GetRedfishV1SFSSAppIpAddressManagements200Response GetRedfishV1SFSSAppIpAddressManagements(ctx).Execute()

Get all interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppIpAddressManagements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppIpAddressManagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppIpAddressManagements`: GetRedfishV1SFSSAppIpAddressManagements200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppIpAddressManagements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppIpAddressManagementsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppIpAddressManagements200Response**](GetRedfishV1SFSSAppIpAddressManagements200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppIpAddressManagementsInterface

> GetRedfishV1SFSSAppIpAddressManagementsInterface200Response GetRedfishV1SFSSAppIpAddressManagementsInterface(ctx, interface_).Execute()

Get specific interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	interface_ := "interface__example" // string | Interface identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppIpAddressManagementsInterface(context.Background(), interface_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppIpAddressManagementsInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppIpAddressManagementsInterface`: GetRedfishV1SFSSAppIpAddressManagementsInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppIpAddressManagementsInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interface_** | **string** | Interface identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppIpAddressManagementsInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppIpAddressManagementsInterface200Response**](GetRedfishV1SFSSAppIpAddressManagementsInterface200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppLicenses

> GetRedfishV1SFSSAppLicenses200Response GetRedfishV1SFSSAppLicenses(ctx).Execute()

Get license count



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppLicenses`: GetRedfishV1SFSSAppLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppLicensesRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppLicenses200Response**](GetRedfishV1SFSSAppLicenses200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppLicensesLicenseId

> GetRedfishV1SFSSAppLicensesLicenseId200Response GetRedfishV1SFSSAppLicensesLicenseId(ctx, licenseId).Execute()

Get specific license



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	licenseId := "licenseId_example" // string | License identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppLicensesLicenseId(context.Background(), licenseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppLicensesLicenseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppLicensesLicenseId`: GetRedfishV1SFSSAppLicensesLicenseId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppLicensesLicenseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseId** | **string** | License identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppLicensesLicenseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppLicensesLicenseId200Response**](GetRedfishV1SFSSAppLicensesLicenseId200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppLicensesexpandLicenses

> GetRedfishV1SFSSAppLicensesExpandLicenses200Response GetRedfishV1SFSSAppLicensesexpandLicenses(ctx).Execute()

Get detailed license information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppLicensesexpandLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppLicensesexpandLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppLicensesexpandLicenses`: GetRedfishV1SFSSAppLicensesExpandLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppLicensesexpandLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppLicensesexpandLicensesRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppLicensesExpandLicenses200Response**](GetRedfishV1SFSSAppLicensesExpandLicenses200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppNTP

> GetRedfishV1SFSSAppNTP200Response GetRedfishV1SFSSAppNTP(ctx).Execute()

Get NTP server information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppNTP(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppNTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppNTP`: GetRedfishV1SFSSAppNTP200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppNTP`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppNTPRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppNTP200Response**](GetRedfishV1SFSSAppNTP200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppRadiusServers

> GetRedfishV1SFSSAppRadiusServers200Response GetRedfishV1SFSSAppRadiusServers(ctx).Execute()

Get all RADIUS servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppRadiusServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppRadiusServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppRadiusServers`: GetRedfishV1SFSSAppRadiusServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppRadiusServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppRadiusServersRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppRadiusServers200Response**](GetRedfishV1SFSSAppRadiusServers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppRadiusServersSequence

> GetRedfishV1SFSSAppRadiusServersSequence(ctx).Execute()

Get all RADIUS servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppRadiusServersSequence(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppRadiusServersSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppRadiusServersSequenceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppRestores

> GetRedfishV1SFSSAppRestores200Response GetRedfishV1SFSSAppRestores(ctx).Execute()

Get all restores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppRestores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppRestores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppRestores`: GetRedfishV1SFSSAppRestores200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppRestores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppRestoresRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppRestores200Response**](GetRedfishV1SFSSAppRestores200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppRestoresID

> GetRedfishV1SFSSAppRestoresID200Response GetRedfishV1SFSSAppRestoresID(ctx, iD).Execute()

Get specific restore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iD := "iD_example" // string | Restore identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppRestoresID(context.Background(), iD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppRestoresID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppRestoresID`: GetRedfishV1SFSSAppRestoresID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppRestoresID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** | Restore identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppRestoresIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppRestoresID200Response**](GetRedfishV1SFSSAppRestoresID200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppSFSSHealthStatus

> GetRedfishV1SFSSAppSFSSHealthStatus200Response GetRedfishV1SFSSAppSFSSHealthStatus(ctx).Execute()

Get system health



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppSFSSHealthStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppSFSSHealthStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppSFSSHealthStatus`: GetRedfishV1SFSSAppSFSSHealthStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppSFSSHealthStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppSFSSHealthStatusRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppSFSSHealthStatus200Response**](GetRedfishV1SFSSAppSFSSHealthStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppSosReports

> GetRedfishV1SFSSAppSosReports200Response GetRedfishV1SFSSAppSosReports(ctx).Execute()

Get SOS report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppSosReports(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppSosReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppSosReports`: GetRedfishV1SFSSAppSosReports200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppSosReports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppSosReportsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppSosReports200Response**](GetRedfishV1SFSSAppSosReports200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppSosReportsID

> GetRedfishV1SFSSAppSosReportsID200Response GetRedfishV1SFSSAppSosReportsID(ctx, iD).Execute()

Get specific SOS report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iD := "iD_example" // string | SOS report ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppSosReportsID(context.Background(), iD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppSosReportsID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppSosReportsID`: GetRedfishV1SFSSAppSosReportsID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppSosReportsID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** | SOS report ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppSosReportsIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppSosReportsID200Response**](GetRedfishV1SFSSAppSosReportsID200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppSosReportsexpandSosReports

> GetRedfishV1SFSSAppSosReportsExpandSosReports200Response GetRedfishV1SFSSAppSosReportsexpandSosReports(ctx).Execute()

Get all SOS reports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppSosReportsexpandSosReports(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppSosReportsexpandSosReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppSosReportsexpandSosReports`: GetRedfishV1SFSSAppSosReportsExpandSosReports200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppSosReportsexpandSosReports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppSosReportsexpandSosReportsRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppSosReportsExpandSosReports200Response**](GetRedfishV1SFSSAppSosReportsExpandSosReports200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppTacacsServers

> GetRedfishV1SFSSAppTacacsServers200Response GetRedfishV1SFSSAppTacacsServers(ctx).Execute()

Get all TACACS+ servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppTacacsServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppTacacsServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppTacacsServers`: GetRedfishV1SFSSAppTacacsServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppTacacsServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppTacacsServersRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppTacacsServers200Response**](GetRedfishV1SFSSAppTacacsServers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppTacacsServersSequence

> GetRedfishV1SFSSAppTacacsServersSequence(ctx).Execute()

Get all TACACS+ servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppTacacsServersSequence(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppTacacsServersSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppTacacsServersSequenceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppUserActivityAudit

> GetRedfishV1SFSSAppUserActivityAudit200Response GetRedfishV1SFSSAppUserActivityAudit(ctx).Execute()

Get user activities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppUserActivityAudit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppUserActivityAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppUserActivityAudit`: GetRedfishV1SFSSAppUserActivityAudit200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppUserActivityAudit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppUserActivityAuditRequest struct via the builder pattern


### Return type

[**GetRedfishV1SFSSAppUserActivityAudit200Response**](GetRedfishV1SFSSAppUserActivityAudit200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSAppUserActivityAuditID

> GetRedfishV1SFSSAppUserActivityAuditID200Response GetRedfishV1SFSSAppUserActivityAuditID(ctx, iD).Execute()

Get specific user activity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iD := "iD_example" // string | User activity audit identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSAppUserActivityAuditID(context.Background(), iD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSAppUserActivityAuditID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSAppUserActivityAuditID`: GetRedfishV1SFSSAppUserActivityAuditID200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSAppUserActivityAuditID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** | User activity audit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSAppUserActivityAuditIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSAppUserActivityAuditID200Response**](GetRedfishV1SFSSAppUserActivityAuditID200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSApp_0

> GetRedfishV1SFSSApp200Response GetRedfishV1SFSSApp_0(ctx).Execute()

Get all images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSApp_0(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSApp_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSApp_0`: GetRedfishV1SFSSApp200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSApp_0`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSApp_1Request struct via the builder pattern


### Return type

[**GetRedfishV1SFSSApp200Response**](GetRedfishV1SFSSApp200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppAlerts

> PostRedfishV1SFSSAppAlerts200Response PostRedfishV1SFSSAppAlerts(ctx).PostRedfishV1SFSSAppAlertsRequest(postRedfishV1SFSSAppAlertsRequest).Execute()

Add alert



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppAlertsRequest := *openapiclient.NewPostRedfishV1SFSSAppAlertsRequest("Protocol_example", "Context_example", []string{"EventTypes_example"}, []string{"CdcInstances_example"}, []string{"HttpHeaders_example"}, "Destination_example") // PostRedfishV1SFSSAppAlertsRequest | {     \"Protocol\": \"redfish\",     \"Context\": \"SomeSubscription\",     \"EventTypes\": [          \"Alert\"      ],     \"CdcInstances\": [         \"APP\"     ],     \"HttpHeaders\": [         \"Authorization: Basic ZG52dXNlcjpAIThwSU1vSQ==\",         \"ExternalServerRequiredHeader: ItsValue\"     ],     \"Destination\": \"https://[ipv4/ipv6]/external/Server/eventHandler\"  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppAlerts(context.Background()).PostRedfishV1SFSSAppAlertsRequest(postRedfishV1SFSSAppAlertsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppAlerts`: PostRedfishV1SFSSAppAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppAlertsRequest** | [**PostRedfishV1SFSSAppAlertsRequest**](PostRedfishV1SFSSAppAlertsRequest.md) | {     \&quot;Protocol\&quot;: \&quot;redfish\&quot;,     \&quot;Context\&quot;: \&quot;SomeSubscription\&quot;,     \&quot;EventTypes\&quot;: [          \&quot;Alert\&quot;      ],     \&quot;CdcInstances\&quot;: [         \&quot;APP\&quot;     ],     \&quot;HttpHeaders\&quot;: [         \&quot;Authorization: Basic ZG52dXNlcjpAIThwSU1vSQ&#x3D;&#x3D;\&quot;,         \&quot;ExternalServerRequiredHeader: ItsValue\&quot;     ],     \&quot;Destination\&quot;: \&quot;https://[ipv4/ipv6]/external/Server/eventHandler\&quot;  } | 

### Return type

[**PostRedfishV1SFSSAppAlerts200Response**](PostRedfishV1SFSSAppAlerts200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppAuthenticationSequence

> PostRedfishV1SFSSAppAuthenticationSequence200Response PostRedfishV1SFSSAppAuthenticationSequence(ctx).PostRedfishV1SFSSAppAuthenticationSequenceRequest(postRedfishV1SFSSAppAuthenticationSequenceRequest).Execute()

Add authentication sequence



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppAuthenticationSequenceRequest := *openapiclient.NewPostRedfishV1SFSSAppAuthenticationSequenceRequest([]string{"AuthenticationSequence_example"}) // PostRedfishV1SFSSAppAuthenticationSequenceRequest | {  \"AuthenticationSequence\": [         \"tacacs+\",         \"local\"     ]  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppAuthenticationSequence(context.Background()).PostRedfishV1SFSSAppAuthenticationSequenceRequest(postRedfishV1SFSSAppAuthenticationSequenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppAuthenticationSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppAuthenticationSequence`: PostRedfishV1SFSSAppAuthenticationSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppAuthenticationSequence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppAuthenticationSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppAuthenticationSequenceRequest** | [**PostRedfishV1SFSSAppAuthenticationSequenceRequest**](PostRedfishV1SFSSAppAuthenticationSequenceRequest.md) | {  \&quot;AuthenticationSequence\&quot;: [         \&quot;tacacs+\&quot;,         \&quot;local\&quot;     ]  } | 

### Return type

[**PostRedfishV1SFSSAppAuthenticationSequence200Response**](PostRedfishV1SFSSAppAuthenticationSequence200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppBackups

> PostRedfishV1SFSSAppBackups200Response PostRedfishV1SFSSAppBackups(ctx).PostRedfishV1SFSSAppBackupsRequest(postRedfishV1SFSSAppBackupsRequest).Execute()

Perform backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppBackupsRequest := *openapiclient.NewPostRedfishV1SFSSAppBackupsRequest("ImageServerLocation_example", "ImageServerPassword_example", "TransportType_example", "ImageServerUserName_example") // PostRedfishV1SFSSAppBackupsRequest |         {             \"ImageServerLocation\": \"100.94.72.166:/home/dell/temp_images/\",             \"ImageServerPassword\": \"Force10\",             \"TransportType\": \"SCP\",             \"ImageServerUserName\": \"Dell\"         } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppBackups(context.Background()).PostRedfishV1SFSSAppBackupsRequest(postRedfishV1SFSSAppBackupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppBackups`: PostRedfishV1SFSSAppBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppBackupsRequest** | [**PostRedfishV1SFSSAppBackupsRequest**](PostRedfishV1SFSSAppBackupsRequest.md) |         {             \&quot;ImageServerLocation\&quot;: \&quot;100.94.72.166:/home/dell/temp_images/\&quot;,             \&quot;ImageServerPassword\&quot;: \&quot;Force10\&quot;,             \&quot;TransportType\&quot;: \&quot;SCP\&quot;,             \&quot;ImageServerUserName\&quot;: \&quot;Dell\&quot;         } | 

### Return type

[**PostRedfishV1SFSSAppBackups200Response**](PostRedfishV1SFSSAppBackups200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppChangePassword

> PostRedfishV1SFSSAppChangePassword(ctx).PostRedfishV1SFSSAppChangePasswordRequest(postRedfishV1SFSSAppChangePasswordRequest).Execute()

Change admin password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppChangePasswordRequest := *openapiclient.NewPostRedfishV1SFSSAppChangePasswordRequest("OldPassword_example", "NewPassword_example") // PostRedfishV1SFSSAppChangePasswordRequest | { \"OldPassword\":\"admin\", \"NewPassword\" : \"Xskfdj@fdk10\" } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppChangePassword(context.Background()).PostRedfishV1SFSSAppChangePasswordRequest(postRedfishV1SFSSAppChangePasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppChangePasswordRequest** | [**PostRedfishV1SFSSAppChangePasswordRequest**](PostRedfishV1SFSSAppChangePasswordRequest.md) | { \&quot;OldPassword\&quot;:\&quot;admin\&quot;, \&quot;NewPassword\&quot; : \&quot;Xskfdj@fdk10\&quot; } | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppFabricManagerInfoPost

> PostRedfishV1SFSSAppFabricManagerInfoPost200Response PostRedfishV1SFSSAppFabricManagerInfoPost(ctx).PostRedfishV1SFSSAppFabricManagerInfoPostRequest(postRedfishV1SFSSAppFabricManagerInfoPostRequest).Execute()

Configure CDC instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppFabricManagerInfoPostRequest := *openapiclient.NewPostRedfishV1SFSSAppFabricManagerInfoPostRequest("InstanceIdentifier_example", []string{"Interfaces_example"}, "CDCAdminState_example", "DiscoverySvcAdminState_example") // PostRedfishV1SFSSAppFabricManagerInfoPostRequest | {     \"InstanceIdentifier\": \"1\",     \"Interfaces\": [\"ens192\"],     \"CDCAdminState\":\"Enable\",     \"DiscoverySvcAdminState\":\"Enable\" } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppFabricManagerInfoPost(context.Background()).PostRedfishV1SFSSAppFabricManagerInfoPostRequest(postRedfishV1SFSSAppFabricManagerInfoPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppFabricManagerInfoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppFabricManagerInfoPost`: PostRedfishV1SFSSAppFabricManagerInfoPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppFabricManagerInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppFabricManagerInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppFabricManagerInfoPostRequest** | [**PostRedfishV1SFSSAppFabricManagerInfoPostRequest**](PostRedfishV1SFSSAppFabricManagerInfoPostRequest.md) | {     \&quot;InstanceIdentifier\&quot;: \&quot;1\&quot;,     \&quot;Interfaces\&quot;: [\&quot;ens192\&quot;],     \&quot;CDCAdminState\&quot;:\&quot;Enable\&quot;,     \&quot;DiscoverySvcAdminState\&quot;:\&quot;Enable\&quot; } | 

### Return type

[**PostRedfishV1SFSSAppFabricManagerInfoPost200Response**](PostRedfishV1SFSSAppFabricManagerInfoPost200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppGlobalSettings

> PostRedfishV1SFSSAppGlobalSettings(ctx).PostRedfishV1SFSSAppGlobalSettingsRequest(postRedfishV1SFSSAppGlobalSettingsRequest).Execute()

Configure global settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppGlobalSettingsRequest := *openapiclient.NewPostRedfishV1SFSSAppGlobalSettingsRequest("HostName_example", "ReservedIPV4SubnetPrefix_example", "ReservedIPV6SubnetPrefix_example", float32(123)) // PostRedfishV1SFSSAppGlobalSettingsRequest | {     \"HostName\": \"dellemc-networkappliance\",     \"ReservedIPV4SubnetPrefix\": \"172.20\",     \"ReservedIPV6SubnetPrefix\": \"fd02\",     \"StorageInterfaceMTU\": 7000  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppGlobalSettings(context.Background()).PostRedfishV1SFSSAppGlobalSettingsRequest(postRedfishV1SFSSAppGlobalSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppGlobalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppGlobalSettingsRequest** | [**PostRedfishV1SFSSAppGlobalSettingsRequest**](PostRedfishV1SFSSAppGlobalSettingsRequest.md) | {     \&quot;HostName\&quot;: \&quot;dellemc-networkappliance\&quot;,     \&quot;ReservedIPV4SubnetPrefix\&quot;: \&quot;172.20\&quot;,     \&quot;ReservedIPV6SubnetPrefix\&quot;: \&quot;fd02\&quot;,     \&quot;StorageInterfaceMTU\&quot;: 7000  } | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppIpAddressManagements

> PostRedfishV1SFSSAppIpAddressManagements(ctx).PostRedfishV1SFSSAppIpAddressManagementsRequest(postRedfishV1SFSSAppIpAddressManagementsRequest).Execute()

Configure interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppIpAddressManagementsRequest := *openapiclient.NewPostRedfishV1SFSSAppIpAddressManagementsRequest([]string{"IPV4Address_example"}, "IPV4Config_example", "IPV4Gateway_example", float32(123), "IPV6Config_example", []string{"IPV6Address_example"}, "IPV6Gateway_example", float32(123), "ParentInterface_example", float32(123)) // PostRedfishV1SFSSAppIpAddressManagementsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppIpAddressManagements(context.Background()).PostRedfishV1SFSSAppIpAddressManagementsRequest(postRedfishV1SFSSAppIpAddressManagementsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppIpAddressManagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppIpAddressManagementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppIpAddressManagementsRequest** | [**PostRedfishV1SFSSAppIpAddressManagementsRequest**](PostRedfishV1SFSSAppIpAddressManagementsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppLicenses

> PostRedfishV1SFSSAppLicenses200Response PostRedfishV1SFSSAppLicenses(ctx).PostRedfishV1SFSSAppLicensesRequest(postRedfishV1SFSSAppLicensesRequest).Execute()

Install a license



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppLicensesRequest := *openapiclient.NewPostRedfishV1SFSSAppLicensesRequest("LicenseContent_example", "LicenseFileName_example") // PostRedfishV1SFSSAppLicensesRequest | {    \"LicenseContent\" :\"Actual XML-LicenseContent\",     \"LicenseFileName\": \"BaseLicense.lic\"  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppLicenses(context.Background()).PostRedfishV1SFSSAppLicensesRequest(postRedfishV1SFSSAppLicensesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppLicenses`: PostRedfishV1SFSSAppLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppLicensesRequest** | [**PostRedfishV1SFSSAppLicensesRequest**](PostRedfishV1SFSSAppLicensesRequest.md) | {    \&quot;LicenseContent\&quot; :\&quot;Actual XML-LicenseContent\&quot;,     \&quot;LicenseFileName\&quot;: \&quot;BaseLicense.lic\&quot;  } | 

### Return type

[**PostRedfishV1SFSSAppLicenses200Response**](PostRedfishV1SFSSAppLicenses200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppNTP1

> PostRedfishV1SFSSAppNTP1200Response PostRedfishV1SFSSAppNTP1(ctx).PostRedfishV1SFSSAppNTP1Request(postRedfishV1SFSSAppNTP1Request).Execute()

Configure NTP server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppNTP1Request := *openapiclient.NewPostRedfishV1SFSSAppNTP1Request() // PostRedfishV1SFSSAppNTP1Request |   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppNTP1(context.Background()).PostRedfishV1SFSSAppNTP1Request(postRedfishV1SFSSAppNTP1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppNTP1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppNTP1`: PostRedfishV1SFSSAppNTP1200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppNTP1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppNTP1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppNTP1Request** | [**PostRedfishV1SFSSAppNTP1Request**](PostRedfishV1SFSSAppNTP1Request.md) |   | 

### Return type

[**PostRedfishV1SFSSAppNTP1200Response**](PostRedfishV1SFSSAppNTP1200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppRadiusServers

> PostRedfishV1SFSSAppRadiusServers200Response PostRedfishV1SFSSAppRadiusServers(ctx).PostRedfishV1SFSSAppRadiusServersRequest(postRedfishV1SFSSAppRadiusServersRequest).Execute()

Configure RADIUS server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppRadiusServersRequest := *openapiclient.NewPostRedfishV1SFSSAppRadiusServersRequest("ServerIp_example", "ServerPass_example") // PostRedfishV1SFSSAppRadiusServersRequest | {  \"ServerIp\": \"200.1.1.1\", \"ServerPass\": \"xxxxxx\"  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppRadiusServers(context.Background()).PostRedfishV1SFSSAppRadiusServersRequest(postRedfishV1SFSSAppRadiusServersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppRadiusServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppRadiusServers`: PostRedfishV1SFSSAppRadiusServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppRadiusServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppRadiusServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppRadiusServersRequest** | [**PostRedfishV1SFSSAppRadiusServersRequest**](PostRedfishV1SFSSAppRadiusServersRequest.md) | {  \&quot;ServerIp\&quot;: \&quot;200.1.1.1\&quot;, \&quot;ServerPass\&quot;: \&quot;xxxxxx\&quot;  } | 

### Return type

[**PostRedfishV1SFSSAppRadiusServers200Response**](PostRedfishV1SFSSAppRadiusServers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppRestores

> PostRedfishV1SFSSAppRestores200Response PostRedfishV1SFSSAppRestores(ctx).PostRedfishV1SFSSAppRestoresRequest(postRedfishV1SFSSAppRestoresRequest).Execute()

Restore a backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppRestoresRequest := *openapiclient.NewPostRedfishV1SFSSAppRestoresRequest("ImageServerLocation_example", "ImageServerPassword_example", "TransportType_example", "ImageServerUserName_example") // PostRedfishV1SFSSAppRestoresRequest |  {             \"ImageServerLocation\": \"100.94.72.166:/home/dell/temp_images/backup_file.tar.gz\",             \"ImageServerPassword\": \"Force10\",             \"TransportType\": \"SCP\",             \"ImageServerUserName\": \"Dell\"         } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppRestores(context.Background()).PostRedfishV1SFSSAppRestoresRequest(postRedfishV1SFSSAppRestoresRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppRestores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppRestores`: PostRedfishV1SFSSAppRestores200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppRestores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppRestoresRequest** | [**PostRedfishV1SFSSAppRestoresRequest**](PostRedfishV1SFSSAppRestoresRequest.md) |  {             \&quot;ImageServerLocation\&quot;: \&quot;100.94.72.166:/home/dell/temp_images/backup_file.tar.gz\&quot;,             \&quot;ImageServerPassword\&quot;: \&quot;Force10\&quot;,             \&quot;TransportType\&quot;: \&quot;SCP\&quot;,             \&quot;ImageServerUserName\&quot;: \&quot;Dell\&quot;         } | 

### Return type

[**PostRedfishV1SFSSAppRestores200Response**](PostRedfishV1SFSSAppRestores200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppSFSSImages

> PutRedfishV1SFSSApp200Response PostRedfishV1SFSSAppSFSSImages(ctx).PostRedfishV1SFSSAppSFSSImagesRequest(postRedfishV1SFSSAppSFSSImagesRequest).Execute()

Add image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppSFSSImagesRequest := *openapiclient.NewPostRedfishV1SFSSAppSFSSImagesRequest("ImageServerUserName_example", "ImageServerPassword_example", "ImageServerLocation_example", "TransportType_example") // PostRedfishV1SFSSAppSFSSImagesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppSFSSImages(context.Background()).PostRedfishV1SFSSAppSFSSImagesRequest(postRedfishV1SFSSAppSFSSImagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppSFSSImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppSFSSImages`: PutRedfishV1SFSSApp200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppSFSSImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppSFSSImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppSFSSImagesRequest** | [**PostRedfishV1SFSSAppSFSSImagesRequest**](PostRedfishV1SFSSAppSFSSImagesRequest.md) |  | 

### Return type

[**PutRedfishV1SFSSApp200Response**](PutRedfishV1SFSSApp200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSAppTacacsServers

> PostRedfishV1SFSSAppTacacsServers200Response PostRedfishV1SFSSAppTacacsServers(ctx).PostRedfishV1SFSSAppTacacsServersRequest(postRedfishV1SFSSAppTacacsServersRequest).Execute()

Configure TACACS+ server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	postRedfishV1SFSSAppTacacsServersRequest := *openapiclient.NewPostRedfishV1SFSSAppTacacsServersRequest("ServerIp_example", "ServerPass_example") // PostRedfishV1SFSSAppTacacsServersRequest | {  \"ServerIp\": \"200.1.1.1\", \"ServerPass\": \"xxxxxx\"  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSAppTacacsServers(context.Background()).PostRedfishV1SFSSAppTacacsServersRequest(postRedfishV1SFSSAppTacacsServersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSAppTacacsServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSAppTacacsServers`: PostRedfishV1SFSSAppTacacsServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSAppTacacsServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSAppTacacsServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedfishV1SFSSAppTacacsServersRequest** | [**PostRedfishV1SFSSAppTacacsServersRequest**](PostRedfishV1SFSSAppTacacsServersRequest.md) | {  \&quot;ServerIp\&quot;: \&quot;200.1.1.1\&quot;, \&quot;ServerPass\&quot;: \&quot;xxxxxx\&quot;  } | 

### Return type

[**PostRedfishV1SFSSAppTacacsServers200Response**](PostRedfishV1SFSSAppTacacsServers200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSApp

> PutRedfishV1SFSSAppRequest PutRedfishV1SFSSApp(ctx).PutRedfishV1SFSSAppRequest(putRedfishV1SFSSAppRequest).Execute()

Upgrade SFSS application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	putRedfishV1SFSSAppRequest := *openapiclient.NewPutRedfishV1SFSSAppRequest("ImageServerUserName_example", "ImageServerPassword_example", "ImageServerLocation_example", "TransportType_example") // PutRedfishV1SFSSAppRequest | {     \"Version\": \"1.2.0\" } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSApp(context.Background()).PutRedfishV1SFSSAppRequest(putRedfishV1SFSSAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSApp`: PutRedfishV1SFSSAppRequest
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putRedfishV1SFSSAppRequest** | [**PutRedfishV1SFSSAppRequest**](PutRedfishV1SFSSAppRequest.md) | {     \&quot;Version\&quot;: \&quot;1.2.0\&quot; } | 

### Return type

[**PutRedfishV1SFSSAppRequest**](PutRedfishV1SFSSAppRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSAppAlerts

> PostRedfishV1SFSSAppAlerts200Response PutRedfishV1SFSSAppAlerts(ctx, uuid).PutRedfishV1SFSSAppAlertsRequest(putRedfishV1SFSSAppAlertsRequest).Execute()

Update alert



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert UUID
	putRedfishV1SFSSAppAlertsRequest := *openapiclient.NewPutRedfishV1SFSSAppAlertsRequest("Protocol_example", "Identifier_example", "Context_example", []interface{}{nil}, []interface{}{nil}, []interface{}{nil}, "Destination_example") // PutRedfishV1SFSSAppAlertsRequest | {     \"Protocol\": \"redfish\",      \"Identifier\": \"uuid\",      \"Context\": \"NewSubscription\",     \"EventTypes\": [          \"Alert\"      ],    \"CdcInstances\": [         \"APP\"     ],      \"HttpHeaders\": [         \"Authorization: Basic ZZZZZZZZZZZZZZZZZZZZZZZZZZ\",         \"ExternalServerRequiredHeader: ItsValue\"     ],     \"Destination\": \"https://[ipv4/ipv6]/external/Server/NeweventHandler\"  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSAppAlerts(context.Background(), uuid).PutRedfishV1SFSSAppAlertsRequest(putRedfishV1SFSSAppAlertsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSAppAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSAppAlerts`: PostRedfishV1SFSSAppAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSAppAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Alert UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSAppAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRedfishV1SFSSAppAlertsRequest** | [**PutRedfishV1SFSSAppAlertsRequest**](PutRedfishV1SFSSAppAlertsRequest.md) | {     \&quot;Protocol\&quot;: \&quot;redfish\&quot;,      \&quot;Identifier\&quot;: \&quot;uuid\&quot;,      \&quot;Context\&quot;: \&quot;NewSubscription\&quot;,     \&quot;EventTypes\&quot;: [          \&quot;Alert\&quot;      ],    \&quot;CdcInstances\&quot;: [         \&quot;APP\&quot;     ],      \&quot;HttpHeaders\&quot;: [         \&quot;Authorization: Basic ZZZZZZZZZZZZZZZZZZZZZZZZZZ\&quot;,         \&quot;ExternalServerRequiredHeader: ItsValue\&quot;     ],     \&quot;Destination\&quot;: \&quot;https://[ipv4/ipv6]/external/Server/NeweventHandler\&quot;  } | 

### Return type

[**PostRedfishV1SFSSAppAlerts200Response**](PostRedfishV1SFSSAppAlerts200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSAppCDCInstanceManagers

> PostRedfishV1SFSSAppFabricManagerInfoPost200Response PutRedfishV1SFSSAppCDCInstanceManagers(ctx, instanceId).PutRedfishV1SFSSAppCDCInstanceManagersRequest(putRedfishV1SFSSAppCDCInstanceManagersRequest).Execute()

Update CDC instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | InstanceIdentifier
	putRedfishV1SFSSAppCDCInstanceManagersRequest := *openapiclient.NewPutRedfishV1SFSSAppCDCInstanceManagersRequest("InstanceIdentifier_example", []string{"Interfaces_example"}, "CDCAdminState_example", "DiscoverySvcAdminState_example") // PutRedfishV1SFSSAppCDCInstanceManagersRequest | {     \"InstanceIdentifier\": \"1\",     \"Interfaces\": [\"ens192\"],     \"CDCAdminState\":\"Enable\",     \"DiscoverySvcAdminState\":\"Disable\" } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSAppCDCInstanceManagers(context.Background(), instanceId).PutRedfishV1SFSSAppCDCInstanceManagersRequest(putRedfishV1SFSSAppCDCInstanceManagersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSAppCDCInstanceManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSAppCDCInstanceManagers`: PostRedfishV1SFSSAppFabricManagerInfoPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSAppCDCInstanceManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | InstanceIdentifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSAppCDCInstanceManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRedfishV1SFSSAppCDCInstanceManagersRequest** | [**PutRedfishV1SFSSAppCDCInstanceManagersRequest**](PutRedfishV1SFSSAppCDCInstanceManagersRequest.md) | {     \&quot;InstanceIdentifier\&quot;: \&quot;1\&quot;,     \&quot;Interfaces\&quot;: [\&quot;ens192\&quot;],     \&quot;CDCAdminState\&quot;:\&quot;Enable\&quot;,     \&quot;DiscoverySvcAdminState\&quot;:\&quot;Disable\&quot; } | 

### Return type

[**PostRedfishV1SFSSAppFabricManagerInfoPost200Response**](PostRedfishV1SFSSAppFabricManagerInfoPost200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSAppIpAddressManagements

> PutRedfishV1SFSSAppIpAddressManagements(ctx, interfaceId).PutRedfishV1SFSSAppIpAddressManagementsRequest(putRedfishV1SFSSAppIpAddressManagementsRequest).Execute()

Update interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	interfaceId := "interfaceId_example" // string | Interface Name
	putRedfishV1SFSSAppIpAddressManagementsRequest := *openapiclient.NewPutRedfishV1SFSSAppIpAddressManagementsRequest([]string{"IPV4Address_example"}, "IPV4Config_example", "IPV4Gateway_example", float32(123), "IPV6Config_example", []string{"IPV6Address_example"}, "IPV6Gateway_example", float32(123)) // PutRedfishV1SFSSAppIpAddressManagementsRequest | {      \"IPV4Address\": [         \"30.1.1.1\"                                           ],     \"IPV4Config\": \"MANUAL\",                    \"IPV4Gateway\": \"30.1.1.2\",     \"IPV4PrefixLength\": 16,     \"IPV6Config\": \"MANUAL\",      \"Name\": \"Name1\",                     \"IPV6Address\": [           \"fe80::1699:6fff:43dd:56c1\"               ],      \"IPV6Gateway\":   \"fe80::1699:6f09:43dd:ffff\",     \"IPV6PrefixLength\": 64,      \"MTU\": 7000 ## If this field is not present, the MTU is chosen as auto  } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PutRedfishV1SFSSAppIpAddressManagements(context.Background(), interfaceId).PutRedfishV1SFSSAppIpAddressManagementsRequest(putRedfishV1SFSSAppIpAddressManagementsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSAppIpAddressManagements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | Interface Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSAppIpAddressManagementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRedfishV1SFSSAppIpAddressManagementsRequest** | [**PutRedfishV1SFSSAppIpAddressManagementsRequest**](PutRedfishV1SFSSAppIpAddressManagementsRequest.md) | {      \&quot;IPV4Address\&quot;: [         \&quot;30.1.1.1\&quot;                                           ],     \&quot;IPV4Config\&quot;: \&quot;MANUAL\&quot;,                    \&quot;IPV4Gateway\&quot;: \&quot;30.1.1.2\&quot;,     \&quot;IPV4PrefixLength\&quot;: 16,     \&quot;IPV6Config\&quot;: \&quot;MANUAL\&quot;,      \&quot;Name\&quot;: \&quot;Name1\&quot;,                     \&quot;IPV6Address\&quot;: [           \&quot;fe80::1699:6fff:43dd:56c1\&quot;               ],      \&quot;IPV6Gateway\&quot;:   \&quot;fe80::1699:6f09:43dd:ffff\&quot;,     \&quot;IPV6PrefixLength\&quot;: 64,      \&quot;MTU\&quot;: 7000 ## If this field is not present, the MTU is chosen as auto  } | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/xml, multipart/form-data, text/html
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSAppLicenses

> PutRedfishV1SFSSAppLicenses200Response PutRedfishV1SFSSAppLicenses(ctx).PutRedfishV1SFSSAppLicensesRequest(putRedfishV1SFSSAppLicensesRequest).Execute()

Accept EULA



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	putRedfishV1SFSSAppLicensesRequest := *openapiclient.NewPutRedfishV1SFSSAppLicensesRequest("Identifier_example", "EULA_example") // PutRedfishV1SFSSAppLicensesRequest | { \"Identifier\": \"2\", \"EULA\": \"Agreed\" } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSAppLicenses(context.Background()).PutRedfishV1SFSSAppLicensesRequest(putRedfishV1SFSSAppLicensesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSAppLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSAppLicenses`: PutRedfishV1SFSSAppLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSAppLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSAppLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putRedfishV1SFSSAppLicensesRequest** | [**PutRedfishV1SFSSAppLicensesRequest**](PutRedfishV1SFSSAppLicensesRequest.md) | { \&quot;Identifier\&quot;: \&quot;2\&quot;, \&quot;EULA\&quot;: \&quot;Agreed\&quot; } | 

### Return type

[**PutRedfishV1SFSSAppLicenses200Response**](PutRedfishV1SFSSAppLicenses200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSAppNTP

> PutRedfishV1SFSSAppNTP200Response PutRedfishV1SFSSAppNTP(ctx).PutRedfishV1SFSSAppNTPRequest(putRedfishV1SFSSAppNTPRequest).Execute()

Enable or disable NTP service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	putRedfishV1SFSSAppNTPRequest := *openapiclient.NewPutRedfishV1SFSSAppNTPRequest() // PutRedfishV1SFSSAppNTPRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSAppNTP(context.Background()).PutRedfishV1SFSSAppNTPRequest(putRedfishV1SFSSAppNTPRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSAppNTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSAppNTP`: PutRedfishV1SFSSAppNTP200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSAppNTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSAppNTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putRedfishV1SFSSAppNTPRequest** | [**PutRedfishV1SFSSAppNTPRequest**](PutRedfishV1SFSSAppNTPRequest.md) |  | 

### Return type

[**PutRedfishV1SFSSAppNTP200Response**](PutRedfishV1SFSSAppNTP200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSApp_0

> PutRedfishV1SFSSApp200Response PutRedfishV1SFSSApp_0(ctx).Versiontbd(versiontbd).PutRedfishV1SFSSAppRequest(putRedfishV1SFSSAppRequest).Execute()

Upgrade application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	versiontbd := "versiontbd_example" // string | ImageId  (optional)
	putRedfishV1SFSSAppRequest := *openapiclient.NewPutRedfishV1SFSSAppRequest("ImageServerUserName_example", "ImageServerPassword_example", "ImageServerLocation_example", "TransportType_example") // PutRedfishV1SFSSAppRequest | {     \"ImageServerUserName\" : \"dell\",     \"ImageServerPassword\" : \"New_Password\",     \"ImageServerLocation\" : \"100.94.72.166:/home/dell/new_location/SFSS-1.2.0.deb\",     \"TransportType\" : \"SCP\" } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSApp_0(context.Background()).Versiontbd(versiontbd).PutRedfishV1SFSSAppRequest(putRedfishV1SFSSAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSApp_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSApp_0`: PutRedfishV1SFSSApp200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSApp_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSApp_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versiontbd** | **string** | ImageId  | 
 **putRedfishV1SFSSAppRequest** | [**PutRedfishV1SFSSAppRequest**](PutRedfishV1SFSSAppRequest.md) | {     \&quot;ImageServerUserName\&quot; : \&quot;dell\&quot;,     \&quot;ImageServerPassword\&quot; : \&quot;New_Password\&quot;,     \&quot;ImageServerLocation\&quot; : \&quot;100.94.72.166:/home/dell/new_location/SFSS-1.2.0.deb\&quot;,     \&quot;TransportType\&quot; : \&quot;SCP\&quot; } | 

### Return type

[**PutRedfishV1SFSSApp200Response**](PutRedfishV1SFSSApp200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

