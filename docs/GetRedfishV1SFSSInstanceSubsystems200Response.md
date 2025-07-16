# GetRedfishV1SFSSInstanceSubsystems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subsystems** | Pointer to [**[]GetRedfishV1SFSSInstanceSubsystems200ResponseSubsystemsInner**](GetRedfishV1SFSSInstanceSubsystems200ResponseSubsystemsInner.md) | Subsystem information | [optional] 
**SubsystemsodataCount** | **float32** | Number of subsystems registered with the CDC instance | 
**OdataId** | **string** |  | 
**OdataContext** | **string** |  | 
**OdataType** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSInstanceSubsystems200Response

`func NewGetRedfishV1SFSSInstanceSubsystems200Response(subsystemsodataCount float32, odataId string, odataContext string, odataType string, ) *GetRedfishV1SFSSInstanceSubsystems200Response`

NewGetRedfishV1SFSSInstanceSubsystems200Response instantiates a new GetRedfishV1SFSSInstanceSubsystems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSInstanceSubsystems200ResponseWithDefaults

`func NewGetRedfishV1SFSSInstanceSubsystems200ResponseWithDefaults() *GetRedfishV1SFSSInstanceSubsystems200Response`

NewGetRedfishV1SFSSInstanceSubsystems200ResponseWithDefaults instantiates a new GetRedfishV1SFSSInstanceSubsystems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubsystems

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetSubsystems() []GetRedfishV1SFSSInstanceSubsystems200ResponseSubsystemsInner`

GetSubsystems returns the Subsystems field if non-nil, zero value otherwise.

### GetSubsystemsOk

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetSubsystemsOk() (*[]GetRedfishV1SFSSInstanceSubsystems200ResponseSubsystemsInner, bool)`

GetSubsystemsOk returns a tuple with the Subsystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystems

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) SetSubsystems(v []GetRedfishV1SFSSInstanceSubsystems200ResponseSubsystemsInner)`

SetSubsystems sets Subsystems field to given value.

### HasSubsystems

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) HasSubsystems() bool`

HasSubsystems returns a boolean if a field has been set.

### GetSubsystemsodataCount

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetSubsystemsodataCount() float32`

GetSubsystemsodataCount returns the SubsystemsodataCount field if non-nil, zero value otherwise.

### GetSubsystemsodataCountOk

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetSubsystemsodataCountOk() (*float32, bool)`

GetSubsystemsodataCountOk returns a tuple with the SubsystemsodataCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystemsodataCount

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) SetSubsystemsodataCount(v float32)`

SetSubsystemsodataCount sets SubsystemsodataCount field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSInstanceSubsystems200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


