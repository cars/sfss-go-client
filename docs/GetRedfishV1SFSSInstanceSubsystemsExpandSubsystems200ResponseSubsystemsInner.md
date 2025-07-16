# GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identity of the Entity - Fully Qualified Name (FQN) of the subsystem | 
**NQN** | **string** | NVMe Qualified Name (NQN) of the subsystem | 
**TransportAddress** | **string** | IP address of the subsystem | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**PortId** | **float32** | Internal Port ID that is mapped to the NVMe Subsystem Port | 
**ControllerId** | **float32** | TCP controller ID, 65535 | 
**TransportType** | **string** | Supported transport types that can be used for communication with the controller; possible values are TCP, RoCE, FC | 
**SubType** | **string** | Possible values include Reserved, Discovery Service, and NVM Subsystem | 
**TREQ** | **string** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**ASQZ** | **float32** | Admin Queue Size | 
**TransportServiceId** | **string** | NVMe Transport Service Identifier as an ASCII string | 
**TSAS** | **string** | NVMe Transport Service Address SubType | 
**RcvdGenCounter** | **float32** | Generation Counter that is maintained for the last received Get Log Page command | 
**RegistrationType** | **string** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **string** | Status of the TCP connection between the subsystem and the CDC instance; possible values include Online and Offline | 
**EKType** | **string** | Entity Key Type; possible values are Unknown, Port, and TRADDR | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner

`func NewGetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner(id string, nQN string, transportAddress string, transportAddressFamily string, portId float32, controllerId float32, transportType string, subType string, tREQ string, aSQZ float32, transportServiceId string, tSAS string, rcvdGenCounter float32, registrationType string, connectionStatus string, eKType string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner`

NewGetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner instantiates a new GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInnerWithDefaults

`func NewGetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInnerWithDefaults() *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner`

NewGetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInnerWithDefaults instantiates a new GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetNQN

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetNQN() string`

GetNQN returns the NQN field if non-nil, zero value otherwise.

### GetNQNOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetNQNOk() (*string, bool)`

GetNQNOk returns a tuple with the NQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNQN

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetNQN(v string)`

SetNQN sets NQN field to given value.


### GetTransportAddress

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetPortId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetPortId() float32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetPortIdOk() (*float32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetPortId(v float32)`

SetPortId sets PortId field to given value.


### GetControllerId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetControllerId() float32`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetControllerIdOk() (*float32, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetControllerId(v float32)`

SetControllerId sets ControllerId field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetSubType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetTREQ

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTREQ() string`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTREQOk() (*string, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetTREQ(v string)`

SetTREQ sets TREQ field to given value.


### GetASQZ

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetASQZ() float32`

GetASQZ returns the ASQZ field if non-nil, zero value otherwise.

### GetASQZOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetASQZOk() (*float32, bool)`

GetASQZOk returns a tuple with the ASQZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASQZ

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetASQZ(v float32)`

SetASQZ sets ASQZ field to given value.


### GetTransportServiceId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportServiceId() string`

GetTransportServiceId returns the TransportServiceId field if non-nil, zero value otherwise.

### GetTransportServiceIdOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTransportServiceIdOk() (*string, bool)`

GetTransportServiceIdOk returns a tuple with the TransportServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportServiceId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetTransportServiceId(v string)`

SetTransportServiceId sets TransportServiceId field to given value.


### GetTSAS

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTSAS() string`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetTSASOk() (*string, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetTSAS(v string)`

SetTSAS sets TSAS field to given value.


### GetRcvdGenCounter

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetRcvdGenCounter() float32`

GetRcvdGenCounter returns the RcvdGenCounter field if non-nil, zero value otherwise.

### GetRcvdGenCounterOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetRcvdGenCounterOk() (*float32, bool)`

GetRcvdGenCounterOk returns a tuple with the RcvdGenCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvdGenCounter

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetRcvdGenCounter(v float32)`

SetRcvdGenCounter sets RcvdGenCounter field to given value.


### GetRegistrationType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetRegistrationType() string`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetRegistrationTypeOk() (*string, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetRegistrationType(v string)`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetEKType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetEKType() string`

GetEKType returns the EKType field if non-nil, zero value otherwise.

### GetEKTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetEKTypeOk() (*string, bool)`

GetEKTypeOk returns a tuple with the EKType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEKType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetEKType(v string)`

SetEKType sets EKType field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSInstanceSubsystemsExpandSubsystems200ResponseSubsystemsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


