# GetRedfishV1SFSSAppEvents200ResponseEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Required arguments | [optional] 
**CDCInstance** | **string** | CDC instance identifier | 
**EEMI** | **string** | Enhanced Event Message Initiative (EEMI) Message Code | 
**HostName** | **string** | Hostname of the CDC instance | 
**Message** | **string** | Message summary | 
**OriginOfCondition** | Pointer to **[]string** |  | [optional] 
**Severity** | **string** | Severity of the event; possible values include Critical, Warning, and Informational | 
**Source** | **string** | SFSS service that is responsible for handling the event | 
**SourceSubType** | **string** | The module within the SFSS service that is responsible for handling the event | 
**TimeStamp** | **string** | Date and time at which the event occurred; date in dd/mm/yyyy format and time in hh:mm:ss format | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppEvents200ResponseEventsInner

`func NewGetRedfishV1SFSSAppEvents200ResponseEventsInner(cDCInstance string, eEMI string, hostName string, message string, severity string, source string, sourceSubType string, timeStamp string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppEvents200ResponseEventsInner`

NewGetRedfishV1SFSSAppEvents200ResponseEventsInner instantiates a new GetRedfishV1SFSSAppEvents200ResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppEvents200ResponseEventsInnerWithDefaults

`func NewGetRedfishV1SFSSAppEvents200ResponseEventsInnerWithDefaults() *GetRedfishV1SFSSAppEvents200ResponseEventsInner`

NewGetRedfishV1SFSSAppEvents200ResponseEventsInnerWithDefaults instantiates a new GetRedfishV1SFSSAppEvents200ResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCDCInstance

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetEEMI

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetEEMI() string`

GetEEMI returns the EEMI field if non-nil, zero value otherwise.

### GetEEMIOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetEEMIOk() (*string, bool)`

GetEEMIOk returns a tuple with the EEMI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEMI

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetEEMI(v string)`

SetEEMI sets EEMI field to given value.


### GetHostName

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetMessage

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOriginOfCondition

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOriginOfCondition() []string`

GetOriginOfCondition returns the OriginOfCondition field if non-nil, zero value otherwise.

### GetOriginOfConditionOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOriginOfConditionOk() (*[]string, bool)`

GetOriginOfConditionOk returns a tuple with the OriginOfCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOfCondition

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetOriginOfCondition(v []string)`

SetOriginOfCondition sets OriginOfCondition field to given value.

### HasOriginOfCondition

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) HasOriginOfCondition() bool`

HasOriginOfCondition returns a boolean if a field has been set.

### GetSeverity

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSource

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceSubType

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetSourceSubType() string`

GetSourceSubType returns the SourceSubType field if non-nil, zero value otherwise.

### GetSourceSubTypeOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetSourceSubTypeOk() (*string, bool)`

GetSourceSubTypeOk returns a tuple with the SourceSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSubType

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetSourceSubType(v string)`

SetSourceSubType sets SourceSubType field to given value.


### GetTimeStamp

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppEvents200ResponseEventsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


