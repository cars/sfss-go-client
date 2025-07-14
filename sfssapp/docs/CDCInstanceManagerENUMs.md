# CDCInstanceManagerENUMs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CDCAdminState** | **[]interface{}** | Administrative state of the CDC instance; possible values include Enabled and Disabled | 
**DiscoverySvcAdminState** | **[]interface{}** | Administrative state of the discovery service; possible values include Enabled and Disabled | 
**Status** | **[]interface{}** | Status of the CDC instance; possible values include Init, InProgress, Success, Fail, and Abort | 

## Methods

### NewCDCInstanceManagerENUMs

`func NewCDCInstanceManagerENUMs(cDCAdminState []interface{}, discoverySvcAdminState []interface{}, status []interface{}, ) *CDCInstanceManagerENUMs`

NewCDCInstanceManagerENUMs instantiates a new CDCInstanceManagerENUMs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDCInstanceManagerENUMsWithDefaults

`func NewCDCInstanceManagerENUMsWithDefaults() *CDCInstanceManagerENUMs`

NewCDCInstanceManagerENUMsWithDefaults instantiates a new CDCInstanceManagerENUMs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCDCAdminState

`func (o *CDCInstanceManagerENUMs) GetCDCAdminState() []interface{}`

GetCDCAdminState returns the CDCAdminState field if non-nil, zero value otherwise.

### GetCDCAdminStateOk

`func (o *CDCInstanceManagerENUMs) GetCDCAdminStateOk() (*[]interface{}, bool)`

GetCDCAdminStateOk returns a tuple with the CDCAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCAdminState

`func (o *CDCInstanceManagerENUMs) SetCDCAdminState(v []interface{})`

SetCDCAdminState sets CDCAdminState field to given value.


### GetDiscoverySvcAdminState

`func (o *CDCInstanceManagerENUMs) GetDiscoverySvcAdminState() []interface{}`

GetDiscoverySvcAdminState returns the DiscoverySvcAdminState field if non-nil, zero value otherwise.

### GetDiscoverySvcAdminStateOk

`func (o *CDCInstanceManagerENUMs) GetDiscoverySvcAdminStateOk() (*[]interface{}, bool)`

GetDiscoverySvcAdminStateOk returns a tuple with the DiscoverySvcAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySvcAdminState

`func (o *CDCInstanceManagerENUMs) SetDiscoverySvcAdminState(v []interface{})`

SetDiscoverySvcAdminState sets DiscoverySvcAdminState field to given value.


### GetStatus

`func (o *CDCInstanceManagerENUMs) GetStatus() []interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CDCInstanceManagerENUMs) GetStatusOk() (*[]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CDCInstanceManagerENUMs) SetStatus(v []interface{})`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


