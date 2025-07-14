# PostRedfishV1SFSSAppFabricManagerInfoPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceIdentifier** | **string** | CDC instance identifier; possible values are 1 through 16 | 
**Interfaces** | **[]string** | Interface(s) assigned to the CDC instance | 
**CDCAdminState** | **string** | Administrative state of the CDC instance | 
**DiscoverySvcAdminState** | **string** | Administrative state of the discovery service | 

## Methods

### NewPostRedfishV1SFSSAppFabricManagerInfoPostRequest

`func NewPostRedfishV1SFSSAppFabricManagerInfoPostRequest(instanceIdentifier string, interfaces []string, cDCAdminState string, discoverySvcAdminState string, ) *PostRedfishV1SFSSAppFabricManagerInfoPostRequest`

NewPostRedfishV1SFSSAppFabricManagerInfoPostRequest instantiates a new PostRedfishV1SFSSAppFabricManagerInfoPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppFabricManagerInfoPostRequestWithDefaults

`func NewPostRedfishV1SFSSAppFabricManagerInfoPostRequestWithDefaults() *PostRedfishV1SFSSAppFabricManagerInfoPostRequest`

NewPostRedfishV1SFSSAppFabricManagerInfoPostRequestWithDefaults instantiates a new PostRedfishV1SFSSAppFabricManagerInfoPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceIdentifier

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetInstanceIdentifier() string`

GetInstanceIdentifier returns the InstanceIdentifier field if non-nil, zero value otherwise.

### GetInstanceIdentifierOk

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetInstanceIdentifierOk() (*string, bool)`

GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIdentifier

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) SetInstanceIdentifier(v string)`

SetInstanceIdentifier sets InstanceIdentifier field to given value.


### GetInterfaces

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.


### GetCDCAdminState

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetCDCAdminState() string`

GetCDCAdminState returns the CDCAdminState field if non-nil, zero value otherwise.

### GetCDCAdminStateOk

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetCDCAdminStateOk() (*string, bool)`

GetCDCAdminStateOk returns a tuple with the CDCAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCAdminState

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) SetCDCAdminState(v string)`

SetCDCAdminState sets CDCAdminState field to given value.


### GetDiscoverySvcAdminState

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetDiscoverySvcAdminState() string`

GetDiscoverySvcAdminState returns the DiscoverySvcAdminState field if non-nil, zero value otherwise.

### GetDiscoverySvcAdminStateOk

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) GetDiscoverySvcAdminStateOk() (*string, bool)`

GetDiscoverySvcAdminStateOk returns a tuple with the DiscoverySvcAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySvcAdminState

`func (o *PostRedfishV1SFSSAppFabricManagerInfoPostRequest) SetDiscoverySvcAdminState(v string)`

SetDiscoverySvcAdminState sets DiscoverySvcAdminState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


