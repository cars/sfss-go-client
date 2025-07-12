# GetRedfishV1SFSSAppIpAddressManagementsInterface200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | **string** | Interface identifier | 
**Type** | **string** | Interface type; possible values include Ethernet and VLAN | 
**IPV4Config** | **string** | IPv4 configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV4Route** | Pointer to [**[]GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner**](GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner.md) | IPv4 static route configuration | [optional] 
**IPV6Config** | **string** | IPv6 configuration type; possible values include Manual, Automatic, and Disabled | 
**IPV6Route** | Pointer to [**[]GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner**](GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner.md) | IPv6 static route configuration | [optional] 
**Name** | Pointer to **string** | Friendly name for the interface | [optional] 
**ParentInterface** | **string** | Ethernet interface on which the VLAN interface is created | 
**VlanId** | **float32** | VLAN ID | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppIpAddressManagementsInterface200Response

`func NewGetRedfishV1SFSSAppIpAddressManagementsInterface200Response(interface_ string, type_ string, iPV4Config string, iPV6Config string, parentInterface string, vlanId float32, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response`

NewGetRedfishV1SFSSAppIpAddressManagementsInterface200Response instantiates a new GetRedfishV1SFSSAppIpAddressManagementsInterface200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseWithDefaults() *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response`

NewGetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppIpAddressManagementsInterface200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetType

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetType(v string)`

SetType sets Type field to given value.


### GetIPV4Config

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4Config() string`

GetIPV4Config returns the IPV4Config field if non-nil, zero value otherwise.

### GetIPV4ConfigOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4ConfigOk() (*string, bool)`

GetIPV4ConfigOk returns a tuple with the IPV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Config

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV4Config(v string)`

SetIPV4Config sets IPV4Config field to given value.


### GetIPV4Route

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4Route() []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner`

GetIPV4Route returns the IPV4Route field if non-nil, zero value otherwise.

### GetIPV4RouteOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4RouteOk() (*[]GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner, bool)`

GetIPV4RouteOk returns a tuple with the IPV4Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV4Route

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV4Route(v []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner)`

SetIPV4Route sets IPV4Route field to given value.

### HasIPV4Route

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) HasIPV4Route() bool`

HasIPV4Route returns a boolean if a field has been set.

### GetIPV6Config

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6Config() string`

GetIPV6Config returns the IPV6Config field if non-nil, zero value otherwise.

### GetIPV6ConfigOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6ConfigOk() (*string, bool)`

GetIPV6ConfigOk returns a tuple with the IPV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Config

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV6Config(v string)`

SetIPV6Config sets IPV6Config field to given value.


### GetIPV6Route

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6Route() []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner`

GetIPV6Route returns the IPV6Route field if non-nil, zero value otherwise.

### GetIPV6RouteOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6RouteOk() (*[]GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner, bool)`

GetIPV6RouteOk returns a tuple with the IPV6Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPV6Route

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV6Route(v []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner)`

SetIPV6Route sets IPV6Route field to given value.

### HasIPV6Route

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) HasIPV6Route() bool`

HasIPV6Route returns a boolean if a field has been set.

### GetName

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentInterface

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetParentInterface() string`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetParentInterfaceOk() (*string, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetParentInterface(v string)`

SetParentInterface sets ParentInterface field to given value.


### GetVlanId

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


