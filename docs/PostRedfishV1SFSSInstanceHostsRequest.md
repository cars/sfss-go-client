# PostRedfishV1SFSSInstanceHostsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NQN** | **string** | NVMe Qualified Name (NQN) of the host | 
**TransportAddress** | **string** | IP address of the host | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**TransportType** | **string** | Protocol used for communication with the host; possible value is TCP  | 

## Methods

### NewPostRedfishV1SFSSInstanceHostsRequest

`func NewPostRedfishV1SFSSInstanceHostsRequest(nQN string, transportAddress string, transportAddressFamily string, transportType string, ) *PostRedfishV1SFSSInstanceHostsRequest`

NewPostRedfishV1SFSSInstanceHostsRequest instantiates a new PostRedfishV1SFSSInstanceHostsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSInstanceHostsRequestWithDefaults

`func NewPostRedfishV1SFSSInstanceHostsRequestWithDefaults() *PostRedfishV1SFSSInstanceHostsRequest`

NewPostRedfishV1SFSSInstanceHostsRequestWithDefaults instantiates a new PostRedfishV1SFSSInstanceHostsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNQN

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetNQN() string`

GetNQN returns the NQN field if non-nil, zero value otherwise.

### GetNQNOk

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetNQNOk() (*string, bool)`

GetNQNOk returns a tuple with the NQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNQN

`func (o *PostRedfishV1SFSSInstanceHostsRequest) SetNQN(v string)`

SetNQN sets NQN field to given value.


### GetTransportAddress

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *PostRedfishV1SFSSInstanceHostsRequest) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceHostsRequest) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetTransportType

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *PostRedfishV1SFSSInstanceHostsRequest) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *PostRedfishV1SFSSInstanceHostsRequest) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


