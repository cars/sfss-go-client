# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | A unique NVMe Qualified Name (NQN) that is used to identify the endpoint. | 
**TotalNumOfEndPoints** | **float32** | Total number of endpoints; total number of hosts, subsystems, and DDCs | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 
**NumOfEndPointInUse** | **float32** | Total number of endpoints in use | 

## Methods

### NewDevice

`func NewDevice(deviceId string, totalNumOfEndPoints float32, odataId string, odataType string, odataContext string, numOfEndPointInUse float32, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *Device) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Device) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Device) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetTotalNumOfEndPoints

`func (o *Device) GetTotalNumOfEndPoints() float32`

GetTotalNumOfEndPoints returns the TotalNumOfEndPoints field if non-nil, zero value otherwise.

### GetTotalNumOfEndPointsOk

`func (o *Device) GetTotalNumOfEndPointsOk() (*float32, bool)`

GetTotalNumOfEndPointsOk returns a tuple with the TotalNumOfEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumOfEndPoints

`func (o *Device) SetTotalNumOfEndPoints(v float32)`

SetTotalNumOfEndPoints sets TotalNumOfEndPoints field to given value.


### GetOdataId

`func (o *Device) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *Device) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *Device) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *Device) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *Device) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *Device) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *Device) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *Device) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *Device) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.


### GetNumOfEndPointInUse

`func (o *Device) GetNumOfEndPointInUse() float32`

GetNumOfEndPointInUse returns the NumOfEndPointInUse field if non-nil, zero value otherwise.

### GetNumOfEndPointInUseOk

`func (o *Device) GetNumOfEndPointInUseOk() (*float32, bool)`

GetNumOfEndPointInUseOk returns a tuple with the NumOfEndPointInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfEndPointInUse

`func (o *Device) SetNumOfEndPointInUse(v float32)`

SetNumOfEndPointInUse sets NumOfEndPointInUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


