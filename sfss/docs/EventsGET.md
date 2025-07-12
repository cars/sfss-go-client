# EventsGET

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventsGETEventsInner**](EventsGETEventsInner.md) | A set of events | 
**EventsodataCount** | **float32** | Number of events reported | 
**OdataId** | **string** |  | 
**OdataContext** | **string** |  | 
**OdataType** | **string** |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewEventsGET

`func NewEventsGET(events []EventsGETEventsInner, eventsodataCount float32, odataId string, odataContext string, odataType string, totalCount float32, ) *EventsGET`

NewEventsGET instantiates a new EventsGET object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsGETWithDefaults

`func NewEventsGETWithDefaults() *EventsGET`

NewEventsGETWithDefaults instantiates a new EventsGET object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsGET) GetEvents() []EventsGETEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsGET) GetEventsOk() (*[]EventsGETEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsGET) SetEvents(v []EventsGETEventsInner)`

SetEvents sets Events field to given value.


### GetEventsodataCount

`func (o *EventsGET) GetEventsodataCount() float32`

GetEventsodataCount returns the EventsodataCount field if non-nil, zero value otherwise.

### GetEventsodataCountOk

`func (o *EventsGET) GetEventsodataCountOk() (*float32, bool)`

GetEventsodataCountOk returns a tuple with the EventsodataCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsodataCount

`func (o *EventsGET) SetEventsodataCount(v float32)`

SetEventsodataCount sets EventsodataCount field to given value.


### GetOdataId

`func (o *EventsGET) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *EventsGET) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *EventsGET) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *EventsGET) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *EventsGET) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *EventsGET) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.


### GetOdataType

`func (o *EventsGET) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *EventsGET) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *EventsGET) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetTotalCount

`func (o *EventsGET) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EventsGET) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EventsGET) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


