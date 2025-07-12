# PostRedfishV1SFSSAppGlobalSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** | Hostname of the SFSS VM | 
**ReservedIPV4SubnetPrefix** | **string** | Reserved IPv4 subnet prefix for internal VM networking use; the default value is 172.18.x.x | 
**ReservedIPV6SubnetPrefix** | **string** | Reserved IPv6 subnet prefix for internal VM networking use; the default value is fd01::x | 
**StorageInterfaceMTU** | **float32** | Maximum transmission unit; range is from 1250 to 9000; default value is 1500  | 

## Methods

### NewPostRedfishV1SFSSAppGlobalSettingsRequest

`func NewPostRedfishV1SFSSAppGlobalSettingsRequest(hostName string, reservedIPV4SubnetPrefix string, reservedIPV6SubnetPrefix string, storageInterfaceMTU float32, ) *PostRedfishV1SFSSAppGlobalSettingsRequest`

NewPostRedfishV1SFSSAppGlobalSettingsRequest instantiates a new PostRedfishV1SFSSAppGlobalSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppGlobalSettingsRequestWithDefaults

`func NewPostRedfishV1SFSSAppGlobalSettingsRequestWithDefaults() *PostRedfishV1SFSSAppGlobalSettingsRequest`

NewPostRedfishV1SFSSAppGlobalSettingsRequestWithDefaults instantiates a new PostRedfishV1SFSSAppGlobalSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetReservedIPV4SubnetPrefix

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetReservedIPV4SubnetPrefix() string`

GetReservedIPV4SubnetPrefix returns the ReservedIPV4SubnetPrefix field if non-nil, zero value otherwise.

### GetReservedIPV4SubnetPrefixOk

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetReservedIPV4SubnetPrefixOk() (*string, bool)`

GetReservedIPV4SubnetPrefixOk returns a tuple with the ReservedIPV4SubnetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIPV4SubnetPrefix

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) SetReservedIPV4SubnetPrefix(v string)`

SetReservedIPV4SubnetPrefix sets ReservedIPV4SubnetPrefix field to given value.


### GetReservedIPV6SubnetPrefix

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetReservedIPV6SubnetPrefix() string`

GetReservedIPV6SubnetPrefix returns the ReservedIPV6SubnetPrefix field if non-nil, zero value otherwise.

### GetReservedIPV6SubnetPrefixOk

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetReservedIPV6SubnetPrefixOk() (*string, bool)`

GetReservedIPV6SubnetPrefixOk returns a tuple with the ReservedIPV6SubnetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIPV6SubnetPrefix

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) SetReservedIPV6SubnetPrefix(v string)`

SetReservedIPV6SubnetPrefix sets ReservedIPV6SubnetPrefix field to given value.


### GetStorageInterfaceMTU

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetStorageInterfaceMTU() float32`

GetStorageInterfaceMTU returns the StorageInterfaceMTU field if non-nil, zero value otherwise.

### GetStorageInterfaceMTUOk

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) GetStorageInterfaceMTUOk() (*float32, bool)`

GetStorageInterfaceMTUOk returns a tuple with the StorageInterfaceMTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceMTU

`func (o *PostRedfishV1SFSSAppGlobalSettingsRequest) SetStorageInterfaceMTU(v float32)`

SetStorageInterfaceMTU sets StorageInterfaceMTU field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


