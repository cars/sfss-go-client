# GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | SOS report ID | 
**FileLocation** | **string** | Location of the remote server to download the SOS report | 
**StatusMessage** | **string** | A detailed message on the status of the download operation | 
**ServerPassword** | **string** | Password to access the remote server | 
**Status** | **string** | Status of the SOS report download operation | 
**TimeStamp** | **string** | Date and time of the SOS report download operation; the date is of mm/dd/yyyy format and time is of hh:mm:ss format | 
**TransportType** | **string** | Transport type used for file transfer; possible values include SCP, HTTP, HTTPS, and SFTP | 
**ServerUserName** | **string** | Username to access the remote server | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner

`func NewGetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner(iD string, fileLocation string, statusMessage string, serverPassword string, status string, timeStamp string, transportType string, serverUserName string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner`

NewGetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner instantiates a new GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInnerWithDefaults

`func NewGetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInnerWithDefaults() *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner`

NewGetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInnerWithDefaults instantiates a new GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetID(v string)`

SetID sets ID field to given value.


### GetFileLocation

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.


### GetStatusMessage

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetServerPassword

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetServerPassword() string`

GetServerPassword returns the ServerPassword field if non-nil, zero value otherwise.

### GetServerPasswordOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetServerPasswordOk() (*string, bool)`

GetServerPasswordOk returns a tuple with the ServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPassword

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetServerPassword(v string)`

SetServerPassword sets ServerPassword field to given value.


### GetStatus

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeStamp

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetServerUserName

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetServerUserName() string`

GetServerUserName returns the ServerUserName field if non-nil, zero value otherwise.

### GetServerUserNameOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetServerUserNameOk() (*string, bool)`

GetServerUserNameOk returns a tuple with the ServerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUserName

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetServerUserName(v string)`

SetServerUserName sets ServerUserName field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppSosReportsExpandSosReports200ResponseSosReportsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


