# PostRedfishV1SFSSInstanceDDCsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DDCs** | Pointer to [**[]PostRedfishV1SFSSInstanceDDCsRequestAnyOfDDCsInner**](PostRedfishV1SFSSInstanceDDCsRequestAnyOfDDCsInner.md) |  | [optional] 
**TransportAddress** | Pointer to **string** |  | [optional] 
**TransportAddressFamily** | Pointer to **string** |  | [optional] 
**PortId** | Pointer to **int32** |  | [optional] 
**TransportType** | Pointer to **string** |  | [optional] 
**Activate** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostRedfishV1SFSSInstanceDDCsRequest

`func NewPostRedfishV1SFSSInstanceDDCsRequest() *PostRedfishV1SFSSInstanceDDCsRequest`

NewPostRedfishV1SFSSInstanceDDCsRequest instantiates a new PostRedfishV1SFSSInstanceDDCsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSInstanceDDCsRequestWithDefaults

`func NewPostRedfishV1SFSSInstanceDDCsRequestWithDefaults() *PostRedfishV1SFSSInstanceDDCsRequest`

NewPostRedfishV1SFSSInstanceDDCsRequestWithDefaults instantiates a new PostRedfishV1SFSSInstanceDDCsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDDCs

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetDDCs() []PostRedfishV1SFSSInstanceDDCsRequestAnyOfDDCsInner`

GetDDCs returns the DDCs field if non-nil, zero value otherwise.

### GetDDCsOk

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetDDCsOk() (*[]PostRedfishV1SFSSInstanceDDCsRequestAnyOfDDCsInner, bool)`

GetDDCsOk returns a tuple with the DDCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDDCs

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) SetDDCs(v []PostRedfishV1SFSSInstanceDDCsRequestAnyOfDDCsInner)`

SetDDCs sets DDCs field to given value.

### HasDDCs

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) HasDDCs() bool`

HasDDCs returns a boolean if a field has been set.

### GetTransportAddress

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.

### HasTransportAddress

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) HasTransportAddress() bool`

HasTransportAddress returns a boolean if a field has been set.

### GetTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.

### HasTransportAddressFamily

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) HasTransportAddressFamily() bool`

HasTransportAddressFamily returns a boolean if a field has been set.

### GetPortId

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetTransportType

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.

### HasTransportType

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) HasTransportType() bool`

HasTransportType returns a boolean if a field has been set.

### GetActivate

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *PostRedfishV1SFSSInstanceDDCsRequest) HasActivate() bool`

HasActivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


