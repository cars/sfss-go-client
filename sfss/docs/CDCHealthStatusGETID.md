# CDCHealthStatusGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | **string** | Health status of the CDC instance; possible values include Ok, Warning, and Critical | 
**ReasonCode** | **[]map[string]interface{}** | Reason for the CDC instance to be in Warning or Critical state | 
**InstanceIdentifier** | **string** | CDC instance identifier | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewCDCHealthStatusGETID

`func NewCDCHealthStatusGETID(health string, reasonCode []map[string]interface{}, instanceIdentifier string, odataId string, odataType string, odataContext string, ) *CDCHealthStatusGETID`

NewCDCHealthStatusGETID instantiates a new CDCHealthStatusGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDCHealthStatusGETIDWithDefaults

`func NewCDCHealthStatusGETIDWithDefaults() *CDCHealthStatusGETID`

NewCDCHealthStatusGETIDWithDefaults instantiates a new CDCHealthStatusGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *CDCHealthStatusGETID) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CDCHealthStatusGETID) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CDCHealthStatusGETID) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetReasonCode

`func (o *CDCHealthStatusGETID) GetReasonCode() []map[string]interface{}`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CDCHealthStatusGETID) GetReasonCodeOk() (*[]map[string]interface{}, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CDCHealthStatusGETID) SetReasonCode(v []map[string]interface{})`

SetReasonCode sets ReasonCode field to given value.


### GetInstanceIdentifier

`func (o *CDCHealthStatusGETID) GetInstanceIdentifier() string`

GetInstanceIdentifier returns the InstanceIdentifier field if non-nil, zero value otherwise.

### GetInstanceIdentifierOk

`func (o *CDCHealthStatusGETID) GetInstanceIdentifierOk() (*string, bool)`

GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIdentifier

`func (o *CDCHealthStatusGETID) SetInstanceIdentifier(v string)`

SetInstanceIdentifier sets InstanceIdentifier field to given value.


### GetOdataId

`func (o *CDCHealthStatusGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *CDCHealthStatusGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *CDCHealthStatusGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *CDCHealthStatusGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *CDCHealthStatusGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *CDCHealthStatusGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *CDCHealthStatusGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *CDCHealthStatusGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *CDCHealthStatusGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


