# PutRedfishV1SFSSAppIpAddressManagementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IPV4Address** | **[]string** | Interface IPv4 address  | 
**IPV4Config** | **string** | IP address configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV4Gateway** | **string** | IPv4 gateway address | 
**IPV4PrefixLength** | **float32** | Subnet mask | 
**IPV4Route** | Pointer to [**[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner**](PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner.md) | IPv4 static route configuration | [optional] 
**IPV6Config** | **string** | IP address configuration type; possible values include Manual, Automatic, and Disabled | 
**Name** | Pointer to **string** | Friendly name for the interface | [optional] 
**IPV6Address** | **[]string** | Interface IPv6 address  | 
**IPV6Gateway** | **string** | IPv6 gateway address | 
**IPV6PrefixLength** | **float32** | Subnet mask | 
**MTU** | Pointer to **float32** | The cost assigned to the route; valid range is from 1 to 255 | [optional] 
**IPV6Route** | Pointer to [**[]PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner**](PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner.md) |  | [optional] 

## Methods

### NewPutRedfishV1SFSSAppIpAddressManagementsRequest

`func NewPutRedfishV1SFSSAppIpAddressManagementsRequest(iPV4Address []string, iPV4Config string, iPV4Gateway string, iPV4PrefixLength float32, iPV6Config string, iPV6Address []string, iPV6Gateway string, iPV6PrefixLength float32, ) *PutRedfishV1SFSSAppIpAddressManagementsRequest`

NewPutRedfishV1SFSSAppIpAddressManagementsRequest instantiates a new PutRedfishV1SFSSAppIpAddressManagementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRedfishV1SFSSAppIpAddressManagementsRequestWithDefaults

`func NewPutRedfishV1SFSSAppIpAddressManagementsRequestWithDefaults() *PutRedfishV1SFSSAppIpAddressManagementsRequest`

NewPutRedfishV1SFSSAppIpAddressManagementsRequestWithDefaults instantiates a new PutRedfishV1SFSSAppIpAddressManagementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIPV4Address

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Address() []string`

GetIPV4Address returns the IPV4Address field if non-nil, zero value otherwise.

### GetIPV4AddressOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4AddressOk() (*[]string, bool)`

GetIPV4AddressOk returns a tuple with the IPV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Address

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Address(v []string)`

SetIPV4Address sets IPV4Address field to given value.


### GetIPV4Config

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Config() string`

GetIPV4Config returns the IPV4Config field if non-nil, zero value otherwise.

### GetIPV4ConfigOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4ConfigOk() (*string, bool)`

GetIPV4ConfigOk returns a tuple with the IPV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Config

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Config(v string)`

SetIPV4Config sets IPV4Config field to given value.


### GetIPV4Gateway

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Gateway() string`

GetIPV4Gateway returns the IPV4Gateway field if non-nil, zero value otherwise.

### GetIPV4GatewayOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4GatewayOk() (*string, bool)`

GetIPV4GatewayOk returns a tuple with the IPV4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Gateway

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Gateway(v string)`

SetIPV4Gateway sets IPV4Gateway field to given value.


### GetIPV4PrefixLength

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4PrefixLength() float32`

GetIPV4PrefixLength returns the IPV4PrefixLength field if non-nil, zero value otherwise.

### GetIPV4PrefixLengthOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4PrefixLengthOk() (*float32, bool)`

GetIPV4PrefixLengthOk returns a tuple with the IPV4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4PrefixLength

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4PrefixLength(v float32)`

SetIPV4PrefixLength sets IPV4PrefixLength field to given value.


### GetIPV4Route

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

GetIPV4Route returns the IPV4Route field if non-nil, zero value otherwise.

### GetIPV4RouteOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4RouteOk() (*[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool)`

GetIPV4RouteOk returns a tuple with the IPV4Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Route

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner)`

SetIPV4Route sets IPV4Route field to given value.

### HasIPV4Route

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) HasIPV4Route() bool`

HasIPV4Route returns a boolean if a field has been set.

### GetIPV6Config

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Config() string`

GetIPV6Config returns the IPV6Config field if non-nil, zero value otherwise.

### GetIPV6ConfigOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6ConfigOk() (*string, bool)`

GetIPV6ConfigOk returns a tuple with the IPV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Config

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Config(v string)`

SetIPV6Config sets IPV6Config field to given value.


### GetName

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIPV6Address

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Address() []string`

GetIPV6Address returns the IPV6Address field if non-nil, zero value otherwise.

### GetIPV6AddressOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6AddressOk() (*[]string, bool)`

GetIPV6AddressOk returns a tuple with the IPV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Address

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Address(v []string)`

SetIPV6Address sets IPV6Address field to given value.


### GetIPV6Gateway

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Gateway() string`

GetIPV6Gateway returns the IPV6Gateway field if non-nil, zero value otherwise.

### GetIPV6GatewayOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6GatewayOk() (*string, bool)`

GetIPV6GatewayOk returns a tuple with the IPV6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Gateway

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Gateway(v string)`

SetIPV6Gateway sets IPV6Gateway field to given value.


### GetIPV6PrefixLength

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6PrefixLength() float32`

GetIPV6PrefixLength returns the IPV6PrefixLength field if non-nil, zero value otherwise.

### GetIPV6PrefixLengthOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6PrefixLengthOk() (*float32, bool)`

GetIPV6PrefixLengthOk returns a tuple with the IPV6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6PrefixLength

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6PrefixLength(v float32)`

SetIPV6PrefixLength sets IPV6PrefixLength field to given value.


### GetMTU

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetMTU() float32`

GetMTU returns the MTU field if non-nil, zero value otherwise.

### GetMTUOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetMTUOk() (*float32, bool)`

GetMTUOk returns a tuple with the MTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTU

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetMTU(v float32)`

SetMTU sets MTU field to given value.

### HasMTU

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) HasMTU() bool`

HasMTU returns a boolean if a field has been set.

### GetIPV6Route

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Route() []PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner`

GetIPV6Route returns the IPV6Route field if non-nil, zero value otherwise.

### GetIPV6RouteOk

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6RouteOk() (*[]PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner, bool)`

GetIPV6RouteOk returns a tuple with the IPV6Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Route

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Route(v []PutRedfishV1SFSSAppIpAddressManagementsRequestIPV6RouteInner)`

SetIPV6Route sets IPV6Route field to given value.

### HasIPV6Route

`func (o *PutRedfishV1SFSSAppIpAddressManagementsRequest) HasIPV6Route() bool`

HasIPV6Route returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


