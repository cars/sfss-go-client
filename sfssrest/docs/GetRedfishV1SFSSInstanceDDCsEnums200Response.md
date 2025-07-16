# GetRedfishV1SFSSInstanceDDCsEnums200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportAddressFamily** | **[]string** | IP address family; possible values include IPv4 and IPv6 | 
**ConfigType** | **[]string** | Configuration Type of DDC; possible values are Manual and KickStart | 
**FailureReason** | **[]string** | Connection Failure Reason - Possible Values are NONE, Peer Closure,Dial Failure, KATO | 
**ConnectionStatus** | **[]string** | Status of the TCP connection between the DDC and the CDC instance; possible values include Online, In-Progess:Connecting, Offline | 
**TransportType** | **[]string** | Supported transport types that can be used for communication with the DDC; possible value is TCP and RoCE | 

## Methods

### NewGetRedfishV1SFSSInstanceDDCsEnums200Response

`func NewGetRedfishV1SFSSInstanceDDCsEnums200Response(transportAddressFamily []string, configType []string, failureReason []string, connectionStatus []string, transportType []string, ) *GetRedfishV1SFSSInstanceDDCsEnums200Response`

NewGetRedfishV1SFSSInstanceDDCsEnums200Response instantiates a new GetRedfishV1SFSSInstanceDDCsEnums200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSInstanceDDCsEnums200ResponseWithDefaults

`func NewGetRedfishV1SFSSInstanceDDCsEnums200ResponseWithDefaults() *GetRedfishV1SFSSInstanceDDCsEnums200Response`

NewGetRedfishV1SFSSInstanceDDCsEnums200ResponseWithDefaults instantiates a new GetRedfishV1SFSSInstanceDDCsEnums200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetTransportAddressFamily() []string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetTransportAddressFamilyOk() (*[]string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) SetTransportAddressFamily(v []string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetConfigType

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetConfigType() []string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetConfigTypeOk() (*[]string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) SetConfigType(v []string)`

SetConfigType sets ConfigType field to given value.


### GetFailureReason

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetFailureReason() []string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetFailureReasonOk() (*[]string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) SetFailureReason(v []string)`

SetFailureReason sets FailureReason field to given value.


### GetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetConnectionStatus() []string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetConnectionStatusOk() (*[]string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) SetConnectionStatus(v []string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetTransportType() []string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) GetTransportTypeOk() (*[]string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSInstanceDDCsEnums200Response) SetTransportType(v []string)`

SetTransportType sets TransportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


