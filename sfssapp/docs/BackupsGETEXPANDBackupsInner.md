# BackupsGETEXPANDBackupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Backup identifier | 
**ImageServerLocation** | **string** | Location where the backup file is copied to | 
**StatusMessage** | **string** | Detailed status message of the backup operation | 
**ImageServerPassword** | **string** | Password to access the remote server | 
**Status** | **string** | Status of the backup; possible values include Success, Failue, NotStarted, and InProgress | 
**TimeStamp** | **string** | Date and time at which the backup was taken; date in mm/dd/yyyy format and time in hh:mm:ss format. | 
**TransportType** | **string** | Transport type used to copy the backup file to the remote server; possible values include SCP, HTTP, HTTPS, and SFTP | 
**ImageServerUserName** | **string** | Username to access the remote server | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewBackupsGETEXPANDBackupsInner

`func NewBackupsGETEXPANDBackupsInner(iD string, imageServerLocation string, statusMessage string, imageServerPassword string, status string, timeStamp string, transportType string, imageServerUserName string, odataId string, odataType string, odataContext string, ) *BackupsGETEXPANDBackupsInner`

NewBackupsGETEXPANDBackupsInner instantiates a new BackupsGETEXPANDBackupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupsGETEXPANDBackupsInnerWithDefaults

`func NewBackupsGETEXPANDBackupsInnerWithDefaults() *BackupsGETEXPANDBackupsInner`

NewBackupsGETEXPANDBackupsInnerWithDefaults instantiates a new BackupsGETEXPANDBackupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *BackupsGETEXPANDBackupsInner) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BackupsGETEXPANDBackupsInner) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BackupsGETEXPANDBackupsInner) SetID(v string)`

SetID sets ID field to given value.


### GetImageServerLocation

`func (o *BackupsGETEXPANDBackupsInner) GetImageServerLocation() string`

GetImageServerLocation returns the ImageServerLocation field if non-nil, zero value otherwise.

### GetImageServerLocationOk

`func (o *BackupsGETEXPANDBackupsInner) GetImageServerLocationOk() (*string, bool)`

GetImageServerLocationOk returns a tuple with the ImageServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerLocation

`func (o *BackupsGETEXPANDBackupsInner) SetImageServerLocation(v string)`

SetImageServerLocation sets ImageServerLocation field to given value.


### GetStatusMessage

`func (o *BackupsGETEXPANDBackupsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *BackupsGETEXPANDBackupsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *BackupsGETEXPANDBackupsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetImageServerPassword

`func (o *BackupsGETEXPANDBackupsInner) GetImageServerPassword() string`

GetImageServerPassword returns the ImageServerPassword field if non-nil, zero value otherwise.

### GetImageServerPasswordOk

`func (o *BackupsGETEXPANDBackupsInner) GetImageServerPasswordOk() (*string, bool)`

GetImageServerPasswordOk returns a tuple with the ImageServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerPassword

`func (o *BackupsGETEXPANDBackupsInner) SetImageServerPassword(v string)`

SetImageServerPassword sets ImageServerPassword field to given value.


### GetStatus

`func (o *BackupsGETEXPANDBackupsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupsGETEXPANDBackupsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupsGETEXPANDBackupsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeStamp

`func (o *BackupsGETEXPANDBackupsInner) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *BackupsGETEXPANDBackupsInner) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *BackupsGETEXPANDBackupsInner) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetTransportType

`func (o *BackupsGETEXPANDBackupsInner) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *BackupsGETEXPANDBackupsInner) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *BackupsGETEXPANDBackupsInner) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetImageServerUserName

`func (o *BackupsGETEXPANDBackupsInner) GetImageServerUserName() string`

GetImageServerUserName returns the ImageServerUserName field if non-nil, zero value otherwise.

### GetImageServerUserNameOk

`func (o *BackupsGETEXPANDBackupsInner) GetImageServerUserNameOk() (*string, bool)`

GetImageServerUserNameOk returns a tuple with the ImageServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServerUserName

`func (o *BackupsGETEXPANDBackupsInner) SetImageServerUserName(v string)`

SetImageServerUserName sets ImageServerUserName field to given value.


### GetOdataId

`func (o *BackupsGETEXPANDBackupsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *BackupsGETEXPANDBackupsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *BackupsGETEXPANDBackupsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *BackupsGETEXPANDBackupsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *BackupsGETEXPANDBackupsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *BackupsGETEXPANDBackupsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *BackupsGETEXPANDBackupsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *BackupsGETEXPANDBackupsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *BackupsGETEXPANDBackupsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


