# PostRedfishV1SFSSAppBackupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageServerLocation** | **string** | Remote server location where the backup file is copied to; only IPv4 communication is supported | 
**ImageServerPassword** | **string** | Password to access the remote server | 
**TransportType** | **string** | Transport type used to copy the backup file to the remote server; possible values include SCP, HTTP, HTTPS, and SFTP | 
**ImageServerUserName** | **string** | Username to access the remote server | 

## Methods

### NewPostRedfishV1SFSSAppBackupsRequest

`func NewPostRedfishV1SFSSAppBackupsRequest(imageServerLocation string, imageServerPassword string, transportType string, imageServerUserName string, ) *PostRedfishV1SFSSAppBackupsRequest`

NewPostRedfishV1SFSSAppBackupsRequest instantiates a new PostRedfishV1SFSSAppBackupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppBackupsRequestWithDefaults

`func NewPostRedfishV1SFSSAppBackupsRequestWithDefaults() *PostRedfishV1SFSSAppBackupsRequest`

NewPostRedfishV1SFSSAppBackupsRequestWithDefaults instantiates a new PostRedfishV1SFSSAppBackupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageServerLocation

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetImageServerLocation() string`

GetImageServerLocation returns the ImageServerLocation field if non-nil, zero value otherwise.

### GetImageServerLocationOk

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetImageServerLocationOk() (*string, bool)`

GetImageServerLocationOk returns a tuple with the ImageServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerLocation

`func (o *PostRedfishV1SFSSAppBackupsRequest) SetImageServerLocation(v string)`

SetImageServerLocation sets ImageServerLocation field to given value.


### GetImageServerPassword

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetImageServerPassword() string`

GetImageServerPassword returns the ImageServerPassword field if non-nil, zero value otherwise.

### GetImageServerPasswordOk

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetImageServerPasswordOk() (*string, bool)`

GetImageServerPasswordOk returns a tuple with the ImageServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerPassword

`func (o *PostRedfishV1SFSSAppBackupsRequest) SetImageServerPassword(v string)`

SetImageServerPassword sets ImageServerPassword field to given value.


### GetTransportType

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *PostRedfishV1SFSSAppBackupsRequest) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetImageServerUserName

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetImageServerUserName() string`

GetImageServerUserName returns the ImageServerUserName field if non-nil, zero value otherwise.

### GetImageServerUserNameOk

`func (o *PostRedfishV1SFSSAppBackupsRequest) GetImageServerUserNameOk() (*string, bool)`

GetImageServerUserNameOk returns a tuple with the ImageServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerUserName

`func (o *PostRedfishV1SFSSAppBackupsRequest) SetImageServerUserName(v string)`

SetImageServerUserName sets ImageServerUserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


