# ZoneMemeberENUMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **[]map[string]interface{}** | Type of endpoint; possible values include host and subsystem | 
**ZoneMemberType** | **[]map[string]interface{}** | Possible values include NVMe Qualified Name (NQN) and Fully Qualified Name (FQN) | 

## Methods

### NewZoneMemeberENUMS

`func NewZoneMemeberENUMS(role []map[string]interface{}, zoneMemberType []map[string]interface{}, ) *ZoneMemeberENUMS`

NewZoneMemeberENUMS instantiates a new ZoneMemeberENUMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneMemeberENUMSWithDefaults

`func NewZoneMemeberENUMSWithDefaults() *ZoneMemeberENUMS`

NewZoneMemeberENUMSWithDefaults instantiates a new ZoneMemeberENUMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ZoneMemeberENUMS) GetRole() []map[string]interface{}`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ZoneMemeberENUMS) GetRoleOk() (*[]map[string]interface{}, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ZoneMemeberENUMS) SetRole(v []map[string]interface{})`

SetRole sets Role field to given value.


### GetZoneMemberType

`func (o *ZoneMemeberENUMS) GetZoneMemberType() []map[string]interface{}`

GetZoneMemberType returns the ZoneMemberType field if non-nil, zero value otherwise.

### GetZoneMemberTypeOk

`func (o *ZoneMemeberENUMS) GetZoneMemberTypeOk() (*[]map[string]interface{}, bool)`

GetZoneMemberTypeOk returns a tuple with the ZoneMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMemberType

`func (o *ZoneMemeberENUMS) SetZoneMemberType(v []map[string]interface{})`

SetZoneMemberType sets ZoneMemberType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


