# GetRedfishV1SFSSAppGlobalSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** | Hostname of the SFSS VM | 
**ReservedIPV4SubnetPrefix** | **string** | Reserved IPv4 subnet prefix for internal VM networking use; the default value is 172.18.x.x | 
**ReservedIPV6SubnetPrefix** | **string** | Reserved IPv6 subnet prefix for internal VM networking use; the default value is fd01::x | 
**StorageInterfaceMTU** | **float32** | Maximum transmission unit; range is from 1250 to 9000; default value is 1500 | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppGlobalSettings200Response

`func NewGetRedfishV1SFSSAppGlobalSettings200Response(hostName string, reservedIPV4SubnetPrefix string, reservedIPV6SubnetPrefix string, storageInterfaceMTU float32, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppGlobalSettings200Response`

NewGetRedfishV1SFSSAppGlobalSettings200Response instantiates a new GetRedfishV1SFSSAppGlobalSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppGlobalSettings200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppGlobalSettings200ResponseWithDefaults() *GetRedfishV1SFSSAppGlobalSettings200Response`

NewGetRedfishV1SFSSAppGlobalSettings200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppGlobalSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetReservedIPV4SubnetPrefix

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetReservedIPV4SubnetPrefix() string`

GetReservedIPV4SubnetPrefix returns the ReservedIPV4SubnetPrefix field if non-nil, zero value otherwise.

### GetReservedIPV4SubnetPrefixOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetReservedIPV4SubnetPrefixOk() (*string, bool)`

GetReservedIPV4SubnetPrefixOk returns a tuple with the ReservedIPV4SubnetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIPV4SubnetPrefix

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetReservedIPV4SubnetPrefix(v string)`

SetReservedIPV4SubnetPrefix sets ReservedIPV4SubnetPrefix field to given value.


### GetReservedIPV6SubnetPrefix

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetReservedIPV6SubnetPrefix() string`

GetReservedIPV6SubnetPrefix returns the ReservedIPV6SubnetPrefix field if non-nil, zero value otherwise.

### GetReservedIPV6SubnetPrefixOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetReservedIPV6SubnetPrefixOk() (*string, bool)`

GetReservedIPV6SubnetPrefixOk returns a tuple with the ReservedIPV6SubnetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIPV6SubnetPrefix

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetReservedIPV6SubnetPrefix(v string)`

SetReservedIPV6SubnetPrefix sets ReservedIPV6SubnetPrefix field to given value.


### GetStorageInterfaceMTU

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetStorageInterfaceMTU() float32`

GetStorageInterfaceMTU returns the StorageInterfaceMTU field if non-nil, zero value otherwise.

### GetStorageInterfaceMTUOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetStorageInterfaceMTUOk() (*float32, bool)`

GetStorageInterfaceMTUOk returns a tuple with the StorageInterfaceMTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceMTU

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetStorageInterfaceMTU(v float32)`

SetStorageInterfaceMTU sets StorageInterfaceMTU field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppGlobalSettings200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


