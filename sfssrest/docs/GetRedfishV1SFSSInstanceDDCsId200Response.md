# GetRedfishV1SFSSInstanceDDCsId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identity of the Entity - Fully Qualified Name (FQN) of the Direct Discovery Controller(DDC) | 
**TransportAddress** | **string** | IP address of the DDC | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**PortId** | **float32** | Port on which the DDC listens for CDC to connect to | 
**TransportType** | **string** | Supported transport types that can be used for communication with the DDC; possible value is TCP and RoCE | 
**Activate** | **bool** | Activation of DDC that initiates a pull registration from CDC | 
**ConfigType** | **string** | Configuration Type of DDC; possible values are Manual and KickStart | 
**ConnectionStatus** | **string** | Status of the TCP connection between the DDC and the CDC instance. Possible values include Online/In-Progess:Connecting/Offline | 
**FailureReason** | **string** | Reason for the DDC being offline | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSInstanceDDCsId200Response

`func NewGetRedfishV1SFSSInstanceDDCsId200Response(id string, transportAddress string, transportAddressFamily string, portId float32, transportType string, activate bool, configType string, connectionStatus string, failureReason string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSInstanceDDCsId200Response`

NewGetRedfishV1SFSSInstanceDDCsId200Response instantiates a new GetRedfishV1SFSSInstanceDDCsId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSInstanceDDCsId200ResponseWithDefaults

`func NewGetRedfishV1SFSSInstanceDDCsId200ResponseWithDefaults() *GetRedfishV1SFSSInstanceDDCsId200Response`

NewGetRedfishV1SFSSInstanceDDCsId200ResponseWithDefaults instantiates a new GetRedfishV1SFSSInstanceDDCsId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetId(v string)`

SetId sets Id field to given value.


### GetTransportAddress

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetPortId

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetPortId() float32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetPortIdOk() (*float32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetPortId(v float32)`

SetPortId sets PortId field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetActivate

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetActivate(v bool)`

SetActivate sets Activate field to given value.


### GetConfigType

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetFailureReason

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSInstanceDDCsId200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


