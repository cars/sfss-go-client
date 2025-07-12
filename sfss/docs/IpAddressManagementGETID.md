# IpAddressManagementGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | **string** | Interface identifier | 
**IPV4Address** | **[]map[string]interface{}** | IPv4 address configured on the interface | 
**Type** | **string** | Interface type; possible values include Ethernet and VLAN | 
**IPV4Config** | **string** | IPv4 configuration type; possible values include Manual, Automatic, and Disabled | 
**Name** | **string** | Friendly name for the configured interface | 
**IPV4Gateway** | **string** | IPv4 gateway address | 
**IPV4PrefixLength** | **float32** | Subnet mask | 
**IPV4Route** | Pointer to [**[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner**](PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner.md) | IPv4 static route configuration | [optional] 
**IPV6Address** | **[]map[string]interface{}** | IPv6 address configured on the interface | 
**IPV6Config** | **string** | IPv6 configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV6PrefixLength** | **float32** | IPv6 address subnet mask | 
**MTU** | **float32** | Maximum transmission unit for the interface; default value is 1500 | 
**IPV6Route** | Pointer to [**[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner**](PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner.md) | IPv6 static route configuration | [optional] 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewIpAddressManagementGETID

`func NewIpAddressManagementGETID(interface_ string, iPV4Address []map[string]interface{}, type_ string, iPV4Config string, name string, iPV4Gateway string, iPV4PrefixLength float32, iPV6Address []map[string]interface{}, iPV6Config string, iPV6PrefixLength float32, mTU float32, odataId string, odataType string, odataContext string, ) *IpAddressManagementGETID`

NewIpAddressManagementGETID instantiates a new IpAddressManagementGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressManagementGETIDWithDefaults

`func NewIpAddressManagementGETIDWithDefaults() *IpAddressManagementGETID`

NewIpAddressManagementGETIDWithDefaults instantiates a new IpAddressManagementGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *IpAddressManagementGETID) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IpAddressManagementGETID) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IpAddressManagementGETID) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetIPV4Address

`func (o *IpAddressManagementGETID) GetIPV4Address() []map[string]interface{}`

GetIPV4Address returns the IPV4Address field if non-nil, zero value otherwise.

### GetIPV4AddressOk

`func (o *IpAddressManagementGETID) GetIPV4AddressOk() (*[]map[string]interface{}, bool)`

GetIPV4AddressOk returns a tuple with the IPV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Address

`func (o *IpAddressManagementGETID) SetIPV4Address(v []map[string]interface{})`

SetIPV4Address sets IPV4Address field to given value.


### GetType

`func (o *IpAddressManagementGETID) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpAddressManagementGETID) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpAddressManagementGETID) SetType(v string)`

SetType sets Type field to given value.


### GetIPV4Config

`func (o *IpAddressManagementGETID) GetIPV4Config() string`

GetIPV4Config returns the IPV4Config field if non-nil, zero value otherwise.

### GetIPV4ConfigOk

`func (o *IpAddressManagementGETID) GetIPV4ConfigOk() (*string, bool)`

GetIPV4ConfigOk returns a tuple with the IPV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Config

`func (o *IpAddressManagementGETID) SetIPV4Config(v string)`

SetIPV4Config sets IPV4Config field to given value.


### GetName

`func (o *IpAddressManagementGETID) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpAddressManagementGETID) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpAddressManagementGETID) SetName(v string)`

SetName sets Name field to given value.


### GetIPV4Gateway

`func (o *IpAddressManagementGETID) GetIPV4Gateway() string`

GetIPV4Gateway returns the IPV4Gateway field if non-nil, zero value otherwise.

### GetIPV4GatewayOk

`func (o *IpAddressManagementGETID) GetIPV4GatewayOk() (*string, bool)`

GetIPV4GatewayOk returns a tuple with the IPV4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Gateway

`func (o *IpAddressManagementGETID) SetIPV4Gateway(v string)`

SetIPV4Gateway sets IPV4Gateway field to given value.


### GetIPV4PrefixLength

`func (o *IpAddressManagementGETID) GetIPV4PrefixLength() float32`

GetIPV4PrefixLength returns the IPV4PrefixLength field if non-nil, zero value otherwise.

### GetIPV4PrefixLengthOk

`func (o *IpAddressManagementGETID) GetIPV4PrefixLengthOk() (*float32, bool)`

GetIPV4PrefixLengthOk returns a tuple with the IPV4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4PrefixLength

`func (o *IpAddressManagementGETID) SetIPV4PrefixLength(v float32)`

SetIPV4PrefixLength sets IPV4PrefixLength field to given value.


### GetIPV4Route

`func (o *IpAddressManagementGETID) GetIPV4Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

GetIPV4Route returns the IPV4Route field if non-nil, zero value otherwise.

### GetIPV4RouteOk

`func (o *IpAddressManagementGETID) GetIPV4RouteOk() (*[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool)`

GetIPV4RouteOk returns a tuple with the IPV4Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Route

`func (o *IpAddressManagementGETID) SetIPV4Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner)`

SetIPV4Route sets IPV4Route field to given value.

### HasIPV4Route

`func (o *IpAddressManagementGETID) HasIPV4Route() bool`

HasIPV4Route returns a boolean if a field has been set.

### GetIPV6Address

`func (o *IpAddressManagementGETID) GetIPV6Address() []map[string]interface{}`

GetIPV6Address returns the IPV6Address field if non-nil, zero value otherwise.

### GetIPV6AddressOk

`func (o *IpAddressManagementGETID) GetIPV6AddressOk() (*[]map[string]interface{}, bool)`

GetIPV6AddressOk returns a tuple with the IPV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Address

`func (o *IpAddressManagementGETID) SetIPV6Address(v []map[string]interface{})`

SetIPV6Address sets IPV6Address field to given value.


### GetIPV6Config

`func (o *IpAddressManagementGETID) GetIPV6Config() string`

GetIPV6Config returns the IPV6Config field if non-nil, zero value otherwise.

### GetIPV6ConfigOk

`func (o *IpAddressManagementGETID) GetIPV6ConfigOk() (*string, bool)`

GetIPV6ConfigOk returns a tuple with the IPV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Config

`func (o *IpAddressManagementGETID) SetIPV6Config(v string)`

SetIPV6Config sets IPV6Config field to given value.


### GetIPV6PrefixLength

`func (o *IpAddressManagementGETID) GetIPV6PrefixLength() float32`

GetIPV6PrefixLength returns the IPV6PrefixLength field if non-nil, zero value otherwise.

### GetIPV6PrefixLengthOk

`func (o *IpAddressManagementGETID) GetIPV6PrefixLengthOk() (*float32, bool)`

GetIPV6PrefixLengthOk returns a tuple with the IPV6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6PrefixLength

`func (o *IpAddressManagementGETID) SetIPV6PrefixLength(v float32)`

SetIPV6PrefixLength sets IPV6PrefixLength field to given value.


### GetMTU

`func (o *IpAddressManagementGETID) GetMTU() float32`

GetMTU returns the MTU field if non-nil, zero value otherwise.

### GetMTUOk

`func (o *IpAddressManagementGETID) GetMTUOk() (*float32, bool)`

GetMTUOk returns a tuple with the MTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTU

`func (o *IpAddressManagementGETID) SetMTU(v float32)`

SetMTU sets MTU field to given value.


### GetIPV6Route

`func (o *IpAddressManagementGETID) GetIPV6Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner`

GetIPV6Route returns the IPV6Route field if non-nil, zero value otherwise.

### GetIPV6RouteOk

`func (o *IpAddressManagementGETID) GetIPV6RouteOk() (*[]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool)`

GetIPV6RouteOk returns a tuple with the IPV6Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Route

`func (o *IpAddressManagementGETID) SetIPV6Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner)`

SetIPV6Route sets IPV6Route field to given value.

### HasIPV6Route

`func (o *IpAddressManagementGETID) HasIPV6Route() bool`

HasIPV6Route returns a boolean if a field has been set.

### GetOdataId

`func (o *IpAddressManagementGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *IpAddressManagementGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *IpAddressManagementGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *IpAddressManagementGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *IpAddressManagementGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *IpAddressManagementGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *IpAddressManagementGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *IpAddressManagementGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *IpAddressManagementGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


