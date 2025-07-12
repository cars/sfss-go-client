# PostRedfishV1SFSSAppIpAddressManagementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IPV4Address** | **[]string** | Interface IPv4 address  | 
**IPV4Config** | **string** | IP address configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV4Gateway** | **string** | IPv4 gateway address | 
**IPV4PrefixLength** | **float32** | Subnet mask | 
**Name** | Pointer to **string** | Friendly name for the interface | [optional] 
**IPV4Route** | Pointer to [**[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner**](PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner.md) | IPv4 static route configuration | [optional] 
**IPV6Config** | **string** | IPv6 configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV6Address** | **[]string** | Interface IPv6 address | 
**IPV6Gateway** | **string** | IPv6 gateway address | 
**IPV6PrefixLength** | **float32** | IPv6 address subnet mask | 
**ParentInterface** | **string** | The parent interface on which the VLAN subinterface is created; possible values include ens160, ens192, ens224, ens256, ens161, ens193, ens225, ens257, ens162, and ens194  | 
**VlanId** | **float32** | VLAN identifier | 
**MTU** | Pointer to **float32** | Maximum transmission unit for the interface; default value is 1500 | [optional] 
**IPV6Route** | Pointer to [**[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner**](PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner.md) | IPv6 static route configuration | [optional] 

## Methods

### NewPostRedfishV1SFSSAppIpAddressManagementsRequest

`func NewPostRedfishV1SFSSAppIpAddressManagementsRequest(iPV4Address []string, iPV4Config string, iPV4Gateway string, iPV4PrefixLength float32, iPV6Config string, iPV6Address []string, iPV6Gateway string, iPV6PrefixLength float32, parentInterface string, vlanId float32, ) *PostRedfishV1SFSSAppIpAddressManagementsRequest`

NewPostRedfishV1SFSSAppIpAddressManagementsRequest instantiates a new PostRedfishV1SFSSAppIpAddressManagementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppIpAddressManagementsRequestWithDefaults

`func NewPostRedfishV1SFSSAppIpAddressManagementsRequestWithDefaults() *PostRedfishV1SFSSAppIpAddressManagementsRequest`

NewPostRedfishV1SFSSAppIpAddressManagementsRequestWithDefaults instantiates a new PostRedfishV1SFSSAppIpAddressManagementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIPV4Address

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Address() []string`

GetIPV4Address returns the IPV4Address field if non-nil, zero value otherwise.

### GetIPV4AddressOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4AddressOk() (*[]string, bool)`

GetIPV4AddressOk returns a tuple with the IPV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Address

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Address(v []string)`

SetIPV4Address sets IPV4Address field to given value.


### GetIPV4Config

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Config() string`

GetIPV4Config returns the IPV4Config field if non-nil, zero value otherwise.

### GetIPV4ConfigOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4ConfigOk() (*string, bool)`

GetIPV4ConfigOk returns a tuple with the IPV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Config

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Config(v string)`

SetIPV4Config sets IPV4Config field to given value.


### GetIPV4Gateway

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Gateway() string`

GetIPV4Gateway returns the IPV4Gateway field if non-nil, zero value otherwise.

### GetIPV4GatewayOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4GatewayOk() (*string, bool)`

GetIPV4GatewayOk returns a tuple with the IPV4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Gateway

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Gateway(v string)`

SetIPV4Gateway sets IPV4Gateway field to given value.


### GetIPV4PrefixLength

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4PrefixLength() float32`

GetIPV4PrefixLength returns the IPV4PrefixLength field if non-nil, zero value otherwise.

### GetIPV4PrefixLengthOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4PrefixLengthOk() (*float32, bool)`

GetIPV4PrefixLengthOk returns a tuple with the IPV4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4PrefixLength

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4PrefixLength(v float32)`

SetIPV4PrefixLength sets IPV4PrefixLength field to given value.


### GetName

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIPV4Route

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

GetIPV4Route returns the IPV4Route field if non-nil, zero value otherwise.

### GetIPV4RouteOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV4RouteOk() (*[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool)`

GetIPV4RouteOk returns a tuple with the IPV4Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Route

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV4Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner)`

SetIPV4Route sets IPV4Route field to given value.

### HasIPV4Route

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) HasIPV4Route() bool`

HasIPV4Route returns a boolean if a field has been set.

### GetIPV6Config

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Config() string`

GetIPV6Config returns the IPV6Config field if non-nil, zero value otherwise.

### GetIPV6ConfigOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6ConfigOk() (*string, bool)`

GetIPV6ConfigOk returns a tuple with the IPV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Config

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Config(v string)`

SetIPV6Config sets IPV6Config field to given value.


### GetIPV6Address

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Address() []string`

GetIPV6Address returns the IPV6Address field if non-nil, zero value otherwise.

### GetIPV6AddressOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6AddressOk() (*[]string, bool)`

GetIPV6AddressOk returns a tuple with the IPV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Address

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Address(v []string)`

SetIPV6Address sets IPV6Address field to given value.


### GetIPV6Gateway

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Gateway() string`

GetIPV6Gateway returns the IPV6Gateway field if non-nil, zero value otherwise.

### GetIPV6GatewayOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6GatewayOk() (*string, bool)`

GetIPV6GatewayOk returns a tuple with the IPV6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Gateway

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Gateway(v string)`

SetIPV6Gateway sets IPV6Gateway field to given value.


### GetIPV6PrefixLength

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6PrefixLength() float32`

GetIPV6PrefixLength returns the IPV6PrefixLength field if non-nil, zero value otherwise.

### GetIPV6PrefixLengthOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6PrefixLengthOk() (*float32, bool)`

GetIPV6PrefixLengthOk returns a tuple with the IPV6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6PrefixLength

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6PrefixLength(v float32)`

SetIPV6PrefixLength sets IPV6PrefixLength field to given value.


### GetParentInterface

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetParentInterface() string`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetParentInterfaceOk() (*string, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetParentInterface(v string)`

SetParentInterface sets ParentInterface field to given value.


### GetVlanId

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.


### GetMTU

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetMTU() float32`

GetMTU returns the MTU field if non-nil, zero value otherwise.

### GetMTUOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetMTUOk() (*float32, bool)`

GetMTUOk returns a tuple with the MTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTU

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetMTU(v float32)`

SetMTU sets MTU field to given value.

### HasMTU

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) HasMTU() bool`

HasMTU returns a boolean if a field has been set.

### GetIPV6Route

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

GetIPV6Route returns the IPV6Route field if non-nil, zero value otherwise.

### GetIPV6RouteOk

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) GetIPV6RouteOk() (*[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool)`

GetIPV6RouteOk returns a tuple with the IPV6Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Route

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) SetIPV6Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner)`

SetIPV6Route sets IPV6Route field to given value.

### HasIPV6Route

`func (o *PostRedfishV1SFSSAppIpAddressManagementsRequest) HasIPV6Route() bool`

HasIPV6Route returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


