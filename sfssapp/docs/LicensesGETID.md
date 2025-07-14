# LicensesGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceTag** | **string** | License service tag | 
**TotalNumOfEndPoints** | **float32** | Total number of endpoints the license supports | 
**Identifier** | **string** | License identifier | 
**LicenseType** | **string** | Type of license; possible values include Evaluation, Enterprise, Partner, and Expansion licenses | 
**LicenseExpiry** | **string** | License expiry date | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 
**DeviceId** | **string** | A unique NVMe Qualified Name (NQN) that is used to identify the SFSS VM. | 

## Methods

### NewLicensesGETID

`func NewLicensesGETID(serviceTag string, totalNumOfEndPoints float32, identifier string, licenseType string, licenseExpiry string, odataId string, odataType string, odataContext string, deviceId string, ) *LicensesGETID`

NewLicensesGETID instantiates a new LicensesGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensesGETIDWithDefaults

`func NewLicensesGETIDWithDefaults() *LicensesGETID`

NewLicensesGETIDWithDefaults instantiates a new LicensesGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceTag

`func (o *LicensesGETID) GetServiceTag() string`

GetServiceTag returns the ServiceTag field if non-nil, zero value otherwise.

### GetServiceTagOk

`func (o *LicensesGETID) GetServiceTagOk() (*string, bool)`

GetServiceTagOk returns a tuple with the ServiceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTag

`func (o *LicensesGETID) SetServiceTag(v string)`

SetServiceTag sets ServiceTag field to given value.


### GetTotalNumOfEndPoints

`func (o *LicensesGETID) GetTotalNumOfEndPoints() float32`

GetTotalNumOfEndPoints returns the TotalNumOfEndPoints field if non-nil, zero value otherwise.

### GetTotalNumOfEndPointsOk

`func (o *LicensesGETID) GetTotalNumOfEndPointsOk() (*float32, bool)`

GetTotalNumOfEndPointsOk returns a tuple with the TotalNumOfEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumOfEndPoints

`func (o *LicensesGETID) SetTotalNumOfEndPoints(v float32)`

SetTotalNumOfEndPoints sets TotalNumOfEndPoints field to given value.


### GetIdentifier

`func (o *LicensesGETID) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LicensesGETID) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LicensesGETID) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLicenseType

`func (o *LicensesGETID) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicensesGETID) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicensesGETID) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetLicenseExpiry

`func (o *LicensesGETID) GetLicenseExpiry() string`

GetLicenseExpiry returns the LicenseExpiry field if non-nil, zero value otherwise.

### GetLicenseExpiryOk

`func (o *LicensesGETID) GetLicenseExpiryOk() (*string, bool)`

GetLicenseExpiryOk returns a tuple with the LicenseExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiry

`func (o *LicensesGETID) SetLicenseExpiry(v string)`

SetLicenseExpiry sets LicenseExpiry field to given value.


### GetOdataId

`func (o *LicensesGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *LicensesGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *LicensesGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *LicensesGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *LicensesGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *LicensesGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *LicensesGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *LicensesGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *LicensesGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.


### GetDeviceId

`func (o *LicensesGETID) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *LicensesGETID) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *LicensesGETID) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


