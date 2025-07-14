# SosReportsGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | SOS report ID | 
**FileLocation** | **string** | Location of the remote server to download the SOS report  | 
**StatusMessage** | **string** | A detailed status message of the SOS report download operation | 
**ServerPassword** | **string** | Password to access the remote server | 
**Status** | **string** | SOS report download status | 
**TimeStamp** | **string** | Date and time at which the SOS report was downloaded; the date is of the mm/dd/yyyy and time is of the hh:mm:ss format | 
**TransportType** | **string** | Protocol used for file transfer; possible values include SCP, SFTP, HTTP, and HTTPS | 
**ServerUserName** | **string** | Username to access the remote server | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewSosReportsGETID

`func NewSosReportsGETID(iD string, fileLocation string, statusMessage string, serverPassword string, status string, timeStamp string, transportType string, serverUserName string, odataId string, odataType string, odataContext string, ) *SosReportsGETID`

NewSosReportsGETID instantiates a new SosReportsGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSosReportsGETIDWithDefaults

`func NewSosReportsGETIDWithDefaults() *SosReportsGETID`

NewSosReportsGETIDWithDefaults instantiates a new SosReportsGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *SosReportsGETID) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *SosReportsGETID) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *SosReportsGETID) SetID(v string)`

SetID sets ID field to given value.


### GetFileLocation

`func (o *SosReportsGETID) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *SosReportsGETID) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *SosReportsGETID) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.


### GetStatusMessage

`func (o *SosReportsGETID) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SosReportsGETID) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SosReportsGETID) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetServerPassword

`func (o *SosReportsGETID) GetServerPassword() string`

GetServerPassword returns the ServerPassword field if non-nil, zero value otherwise.

### GetServerPasswordOk

`func (o *SosReportsGETID) GetServerPasswordOk() (*string, bool)`

GetServerPasswordOk returns a tuple with the ServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPassword

`func (o *SosReportsGETID) SetServerPassword(v string)`

SetServerPassword sets ServerPassword field to given value.


### GetStatus

`func (o *SosReportsGETID) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SosReportsGETID) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SosReportsGETID) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeStamp

`func (o *SosReportsGETID) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *SosReportsGETID) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *SosReportsGETID) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetTransportType

`func (o *SosReportsGETID) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *SosReportsGETID) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *SosReportsGETID) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetServerUserName

`func (o *SosReportsGETID) GetServerUserName() string`

GetServerUserName returns the ServerUserName field if non-nil, zero value otherwise.

### GetServerUserNameOk

`func (o *SosReportsGETID) GetServerUserNameOk() (*string, bool)`

GetServerUserNameOk returns a tuple with the ServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUserName

`func (o *SosReportsGETID) SetServerUserName(v string)`

SetServerUserName sets ServerUserName field to given value.


### GetOdataId

`func (o *SosReportsGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *SosReportsGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *SosReportsGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *SosReportsGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *SosReportsGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *SosReportsGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *SosReportsGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *SosReportsGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *SosReportsGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


