# \DefaultAPI

All URIs are relative to *http://IPAddress*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRedfishV1SFSSInstanceHostsId**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceHostsId) | **Delete** /redfish/v1/SFSS/{InstanceID}/Hosts({Id}) | Delete host
[**DeleteRedfishV1SFSSInstanceIDDDCs**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceIDDDCs) | **Delete** /redfish/v1/SFSS/{InstanceID}/DDCs | Delete DDCs
[**DeleteRedfishV1SFSSInstanceIDDDCsId**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceIDDDCsId) | **Delete** /redfish/v1/SFSS/{InstanceID}/DDCs({Id}) | Delete DDC
[**DeleteRedfishV1SFSSInstanceSubsystemsId**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceSubsystemsId) | **Delete** /redfish/v1/SFSS/{InstanceID}/Subsystems({Id}) | Delete subsystem
[**DeleteRedfishV1SFSSInstanceZoneDBsConfig**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceZoneDBsConfig) | **Delete** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;) | Delete zoning configuration
[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias) | **Delete** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId}) | Delete a zone alias
[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers) | **Delete** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId})/ZoneAliasMembers({ZoneAliasMemberId}) | Delete zone alias member
[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups) | **Delete** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ID}) | Delete zone group
[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers) | **Delete** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGroupId})/Zones({ZoneID})/ZoneMembers({ZoneMemberID}) | Delete a zone member
[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones**](DefaultAPI.md#DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones) | **Delete** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGrpID})/Zones({ZoneID}) | Delete zone
[**GetHttpsIPRedfishV1SFSSInstanceDDCs**](DefaultAPI.md#GetHttpsIPRedfishV1SFSSInstanceDDCs) | **Get** /redfish/v1/SFSS/{InstanceID}/DDCs | Get DDC
[**GetHttpsIPRedfishV1SFSSInstanceHostsEnums**](DefaultAPI.md#GetHttpsIPRedfishV1SFSSInstanceHostsEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/Hosts/Enums | Get host enums
[**GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts**](DefaultAPI.md#GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts) | **Get** /redfish/v1/SFSS/{InstanceID}/Hosts2 | Get detailed host information
[**GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig**](DefaultAPI.md#GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGroupId})/Zones({ZoneID})/ZoneMembers({ZoneMemberID}) | Get zone members
[**GetRedfishV1SFSSInstanceCDCInstancesourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceCDCInstancesourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/CDCInstance?$source&#x3D;config | Get detailed CDC information
[**GetRedfishV1SFSSInstanceDDCsEnums**](DefaultAPI.md#GetRedfishV1SFSSInstanceDDCsEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/DDCs/Enums | Get DDC enums
[**GetRedfishV1SFSSInstanceDDCsId**](DefaultAPI.md#GetRedfishV1SFSSInstanceDDCsId) | **Get** /redfish/v1/SFSS/{InstanceID}/DDCs({Id}) | Get specific DDC
[**GetRedfishV1SFSSInstanceGlobalPolicies**](DefaultAPI.md#GetRedfishV1SFSSInstanceGlobalPolicies) | **Get** /redfish/v1/SFSS/{InstanceID}/GlobalPolicies | Get global policies
[**GetRedfishV1SFSSInstanceGlobalPoliciesEnums**](DefaultAPI.md#GetRedfishV1SFSSInstanceGlobalPoliciesEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/GlobalPolicies/Enums | Get global policy enums
[**GetRedfishV1SFSSInstanceHosts**](DefaultAPI.md#GetRedfishV1SFSSInstanceHosts) | **Get** /redfish/v1/SFSS/{InstanceID}/Hosts | Get all hosts
[**GetRedfishV1SFSSInstanceHostsId**](DefaultAPI.md#GetRedfishV1SFSSInstanceHostsId) | **Get** /redfish/v1/SFSS/{InstanceID}/Hosts({Id}) | Get specific host
[**GetRedfishV1SFSSInstanceSubsystems**](DefaultAPI.md#GetRedfishV1SFSSInstanceSubsystems) | **Get** /redfish/v1/SFSS/{InstanceID}/Subsystems | Get all subsystems
[**GetRedfishV1SFSSInstanceSubsystemsEnums**](DefaultAPI.md#GetRedfishV1SFSSInstanceSubsystemsEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/Subsystems/Enums | Get subsystem enums
[**GetRedfishV1SFSSInstanceSubsystemsId**](DefaultAPI.md#GetRedfishV1SFSSInstanceSubsystemsId) | **Get** /redfish/v1/SFSS/{InstanceID}/Subsystems({Id}) | Get specific subsystem
[**GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems**](DefaultAPI.md#GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems) | **Get** /redfish/v1/SFSS/{InstanceID}/Subsystems?$expand&#x3D;Subsystems | Get detailed subsystem information
[**GetRedfishV1SFSSInstanceZoneDBs**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBs) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs | Get zone database
[**GetRedfishV1SFSSInstanceZoneDBsActive**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsActive) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;active&#39;) | Get active database
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias?$source&#x3D;config&amp;$expand&#x3D;ZoneAlias | Get a zone alias
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId})/ZoneAliasMembers({ZoneAliasMemberId}) | Get zone alias member
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId)/ZoneAliasMembers | Get zone alias member
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId})/ZoneAliasMembers2({ZoneAliasMemberId}) | Get zone alias members
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias | Get zone alias
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGroupID}) | Get specific zone group
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGrpID})/Zones({ZoneID}) | Get all zones
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig) | **Get** /redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)/ZoneGroups(&#39;ZoneGroupId&#39;)/Zones(&#39;ZoneID&#39;)/ZoneMembers(&#39;ZoneMemberID&#39;) | Get a zone member
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups(&#39;ZoneGroupId&#39;)/Zones(&#39;ZoneID&#39;)/ZoneMembers{config}{zonemem} | Get zone members (detailed)
[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ID}) | Get specific zone group
[**GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;) | Get config database
[**GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs/ZoneAlias/ZoneAliasMembers/Enums | Get zone alias member enums
[**GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs/ZoneGroups/Enums | Get zone group enums
[**GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums**](DefaultAPI.md#GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums) | **Get** /redfish/v1/SFSS/{InstanceID}/ZoneDBs/ZoneGroups/Zones/ZoneMembers/Enums | Get zone member enums
[**PostRedfishV1SFSSInstanceDDCs**](DefaultAPI.md#PostRedfishV1SFSSInstanceDDCs) | **Post** /redfish/v1/SFSS/{InstanceID}/DDCs | Add DDC
[**PostRedfishV1SFSSInstanceGlobalPolicies**](DefaultAPI.md#PostRedfishV1SFSSInstanceGlobalPolicies) | **Post** /redfish/v1/SFSS/{InstanceID}/GlobalPolicies | Configure global policies
[**PostRedfishV1SFSSInstanceHosts**](DefaultAPI.md#PostRedfishV1SFSSInstanceHosts) | **Post** /redfish/v1/SFSS/{InstanceID}/Hosts | Add host
[**PostRedfishV1SFSSInstanceSubsystems**](DefaultAPI.md#PostRedfishV1SFSSInstanceSubsystems) | **Post** /redfish/v1/SFSS/{InstanceID}/Subsystems | Add subsystem
[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig**](DefaultAPI.md#PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig) | **Post** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId)/ZoneAliasMembers | Add zone alias member
[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias**](DefaultAPI.md#PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias) | **Post** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias | Add a zone alias
[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups**](DefaultAPI.md#PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups) | **Post** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ID}) | Add zone group
[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers**](DefaultAPI.md#PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers) | **Post** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGroupId})/Zones({ZoneID})/ZoneMembers({ZoneMemberID}) | Add zone member
[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig**](DefaultAPI.md#PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig) | **Post** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGrpID})/Zones({ZoneID}) | Add zone
[**PutRedfishV1SFSSInstanceIDDDCs**](DefaultAPI.md#PutRedfishV1SFSSInstanceIDDDCs) | **Put** /redfish/v1/SFSS/{InstanceID}/DDCs | Update DDC
[**PutRedfishV1SFSSInstanceIDDDCsId**](DefaultAPI.md#PutRedfishV1SFSSInstanceIDDDCsId) | **Put** /redfish/v1/SFSS/{InstanceID}/DDCs({Id}) | Update DDC
[**PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig**](DefaultAPI.md#PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig) | **Put** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneAlias({ZoneAliasId})/ZoneAliasMembers({ZoneAliasMemberId}) | Update zone alias member
[**PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups**](DefaultAPI.md#PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups) | **Put** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ID}) | Activate (or deactivate) zone group
[**PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers**](DefaultAPI.md#PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers) | **Put** /redfish/v1/SFSS/{InstanceID}/ZoneDBs(&#39;config&#39;)/ZoneGroups({ZoneGroupId})/Zones({ZoneID})/ZoneMembers({ZoneMemberID}) | Update a zone member



## DeleteRedfishV1SFSSInstanceHostsId

> DeleteRedfishV1SFSSInstanceHostsId(ctx, instanceID, id).Execute()

Delete host



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
	instanceID := int32(56) // int32 |  (default to 1)
	id := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceHostsId(context.Background(), instanceID, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceHostsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceHostsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceIDDDCs

> DeleteRedfishV1SFSSInstanceIDDDCs200Response DeleteRedfishV1SFSSInstanceIDDDCs(ctx, instanceID).Execute()

Delete DDCs



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceIDDDCs(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceIDDDCs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceIDDDCs`: DeleteRedfishV1SFSSInstanceIDDDCs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceIDDDCs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceIDDDCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRedfishV1SFSSInstanceIDDDCs200Response**](DeleteRedfishV1SFSSInstanceIDDDCs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceIDDDCsId

> PutRedfishV1SFSSInstanceIDDDCsId200Response DeleteRedfishV1SFSSInstanceIDDDCsId(ctx, instanceID, id2).Id(id).Execute()

Delete DDC



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
	instanceID := int32(56) // int32 |  (default to 1)
	id2 := int32(56) // int32 | 
	id := "id_example" // string | The IP address of the specific DDC to be removed must be given. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceIDDDCsId(context.Background(), instanceID, id2).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceIDDDCsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceIDDDCsId`: PutRedfishV1SFSSInstanceIDDDCsId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceIDDDCsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceIDDDCsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **string** | The IP address of the specific DDC to be removed must be given. | 

### Return type

[**PutRedfishV1SFSSInstanceIDDDCsId200Response**](PutRedfishV1SFSSInstanceIDDDCsId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceSubsystemsId

> DeleteRedfishV1SFSSInstanceSubsystemsId(ctx, instanceID, id).Execute()

Delete subsystem



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
	instanceID := int32(56) // int32 |  (default to 1)
	id := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceSubsystemsId(context.Background(), instanceID, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceSubsystemsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceSubsystemsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceZoneDBsConfig

> DeleteRedfishV1SFSSInstanceZoneDBsConfig200Response DeleteRedfishV1SFSSInstanceZoneDBsConfig(ctx, instanceID).Execute()

Delete zoning configuration



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfig(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceZoneDBsConfig`: DeleteRedfishV1SFSSInstanceZoneDBsConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceZoneDBsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRedfishV1SFSSInstanceZoneDBsConfig200Response**](DeleteRedfishV1SFSSInstanceZoneDBsConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias

> DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias200Response DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias(ctx, zoneAliasId, instanceID).Execute()

Delete a zone alias



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
	zoneAliasId := "zoneAliasId_example" // string | 
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias(context.Background(), zoneAliasId, instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias`: DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneAliasId** | **string** |  | 
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias200Response**](DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAlias200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers(ctx, instanceID, zoneAliasId, zoneAliasMemberId).Execute()

Delete zone alias member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneAliasId := "zoneAliasId_example" // string | 
	zoneAliasMemberId := "zoneAliasMemberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers(context.Background(), instanceID, zoneAliasId, zoneAliasMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneAliasId** | **string** |  | 
**zoneAliasMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups(ctx, instanceID, iD).ZoneGroupID(zoneGroupID).Execute()

Delete zone group



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
	instanceID := int32(56) // int32 |  (default to 1)
	iD := "iD_example" // string | 
	zoneGroupID := "zoneGroupID_example" // string | The Zone Group ID to be deleted should be passed to this API. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups(context.Background(), instanceID, iD).ZoneGroupID(zoneGroupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**iD** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneGroupID** | **string** | The Zone Group ID to be deleted should be passed to this API. | 

### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers

> DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers(ctx, instanceID, zoneGroupId, zoneID, zoneMemberID2).ZoneMemberID(zoneMemberID).Execute()

Delete a zone member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupId := "zoneGroupId_example" // string | 
	zoneID := "zoneID_example" // string | 
	zoneMemberID2 := "zoneMemberID_example" // string | 
	zoneMemberID := "zoneMemberID_example" // string | The ZoneMember ID to be deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers(context.Background(), instanceID, zoneGroupId, zoneID, zoneMemberID2).ZoneMemberID(zoneMemberID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers`: DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGroupId** | **string** |  | 
**zoneID** | **string** |  | 
**zoneMemberID2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **zoneMemberID** | **string** | The ZoneMember ID to be deleted. | 

### Return type

[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response**](DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones

> DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones200Response DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones(ctx, instanceID, zoneGrpID, zoneID2).ZoneID(zoneID).Execute()

Delete zone



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGrpID := "zoneGrpID_example" // string | 
	zoneID2 := "zoneID_example" // string | 
	zoneID := "zoneID_example" // string | The Zone ID to be deleted should be passed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones(context.Background(), instanceID, zoneGrpID, zoneID2).ZoneID(zoneID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones`: DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGrpID** | **string** |  | 
**zoneID2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **zoneID** | **string** | The Zone ID to be deleted should be passed. | 

### Return type

[**DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones200Response**](DeleteRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZones200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHttpsIPRedfishV1SFSSInstanceDDCs

> GetHttpsIPRedfishV1SFSSInstanceDDCs200Response GetHttpsIPRedfishV1SFSSInstanceDDCs(ctx, instanceID).Execute()

Get DDC



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceDDCs(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceDDCs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpsIPRedfishV1SFSSInstanceDDCs`: GetHttpsIPRedfishV1SFSSInstanceDDCs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceDDCs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpsIPRedfishV1SFSSInstanceDDCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHttpsIPRedfishV1SFSSInstanceDDCs200Response**](GetHttpsIPRedfishV1SFSSInstanceDDCs200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHttpsIPRedfishV1SFSSInstanceHostsEnums

> GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response GetHttpsIPRedfishV1SFSSInstanceHostsEnums(ctx, instanceID).Execute()

Get host enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceHostsEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceHostsEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpsIPRedfishV1SFSSInstanceHostsEnums`: GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceHostsEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpsIPRedfishV1SFSSInstanceHostsEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response**](GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts

> GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200Response GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts(ctx, instanceID).Expand(expand).Skip(skip).Filter(filter).Execute()

Get detailed host information



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
	expand := "expand_example" // string |  (default to "Hosts")
	instanceID := int32(56) // int32 |  (default to 1)
	skip := int32(56) // int32 | Helps to skip records (optional)
	filter := "filter_example" // string | Filter based on the  eq conditions: ex: TransportAddress eq 11.22.33.44 or TransportAddress eq 1.2.3.4 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts(context.Background(), instanceID).Expand(expand).Skip(skip).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts`: GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceHostsexpandHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpsIPRedfishV1SFSSInstanceHostsexpandHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** |  | [default to &quot;Hosts&quot;]

 **skip** | **int32** | Helps to skip records | 
 **filter** | **string** | Filter based on the  eq conditions: ex: TransportAddress eq 11.22.33.44 or TransportAddress eq 1.2.3.4 | 

### Return type

[**GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200Response**](GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig

> GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMembersSourceConfig200Response GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig(ctx, instanceID, zoneGroupId, zoneID, zoneMemberID).Sourceconfig(sourceconfig).Execute()

Get zone members



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupId := "zoneGroupId_example" // string | 
	zoneID := "zoneID_example" // string | 
	zoneMemberID := "zoneMemberID_example" // string | 
	sourceconfig := "sourceconfig_example" // string | This query param is mandatory. The request should looks like: /redfish/v1/SFSS/<Instance#>/ZoneDBs('config')/ZoneGroups('ZoneGroupId')/Zones('ZoneID')/ZoneMembers?$source=config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig(context.Background(), instanceID, zoneGroupId, zoneID, zoneMemberID).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig`: GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMembersSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGroupId** | **string** |  | 
**zoneID** | **string** |  | 
**zoneMemberID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMemberssourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sourceconfig** | **string** | This query param is mandatory. The request should looks like: /redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)/ZoneGroups(&#39;ZoneGroupId&#39;)/Zones(&#39;ZoneID&#39;)/ZoneMembers?$source&#x3D;config | 

### Return type

[**GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMembersSourceConfig200Response**](GetHttpsIPRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZonesConfigZoneMembersSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceCDCInstancesourceconfig

> GetRedfishV1SFSSInstanceCDCInstanceSourceConfig200Response GetRedfishV1SFSSInstanceCDCInstancesourceconfig(ctx, instanceID).Execute()

Get detailed CDC information



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceCDCInstancesourceconfig(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceCDCInstancesourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceCDCInstancesourceconfig`: GetRedfishV1SFSSInstanceCDCInstanceSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceCDCInstancesourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceCDCInstancesourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceCDCInstanceSourceConfig200Response**](GetRedfishV1SFSSInstanceCDCInstanceSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceDDCsEnums

> GetRedfishV1SFSSInstanceDDCsEnums200Response GetRedfishV1SFSSInstanceDDCsEnums(ctx, instanceID).Execute()

Get DDC enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceDDCsEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceDDCsEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceDDCsEnums`: GetRedfishV1SFSSInstanceDDCsEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceDDCsEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceDDCsEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceDDCsEnums200Response**](GetRedfishV1SFSSInstanceDDCsEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceDDCsId

> GetRedfishV1SFSSInstanceDDCsId200Response GetRedfishV1SFSSInstanceDDCsId(ctx, instanceID, id).Execute()

Get specific DDC



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
	instanceID := int32(56) // int32 |  (default to 1)
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceDDCsId(context.Background(), instanceID, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceDDCsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceDDCsId`: GetRedfishV1SFSSInstanceDDCsId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceDDCsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceDDCsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRedfishV1SFSSInstanceDDCsId200Response**](GetRedfishV1SFSSInstanceDDCsId200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceGlobalPolicies

> GetRedfishV1SFSSInstanceGlobalPolicies200Response GetRedfishV1SFSSInstanceGlobalPolicies(ctx, instanceID).Execute()

Get global policies



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceGlobalPolicies(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceGlobalPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceGlobalPolicies`: GetRedfishV1SFSSInstanceGlobalPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceGlobalPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceGlobalPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceGlobalPolicies200Response**](GetRedfishV1SFSSInstanceGlobalPolicies200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceGlobalPoliciesEnums

> GetRedfishV1SFSSInstanceGlobalPoliciesEnums200Response GetRedfishV1SFSSInstanceGlobalPoliciesEnums(ctx, instanceID).Execute()

Get global policy enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceGlobalPoliciesEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceGlobalPoliciesEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceGlobalPoliciesEnums`: GetRedfishV1SFSSInstanceGlobalPoliciesEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceGlobalPoliciesEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceGlobalPoliciesEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceGlobalPoliciesEnums200Response**](GetRedfishV1SFSSInstanceGlobalPoliciesEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceHosts

> GetRedfishV1SFSSInstanceHosts200Response GetRedfishV1SFSSInstanceHosts(ctx, instanceID).Sourceconfig(sourceconfig).Execute()

Get all hosts



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
	instanceID := int32(56) // int32 |  (default to 1)
	sourceconfig := "sourceconfig_example" // string | Optional query param to fetch from config DB (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceHosts(context.Background(), instanceID).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceHosts`: GetRedfishV1SFSSInstanceHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceconfig** | **string** | Optional query param to fetch from config DB | 

### Return type

[**GetRedfishV1SFSSInstanceHosts200Response**](GetRedfishV1SFSSInstanceHosts200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceHostsId

> GetRedfishV1SFSSInstanceHostsId200Response GetRedfishV1SFSSInstanceHostsId(ctx, instanceID, id).Execute()

Get specific host



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
	instanceID := int32(56) // int32 |  (default to 1)
	id := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceHostsId(context.Background(), instanceID, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceHostsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceHostsId`: GetRedfishV1SFSSInstanceHostsId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceHostsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceHostsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRedfishV1SFSSInstanceHostsId200Response**](GetRedfishV1SFSSInstanceHostsId200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceSubsystems

> GetRedfishV1SFSSInstanceSubsystems200Response GetRedfishV1SFSSInstanceSubsystems(ctx, instanceID).Sourceconfig(sourceconfig).Execute()

Get all subsystems



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
	instanceID := int32(56) // int32 |  (default to 1)
	sourceconfig := "sourceconfig_example" // string | This is an optional param to fetch the details from config DB. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceSubsystems(context.Background(), instanceID).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceSubsystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceSubsystems`: GetRedfishV1SFSSInstanceSubsystems200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceSubsystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceSubsystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceconfig** | **string** | This is an optional param to fetch the details from config DB. | 

### Return type

[**GetRedfishV1SFSSInstanceSubsystems200Response**](GetRedfishV1SFSSInstanceSubsystems200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceSubsystemsEnums

> GetRedfishV1SFSSInstanceSubsystemsEnums200Response GetRedfishV1SFSSInstanceSubsystemsEnums(ctx, instanceID).Execute()

Get subsystem enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceSubsystemsEnums`: GetRedfishV1SFSSInstanceSubsystemsEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceSubsystemsEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceSubsystemsEnums200Response**](GetRedfishV1SFSSInstanceSubsystemsEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceSubsystemsId

> GetRedfishV1SFSSInstanceSubsystemsId200Response GetRedfishV1SFSSInstanceSubsystemsId(ctx, instanceID, id).Execute()

Get specific subsystem



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
	instanceID := int32(56) // int32 |  (default to 1)
	id := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsId(context.Background(), instanceID, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceSubsystemsId`: GetRedfishV1SFSSInstanceSubsystemsId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceSubsystemsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRedfishV1SFSSInstanceSubsystemsId200Response**](GetRedfishV1SFSSInstanceSubsystemsId200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems

> GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200Response GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems(ctx, instanceID).FilterTransportAddressEq1111OrTransportAddressEq3333(filterTransportAddressEq1111OrTransportAddressEq3333).Execute()

Get detailed subsystem information



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
	instanceID := int32(56) // int32 |  (default to 1)
	filterTransportAddressEq1111OrTransportAddressEq3333 := "filterTransportAddressEq1111OrTransportAddressEq3333_example" // string | Optional query param to filer based on the eq condition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems(context.Background(), instanceID).FilterTransportAddressEq1111OrTransportAddressEq3333(filterTransportAddressEq1111OrTransportAddressEq3333).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems`: GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceSubsystemsexpandSubsystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceSubsystemsexpandSubsystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTransportAddressEq1111OrTransportAddressEq3333** | **string** | Optional query param to filer based on the eq condition. | 

### Return type

[**GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200Response**](GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBs

> GetRedfishV1SFSSInstanceZoneDBs200Response GetRedfishV1SFSSInstanceZoneDBs(ctx, instanceID).Sourceconfig(sourceconfig).Execute()

Get zone database



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
	instanceID := int32(56) // int32 |  (default to 1)
	sourceconfig := "sourceconfig_example" // string | To fetch from configuration database (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBs(context.Background(), instanceID).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBs`: GetRedfishV1SFSSInstanceZoneDBs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceconfig** | **string** | To fetch from configuration database | 

### Return type

[**GetRedfishV1SFSSInstanceZoneDBs200Response**](GetRedfishV1SFSSInstanceZoneDBs200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsActive

> GetRedfishV1SFSSInstanceZoneDBsActive200Response GetRedfishV1SFSSInstanceZoneDBsActive(ctx, instanceID).Execute()

Get active database



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsActive(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsActive`: GetRedfishV1SFSSInstanceZoneDBsActive200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceZoneDBsActive200Response**](GetRedfishV1SFSSInstanceZoneDBsActive200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig(ctx, instanceID).Execute()

Get a zone alias



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestsourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasTestSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig(ctx, instanceID, zoneAliasId, zoneAliasMemberId).Execute()

Get zone alias member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneAliasId := "zoneAliasId_example" // string | 
	zoneAliasMemberId := "zoneAliasMemberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig(context.Background(), instanceID, zoneAliasId, zoneAliasMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneAliasId** | **string** |  | 
**zoneAliasMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostsourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersConfigDhanaSampleAliasNqn201408OrgNvmexpressUuidHostSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig(ctx, instanceID, zoneAliasId).Sourceconfig(sourceconfig).Execute()

Get zone alias member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneAliasId := "zoneAliasId_example" // string | 
	sourceconfig := "sourceconfig_example" // string | This query parameter should be passed to fetch details. The sample query looks like: http://IPAddress//redfish/v1/SFSS/<Instance#>/ZoneDBs('config')/ZoneAlias('config:DhanaSampleAlias')/ZoneAliasMembers?$source=config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig(context.Background(), instanceID, zoneAliasId).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneAliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceconfig** | **string** | This query parameter should be passed to fetch details. The sample query looks like: http://IPAddress//redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)/ZoneAlias(&#39;config:DhanaSampleAlias&#39;)/ZoneAliasMembers?$source&#x3D;config | 

### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigExpandZoneAliasMembers200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers(ctx, config, zonealiasmem, instanceID, zoneAliasId, zoneAliasMemberId).Execute()

Get zone alias members



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
	config := "config_example" // string | ?$source=config
	zonealiasmem := "zonealiasmem_example" // string | &$expand=ZoneAliasMembers
	instanceID := int32(56) // int32 |  (default to 1)
	zoneAliasId := "zoneAliasId_example" // string | 
	zoneAliasMemberId := "zoneAliasMemberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers(context.Background(), config, zonealiasmem, instanceID, zoneAliasId, zoneAliasMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigExpandZoneAliasMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**config** | **string** | ?$source&#x3D;config | 
**zonealiasmem** | **string** | &amp;$expand&#x3D;ZoneAliasMembers | 
**instanceID** | **int32** |  | [default to 1]
**zoneAliasId** | **string** |  | 
**zoneAliasMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigexpandZoneAliasMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigExpandZoneAliasMembers200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigExpandZoneAliasMembers200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias(ctx, instanceID).SourceconfigexpandZoneAlias(sourceconfigexpandZoneAlias).Execute()

Get zone alias



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
	instanceID := int32(56) // int32 |  (default to 1)
	sourceconfigexpandZoneAlias := "sourceconfigexpandZoneAlias_example" // string | The query parameters should be passed. The URL should looks like http://IPAdress/redfish/v1/SFSS/<Instance#>/ZoneDBs('config')/ZoneAlias?$source=config&$expand=ZoneAlias (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias(context.Background(), instanceID).SourceconfigexpandZoneAlias(sourceconfigexpandZoneAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceconfigexpandZoneAlias** | **string** | The query parameters should be passed. The URL should looks like http://IPAdress/redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)/ZoneAlias?$source&#x3D;config&amp;$expand&#x3D;ZoneAlias | 

### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8SourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig(ctx, expand, instanceID, zoneGroupID).Execute()

Get specific zone group



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
	expand := "expand_example" // string | ?$source=config
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupID := "zoneGroupID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig(context.Background(), expand, instanceID, zoneGroupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8SourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expand** | **string** | ?$source&#x3D;config | 
**instanceID** | **int32** |  | [default to 1]
**zoneGroupID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8sourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8SourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGroup1Nqn198811ComDellSFSS120210706164404e8SourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8ZonesSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig(ctx, instanceID, zoneGrpID, zoneID).Sourceconfig(sourceconfig).Execute()

Get all zones



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGrpID := "zoneGrpID_example" // string | 
	zoneID := "zoneID_example" // string | 
	sourceconfig := "sourceconfig_example" // string | The API should looks like: /redfish/v1/SFSS/<Instance#>/ZoneDBs('config')/ZoneGroups('ZoneGrpID')/Zones?$source=config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig(context.Background(), instanceID, zoneGrpID, zoneID).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8ZonesSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8Zonessourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGrpID** | **string** |  | 
**zoneID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8ZonessourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sourceconfig** | **string** | The API should looks like: /redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)/ZoneGroups(&#39;ZoneGrpID&#39;)/Zones?$source&#x3D;config | 

### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8ZonesSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsConfigZoneGrpDhana1111Nqn198811ComDellSFSS120210706164404e8ZonesSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig(ctx, instanceID, zoneGroupId, zoneID, zoneMemberID).Execute()

Get a zone member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupId := "zoneGroupId_example" // string | 
	zoneID := "zoneID_example" // string | 
	zoneMemberID := "zoneMemberID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig(context.Background(), instanceID, zoneGroupId, zoneID, zoneMemberID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGroupId** | **string** |  | 
**zoneID** | **string** |  | 
**zoneMemberID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfigExpandZoneMembers200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers(ctx, config, zonemem, instanceID, zoneGroupId, zoneID).Execute()

Get zone members (detailed)



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
	config := "config_example" // string | ?$source=config
	zonemem := "zonemem_example" // string | &$expand=ZoneMembers
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupId := "zoneGroupId_example" // string | 
	zoneID := "zoneID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers(context.Background(), config, zonemem, instanceID, zoneGroupId, zoneID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfigExpandZoneMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**config** | **string** | ?$source&#x3D;config | 
**zonemem** | **string** | &amp;$expand&#x3D;ZoneMembers | 
**instanceID** | **int32** |  | [default to 1]
**zoneGroupId** | **string** |  | 
**zoneID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDsourceconfigexpandZoneMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfigExpandZoneMembers200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersZoneMemberIDSourceConfigExpandZoneMembers200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig(ctx, instanceID, iD).Sourceconfig(sourceconfig).Execute()

Get specific zone group



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
	instanceID := int32(56) // int32 |  (default to 1)
	iD := "iD_example" // string | 
	sourceconfig := "sourceconfig_example" // string | This parameter should be passed to fetch the Zone Groups (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig(context.Background(), instanceID, iD).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**iD** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupssourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceconfig** | **string** | This parameter should be passed to fetch the Zone Groups | 

### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig

> GetRedfishV1SFSSInstanceZoneDBsConfigSourceConfig200Response GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig(ctx, instanceID).Sourceconfig(sourceconfig).Execute()

Get config database



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
	instanceID := int32(56) // int32 |  (default to 1)
	sourceconfig := "sourceconfig_example" // string | This param is required to fetch the ZoneDB. The URL should looks like: /redfish/v1/SFSS/<Instance#>/ZoneDBs('config')?$source=config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig(context.Background(), instanceID).Sourceconfig(sourceconfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig`: GetRedfishV1SFSSInstanceZoneDBsConfigSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsConfigsourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsConfigsourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceconfig** | **string** | This param is required to fetch the ZoneDB. The URL should looks like: /redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)?$source&#x3D;config | 

### Return type

[**GetRedfishV1SFSSInstanceZoneDBsConfigSourceConfig200Response**](GetRedfishV1SFSSInstanceZoneDBsConfigSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums

> GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums200Response GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums(ctx, instanceID).Execute()

Get zone alias member enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums`: GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums200Response**](GetRedfishV1SFSSInstanceZoneDBsZoneAliasZoneAliasMembersEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums

> GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums200Response GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums(ctx, instanceID).Execute()

Get zone group enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums`: GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums200Response**](GetRedfishV1SFSSInstanceZoneDBsZoneGroupsEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums

> GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums200Response GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums(ctx, instanceID).Execute()

Get zone member enums



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
	instanceID := int32(56) // int32 |  (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums(context.Background(), instanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums`: GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums200Response**](GetRedfishV1SFSSInstanceZoneDBsZoneGroupsZonesZoneMembersEnums200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceDDCs

> PostRedfishV1SFSSInstanceDDCs200Response PostRedfishV1SFSSInstanceDDCs(ctx, instanceID).PostRedfishV1SFSSInstanceDDCsRequest(postRedfishV1SFSSInstanceDDCsRequest).Execute()

Add DDC



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
	instanceID := int32(56) // int32 |  (default to 1)
	postRedfishV1SFSSInstanceDDCsRequest := *openapiclient.NewPostRedfishV1SFSSInstanceDDCsRequest() // PostRedfishV1SFSSInstanceDDCsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceDDCs(context.Background(), instanceID).PostRedfishV1SFSSInstanceDDCsRequest(postRedfishV1SFSSInstanceDDCsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceDDCs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceDDCs`: PostRedfishV1SFSSInstanceDDCs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceDDCs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceDDCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRedfishV1SFSSInstanceDDCsRequest** | [**PostRedfishV1SFSSInstanceDDCsRequest**](PostRedfishV1SFSSInstanceDDCsRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceDDCs200Response**](PostRedfishV1SFSSInstanceDDCs200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceGlobalPolicies

> PostRedfishV1SFSSInstanceGlobalPolicies(ctx, instanceID).PostRedfishV1SFSSInstanceGlobalPoliciesRequest(postRedfishV1SFSSInstanceGlobalPoliciesRequest).Execute()

Configure global policies



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
	instanceID := int32(56) // int32 |  (default to 1)
	postRedfishV1SFSSInstanceGlobalPoliciesRequest := *openapiclient.NewPostRedfishV1SFSSInstanceGlobalPoliciesRequest("ZoningPolicy_example", "NameServerEntityPurgeTOV_example") // PostRedfishV1SFSSInstanceGlobalPoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceGlobalPolicies(context.Background(), instanceID).PostRedfishV1SFSSInstanceGlobalPoliciesRequest(postRedfishV1SFSSInstanceGlobalPoliciesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceGlobalPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceGlobalPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRedfishV1SFSSInstanceGlobalPoliciesRequest** | [**PostRedfishV1SFSSInstanceGlobalPoliciesRequest**](PostRedfishV1SFSSInstanceGlobalPoliciesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceHosts

> PostRedfishV1SFSSInstanceHosts200Response PostRedfishV1SFSSInstanceHosts(ctx, instanceID).PostRedfishV1SFSSInstanceHostsRequest(postRedfishV1SFSSInstanceHostsRequest).Execute()

Add host



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
	instanceID := int32(56) // int32 |  (default to 1)
	postRedfishV1SFSSInstanceHostsRequest := *openapiclient.NewPostRedfishV1SFSSInstanceHostsRequest("NQN_example", "TransportAddress_example", "TransportAddressFamily_example", "TransportType_example") // PostRedfishV1SFSSInstanceHostsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceHosts(context.Background(), instanceID).PostRedfishV1SFSSInstanceHostsRequest(postRedfishV1SFSSInstanceHostsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceHosts`: PostRedfishV1SFSSInstanceHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRedfishV1SFSSInstanceHostsRequest** | [**PostRedfishV1SFSSInstanceHostsRequest**](PostRedfishV1SFSSInstanceHostsRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceHosts200Response**](PostRedfishV1SFSSInstanceHosts200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceSubsystems

> PostRedfishV1SFSSInstanceSubsystems200Response PostRedfishV1SFSSInstanceSubsystems(ctx, instanceID).Sourceconfig(sourceconfig).PostRedfishV1SFSSInstanceSubsystemsRequest(postRedfishV1SFSSInstanceSubsystemsRequest).Execute()

Add subsystem



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
	instanceID := int32(56) // int32 |  (default to 1)
	sourceconfig := "sourceconfig_example" // string | This is an optional param to fetch the details from config DB. (optional)
	postRedfishV1SFSSInstanceSubsystemsRequest := *openapiclient.NewPostRedfishV1SFSSInstanceSubsystemsRequest("NQN_example", "TransportAddress_example", "TransportAddressFamily_example", float32(123), "TransportType_example", "TransportServiceId_example") // PostRedfishV1SFSSInstanceSubsystemsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceSubsystems(context.Background(), instanceID).Sourceconfig(sourceconfig).PostRedfishV1SFSSInstanceSubsystemsRequest(postRedfishV1SFSSInstanceSubsystemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceSubsystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceSubsystems`: PostRedfishV1SFSSInstanceSubsystems200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceSubsystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceSubsystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceconfig** | **string** | This is an optional param to fetch the details from config DB. | 
 **postRedfishV1SFSSInstanceSubsystemsRequest** | [**PostRedfishV1SFSSInstanceSubsystemsRequest**](PostRedfishV1SFSSInstanceSubsystemsRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceSubsystems200Response**](PostRedfishV1SFSSInstanceSubsystems200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig(ctx, instanceID, zoneAliasId).PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest).Execute()

Add zone alias member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneAliasId := "zoneAliasId_example" // string | 
	postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest := *openapiclient.NewPostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest("ZoneAliasMemberId_example", "ZoneAliasMemberType_example", "Role_example") // PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig(context.Background(), instanceID, zoneAliasId).PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneAliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest** | [**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias(ctx, instanceID).PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest).Execute()

Add a zone alias



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
	instanceID := int32(56) // int32 |  (default to 1)
	postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest := *openapiclient.NewPostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest("ZoneDBType_example", "ZoneAliasName_example") // PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias(context.Background(), instanceID).PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceZoneDBsConfigZoneAliassourceconfigexpandZoneAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest** | [**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAliasRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneAliasSourceConfigExpandZoneAlias200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups(ctx, instanceID, iD).PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest).Execute()

Add zone group



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
	instanceID := int32(56) // int32 |  (default to 1)
	iD := "iD_example" // string | 
	postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest := *openapiclient.NewPostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest("ZoneDBType_example", "ZoneGroupName_example") // PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups(context.Background(), instanceID, iD).PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**iD** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest** | [**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers(ctx, instanceID, zoneGroupId, zoneID, zoneMemberID).PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest).Execute()

Add zone member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupId := "zoneGroupId_example" // string | 
	zoneID := "zoneID_example" // string | 
	zoneMemberID := "zoneMemberID_example" // string | 
	postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest := *openapiclient.NewPostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest("ZoneMemberId_example", "ZoneMemberType_example") // PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers(context.Background(), instanceID, zoneGroupId, zoneID, zoneMemberID).PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGroupId** | **string** |  | 
**zoneID** | **string** |  | 
**zoneMemberID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest** | [**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig

> PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfig200Response PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig(ctx, instanceID, zoneGrpID, zoneID).PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest).Execute()

Add zone



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGrpID := "zoneGrpID_example" // string | 
	zoneID := "zoneID_example" // string | 
	postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest := *openapiclient.NewPostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest("ZoneName_example") // PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig(context.Background(), instanceID, zoneGrpID, zoneID).PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest(postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig`: PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGrpID** | **string** |  | 
**zoneID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonessourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest** | [**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfigRequest.md) |  | 

### Return type

[**PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfig200Response**](PostRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGrpIDZonesSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSInstanceIDDDCs

> PutRedfishV1SFSSInstanceIDDDCs200Response PutRedfishV1SFSSInstanceIDDDCs(ctx, instanceID).PutRedfishV1SFSSInstanceIDDDCsRequest(putRedfishV1SFSSInstanceIDDDCsRequest).Execute()

Update DDC



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
	instanceID := int32(56) // int32 |  (default to 1)
	putRedfishV1SFSSInstanceIDDDCsRequest := *openapiclient.NewPutRedfishV1SFSSInstanceIDDDCsRequest() // PutRedfishV1SFSSInstanceIDDDCsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSInstanceIDDDCs(context.Background(), instanceID).PutRedfishV1SFSSInstanceIDDDCsRequest(putRedfishV1SFSSInstanceIDDDCsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSInstanceIDDDCs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSInstanceIDDDCs`: PutRedfishV1SFSSInstanceIDDDCs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSInstanceIDDDCs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSInstanceIDDDCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRedfishV1SFSSInstanceIDDDCsRequest** | [**PutRedfishV1SFSSInstanceIDDDCsRequest**](PutRedfishV1SFSSInstanceIDDDCsRequest.md) |  | 

### Return type

[**PutRedfishV1SFSSInstanceIDDDCs200Response**](PutRedfishV1SFSSInstanceIDDDCs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSInstanceIDDDCsId

> PutRedfishV1SFSSInstanceIDDDCsId200Response PutRedfishV1SFSSInstanceIDDDCsId(ctx, instanceID, id2).Id(id).PutRedfishV1SFSSInstanceIDDDCsIdRequest(putRedfishV1SFSSInstanceIDDDCsIdRequest).Execute()

Update DDC



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
	instanceID := int32(56) // int32 |  (default to 1)
	id2 := int32(56) // int32 | 
	id := "id_example" // string | The DDC ID to be updated should be given. (optional)
	putRedfishV1SFSSInstanceIDDDCsIdRequest := *openapiclient.NewPutRedfishV1SFSSInstanceIDDDCsIdRequest() // PutRedfishV1SFSSInstanceIDDDCsIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSInstanceIDDDCsId(context.Background(), instanceID, id2).Id(id).PutRedfishV1SFSSInstanceIDDDCsIdRequest(putRedfishV1SFSSInstanceIDDDCsIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSInstanceIDDDCsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSInstanceIDDDCsId`: PutRedfishV1SFSSInstanceIDDDCsId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSInstanceIDDDCsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**id2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSInstanceIDDDCsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **string** | The DDC ID to be updated should be given. | 
 **putRedfishV1SFSSInstanceIDDDCsIdRequest** | [**PutRedfishV1SFSSInstanceIDDDCsIdRequest**](PutRedfishV1SFSSInstanceIDDDCsIdRequest.md) |  | 

### Return type

[**PutRedfishV1SFSSInstanceIDDDCsId200Response**](PutRedfishV1SFSSInstanceIDDDCsId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig

> PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig(ctx, instanceID, zoneAliasId, zoneAliasMemberId).PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest(putRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest).Execute()

Update zone alias member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneAliasId := "zoneAliasId_example" // string | 
	zoneAliasMemberId := "zoneAliasMemberId_example" // string | 
	putRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest := *openapiclient.NewPutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest("Role_example") // PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig(context.Background(), instanceID, zoneAliasId, zoneAliasMemberId).PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest(putRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig`: PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneAliasId** | **string** |  | 
**zoneAliasMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMemberssourceconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest** | [**PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest**](PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfigRequest.md) |  | 

### Return type

[**PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response**](PutRedfishV1SFSSInstanceZoneDBsConfigZoneAliasConfigDhanaSampleAliasZoneAliasMembersSourceConfig200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups

> PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups(ctx, instanceID, iD2).ID(iD).PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest(putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest).Execute()

Activate (or deactivate) zone group



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
	instanceID := int32(56) // int32 |  (default to 1)
	iD2 := "iD_example" // string | 
	iD := "iD_example" // string | The ID of Zonegroup should be passed for PUT. (optional)
	putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest := *openapiclient.NewPutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest("ActivateStatus_example") // PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest | The URL should looks like: http://IPAddress//redfish/v1/SFSS/<Instance#>/ZoneDBs('config')/ZoneGroups('ID') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups(context.Background(), instanceID, iD2).ID(iD).PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest(putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups`: PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**iD2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iD** | **string** | The ID of Zonegroup should be passed for PUT. | 
 **putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest** | [**PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest**](PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsRequest.md) | The URL should looks like: http://IPAddress//redfish/v1/SFSS/&lt;Instance#&gt;/ZoneDBs(&#39;config&#39;)/ZoneGroups(&#39;ID&#39;) | 

### Return type

[**PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response**](PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroups200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers

> PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers(ctx, instanceID, zoneGroupId, zoneID, zoneMemberID2).ZoneMemberID(zoneMemberID).PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest(putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest).Execute()

Update a zone member



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
	instanceID := int32(56) // int32 |  (default to 1)
	zoneGroupId := "zoneGroupId_example" // string | 
	zoneID := "zoneID_example" // string | 
	zoneMemberID2 := "zoneMemberID_example" // string | 
	zoneMemberID := "zoneMemberID_example" // string | The Zone Member ID to be updated. (optional)
	putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest := *openapiclient.NewPutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest("Role_example") // PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers(context.Background(), instanceID, zoneGroupId, zoneID, zoneMemberID2).ZoneMemberID(zoneMemberID).PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest(putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceID** | **int32** |  | [default to 1]
**zoneGroupId** | **string** |  | 
**zoneID** | **string** |  | 
**zoneMemberID2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **zoneMemberID** | **string** | The Zone Member ID to be updated. | 
 **putRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest** | [**PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest**](PutRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRedfishV1SFSSInstanceZoneDBsConfigZoneGroupsZoneGroupIdZonesZoneIDZoneMembersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

