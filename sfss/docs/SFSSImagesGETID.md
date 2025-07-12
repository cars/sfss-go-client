# SFSSImagesGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageServerLocation** | **string** | Location of the SFSS image in the remote repository | 
**StatusMessage** | **string** | A detailed status message on the status of the add image operation. | 
**ImageServerPassword** | **string** | Password to access the image in the remote server | 
**Status** | **string** | Status of the add image operation | 
**TransportType** | **string** | Transport type used to copy the image from the remote repository; possible values include SCP, SFTP, HTTP, and HTTPS | 
**ImageServerUserName** | **string** | Username to access the remote repository | 
**Version** | **string** | Version of the SFSS image | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewSFSSImagesGETID

`func NewSFSSImagesGETID(imageServerLocation string, statusMessage string, imageServerPassword string, status string, transportType string, imageServerUserName string, version string, odataId string, odataType string, odataContext string, ) *SFSSImagesGETID`

NewSFSSImagesGETID instantiates a new SFSSImagesGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFSSImagesGETIDWithDefaults

`func NewSFSSImagesGETIDWithDefaults() *SFSSImagesGETID`

NewSFSSImagesGETIDWithDefaults instantiates a new SFSSImagesGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageServerLocation

`func (o *SFSSImagesGETID) GetImageServerLocation() string`

GetImageServerLocation returns the ImageServerLocation field if non-nil, zero value otherwise.

### GetImageServerLocationOk

`func (o *SFSSImagesGETID) GetImageServerLocationOk() (*string, bool)`

GetImageServerLocationOk returns a tuple with the ImageServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerLocation

`func (o *SFSSImagesGETID) SetImageServerLocation(v string)`

SetImageServerLocation sets ImageServerLocation field to given value.


### GetStatusMessage

`func (o *SFSSImagesGETID) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SFSSImagesGETID) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SFSSImagesGETID) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetImageServerPassword

`func (o *SFSSImagesGETID) GetImageServerPassword() string`

GetImageServerPassword returns the ImageServerPassword field if non-nil, zero value otherwise.

### GetImageServerPasswordOk

`func (o *SFSSImagesGETID) GetImageServerPasswordOk() (*string, bool)`

GetImageServerPasswordOk returns a tuple with the ImageServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerPassword

`func (o *SFSSImagesGETID) SetImageServerPassword(v string)`

SetImageServerPassword sets ImageServerPassword field to given value.


### GetStatus

`func (o *SFSSImagesGETID) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SFSSImagesGETID) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SFSSImagesGETID) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransportType

`func (o *SFSSImagesGETID) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *SFSSImagesGETID) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *SFSSImagesGETID) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetImageServerUserName

`func (o *SFSSImagesGETID) GetImageServerUserName() string`

GetImageServerUserName returns the ImageServerUserName field if non-nil, zero value otherwise.

### GetImageServerUserNameOk

`func (o *SFSSImagesGETID) GetImageServerUserNameOk() (*string, bool)`

GetImageServerUserNameOk returns a tuple with the ImageServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerUserName

`func (o *SFSSImagesGETID) SetImageServerUserName(v string)`

SetImageServerUserName sets ImageServerUserName field to given value.


### GetVersion

`func (o *SFSSImagesGETID) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SFSSImagesGETID) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SFSSImagesGETID) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetOdataId

`func (o *SFSSImagesGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *SFSSImagesGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *SFSSImagesGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *SFSSImagesGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *SFSSImagesGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *SFSSImagesGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *SFSSImagesGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *SFSSImagesGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *SFSSImagesGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


