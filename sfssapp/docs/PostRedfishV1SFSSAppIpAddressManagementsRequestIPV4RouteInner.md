# PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Destination IP address | [optional] 
**DestinationPrefix** | Pointer to **int32** | Destination subnet mask | [optional] 
**NextHop** | Pointer to **string** | Next hop IP address | [optional] 
**Metric** | Pointer to **int32** | The cost assigned to the route; valid range is from 1 to 255 | [optional] 

## Methods

### NewPostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner

`func NewPostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner() *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

NewPostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner instantiates a new PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInnerWithDefaults

`func NewPostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInnerWithDefaults() *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

NewPostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInnerWithDefaults instantiates a new PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationPrefix

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetDestinationPrefix() int32`

GetDestinationPrefix returns the DestinationPrefix field if non-nil, zero value otherwise.

### GetDestinationPrefixOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetDestinationPrefixOk() (*int32, bool)`

GetDestinationPrefixOk returns a tuple with the DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPrefix

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) SetDestinationPrefix(v int32)`

SetDestinationPrefix sets DestinationPrefix field to given value.

### HasDestinationPrefix

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) HasDestinationPrefix() bool`

HasDestinationPrefix returns a boolean if a field has been set.

### GetNextHop

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetMetric

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


