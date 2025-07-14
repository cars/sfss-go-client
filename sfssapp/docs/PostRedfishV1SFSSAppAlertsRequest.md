# PostRedfishV1SFSSAppAlertsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Protocol supported for alerts; possible value is redfish | 
**Context** | **string** |  | 
**EventTypes** | **[]string** |  | 
**CdcInstances** | **[]string** |  | 
**HttpHeaders** | **[]string** |  | 
**Destination** | **string** |  | 

## Methods

### NewPostRedfishV1SFSSAppAlertsRequest

`func NewPostRedfishV1SFSSAppAlertsRequest(protocol string, context string, eventTypes []string, cdcInstances []string, httpHeaders []string, destination string, ) *PostRedfishV1SFSSAppAlertsRequest`

NewPostRedfishV1SFSSAppAlertsRequest instantiates a new PostRedfishV1SFSSAppAlertsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppAlertsRequestWithDefaults

`func NewPostRedfishV1SFSSAppAlertsRequestWithDefaults() *PostRedfishV1SFSSAppAlertsRequest`

NewPostRedfishV1SFSSAppAlertsRequestWithDefaults instantiates a new PostRedfishV1SFSSAppAlertsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PostRedfishV1SFSSAppAlertsRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetContext

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PostRedfishV1SFSSAppAlertsRequest) SetContext(v string)`

SetContext sets Context field to given value.


### GetEventTypes

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *PostRedfishV1SFSSAppAlertsRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetCdcInstances

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetCdcInstances() []string`

GetCdcInstances returns the CdcInstances field if non-nil, zero value otherwise.

### GetCdcInstancesOk

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetCdcInstancesOk() (*[]string, bool)`

GetCdcInstancesOk returns a tuple with the CdcInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcInstances

`func (o *PostRedfishV1SFSSAppAlertsRequest) SetCdcInstances(v []string)`

SetCdcInstances sets CdcInstances field to given value.


### GetHttpHeaders

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetHttpHeaders() []string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetHttpHeadersOk() (*[]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *PostRedfishV1SFSSAppAlertsRequest) SetHttpHeaders(v []string)`

SetHttpHeaders sets HttpHeaders field to given value.


### GetDestination

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PostRedfishV1SFSSAppAlertsRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PostRedfishV1SFSSAppAlertsRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


