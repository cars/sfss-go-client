# GetIDRedfishV1SFSSApp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageServerLocation** | **string** | Location of the SFSS image in the remote server; SFSS supports only IPv4 communication with remote servers | 
**StatusMessage** | **string** | A detailed status message about the add image operation | 
**ImageServerPassword** | **string** | Password to access the image in the remote server | 
**Status** | **string** | Status of the add image operation | 
**TransportType** | **string** | Transport type used to copy the image from the remote server; possible values include SCP, SFTP, HTTP, and HTTPS | 
**ImageServerUserName** | **string** | Username to access the remote server | 
**Version** | **string** | Version of the SFSS image | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetIDRedfishV1SFSSApp200Response

`func NewGetIDRedfishV1SFSSApp200Response(imageServerLocation string, statusMessage string, imageServerPassword string, status string, transportType string, imageServerUserName string, version string, odataId string, odataType string, odataContext string, ) *GetIDRedfishV1SFSSApp200Response`

NewGetIDRedfishV1SFSSApp200Response instantiates a new GetIDRedfishV1SFSSApp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIDRedfishV1SFSSApp200ResponseWithDefaults

`func NewGetIDRedfishV1SFSSApp200ResponseWithDefaults() *GetIDRedfishV1SFSSApp200Response`

NewGetIDRedfishV1SFSSApp200ResponseWithDefaults instantiates a new GetIDRedfishV1SFSSApp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageServerLocation

`func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerLocation() string`

GetImageServerLocation returns the ImageServerLocation field if non-nil, zero value otherwise.

### GetImageServerLocationOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerLocationOk() (*string, bool)`

GetImageServerLocationOk returns a tuple with the ImageServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerLocation

`func (o *GetIDRedfishV1SFSSApp200Response) SetImageServerLocation(v string)`

SetImageServerLocation sets ImageServerLocation field to given value.


### GetStatusMessage

`func (o *GetIDRedfishV1SFSSApp200Response) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetIDRedfishV1SFSSApp200Response) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetImageServerPassword

`func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerPassword() string`

GetImageServerPassword returns the ImageServerPassword field if non-nil, zero value otherwise.

### GetImageServerPasswordOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerPasswordOk() (*string, bool)`

GetImageServerPasswordOk returns a tuple with the ImageServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerPassword

`func (o *GetIDRedfishV1SFSSApp200Response) SetImageServerPassword(v string)`

SetImageServerPassword sets ImageServerPassword field to given value.


### GetStatus

`func (o *GetIDRedfishV1SFSSApp200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIDRedfishV1SFSSApp200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransportType

`func (o *GetIDRedfishV1SFSSApp200Response) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetIDRedfishV1SFSSApp200Response) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetImageServerUserName

`func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerUserName() string`

GetImageServerUserName returns the ImageServerUserName field if non-nil, zero value otherwise.

### GetImageServerUserNameOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerUserNameOk() (*string, bool)`

GetImageServerUserNameOk returns a tuple with the ImageServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerUserName

`func (o *GetIDRedfishV1SFSSApp200Response) SetImageServerUserName(v string)`

SetImageServerUserName sets ImageServerUserName field to given value.


### GetVersion

`func (o *GetIDRedfishV1SFSSApp200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetIDRedfishV1SFSSApp200Response) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetOdataId

`func (o *GetIDRedfishV1SFSSApp200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetIDRedfishV1SFSSApp200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetIDRedfishV1SFSSApp200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetIDRedfishV1SFSSApp200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetIDRedfishV1SFSSApp200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetIDRedfishV1SFSSApp200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetIDRedfishV1SFSSApp200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


