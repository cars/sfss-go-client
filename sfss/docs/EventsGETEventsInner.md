# EventsGETEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]map[string]interface{}** | Required arguments | [optional] 
**CDCInstance** | **string** | CDC instance identifier | 
**EEMI** | **string** | Enhanced Event Message Initiative (EEMI) Message Code | 
**HostName** | **string** | Hostname of the CDC instance | 
**Message** | **string** | Message summary | 
**OriginOfCondition** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Severity** | **string** | Severity of the event; possible values include Critical, Warning, and Informational | 
**Source** | **string** | SFSS service that is responsible for handling the event | 
**SourceSubType** | **string** | The module within the SFSS service that is responsible for handling the event  | 
**TimeStamp** | **string** | Date and time at which the event occurred; date is of the mm/dd/yyyy format and time is of hh:mm:ss format | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewEventsGETEventsInner

`func NewEventsGETEventsInner(cDCInstance string, eEMI string, hostName string, message string, severity string, source string, sourceSubType string, timeStamp string, odataId string, odataType string, odataContext string, ) *EventsGETEventsInner`

NewEventsGETEventsInner instantiates a new EventsGETEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsGETEventsInnerWithDefaults

`func NewEventsGETEventsInnerWithDefaults() *EventsGETEventsInner`

NewEventsGETEventsInnerWithDefaults instantiates a new EventsGETEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *EventsGETEventsInner) GetArgs() []map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *EventsGETEventsInner) GetArgsOk() (*[]map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *EventsGETEventsInner) SetArgs(v []map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *EventsGETEventsInner) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCDCInstance

`func (o *EventsGETEventsInner) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *EventsGETEventsInner) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *EventsGETEventsInner) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetEEMI

`func (o *EventsGETEventsInner) GetEEMI() string`

GetEEMI returns the EEMI field if non-nil, zero value otherwise.

### GetEEMIOk

`func (o *EventsGETEventsInner) GetEEMIOk() (*string, bool)`

GetEEMIOk returns a tuple with the EEMI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEMI

`func (o *EventsGETEventsInner) SetEEMI(v string)`

SetEEMI sets EEMI field to given value.


### GetHostName

`func (o *EventsGETEventsInner) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *EventsGETEventsInner) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *EventsGETEventsInner) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetMessage

`func (o *EventsGETEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventsGETEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventsGETEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOriginOfCondition

`func (o *EventsGETEventsInner) GetOriginOfCondition() []map[string]interface{}`

GetOriginOfCondition returns the OriginOfCondition field if non-nil, zero value otherwise.

### GetOriginOfConditionOk

`func (o *EventsGETEventsInner) GetOriginOfConditionOk() (*[]map[string]interface{}, bool)`

GetOriginOfConditionOk returns a tuple with the OriginOfCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOfCondition

`func (o *EventsGETEventsInner) SetOriginOfCondition(v []map[string]interface{})`

SetOriginOfCondition sets OriginOfCondition field to given value.

### HasOriginOfCondition

`func (o *EventsGETEventsInner) HasOriginOfCondition() bool`

HasOriginOfCondition returns a boolean if a field has been set.

### GetSeverity

`func (o *EventsGETEventsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventsGETEventsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventsGETEventsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSource

`func (o *EventsGETEventsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventsGETEventsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventsGETEventsInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceSubType

`func (o *EventsGETEventsInner) GetSourceSubType() string`

GetSourceSubType returns the SourceSubType field if non-nil, zero value otherwise.

### GetSourceSubTypeOk

`func (o *EventsGETEventsInner) GetSourceSubTypeOk() (*string, bool)`

GetSourceSubTypeOk returns a tuple with the SourceSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSubType

`func (o *EventsGETEventsInner) SetSourceSubType(v string)`

SetSourceSubType sets SourceSubType field to given value.


### GetTimeStamp

`func (o *EventsGETEventsInner) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *EventsGETEventsInner) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *EventsGETEventsInner) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetOdataId

`func (o *EventsGETEventsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *EventsGETEventsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *EventsGETEventsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *EventsGETEventsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *EventsGETEventsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *EventsGETEventsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *EventsGETEventsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *EventsGETEventsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *EventsGETEventsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


