# ZoneMemberGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | Type of endpoint; possible value includes host and subsystem | 
**ZoneMemberId** | **string** | Unique zone member identifier | 
**ZoneMemberType** | **string** | Possible values include NVMe qualified name (NQN) and Fully Qualified Name (FQN) | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewZoneMemberGETID

`func NewZoneMemberGETID(role string, zoneMemberId string, zoneMemberType string, odataId string, odataType string, odataContext string, ) *ZoneMemberGETID`

NewZoneMemberGETID instantiates a new ZoneMemberGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneMemberGETIDWithDefaults

`func NewZoneMemberGETIDWithDefaults() *ZoneMemberGETID`

NewZoneMemberGETIDWithDefaults instantiates a new ZoneMemberGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ZoneMemberGETID) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ZoneMemberGETID) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ZoneMemberGETID) SetRole(v string)`

SetRole sets Role field to given value.


### GetZoneMemberId

`func (o *ZoneMemberGETID) GetZoneMemberId() string`

GetZoneMemberId returns the ZoneMemberId field if non-nil, zero value otherwise.

### GetZoneMemberIdOk

`func (o *ZoneMemberGETID) GetZoneMemberIdOk() (*string, bool)`

GetZoneMemberIdOk returns a tuple with the ZoneMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMemberId

`func (o *ZoneMemberGETID) SetZoneMemberId(v string)`

SetZoneMemberId sets ZoneMemberId field to given value.


### GetZoneMemberType

`func (o *ZoneMemberGETID) GetZoneMemberType() string`

GetZoneMemberType returns the ZoneMemberType field if non-nil, zero value otherwise.

### GetZoneMemberTypeOk

`func (o *ZoneMemberGETID) GetZoneMemberTypeOk() (*string, bool)`

GetZoneMemberTypeOk returns a tuple with the ZoneMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMemberType

`func (o *ZoneMemberGETID) SetZoneMemberType(v string)`

SetZoneMemberType sets ZoneMemberType field to given value.


### GetOdataId

`func (o *ZoneMemberGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *ZoneMemberGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *ZoneMemberGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *ZoneMemberGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *ZoneMemberGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *ZoneMemberGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *ZoneMemberGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *ZoneMemberGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *ZoneMemberGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


