# GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | License identifier | 
**LicenseType** | **string** | Type of license; possible values include Evaluation, Enterprise, Partner, and Expansion licenses | 
**TotalNumOfEndPoints** | **float32** | Total number of endpoints that the license supports | 
**EULA** | **string** | End User License Agreement | 
**ServiceTag** | **string** | License service tag | 
**LicenseExpiry** | **string** | License expiry date | 
**DeviceId** | **string** | A unique NVMe Qualified Name (NQN) that is used to identify the SFSS VM | 

## Methods

### NewGetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner

`func NewGetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner(identifier string, licenseType string, totalNumOfEndPoints float32, eULA string, serviceTag string, licenseExpiry string, deviceId string, ) *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner`

NewGetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner instantiates a new GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInnerWithDefaults

`func NewGetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInnerWithDefaults() *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner`

NewGetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInnerWithDefaults instantiates a new GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLicenseType

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetTotalNumOfEndPoints

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetTotalNumOfEndPoints() float32`

GetTotalNumOfEndPoints returns the TotalNumOfEndPoints field if non-nil, zero value otherwise.

### GetTotalNumOfEndPointsOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetTotalNumOfEndPointsOk() (*float32, bool)`

GetTotalNumOfEndPointsOk returns a tuple with the TotalNumOfEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumOfEndPoints

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetTotalNumOfEndPoints(v float32)`

SetTotalNumOfEndPoints sets TotalNumOfEndPoints field to given value.


### GetEULA

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetEULA() string`

GetEULA returns the EULA field if non-nil, zero value otherwise.

### GetEULAOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetEULAOk() (*string, bool)`

GetEULAOk returns a tuple with the EULA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEULA

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetEULA(v string)`

SetEULA sets EULA field to given value.


### GetServiceTag

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetServiceTag() string`

GetServiceTag returns the ServiceTag field if non-nil, zero value otherwise.

### GetServiceTagOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetServiceTagOk() (*string, bool)`

GetServiceTagOk returns a tuple with the ServiceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTag

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetServiceTag(v string)`

SetServiceTag sets ServiceTag field to given value.


### GetLicenseExpiry

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetLicenseExpiry() string`

GetLicenseExpiry returns the LicenseExpiry field if non-nil, zero value otherwise.

### GetLicenseExpiryOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetLicenseExpiryOk() (*string, bool)`

GetLicenseExpiryOk returns a tuple with the LicenseExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiry

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetLicenseExpiry(v string)`

SetLicenseExpiry sets LicenseExpiry field to given value.


### GetDeviceId

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetRedfishV1SFSSAppLicensesExpandLicenses200ResponseLicensesInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


