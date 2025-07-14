# EventsGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | **[]map[string]interface{}** | Required arguments for the specified event  | 
**CDCInstance** | **string** | CDC instance identifier | 
**EEMI** | **string** | Enhanced Event Message Initiative (EEMI) Message Code | 
**HostName** | **string** | Hostname of the SFSS VM | 
**Message** | **string** | Message summary | 
**OriginOfCondition** | **[]map[string]interface{}** |  | 
**Severity** | **string** | Severity of the system message; possible values include Critical, Warning, and Informational | 
**Source** | **string** | SFSS service that is responsible for handling the event | 
**SourceSubType** | **string** | The module within the SFSS service that is responsible for handling the event  | 
**TimeStamp** | **string** | Date and time at which the event took place; the date is of the mm/dd/yyyy format and time is of the hh:mm:ss format | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewEventsGETID

`func NewEventsGETID(args []map[string]interface{}, cDCInstance string, eEMI string, hostName string, message string, originOfCondition []map[string]interface{}, severity string, source string, sourceSubType string, timeStamp string, odataId string, odataType string, odataContext string, ) *EventsGETID`

NewEventsGETID instantiates a new EventsGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsGETIDWithDefaults

`func NewEventsGETIDWithDefaults() *EventsGETID`

NewEventsGETIDWithDefaults instantiates a new EventsGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *EventsGETID) GetArgs() []map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *EventsGETID) GetArgsOk() (*[]map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *EventsGETID) SetArgs(v []map[string]interface{})`

SetArgs sets Args field to given value.


### GetCDCInstance

`func (o *EventsGETID) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *EventsGETID) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *EventsGETID) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetEEMI

`func (o *EventsGETID) GetEEMI() string`

GetEEMI returns the EEMI field if non-nil, zero value otherwise.

### GetEEMIOk

`func (o *EventsGETID) GetEEMIOk() (*string, bool)`

GetEEMIOk returns a tuple with the EEMI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEMI

`func (o *EventsGETID) SetEEMI(v string)`

SetEEMI sets EEMI field to given value.


### GetHostName

`func (o *EventsGETID) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *EventsGETID) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *EventsGETID) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetMessage

`func (o *EventsGETID) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventsGETID) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventsGETID) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOriginOfCondition

`func (o *EventsGETID) GetOriginOfCondition() []map[string]interface{}`

GetOriginOfCondition returns the OriginOfCondition field if non-nil, zero value otherwise.

### GetOriginOfConditionOk

`func (o *EventsGETID) GetOriginOfConditionOk() (*[]map[string]interface{}, bool)`

GetOriginOfConditionOk returns a tuple with the OriginOfCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOfCondition

`func (o *EventsGETID) SetOriginOfCondition(v []map[string]interface{})`

SetOriginOfCondition sets OriginOfCondition field to given value.


### GetSeverity

`func (o *EventsGETID) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventsGETID) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventsGETID) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSource

`func (o *EventsGETID) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventsGETID) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventsGETID) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceSubType

`func (o *EventsGETID) GetSourceSubType() string`

GetSourceSubType returns the SourceSubType field if non-nil, zero value otherwise.

### GetSourceSubTypeOk

`func (o *EventsGETID) GetSourceSubTypeOk() (*string, bool)`

GetSourceSubTypeOk returns a tuple with the SourceSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSubType

`func (o *EventsGETID) SetSourceSubType(v string)`

SetSourceSubType sets SourceSubType field to given value.


### GetTimeStamp

`func (o *EventsGETID) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *EventsGETID) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *EventsGETID) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetOdataId

`func (o *EventsGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *EventsGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *EventsGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *EventsGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *EventsGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *EventsGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *EventsGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *EventsGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *EventsGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


