# GetRedfishV1SFSSAppSFSSHealthStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | **string** | Health status of the SFSS VM; possible values include Ok, Warning, and Critical | 
**ReasonCode** | **[]string** |  Reason for the SFSS VM health to be in Warning or Critical state | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppSFSSHealthStatus200Response

`func NewGetRedfishV1SFSSAppSFSSHealthStatus200Response(health string, reasonCode []string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppSFSSHealthStatus200Response`

NewGetRedfishV1SFSSAppSFSSHealthStatus200Response instantiates a new GetRedfishV1SFSSAppSFSSHealthStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppSFSSHealthStatus200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppSFSSHealthStatus200ResponseWithDefaults() *GetRedfishV1SFSSAppSFSSHealthStatus200Response`

NewGetRedfishV1SFSSAppSFSSHealthStatus200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppSFSSHealthStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetReasonCode

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetReasonCode() []string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetReasonCodeOk() (*[]string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) SetReasonCode(v []string)`

SetReasonCode sets ReasonCode field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppSFSSHealthStatus200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


