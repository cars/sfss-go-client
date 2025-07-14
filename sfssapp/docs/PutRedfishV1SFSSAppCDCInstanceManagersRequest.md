# PutRedfishV1SFSSAppCDCInstanceManagersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceIdentifier** | **string** | CDC instance identifier | 
**Interfaces** | **[]string** | Interface(s) assigned to the CDC instance | 
**CDCAdminState** | **string** | Administrative state of the CDC instance | 
**DiscoverySvcAdminState** | **string** | Administrative state of the discovery service | 

## Methods

### NewPutRedfishV1SFSSAppCDCInstanceManagersRequest

`func NewPutRedfishV1SFSSAppCDCInstanceManagersRequest(instanceIdentifier string, interfaces []string, cDCAdminState string, discoverySvcAdminState string, ) *PutRedfishV1SFSSAppCDCInstanceManagersRequest`

NewPutRedfishV1SFSSAppCDCInstanceManagersRequest instantiates a new PutRedfishV1SFSSAppCDCInstanceManagersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRedfishV1SFSSAppCDCInstanceManagersRequestWithDefaults

`func NewPutRedfishV1SFSSAppCDCInstanceManagersRequestWithDefaults() *PutRedfishV1SFSSAppCDCInstanceManagersRequest`

NewPutRedfishV1SFSSAppCDCInstanceManagersRequestWithDefaults instantiates a new PutRedfishV1SFSSAppCDCInstanceManagersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceIdentifier

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetInstanceIdentifier() string`

GetInstanceIdentifier returns the InstanceIdentifier field if non-nil, zero value otherwise.

### GetInstanceIdentifierOk

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetInstanceIdentifierOk() (*string, bool)`

GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIdentifier

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) SetInstanceIdentifier(v string)`

SetInstanceIdentifier sets InstanceIdentifier field to given value.


### GetInterfaces

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.


### GetCDCAdminState

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetCDCAdminState() string`

GetCDCAdminState returns the CDCAdminState field if non-nil, zero value otherwise.

### GetCDCAdminStateOk

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetCDCAdminStateOk() (*string, bool)`

GetCDCAdminStateOk returns a tuple with the CDCAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCAdminState

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) SetCDCAdminState(v string)`

SetCDCAdminState sets CDCAdminState field to given value.


### GetDiscoverySvcAdminState

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetDiscoverySvcAdminState() string`

GetDiscoverySvcAdminState returns the DiscoverySvcAdminState field if non-nil, zero value otherwise.

### GetDiscoverySvcAdminStateOk

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) GetDiscoverySvcAdminStateOk() (*string, bool)`

GetDiscoverySvcAdminStateOk returns a tuple with the DiscoverySvcAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySvcAdminState

`func (o *PutRedfishV1SFSSAppCDCInstanceManagersRequest) SetDiscoverySvcAdminState(v string)`

SetDiscoverySvcAdminState sets DiscoverySvcAdminState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


