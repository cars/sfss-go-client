# PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Destination IP address | [optional] 
**DestinationPrefix** | Pointer to **int32** | Subnet mask | [optional] 
**NextHop** | Pointer to **string** | Next hop IP address | [optional] 
**Metric** | Pointer to **int32** | The cost assigned to the interface; valid range is from 1 to 255 | [optional] 

## Methods

### NewPutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner

`func NewPutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner() *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner`

NewPutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner instantiates a new PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInnerWithDefaults

`func NewPutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInnerWithDefaults() *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner`

NewPutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInnerWithDefaults instantiates a new PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationPrefix

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetDestinationPrefix() int32`

GetDestinationPrefix returns the DestinationPrefix field if non-nil, zero value otherwise.

### GetDestinationPrefixOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetDestinationPrefixOk() (*int32, bool)`

GetDestinationPrefixOk returns a tuple with the DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPrefix

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) SetDestinationPrefix(v int32)`

SetDestinationPrefix sets DestinationPrefix field to given value.

### HasDestinationPrefix

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) HasDestinationPrefix() bool`

HasDestinationPrefix returns a boolean if a field has been set.

### GetNextHop

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetMetric

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


