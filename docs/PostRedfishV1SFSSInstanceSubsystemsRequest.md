# PostRedfishV1SFSSInstanceSubsystemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NQN** | **string** | NVMe Qualified Name (NQN) of the subsystem | 
**TransportAddress** | **string** | IP address of the subsystem | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**PortId** | **float32** | Internal Port ID that is mapped to the NVMe Subsystem Port | 
**TransportType** | **string** | Supported transport types that can be used for communication with the controller; possible values are TCP, RoCE, FC | 
**TransportServiceId** | **string** | NVMe Transport Service Identifier as an ASCII string | 

## Methods

### NewPostRedfishV1SFSSInstanceSubsystemsRequest

`func NewPostRedfishV1SFSSInstanceSubsystemsRequest(nQN string, transportAddress string, transportAddressFamily string, portId float32, transportType string, transportServiceId string, ) *PostRedfishV1SFSSInstanceSubsystemsRequest`

NewPostRedfishV1SFSSInstanceSubsystemsRequest instantiates a new PostRedfishV1SFSSInstanceSubsystemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSInstanceSubsystemsRequestWithDefaults

`func NewPostRedfishV1SFSSInstanceSubsystemsRequestWithDefaults() *PostRedfishV1SFSSInstanceSubsystemsRequest`

NewPostRedfishV1SFSSInstanceSubsystemsRequestWithDefaults instantiates a new PostRedfishV1SFSSInstanceSubsystemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNQN

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetNQN() string`

GetNQN returns the NQN field if non-nil, zero value otherwise.

### GetNQNOk

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetNQNOk() (*string, bool)`

GetNQNOk returns a tuple with the NQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNQN

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) SetNQN(v string)`

SetNQN sets NQN field to given value.


### GetTransportAddress

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetPortId

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetPortId() float32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetPortIdOk() (*float32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) SetPortId(v float32)`

SetPortId sets PortId field to given value.


### GetTransportType

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetTransportServiceId

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportServiceId() string`

GetTransportServiceId returns the TransportServiceId field if non-nil, zero value otherwise.

### GetTransportServiceIdOk

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) GetTransportServiceIdOk() (*string, bool)`

GetTransportServiceIdOk returns a tuple with the TransportServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportServiceId

`func (o *PostRedfishV1SFSSInstanceSubsystemsRequest) SetTransportServiceId(v string)`

SetTransportServiceId sets TransportServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


