# IpAddressManagementGETEXPANDIpAddressManagementsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | **string** | Interface identifier | 
**Type** | **string** | Interface type; possible values include Ethernet and VLAN | 
**IPV4Address** | Pointer to **[]map[string]interface{}** | IPv4 address configured on the interface | [optional] 
**IPV4Config** | **string** | IPv4 configuration type; possible values include Manual, Automatic, and Disabled  | 
**IPV4Gateway** | **string** | IPv4 gateway address | 
**IPV4PrefixLength** | **float32** | Subnet mask | 
**IPV6Address** | Pointer to **[]map[string]interface{}** | IPv6 address configured on the interface | [optional] 
**IPV6Config** | **string** | IPv6 configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV6Gateway** | **string** | IPv6 gateway address | 
**IPV6PrefixLength** | **float32** | Subnet mask | 
**MACAddress** | **string** | MAC address of the interface | 
**MTU** | **float32** | Maximum transmission unit of the interface; default value is 1500 | 
**Name** | **string** | Friendly name configured for the interface | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewIpAddressManagementGETEXPANDIpAddressManagementsInner

`func NewIpAddressManagementGETEXPANDIpAddressManagementsInner(interface_ string, type_ string, iPV4Config string, iPV4Gateway string, iPV4PrefixLength float32, iPV6Config string, iPV6Gateway string, iPV6PrefixLength float32, mACAddress string, mTU float32, name string, odataId string, odataType string, odataContext string, ) *IpAddressManagementGETEXPANDIpAddressManagementsInner`

NewIpAddressManagementGETEXPANDIpAddressManagementsInner instantiates a new IpAddressManagementGETEXPANDIpAddressManagementsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressManagementGETEXPANDIpAddressManagementsInnerWithDefaults

`func NewIpAddressManagementGETEXPANDIpAddressManagementsInnerWithDefaults() *IpAddressManagementGETEXPANDIpAddressManagementsInner`

NewIpAddressManagementGETEXPANDIpAddressManagementsInnerWithDefaults instantiates a new IpAddressManagementGETEXPANDIpAddressManagementsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetType

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetType(v string)`

SetType sets Type field to given value.


### GetIPV4Address

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4Address() []map[string]interface{}`

GetIPV4Address returns the IPV4Address field if non-nil, zero value otherwise.

### GetIPV4AddressOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4AddressOk() (*[]map[string]interface{}, bool)`

GetIPV4AddressOk returns a tuple with the IPV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Address

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV4Address(v []map[string]interface{})`

SetIPV4Address sets IPV4Address field to given value.

### HasIPV4Address

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) HasIPV4Address() bool`

HasIPV4Address returns a boolean if a field has been set.

### GetIPV4Config

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4Config() string`

GetIPV4Config returns the IPV4Config field if non-nil, zero value otherwise.

### GetIPV4ConfigOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4ConfigOk() (*string, bool)`

GetIPV4ConfigOk returns a tuple with the IPV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Config

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV4Config(v string)`

SetIPV4Config sets IPV4Config field to given value.


### GetIPV4Gateway

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4Gateway() string`

GetIPV4Gateway returns the IPV4Gateway field if non-nil, zero value otherwise.

### GetIPV4GatewayOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4GatewayOk() (*string, bool)`

GetIPV4GatewayOk returns a tuple with the IPV4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Gateway

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV4Gateway(v string)`

SetIPV4Gateway sets IPV4Gateway field to given value.


### GetIPV4PrefixLength

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4PrefixLength() float32`

GetIPV4PrefixLength returns the IPV4PrefixLength field if non-nil, zero value otherwise.

### GetIPV4PrefixLengthOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV4PrefixLengthOk() (*float32, bool)`

GetIPV4PrefixLengthOk returns a tuple with the IPV4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4PrefixLength

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV4PrefixLength(v float32)`

SetIPV4PrefixLength sets IPV4PrefixLength field to given value.


### GetIPV6Address

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6Address() []map[string]interface{}`

GetIPV6Address returns the IPV6Address field if non-nil, zero value otherwise.

### GetIPV6AddressOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6AddressOk() (*[]map[string]interface{}, bool)`

GetIPV6AddressOk returns a tuple with the IPV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Address

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV6Address(v []map[string]interface{})`

SetIPV6Address sets IPV6Address field to given value.

### HasIPV6Address

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) HasIPV6Address() bool`

HasIPV6Address returns a boolean if a field has been set.

### GetIPV6Config

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6Config() string`

GetIPV6Config returns the IPV6Config field if non-nil, zero value otherwise.

### GetIPV6ConfigOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6ConfigOk() (*string, bool)`

GetIPV6ConfigOk returns a tuple with the IPV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Config

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV6Config(v string)`

SetIPV6Config sets IPV6Config field to given value.


### GetIPV6Gateway

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6Gateway() string`

GetIPV6Gateway returns the IPV6Gateway field if non-nil, zero value otherwise.

### GetIPV6GatewayOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6GatewayOk() (*string, bool)`

GetIPV6GatewayOk returns a tuple with the IPV6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Gateway

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV6Gateway(v string)`

SetIPV6Gateway sets IPV6Gateway field to given value.


### GetIPV6PrefixLength

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6PrefixLength() float32`

GetIPV6PrefixLength returns the IPV6PrefixLength field if non-nil, zero value otherwise.

### GetIPV6PrefixLengthOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetIPV6PrefixLengthOk() (*float32, bool)`

GetIPV6PrefixLengthOk returns a tuple with the IPV6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6PrefixLength

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetIPV6PrefixLength(v float32)`

SetIPV6PrefixLength sets IPV6PrefixLength field to given value.


### GetMACAddress

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetMACAddress() string`

GetMACAddress returns the MACAddress field if non-nil, zero value otherwise.

### GetMACAddressOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetMACAddressOk() (*string, bool)`

GetMACAddressOk returns a tuple with the MACAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMACAddress

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetMACAddress(v string)`

SetMACAddress sets MACAddress field to given value.


### GetMTU

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetMTU() float32`

GetMTU returns the MTU field if non-nil, zero value otherwise.

### GetMTUOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetMTUOk() (*float32, bool)`

GetMTUOk returns a tuple with the MTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMTU

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetMTU(v float32)`

SetMTU sets MTU field to given value.


### GetName

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetName(v string)`

SetName sets Name field to given value.


### GetOdataId

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *IpAddressManagementGETEXPANDIpAddressManagementsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


