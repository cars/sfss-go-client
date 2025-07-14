# GetRedfishV1SFSSAppCDCHealthStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CDCHealthStatus** | [**[]GetRedfishV1SFSSAppLicenses200ResponseLicensesInner**](GetRedfishV1SFSSAppLicenses200ResponseLicensesInner.md) | Health status of the CDC instance; possible values include Ok, Warning, and Critical | 
**CDCHealthStatusodataCount** | **float32** | Number of health status records (one for each CDC instance) | 
**OdataId** | **string** |  | 
**OdataContext** | **string** |  | 
**OdataType** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppCDCHealthStatus200Response

`func NewGetRedfishV1SFSSAppCDCHealthStatus200Response(cDCHealthStatus []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, cDCHealthStatusodataCount float32, odataId string, odataContext string, odataType string, ) *GetRedfishV1SFSSAppCDCHealthStatus200Response`

NewGetRedfishV1SFSSAppCDCHealthStatus200Response instantiates a new GetRedfishV1SFSSAppCDCHealthStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppCDCHealthStatus200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppCDCHealthStatus200ResponseWithDefaults() *GetRedfishV1SFSSAppCDCHealthStatus200Response`

NewGetRedfishV1SFSSAppCDCHealthStatus200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppCDCHealthStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCDCHealthStatus

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetCDCHealthStatus() []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner`

GetCDCHealthStatus returns the CDCHealthStatus field if non-nil, zero value otherwise.

### GetCDCHealthStatusOk

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetCDCHealthStatusOk() (*[]GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, bool)`

GetCDCHealthStatusOk returns a tuple with the CDCHealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCHealthStatus

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) SetCDCHealthStatus(v []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner)`

SetCDCHealthStatus sets CDCHealthStatus field to given value.


### GetCDCHealthStatusodataCount

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetCDCHealthStatusodataCount() float32`

GetCDCHealthStatusodataCount returns the CDCHealthStatusodataCount field if non-nil, zero value otherwise.

### GetCDCHealthStatusodataCountOk

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetCDCHealthStatusodataCountOk() (*float32, bool)`

GetCDCHealthStatusodataCountOk returns a tuple with the CDCHealthStatusodataCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCHealthStatusodataCount

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) SetCDCHealthStatusodataCount(v float32)`

SetCDCHealthStatusodataCount sets CDCHealthStatusodataCount field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppCDCHealthStatus200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


