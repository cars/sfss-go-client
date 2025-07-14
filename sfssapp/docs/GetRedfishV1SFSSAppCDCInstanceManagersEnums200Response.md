# GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CDCAdminState** | **[]string** | Administrative state of the CDC instance | 
**DiscoverySvcAdminState** | **[]string** | Administrative state of the discovery service | 
**Status** | **[]string** | Status of Add CDC operation; possible values include INIT, INPROGRESS, SUCCESS, FAIL, and ABORT  | 

## Methods

### NewGetRedfishV1SFSSAppCDCInstanceManagersEnums200Response

`func NewGetRedfishV1SFSSAppCDCInstanceManagersEnums200Response(cDCAdminState []string, discoverySvcAdminState []string, status []string, ) *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response`

NewGetRedfishV1SFSSAppCDCInstanceManagersEnums200Response instantiates a new GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppCDCInstanceManagersEnums200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppCDCInstanceManagersEnums200ResponseWithDefaults() *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response`

NewGetRedfishV1SFSSAppCDCInstanceManagersEnums200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCDCAdminState

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) GetCDCAdminState() []string`

GetCDCAdminState returns the CDCAdminState field if non-nil, zero value otherwise.

### GetCDCAdminStateOk

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) GetCDCAdminStateOk() (*[]string, bool)`

GetCDCAdminStateOk returns a tuple with the CDCAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCAdminState

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) SetCDCAdminState(v []string)`

SetCDCAdminState sets CDCAdminState field to given value.


### GetDiscoverySvcAdminState

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) GetDiscoverySvcAdminState() []string`

GetDiscoverySvcAdminState returns the DiscoverySvcAdminState field if non-nil, zero value otherwise.

### GetDiscoverySvcAdminStateOk

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) GetDiscoverySvcAdminStateOk() (*[]string, bool)`

GetDiscoverySvcAdminStateOk returns a tuple with the DiscoverySvcAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySvcAdminState

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) SetDiscoverySvcAdminState(v []string)`

SetDiscoverySvcAdminState sets DiscoverySvcAdminState field to given value.


### GetStatus

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRedfishV1SFSSAppCDCInstanceManagersEnums200Response) SetStatus(v []string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


