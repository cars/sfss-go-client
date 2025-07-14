# SFSSHealthStatusGET

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | **string** | Health status of the SFSS VM; possible values include Ok, Warning, and Critical | 
**ReasonCode** | **[]map[string]interface{}** | Reason for the SFSS VM health to be in Warning or Critical state | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewSFSSHealthStatusGET

`func NewSFSSHealthStatusGET(health string, reasonCode []map[string]interface{}, odataId string, odataType string, odataContext string, ) *SFSSHealthStatusGET`

NewSFSSHealthStatusGET instantiates a new SFSSHealthStatusGET object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFSSHealthStatusGETWithDefaults

`func NewSFSSHealthStatusGETWithDefaults() *SFSSHealthStatusGET`

NewSFSSHealthStatusGETWithDefaults instantiates a new SFSSHealthStatusGET object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *SFSSHealthStatusGET) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *SFSSHealthStatusGET) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *SFSSHealthStatusGET) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetReasonCode

`func (o *SFSSHealthStatusGET) GetReasonCode() []map[string]interface{}`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *SFSSHealthStatusGET) GetReasonCodeOk() (*[]map[string]interface{}, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *SFSSHealthStatusGET) SetReasonCode(v []map[string]interface{})`

SetReasonCode sets ReasonCode field to given value.


### GetOdataId

`func (o *SFSSHealthStatusGET) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *SFSSHealthStatusGET) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *SFSSHealthStatusGET) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *SFSSHealthStatusGET) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *SFSSHealthStatusGET) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *SFSSHealthStatusGET) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *SFSSHealthStatusGET) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *SFSSHealthStatusGET) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *SFSSHealthStatusGET) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


