# GetRedfishV1SFSSAppEventsID200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | **[]string** | Required arguments for the specified event | 
**CDCInstance** | **string** | CDC instance identifier | 
**EEMI** | **string** | Enhanced Event Message Initiative (EEMI) Message Code  | 
**HostName** | **string** | Hostname of the SFSS VM | 
**Message** | **string** | Message summary | 
**OriginOfCondition** | **[]string** |  | 
**Severity** | **string** | Severity of the system message; possible values include Critical, Warning, and Informational | 
**Source** | **string** | SFSS service that is responsible for handling the event | 
**SourceSubType** | **string** | The module within the SFSS service that is responsible for handling the event | 
**TimeStamp** | **string** | Date and time at which the event occurred; date in dd/mm/yyyy format and time in hh:mm:ss format | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppEventsID200Response

`func NewGetRedfishV1SFSSAppEventsID200Response(args []string, cDCInstance string, eEMI string, hostName string, message string, originOfCondition []string, severity string, source string, sourceSubType string, timeStamp string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppEventsID200Response`

NewGetRedfishV1SFSSAppEventsID200Response instantiates a new GetRedfishV1SFSSAppEventsID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppEventsID200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppEventsID200ResponseWithDefaults() *GetRedfishV1SFSSAppEventsID200Response`

NewGetRedfishV1SFSSAppEventsID200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppEventsID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetArgs(v []string)`

SetArgs sets Args field to given value.


### GetCDCInstance

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetEEMI

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetEEMI() string`

GetEEMI returns the EEMI field if non-nil, zero value otherwise.

### GetEEMIOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetEEMIOk() (*string, bool)`

GetEEMIOk returns a tuple with the EEMI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEMI

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetEEMI(v string)`

SetEEMI sets EEMI field to given value.


### GetHostName

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetMessage

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOriginOfCondition

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOriginOfCondition() []string`

GetOriginOfCondition returns the OriginOfCondition field if non-nil, zero value otherwise.

### GetOriginOfConditionOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOriginOfConditionOk() (*[]string, bool)`

GetOriginOfConditionOk returns a tuple with the OriginOfCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOfCondition

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetOriginOfCondition(v []string)`

SetOriginOfCondition sets OriginOfCondition field to given value.


### GetSeverity

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSource

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceSubType

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetSourceSubType() string`

GetSourceSubType returns the SourceSubType field if non-nil, zero value otherwise.

### GetSourceSubTypeOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetSourceSubTypeOk() (*string, bool)`

GetSourceSubTypeOk returns a tuple with the SourceSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSubType

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetSourceSubType(v string)`

SetSourceSubType sets SourceSubType field to given value.


### GetTimeStamp

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppEventsID200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppEventsID200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


