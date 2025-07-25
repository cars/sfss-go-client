# Go API client for sfss

REST APIs used for managing SFSS application are captured in this section.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.14.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import sfss "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sfss.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sfss.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sfss.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sfss.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sfss.ContextOperationServerIndices` and `sfss.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), sfss.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sfss.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://IPAddress*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers**](docs/DefaultAPI.md#deleteredfishv1sfssappalertscdcinstancemanagers) | **Delete** /redfish/v1/SFSSApp/Alerts({uuid}) | Delete alert
*DefaultAPI* | [**DeleteRedfishV1SFSSAppCDCInstanceManagers**](docs/DefaultAPI.md#deleteredfishv1sfssappcdcinstancemanagers) | **Delete** /redfish/v1/SFSSApp/CDCInstanceManagers({InstanceId}) | Delete CDC instance
*DefaultAPI* | [**DeleteRedfishV1SFSSAppIpAddressManagements**](docs/DefaultAPI.md#deleteredfishv1sfssappipaddressmanagements) | **Delete** /redfish/v1/SFSSApp/IpAddressManagements({InterfaceId}) | Delete VLAN interface
*DefaultAPI* | [**DeleteRedfishV1SFSSAppNTPServerIP**](docs/DefaultAPI.md#deleteredfishv1sfssappntpserverip) | **Delete** /redfish/v1/SFSSApp/NTP({ServerIP}) | Remove NTP configuration
*DefaultAPI* | [**DeleteRedfishV1SFSSAppSFSSImages**](docs/DefaultAPI.md#deleteredfishv1sfssappsfssimages) | **Delete** /redfish/v1/SFSSApp/SFSSImages | Delete image
*DefaultAPI* | [**GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums**](docs/DefaultAPI.md#getenumsredfishv1sfssappipaddressmanagementsenums) | **Get** /redfish/v1/SFSSApp/IpAddressManagements/Enums | Get interface enums
*DefaultAPI* | [**GetExpandRedfishV1SFSSAppBackups**](docs/DefaultAPI.md#getexpandredfishv1sfssappbackups) | **Get** /redfish/v1/SFSSApp/Backups?$expand&#x3D;Backups | Get all backups
*DefaultAPI* | [**GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand**](docs/DefaultAPI.md#getexpandredfishv1sfssappfoundationalconfigsgetexpand) | **Get** /redfish/v1/SFSSApp/FoundationalConfigs?$expand&#x3D;FoundationalConfigs | Get all foundational configuration
*DefaultAPI* | [**GetExpandRedfishV1SFSSAppIpAddressManagements**](docs/DefaultAPI.md#getexpandredfishv1sfssappipaddressmanagements) | **Get** /redfish/v1/SFSSApp/IpAddressManagements?$expand&#x3D;IpAddressManagements | Get all interfaces
*DefaultAPI* | [**GetExpandRedfishV1SFSSAppRestores**](docs/DefaultAPI.md#getexpandredfishv1sfssapprestores) | **Get** /redfish/v1/SFSSApp/Restores?$expand&#x3D;Restores | Get detailed restore information
*DefaultAPI* | [**GetIDRedfishV1SFSSApp**](docs/DefaultAPI.md#getidredfishv1sfssapp) | **Get** /redfish/v1/SFSSApp/SFSSImages({version}) | Get specific image
*DefaultAPI* | [**GetIDRedfishV1SFSSAppRadiusServers**](docs/DefaultAPI.md#getidredfishv1sfssappradiusservers) | **Get** /redfish/v1/SFSSApp/RadiusServers({IP}) | Get specific RADIUS server
*DefaultAPI* | [**GetIDRedfishV1SFSSAppTacacsServers**](docs/DefaultAPI.md#getidredfishv1sfssapptacacsservers) | **Get** /redfish/v1/SFSSApp/TacacsServers({IP}) | Get specific TACACS+ server
*DefaultAPI* | [**GetIDredfishV1SFSSAppSFSSInterfaceList**](docs/DefaultAPI.md#getidredfishv1sfssappsfssinterfacelist) | **Get** /redfish/v1/SFSSApp/SFSSInterfaceList | Get all interfaces
*DefaultAPI* | [**GetRedfishV1SFSSApp**](docs/DefaultAPI.md#getredfishv1sfssapp) | **Get** /redfish/v1/SFSSApp | Get SFSS application details
*DefaultAPI* | [**GetRedfishV1SFSSAppAlertsAlertId**](docs/DefaultAPI.md#getredfishv1sfssappalertsalertid) | **Get** /redfish/v1/SFSSApp/Alerts({uuid}) | Get specific alert
*DefaultAPI* | [**GetRedfishV1SFSSAppAuthenticationSequence**](docs/DefaultAPI.md#getredfishv1sfssappauthenticationsequence) | **Get** /redfish/v1/SFSSApp/AuthenticationSequence | Get authentication sequence
*DefaultAPI* | [**GetRedfishV1SFSSAppAuthenticationSequenceEnums**](docs/DefaultAPI.md#getredfishv1sfssappauthenticationsequenceenums) | **Get** /redfish/v1/SFSSApp/AuthenticationSequence/Enums | Get authentication sequence enums
*DefaultAPI* | [**GetRedfishV1SFSSAppBackups**](docs/DefaultAPI.md#getredfishv1sfssappbackups) | **Get** /redfish/v1/SFSSApp/Backups | Get all backups
*DefaultAPI* | [**GetRedfishV1SFSSAppBackupsID**](docs/DefaultAPI.md#getredfishv1sfssappbackupsid) | **Get** /redfish/v1/SFSSApp/Backups({ID}) | Get specific backup
*DefaultAPI* | [**GetRedfishV1SFSSAppCDCHealthStatus**](docs/DefaultAPI.md#getredfishv1sfssappcdchealthstatus) | **Get** /redfish/v1/SFSSApp/CDCHealthStatus | Get CDC health
*DefaultAPI* | [**GetRedfishV1SFSSAppCDCHealthStatusID**](docs/DefaultAPI.md#getredfishv1sfssappcdchealthstatusid) | **Get** /redfish/v1/SFSSApp/CDCHealthStatus({ID}) | Get specific CDC health
*DefaultAPI* | [**GetRedfishV1SFSSAppCDCInstanceManagers**](docs/DefaultAPI.md#getredfishv1sfssappcdcinstancemanagers) | **Get** /redfish/v1/SFSSApp/CDCInstanceManagers | Get all CDCs
*DefaultAPI* | [**GetRedfishV1SFSSAppCDCInstanceManagersEnums**](docs/DefaultAPI.md#getredfishv1sfssappcdcinstancemanagersenums) | **Get** /redfish/v1/SFSSApp/CDCInstanceManagers/Enums | Get CDC enums
*DefaultAPI* | [**GetRedfishV1SFSSAppCDCInstanceManagersID**](docs/DefaultAPI.md#getredfishv1sfssappcdcinstancemanagersid) | **Get** /redfish/v1/SFSSApp/CDCInstanceManagers({InstanceId}) | Get specific CDC information
*DefaultAPI* | [**GetRedfishV1SFSSAppDevice**](docs/DefaultAPI.md#getredfishv1sfssappdevice) | **Get** /redfish/v1/SFSSApp/Device | Get device details
*DefaultAPI* | [**GetRedfishV1SFSSAppEnums**](docs/DefaultAPI.md#getredfishv1sfssappenums) | **Get** /redfish/v1/SFSSApp/SFSSImages/Enums | Get image enums
*DefaultAPI* | [**GetRedfishV1SFSSAppEvents**](docs/DefaultAPI.md#getredfishv1sfssappevents) | **Get** /redfish/v1/SFSSApp/Events | Get all events
*DefaultAPI* | [**GetRedfishV1SFSSAppEventsID**](docs/DefaultAPI.md#getredfishv1sfssappeventsid) | **Get** /redfish/v1/SFSSApp/Events({ID}) | Get specific event
*DefaultAPI* | [**GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier**](docs/DefaultAPI.md#getredfishv1sfssappfoundationalconfigsinstanceidentifier) | **Get** /redfish/v1/SFSSApp/FoundationalConfigs({InstanceIdentifer}) | Get specific foundational configuration
*DefaultAPI* | [**GetRedfishV1SFSSAppGlobalSettings**](docs/DefaultAPI.md#getredfishv1sfssappglobalsettings) | **Get** /redfish/v1/SFSSApp/GlobalSettings | Get global settings
*DefaultAPI* | [**GetRedfishV1SFSSAppIpAddressManagements**](docs/DefaultAPI.md#getredfishv1sfssappipaddressmanagements) | **Get** /redfish/v1/SFSSApp/IpAddressManagements | Get all interfaces
*DefaultAPI* | [**GetRedfishV1SFSSAppIpAddressManagementsInterface**](docs/DefaultAPI.md#getredfishv1sfssappipaddressmanagementsinterface) | **Get** /redfish/v1/SFSSApp/IpAddressManagements({interface}) | Get specific interface
*DefaultAPI* | [**GetRedfishV1SFSSAppLicenses**](docs/DefaultAPI.md#getredfishv1sfssapplicenses) | **Get** /redfish/v1/SFSSApp/Licenses | Get license count
*DefaultAPI* | [**GetRedfishV1SFSSAppLicensesLicenseId**](docs/DefaultAPI.md#getredfishv1sfssapplicenseslicenseid) | **Get** /redfish/v1/SFSSApp/Licenses({LicenseId}) | Get specific license
*DefaultAPI* | [**GetRedfishV1SFSSAppLicensesexpandLicenses**](docs/DefaultAPI.md#getredfishv1sfssapplicensesexpandlicenses) | **Get** /redfish/v1/SFSSApp/Licenses?$expand&#x3D;Licenses | Get detailed license information
*DefaultAPI* | [**GetRedfishV1SFSSAppNTP**](docs/DefaultAPI.md#getredfishv1sfssappntp) | **Get** /redfish/v1/SFSSApp/NTP | Get NTP server information
*DefaultAPI* | [**GetRedfishV1SFSSAppRadiusServers**](docs/DefaultAPI.md#getredfishv1sfssappradiusservers) | **Get** /redfish/v1/SFSSApp/RadiusServers | Get all RADIUS servers
*DefaultAPI* | [**GetRedfishV1SFSSAppRadiusServersSequence**](docs/DefaultAPI.md#getredfishv1sfssappradiusserverssequence) | **Get** /redfish/v1/SFSSApp/RadiusServers/Sequence | Get all RADIUS servers
*DefaultAPI* | [**GetRedfishV1SFSSAppRestores**](docs/DefaultAPI.md#getredfishv1sfssapprestores) | **Get** /redfish/v1/SFSSApp/Restores | Get all restores
*DefaultAPI* | [**GetRedfishV1SFSSAppRestoresID**](docs/DefaultAPI.md#getredfishv1sfssapprestoresid) | **Get** /redfish/v1/SFSSApp/Restores({ID}) | Get specific restore
*DefaultAPI* | [**GetRedfishV1SFSSAppSFSSHealthStatus**](docs/DefaultAPI.md#getredfishv1sfssappsfsshealthstatus) | **Get** /redfish/v1/SFSSApp/SFSSHealthStatus | Get system health
*DefaultAPI* | [**GetRedfishV1SFSSAppSosReports**](docs/DefaultAPI.md#getredfishv1sfssappsosreports) | **Get** /redfish/v1/SFSSApp/SosReports | Get SOS report
*DefaultAPI* | [**GetRedfishV1SFSSAppSosReportsID**](docs/DefaultAPI.md#getredfishv1sfssappsosreportsid) | **Get** /redfish/v1/SFSSApp/SosReports({ID}) | Get specific SOS report
*DefaultAPI* | [**GetRedfishV1SFSSAppSosReportsexpandSosReports**](docs/DefaultAPI.md#getredfishv1sfssappsosreportsexpandsosreports) | **Get** /redfish/v1/SFSSApp/SosReports?$expand&#x3D;SosReports | Get all SOS reports
*DefaultAPI* | [**GetRedfishV1SFSSAppTacacsServers**](docs/DefaultAPI.md#getredfishv1sfssapptacacsservers) | **Get** /redfish/v1/SFSSApp/TacacsServers | Get all TACACS+ servers
*DefaultAPI* | [**GetRedfishV1SFSSAppTacacsServersSequence**](docs/DefaultAPI.md#getredfishv1sfssapptacacsserverssequence) | **Get** /redfish/v1/SFSSApp/TacacsServers/Sequence | Get all TACACS+ servers
*DefaultAPI* | [**GetRedfishV1SFSSAppUserActivityAudit**](docs/DefaultAPI.md#getredfishv1sfssappuseractivityaudit) | **Get** /redfish/v1/SFSSApp/UserActivityAudit | Get user activities
*DefaultAPI* | [**GetRedfishV1SFSSAppUserActivityAuditID**](docs/DefaultAPI.md#getredfishv1sfssappuseractivityauditid) | **Get** /redfish/v1/SFSSApp/UserActivityAudit({ID}) | Get specific user activity
*DefaultAPI* | [**GetRedfishV1SFSSApp_0**](docs/DefaultAPI.md#getredfishv1sfssapp_0) | **Get** /redfish/v1/SFSSApp/SFSSImages | Get all images
*DefaultAPI* | [**PostRedfishV1SFSSAppAlerts**](docs/DefaultAPI.md#postredfishv1sfssappalerts) | **Post** /redfish/v1/SFSSApp/Alerts | Add alert
*DefaultAPI* | [**PostRedfishV1SFSSAppAuthenticationSequence**](docs/DefaultAPI.md#postredfishv1sfssappauthenticationsequence) | **Post** /redfish/v1/SFSSApp/AuthenticationSequence | Add authentication sequence
*DefaultAPI* | [**PostRedfishV1SFSSAppBackups**](docs/DefaultAPI.md#postredfishv1sfssappbackups) | **Post** /redfish/v1/SFSSApp/Backups | Perform backup
*DefaultAPI* | [**PostRedfishV1SFSSAppChangePassword**](docs/DefaultAPI.md#postredfishv1sfssappchangepassword) | **Post** /redfish/v1/SFSSApp/ChangePassword | Change admin password
*DefaultAPI* | [**PostRedfishV1SFSSAppFabricManagerInfoPost**](docs/DefaultAPI.md#postredfishv1sfssappfabricmanagerinfopost) | **Post** /redfish/v1/SFSSApp/CDCInstanceManagers | Configure CDC instance
*DefaultAPI* | [**PostRedfishV1SFSSAppGlobalSettings**](docs/DefaultAPI.md#postredfishv1sfssappglobalsettings) | **Post** /redfish/v1/SFSSApp/GlobalSettings | Configure global settings
*DefaultAPI* | [**PostRedfishV1SFSSAppIpAddressManagements**](docs/DefaultAPI.md#postredfishv1sfssappipaddressmanagements) | **Post** /redfish/v1/SFSSApp/IpAddressManagements | Configure interface
*DefaultAPI* | [**PostRedfishV1SFSSAppLicenses**](docs/DefaultAPI.md#postredfishv1sfssapplicenses) | **Post** /redfish/v1/SFSSApp/Licenses | Install a license
*DefaultAPI* | [**PostRedfishV1SFSSAppNTP1**](docs/DefaultAPI.md#postredfishv1sfssappntp1) | **Post** /redfish/v1/SFSSApp/NTP | Configure NTP server
*DefaultAPI* | [**PostRedfishV1SFSSAppRadiusServers**](docs/DefaultAPI.md#postredfishv1sfssappradiusservers) | **Post** /redfish/v1/SFSSApp/RadiusServers | Configure RADIUS server
*DefaultAPI* | [**PostRedfishV1SFSSAppRestores**](docs/DefaultAPI.md#postredfishv1sfssapprestores) | **Post** /redfish/v1/SFSSApp/Restores | Restore a backup
*DefaultAPI* | [**PostRedfishV1SFSSAppSFSSImages**](docs/DefaultAPI.md#postredfishv1sfssappsfssimages) | **Post** /redfish/v1/SFSSApp/SFSSImages | Add image
*DefaultAPI* | [**PostRedfishV1SFSSAppTacacsServers**](docs/DefaultAPI.md#postredfishv1sfssapptacacsservers) | **Post** /redfish/v1/SFSSApp/TacacsServers | Configure TACACS+ server
*DefaultAPI* | [**PutRedfishV1SFSSApp**](docs/DefaultAPI.md#putredfishv1sfssapp) | **Put** /redfish/v1/SFSSApp | Upgrade SFSS application
*DefaultAPI* | [**PutRedfishV1SFSSAppAlerts**](docs/DefaultAPI.md#putredfishv1sfssappalerts) | **Put** /redfish/v1/SFSSApp/Alerts({uuid}) | Update alert
*DefaultAPI* | [**PutRedfishV1SFSSAppCDCInstanceManagers**](docs/DefaultAPI.md#putredfishv1sfssappcdcinstancemanagers) | **Put** /redfish/v1/SFSSApp/CDCInstanceManagers({InstanceId}) | Update CDC instance
*DefaultAPI* | [**PutRedfishV1SFSSAppIpAddressManagements**](docs/DefaultAPI.md#putredfishv1sfssappipaddressmanagements) | **Put** /redfish/v1/SFSSApp/IpAddressManagements({InterfaceId}) | Update interface
*DefaultAPI* | [**PutRedfishV1SFSSAppLicenses**](docs/DefaultAPI.md#putredfishv1sfssapplicenses) | **Put** /redfish/v1/SFSSApp/Licenses | Accept EULA
*DefaultAPI* | [**PutRedfishV1SFSSAppNTP**](docs/DefaultAPI.md#putredfishv1sfssappntp) | **Put** /redfish/v1/SFSSApp/NTP | Enable or disable NTP service
*DefaultAPI* | [**PutRedfishV1SFSSApp_0**](docs/DefaultAPI.md#putredfishv1sfssapp_0) | **Put** /redfish/v1/SFSSApp/SFSSImages | Upgrade application


## Documentation For Models

 - [AlertsGET](docs/AlertsGET.md)
 - [AlertsID](docs/AlertsID.md)
 - [AuthenticationSequenceENUMs](docs/AuthenticationSequenceENUMs.md)
 - [AuthenticationSequenceGET](docs/AuthenticationSequenceGET.md)
 - [BackupsGET](docs/BackupsGET.md)
 - [BackupsGETEXPAND](docs/BackupsGETEXPAND.md)
 - [BackupsGETEXPANDBackupsInner](docs/BackupsGETEXPANDBackupsInner.md)
 - [BackupsGETID](docs/BackupsGETID.md)
 - [CDCHealthStatusGET](docs/CDCHealthStatusGET.md)
 - [CDCHealthStatusGETID](docs/CDCHealthStatusGETID.md)
 - [CDCInstanceManager](docs/CDCInstanceManager.md)
 - [CDCInstanceManagerENUMs](docs/CDCInstanceManagerENUMs.md)
 - [CDCInstanceManagerGETID](docs/CDCInstanceManagerGETID.md)
 - [CDCInstanceManagerPOST](docs/CDCInstanceManagerPOST.md)
 - [CDCInstanceManagerPUT](docs/CDCInstanceManagerPUT.md)
 - [DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers200Response](docs/DeleteRedfishV1SFSSAppAlertsCDCInstanceManagers200Response.md)
 - [DeleteRedfishV1SFSSAppNTPServerIP200Response](docs/DeleteRedfishV1SFSSAppNTPServerIP200Response.md)
 - [Device](docs/Device.md)
 - [EventsGET](docs/EventsGET.md)
 - [EventsGETEventsInner](docs/EventsGETEventsInner.md)
 - [EventsGETID](docs/EventsGETID.md)
 - [FoundationalConfigsGET](docs/FoundationalConfigsGET.md)
 - [FoundationalConfigsGETExpand](docs/FoundationalConfigsGETExpand.md)
 - [FoundationalConfigsGETID](docs/FoundationalConfigsGETID.md)
 - [GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums200Response](docs/GetEnumsRedfishV1SFSSAppIpAddressManagementsEnums200Response.md)
 - [GetExpandRedfishV1SFSSAppBackups200Response](docs/GetExpandRedfishV1SFSSAppBackups200Response.md)
 - [GetExpandRedfishV1SFSSAppBackups200ResponseBackupsInner](docs/GetExpandRedfishV1SFSSAppBackups200ResponseBackupsInner.md)
 - [GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200Response](docs/GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200Response.md)
 - [GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200ResponseFoundationalConfigsInner](docs/GetExpandRedfishV1SFSSAppFoundationalConfigsGETExpand200ResponseFoundationalConfigsInner.md)
 - [GetExpandRedfishV1SFSSAppIpAddressManagements200Response](docs/GetExpandRedfishV1SFSSAppIpAddressManagements200Response.md)
 - [GetExpandRedfishV1SFSSAppRestores200Response](docs/GetExpandRedfishV1SFSSAppRestores200Response.md)
 - [GetExpandRedfishV1SFSSAppRestores200ResponseRestoresInner](docs/GetExpandRedfishV1SFSSAppRestores200ResponseRestoresInner.md)
 - [GetIDRedfishV1SFSSApp200Response](docs/GetIDRedfishV1SFSSApp200Response.md)
 - [GetIDRedfishV1SFSSAppRadiusServers200Response](docs/GetIDRedfishV1SFSSAppRadiusServers200Response.md)
 - [GetIDRedfishV1SFSSAppTacacsServers200Response](docs/GetIDRedfishV1SFSSAppTacacsServers200Response.md)
 - [GetIDredfishV1SFSSAppSFSSInterfaceList200Response](docs/GetIDredfishV1SFSSAppSFSSInterfaceList200Response.md)
 - [GetRedfishV1SFSSApp200Response](docs/GetRedfishV1SFSSApp200Response.md)
 - [GetRedfishV1SFSSAppAlertsAlertId201Response](docs/GetRedfishV1SFSSAppAlertsAlertId201Response.md)
 - [GetRedfishV1SFSSAppAuthenticationSequence200Response](docs/GetRedfishV1SFSSAppAuthenticationSequence200Response.md)
 - [GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response](docs/GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response.md)
 - [GetRedfishV1SFSSAppBackups200Response](docs/GetRedfishV1SFSSAppBackups200Response.md)
 - [GetRedfishV1SFSSAppBackupsID200Response](docs/GetRedfishV1SFSSAppBackupsID200Response.md)
 - [GetRedfishV1SFSSAppCDCHealthStatus200Response](docs/GetRedfishV1SFSSAppCDCHealthStatus200Response.md)
 - [GetRedfishV1SFSSAppCDCHealthStatusID200Response](docs/GetRedfishV1SFSSAppCDCHealthStatusID200Response.md)
 - [GetRedfishV1SFSSAppCDCInstanceManagers200Response](docs/GetRedfishV1SFSSAppCDCInstanceManagers200Response.md)
 - [GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response](docs/GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response.md)
 - [GetRedfishV1SFSSAppCDCInstanceManagersID200Response](docs/GetRedfishV1SFSSAppCDCInstanceManagersID200Response.md)
 - [GetRedfishV1SFSSAppDevice200Response](docs/GetRedfishV1SFSSAppDevice200Response.md)
 - [GetRedfishV1SFSSAppEnums200Response](docs/GetRedfishV1SFSSAppEnums200Response.md)
 - [GetRedfishV1SFSSAppEvents200Response](docs/GetRedfishV1SFSSAppEvents200Response.md)
 - [GetRedfishV1SFSSAppEvents200ResponseEventsInner](docs/GetRedfishV1SFSSAppEvents200ResponseEventsInner.md)
 - [GetRedfishV1SFSSAppEventsID200Response](docs/GetRedfishV1SFSSAppEventsID200Response.md)
 - [GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier200Response](docs/GetRedfishV1SFSSAppFoundationalConfigsInstanceIdentifier200Response.md)
 - [GetRedfishV1SFSSAppGlobalSettings200Response](docs/GetRedfishV1SFSSAppGlobalSettings200Response.md)
 - [GetRedfishV1SFSSAppIpAddressManagements200Response](docs/GetRedfishV1SFSSAppIpAddressManagements200Response.md)
 - [GetRedfishV1SFSSAppIpAddressManagementsInterface200Response](docs/GetRedfishV1SFSSAppIpAddressManagementsInterface200Response.md)
 - [GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner](docs/GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner.md)
 - [GetRedfishV1SFSSAppLicenses200Response](docs/GetRedfishV1SFSSAppLicenses200Response.md)
 - [GetRedfishV1SFSSAppLicenses200ResponseLicensesInner](docs/GetRedfishV1SFSSAppLicenses200ResponseLicensesInner.md)
 - [GetRedfishV1SFSSAppLicensesExpandLicenses200Response](docs/GetRedfishV1SFSSAppLicensesExpandLicenses200Response.md)
 - [GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner](docs/GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner.md)
 - [GetRedfishV1SFSSAppLicensesLicenseId200Response](docs/GetRedfishV1SFSSAppLicensesLicenseId200Response.md)
 - [GetRedfishV1SFSSAppNTP200Response](docs/GetRedfishV1SFSSAppNTP200Response.md)
 - [GetRedfishV1SFSSAppRadiusServers200Response](docs/GetRedfishV1SFSSAppRadiusServers200Response.md)
 - [GetRedfishV1SFSSAppRestores200Response](docs/GetRedfishV1SFSSAppRestores200Response.md)
 - [GetRedfishV1SFSSAppRestoresID200Response](docs/GetRedfishV1SFSSAppRestoresID200Response.md)
 - [GetRedfishV1SFSSAppSFSSHealthStatus200Response](docs/GetRedfishV1SFSSAppSFSSHealthStatus200Response.md)
 - [GetRedfishV1SFSSAppSosReports200Response](docs/GetRedfishV1SFSSAppSosReports200Response.md)
 - [GetRedfishV1SFSSAppSosReportsExpandSosReports200Response](docs/GetRedfishV1SFSSAppSosReportsExpandSosReports200Response.md)
 - [GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner](docs/GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner.md)
 - [GetRedfishV1SFSSAppSosReportsID200Response](docs/GetRedfishV1SFSSAppSosReportsID200Response.md)
 - [GetRedfishV1SFSSAppTacacsServers200Response](docs/GetRedfishV1SFSSAppTacacsServers200Response.md)
 - [GetRedfishV1SFSSAppUserActivityAudit200Response](docs/GetRedfishV1SFSSAppUserActivityAudit200Response.md)
 - [GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner](docs/GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner.md)
 - [GetRedfishV1SFSSAppUserActivityAuditID200Response](docs/GetRedfishV1SFSSAppUserActivityAuditID200Response.md)
 - [GlobalSettingsGET](docs/GlobalSettingsGET.md)
 - [IpAddressManagementGET](docs/IpAddressManagementGET.md)
 - [IpAddressManagementGETEXPAND](docs/IpAddressManagementGETEXPAND.md)
 - [IpAddressManagementGETEXPANDIpAddressManagementsInner](docs/IpAddressManagementGETEXPANDIpAddressManagementsInner.md)
 - [IpAddressManagementGETID](docs/IpAddressManagementGETID.md)
 - [LicensesGET](docs/LicensesGET.md)
 - [LicensesGETExpand](docs/LicensesGETExpand.md)
 - [LicensesGETExpandLicensesInner](docs/LicensesGETExpandLicensesInner.md)
 - [LicensesGETID](docs/LicensesGETID.md)
 - [NTPServerConfig](docs/NTPServerConfig.md)
 - [NTPServerGET](docs/NTPServerGET.md)
 - [PostRedfishV1SFSSAppAlerts200Response](docs/PostRedfishV1SFSSAppAlerts200Response.md)
 - [PostRedfishV1SFSSAppAlertsRequest](docs/PostRedfishV1SFSSAppAlertsRequest.md)
 - [PostRedfishV1SFSSAppAuthenticationSequence200Response](docs/PostRedfishV1SFSSAppAuthenticationSequence200Response.md)
 - [PostRedfishV1SFSSAppAuthenticationSequenceRequest](docs/PostRedfishV1SFSSAppAuthenticationSequenceRequest.md)
 - [PostRedfishV1SFSSAppBackups200Response](docs/PostRedfishV1SFSSAppBackups200Response.md)
 - [PostRedfishV1SFSSAppBackupsRequest](docs/PostRedfishV1SFSSAppBackupsRequest.md)
 - [PostRedfishV1SFSSAppChangePasswordRequest](docs/PostRedfishV1SFSSAppChangePasswordRequest.md)
 - [PostRedfishV1SFSSAppFabricManagerInfoPost200Response](docs/PostRedfishV1SFSSAppFabricManagerInfoPost200Response.md)
 - [PostRedfishV1SFSSAppFabricManagerInfoPostRequest](docs/PostRedfishV1SFSSAppFabricManagerInfoPostRequest.md)
 - [PostRedfishV1SFSSAppGlobalSettingsRequest](docs/PostRedfishV1SFSSAppGlobalSettingsRequest.md)
 - [PostRedfishV1SFSSAppIpAddressManagementsRequest](docs/PostRedfishV1SFSSAppIpAddressManagementsRequest.md)
 - [PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner](docs/PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner.md)
 - [PostRedfishV1SFSSAppLicenses200Response](docs/PostRedfishV1SFSSAppLicenses200Response.md)
 - [PostRedfishV1SFSSAppLicensesRequest](docs/PostRedfishV1SFSSAppLicensesRequest.md)
 - [PostRedfishV1SFSSAppNTP1200Response](docs/PostRedfishV1SFSSAppNTP1200Response.md)
 - [PostRedfishV1SFSSAppNTP1Request](docs/PostRedfishV1SFSSAppNTP1Request.md)
 - [PostRedfishV1SFSSAppRadiusServers200Response](docs/PostRedfishV1SFSSAppRadiusServers200Response.md)
 - [PostRedfishV1SFSSAppRadiusServersRequest](docs/PostRedfishV1SFSSAppRadiusServersRequest.md)
 - [PostRedfishV1SFSSAppRestores200Response](docs/PostRedfishV1SFSSAppRestores200Response.md)
 - [PostRedfishV1SFSSAppRestoresRequest](docs/PostRedfishV1SFSSAppRestoresRequest.md)
 - [PostRedfishV1SFSSAppSFSSImagesRequest](docs/PostRedfishV1SFSSAppSFSSImagesRequest.md)
 - [PostRedfishV1SFSSAppTacacsServers200Response](docs/PostRedfishV1SFSSAppTacacsServers200Response.md)
 - [PostRedfishV1SFSSAppTacacsServersRequest](docs/PostRedfishV1SFSSAppTacacsServersRequest.md)
 - [PutRedfishV1SFSSApp200Response](docs/PutRedfishV1SFSSApp200Response.md)
 - [PutRedfishV1SFSSAppAlertsRequest](docs/PutRedfishV1SFSSAppAlertsRequest.md)
 - [PutRedfishV1SFSSAppCDCInstanceManagersRequest](docs/PutRedfishV1SFSSAppCDCInstanceManagersRequest.md)
 - [PutRedfishV1SFSSAppIpAddressManagementsRequest](docs/PutRedfishV1SFSSAppIpAddressManagementsRequest.md)
 - [PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner](docs/PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner.md)
 - [PutRedfishV1SFSSAppLicenses200Response](docs/PutRedfishV1SFSSAppLicenses200Response.md)
 - [PutRedfishV1SFSSAppLicensesRequest](docs/PutRedfishV1SFSSAppLicensesRequest.md)
 - [PutRedfishV1SFSSAppNTP200Response](docs/PutRedfishV1SFSSAppNTP200Response.md)
 - [PutRedfishV1SFSSAppNTPRequest](docs/PutRedfishV1SFSSAppNTPRequest.md)
 - [PutRedfishV1SFSSAppRequest](docs/PutRedfishV1SFSSAppRequest.md)
 - [RadiusServersGET](docs/RadiusServersGET.md)
 - [RadiusServersGETID](docs/RadiusServersGETID.md)
 - [RestoresGET](docs/RestoresGET.md)
 - [RestoresGETID](docs/RestoresGETID.md)
 - [SFSSApp](docs/SFSSApp.md)
 - [SFSSHealthStatusGET](docs/SFSSHealthStatusGET.md)
 - [SFSSImagesGET](docs/SFSSImagesGET.md)
 - [SFSSImagesGETENUMs](docs/SFSSImagesGETENUMs.md)
 - [SFSSImagesGETID](docs/SFSSImagesGETID.md)
 - [SFSSInterfaceListGET](docs/SFSSInterfaceListGET.md)
 - [SosReportsGET](docs/SosReportsGET.md)
 - [SosReportsGETEXPAND](docs/SosReportsGETEXPAND.md)
 - [SosReportsGETEXPANDSosReportsInner](docs/SosReportsGETEXPANDSosReportsInner.md)
 - [SosReportsGETID](docs/SosReportsGETID.md)
 - [TacacsServersGET](docs/TacacsServersGET.md)
 - [TacacsServersGETID](docs/TacacsServersGETID.md)
 - [UserActivityAuditGET](docs/UserActivityAuditGET.md)
 - [UserActivityAuditGETID](docs/UserActivityAuditGETID.md)
 - [UserActivityAuditGETUserActivityAuditsInner](docs/UserActivityAuditGETUserActivityAuditsInner.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BasicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), sfss.ContextBasicAuth, sfss.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



