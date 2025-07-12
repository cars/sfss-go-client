# GetRedfishV1SFSSAppBackupsID200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Backup identifier | 
**ImageServerLocation** | **string** | Remote server location where the backup file is copied to; only IPv4 communication is supported | 
**StatusMessage** | **string** | Detailed status message of the backup operation | 
**ImageServerPassword** | **string** | Password to access the remote server | 
**Status** | **string** | Status of the backup; possible values include Success, Failure, NotStarted, InProgress | 
**TimeStamp** | **string** | Date and time at which the backup was taken; date in mm/dd/yyyy format and time in hh:mm:ss format | 
**TransportType** | **string** | Transport type used to copy the backup file to the remote server; possible values include SCP, HTTP, HTTPS, and SFTP | 
**ImageServerUserName** | **string** | Username to access the remote server | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppBackupsID200Response

`func NewGetRedfishV1SFSSAppBackupsID200Response(iD string, imageServerLocation string, statusMessage string, imageServerPassword string, status string, timeStamp string, transportType string, imageServerUserName string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppBackupsID200Response`

NewGetRedfishV1SFSSAppBackupsID200Response instantiates a new GetRedfishV1SFSSAppBackupsID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppBackupsID200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppBackupsID200ResponseWithDefaults() *GetRedfishV1SFSSAppBackupsID200Response`

NewGetRedfishV1SFSSAppBackupsID200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppBackupsID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetID(v string)`

SetID sets ID field to given value.


### GetImageServerLocation

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetImageServerLocation() string`

GetImageServerLocation returns the ImageServerLocation field if non-nil, zero value otherwise.

### GetImageServerLocationOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetImageServerLocationOk() (*string, bool)`

GetImageServerLocationOk returns a tuple with the ImageServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerLocation

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetImageServerLocation(v string)`

SetImageServerLocation sets ImageServerLocation field to given value.


### GetStatusMessage

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetImageServerPassword

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetImageServerPassword() string`

GetImageServerPassword returns the ImageServerPassword field if non-nil, zero value otherwise.

### GetImageServerPasswordOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetImageServerPasswordOk() (*string, bool)`

GetImageServerPasswordOk returns a tuple with the ImageServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerPassword

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetImageServerPassword(v string)`

SetImageServerPassword sets ImageServerPassword field to given value.


### GetStatus

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeStamp

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetImageServerUserName

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetImageServerUserName() string`

GetImageServerUserName returns the ImageServerUserName field if non-nil, zero value otherwise.

### GetImageServerUserNameOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetImageServerUserNameOk() (*string, bool)`

GetImageServerUserNameOk returns a tuple with the ImageServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerUserName

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetImageServerUserName(v string)`

SetImageServerUserName sets ImageServerUserName field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppBackupsID200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppBackupsID200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


