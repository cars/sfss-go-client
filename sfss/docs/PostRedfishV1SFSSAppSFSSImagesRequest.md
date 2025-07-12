# PostRedfishV1SFSSAppSFSSImagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageServerUserName** | **string** | Username to access the remote server | 
**ImageServerPassword** | **string** | Password to access the image in the remote server | 
**ImageServerLocation** | **string** | Location of the SFSS image in the remote server; SFSS supports only IPv4 communication with remote servers | 
**TransportType** | **string** | Transport type used to copy the image from the remote server; possible values include SCP, SFTP, HTTP, and HTTPS | 

## Methods

### NewPostRedfishV1SFSSAppSFSSImagesRequest

`func NewPostRedfishV1SFSSAppSFSSImagesRequest(imageServerUserName string, imageServerPassword string, imageServerLocation string, transportType string, ) *PostRedfishV1SFSSAppSFSSImagesRequest`

NewPostRedfishV1SFSSAppSFSSImagesRequest instantiates a new PostRedfishV1SFSSAppSFSSImagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppSFSSImagesRequestWithDefaults

`func NewPostRedfishV1SFSSAppSFSSImagesRequestWithDefaults() *PostRedfishV1SFSSAppSFSSImagesRequest`

NewPostRedfishV1SFSSAppSFSSImagesRequestWithDefaults instantiates a new PostRedfishV1SFSSAppSFSSImagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageServerUserName

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetImageServerUserName() string`

GetImageServerUserName returns the ImageServerUserName field if non-nil, zero value otherwise.

### GetImageServerUserNameOk

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetImageServerUserNameOk() (*string, bool)`

GetImageServerUserNameOk returns a tuple with the ImageServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerUserName

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) SetImageServerUserName(v string)`

SetImageServerUserName sets ImageServerUserName field to given value.


### GetImageServerPassword

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetImageServerPassword() string`

GetImageServerPassword returns the ImageServerPassword field if non-nil, zero value otherwise.

### GetImageServerPasswordOk

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetImageServerPasswordOk() (*string, bool)`

GetImageServerPasswordOk returns a tuple with the ImageServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerPassword

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) SetImageServerPassword(v string)`

SetImageServerPassword sets ImageServerPassword field to given value.


### GetImageServerLocation

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetImageServerLocation() string`

GetImageServerLocation returns the ImageServerLocation field if non-nil, zero value otherwise.

### GetImageServerLocationOk

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetImageServerLocationOk() (*string, bool)`

GetImageServerLocationOk returns a tuple with the ImageServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerLocation

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) SetImageServerLocation(v string)`

SetImageServerLocation sets ImageServerLocation field to given value.


### GetTransportType

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *PostRedfishV1SFSSAppSFSSImagesRequest) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


