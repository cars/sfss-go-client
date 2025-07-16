# ZoneGroupGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneDBType** | **string** | Type of zone database; possible values include Config and Running | 
**ZoneGroupId** | **string** | Unique zone group identifier | 
**ZoneGroupName** | **string** | Zone group name | 
**OriginatorNQN** | **string** | NVMe Qualified Name (NQN) of the CDC instance | 
**Type** | **string** |  | 
**ActivateStatus** | **string** |  | 
**ActivationState** | **string** |  | 
**NumberZones** | **float32** | Number of zones within the zonegroup | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewZoneGroupGETID

`func NewZoneGroupGETID(zoneDBType string, zoneGroupId string, zoneGroupName string, originatorNQN string, type_ string, activateStatus string, activationState string, numberZones float32, odataId string, odataType string, odataContext string, ) *ZoneGroupGETID`

NewZoneGroupGETID instantiates a new ZoneGroupGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneGroupGETIDWithDefaults

`func NewZoneGroupGETIDWithDefaults() *ZoneGroupGETID`

NewZoneGroupGETIDWithDefaults instantiates a new ZoneGroupGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneDBType

`func (o *ZoneGroupGETID) GetZoneDBType() string`

GetZoneDBType returns the ZoneDBType field if non-nil, zero value otherwise.

### GetZoneDBTypeOk

`func (o *ZoneGroupGETID) GetZoneDBTypeOk() (*string, bool)`

GetZoneDBTypeOk returns a tuple with the ZoneDBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDBType

`func (o *ZoneGroupGETID) SetZoneDBType(v string)`

SetZoneDBType sets ZoneDBType field to given value.


### GetZoneGroupId

`func (o *ZoneGroupGETID) GetZoneGroupId() string`

GetZoneGroupId returns the ZoneGroupId field if non-nil, zero value otherwise.

### GetZoneGroupIdOk

`func (o *ZoneGroupGETID) GetZoneGroupIdOk() (*string, bool)`

GetZoneGroupIdOk returns a tuple with the ZoneGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneGroupId

`func (o *ZoneGroupGETID) SetZoneGroupId(v string)`

SetZoneGroupId sets ZoneGroupId field to given value.


### GetZoneGroupName

`func (o *ZoneGroupGETID) GetZoneGroupName() string`

GetZoneGroupName returns the ZoneGroupName field if non-nil, zero value otherwise.

### GetZoneGroupNameOk

`func (o *ZoneGroupGETID) GetZoneGroupNameOk() (*string, bool)`

GetZoneGroupNameOk returns a tuple with the ZoneGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneGroupName

`func (o *ZoneGroupGETID) SetZoneGroupName(v string)`

SetZoneGroupName sets ZoneGroupName field to given value.


### GetOriginatorNQN

`func (o *ZoneGroupGETID) GetOriginatorNQN() string`

GetOriginatorNQN returns the OriginatorNQN field if non-nil, zero value otherwise.

### GetOriginatorNQNOk

`func (o *ZoneGroupGETID) GetOriginatorNQNOk() (*string, bool)`

GetOriginatorNQNOk returns a tuple with the OriginatorNQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorNQN

`func (o *ZoneGroupGETID) SetOriginatorNQN(v string)`

SetOriginatorNQN sets OriginatorNQN field to given value.


### GetType

`func (o *ZoneGroupGETID) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneGroupGETID) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneGroupGETID) SetType(v string)`

SetType sets Type field to given value.


### GetActivateStatus

`func (o *ZoneGroupGETID) GetActivateStatus() string`

GetActivateStatus returns the ActivateStatus field if non-nil, zero value otherwise.

### GetActivateStatusOk

`func (o *ZoneGroupGETID) GetActivateStatusOk() (*string, bool)`

GetActivateStatusOk returns a tuple with the ActivateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateStatus

`func (o *ZoneGroupGETID) SetActivateStatus(v string)`

SetActivateStatus sets ActivateStatus field to given value.


### GetActivationState

`func (o *ZoneGroupGETID) GetActivationState() string`

GetActivationState returns the ActivationState field if non-nil, zero value otherwise.

### GetActivationStateOk

`func (o *ZoneGroupGETID) GetActivationStateOk() (*string, bool)`

GetActivationStateOk returns a tuple with the ActivationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationState

`func (o *ZoneGroupGETID) SetActivationState(v string)`

SetActivationState sets ActivationState field to given value.


### GetNumberZones

`func (o *ZoneGroupGETID) GetNumberZones() float32`

GetNumberZones returns the NumberZones field if non-nil, zero value otherwise.

### GetNumberZonesOk

`func (o *ZoneGroupGETID) GetNumberZonesOk() (*float32, bool)`

GetNumberZonesOk returns a tuple with the NumberZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberZones

`func (o *ZoneGroupGETID) SetNumberZones(v float32)`

SetNumberZones sets NumberZones field to given value.


### GetOdataId

`func (o *ZoneGroupGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *ZoneGroupGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *ZoneGroupGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *ZoneGroupGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *ZoneGroupGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *ZoneGroupGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *ZoneGroupGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *ZoneGroupGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *ZoneGroupGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


