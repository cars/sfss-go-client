# CDCInstanceManagerPOST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceIdentifier** | **string** | CDC instance identifier | 
**Interfaces** | **[]interface{}** | Interface that is assigned to the CDC instance  | 
**CDCAdminState** | Pointer to **string** | Administrative state of the CDC instance | [optional] 
**DiscoverySvcAdminState** | Pointer to **string** | Administrative state of the discovery service | [optional] 

## Methods

### NewCDCInstanceManagerPOST

`func NewCDCInstanceManagerPOST(instanceIdentifier string, interfaces []interface{}, ) *CDCInstanceManagerPOST`

NewCDCInstanceManagerPOST instantiates a new CDCInstanceManagerPOST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDCInstanceManagerPOSTWithDefaults

`func NewCDCInstanceManagerPOSTWithDefaults() *CDCInstanceManagerPOST`

NewCDCInstanceManagerPOSTWithDefaults instantiates a new CDCInstanceManagerPOST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceIdentifier

`func (o *CDCInstanceManagerPOST) GetInstanceIdentifier() string`

GetInstanceIdentifier returns the InstanceIdentifier field if non-nil, zero value otherwise.

### GetInstanceIdentifierOk

`func (o *CDCInstanceManagerPOST) GetInstanceIdentifierOk() (*string, bool)`

GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIdentifier

`func (o *CDCInstanceManagerPOST) SetInstanceIdentifier(v string)`

SetInstanceIdentifier sets InstanceIdentifier field to given value.


### GetInterfaces

`func (o *CDCInstanceManagerPOST) GetInterfaces() []interface{}`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *CDCInstanceManagerPOST) GetInterfacesOk() (*[]interface{}, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *CDCInstanceManagerPOST) SetInterfaces(v []interface{})`

SetInterfaces sets Interfaces field to given value.


### GetCDCAdminState

`func (o *CDCInstanceManagerPOST) GetCDCAdminState() string`

GetCDCAdminState returns the CDCAdminState field if non-nil, zero value otherwise.

### GetCDCAdminStateOk

`func (o *CDCInstanceManagerPOST) GetCDCAdminStateOk() (*string, bool)`

GetCDCAdminStateOk returns a tuple with the CDCAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCAdminState

`func (o *CDCInstanceManagerPOST) SetCDCAdminState(v string)`

SetCDCAdminState sets CDCAdminState field to given value.

### HasCDCAdminState

`func (o *CDCInstanceManagerPOST) HasCDCAdminState() bool`

HasCDCAdminState returns a boolean if a field has been set.

### GetDiscoverySvcAdminState

`func (o *CDCInstanceManagerPOST) GetDiscoverySvcAdminState() string`

GetDiscoverySvcAdminState returns the DiscoverySvcAdminState field if non-nil, zero value otherwise.

### GetDiscoverySvcAdminStateOk

`func (o *CDCInstanceManagerPOST) GetDiscoverySvcAdminStateOk() (*string, bool)`

GetDiscoverySvcAdminStateOk returns a tuple with the DiscoverySvcAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySvcAdminState

`func (o *CDCInstanceManagerPOST) SetDiscoverySvcAdminState(v string)`

SetDiscoverySvcAdminState sets DiscoverySvcAdminState field to given value.

### HasDiscoverySvcAdminState

`func (o *CDCInstanceManagerPOST) HasDiscoverySvcAdminState() bool`

HasDiscoverySvcAdminState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


