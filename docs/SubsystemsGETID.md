# SubsystemsGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identity of the Entity - Fully Qualified Name (FQN) of the subsystem | 
**NQN** | **string** | NVMe Qualified Name (NQN) of the subsystem | 
**TransportAddress** | **string** | IP address of the subsystem | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**PortId** | **float32** | Internal Port ID that is mapped to the NVMe Subsystem Port | 
**ControllerId** | **float32** | TCP controller ID, 65535 | 
**TransportType** | **string** | Protocol used for communication with the host; possible value is TCP | 
**SubType** | **string** | Possible values include Reserved, Discovery Service, and NVM Subsystem | 
**TREQ** | **string** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**ASQZ** | **float32** | Admin Queue Size | 
**TransportServiceId** | **string** | NVMe Transport Service Identifier as an ASCII string | 
**TSAS** | **string** | NVMe Transport Service Address SubType; possible values are No Security and TLS | 
**RcvdGenCounter** | **float32** | Generation Counter that is maintained for the last received Get Log Page command | 
**RegistrationType** | **string** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **string** | Status of the TCP connection between the subsystem and the CDC instance; possible values include Online and Offline | 
**EKType** | **string** | Entity Key Type; possible values are Unknown, Port, and TRADDR | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewSubsystemsGETID

`func NewSubsystemsGETID(id string, nQN string, transportAddress string, transportAddressFamily string, portId float32, controllerId float32, transportType string, subType string, tREQ string, aSQZ float32, transportServiceId string, tSAS string, rcvdGenCounter float32, registrationType string, connectionStatus string, eKType string, odataId string, odataType string, odataContext string, ) *SubsystemsGETID`

NewSubsystemsGETID instantiates a new SubsystemsGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsystemsGETIDWithDefaults

`func NewSubsystemsGETIDWithDefaults() *SubsystemsGETID`

NewSubsystemsGETIDWithDefaults instantiates a new SubsystemsGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubsystemsGETID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubsystemsGETID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubsystemsGETID) SetId(v string)`

SetId sets Id field to given value.


### GetNQN

`func (o *SubsystemsGETID) GetNQN() string`

GetNQN returns the NQN field if non-nil, zero value otherwise.

### GetNQNOk

`func (o *SubsystemsGETID) GetNQNOk() (*string, bool)`

GetNQNOk returns a tuple with the NQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNQN

`func (o *SubsystemsGETID) SetNQN(v string)`

SetNQN sets NQN field to given value.


### GetTransportAddress

`func (o *SubsystemsGETID) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *SubsystemsGETID) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *SubsystemsGETID) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *SubsystemsGETID) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *SubsystemsGETID) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *SubsystemsGETID) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetPortId

`func (o *SubsystemsGETID) GetPortId() float32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SubsystemsGETID) GetPortIdOk() (*float32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SubsystemsGETID) SetPortId(v float32)`

SetPortId sets PortId field to given value.


### GetControllerId

`func (o *SubsystemsGETID) GetControllerId() float32`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *SubsystemsGETID) GetControllerIdOk() (*float32, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *SubsystemsGETID) SetControllerId(v float32)`

SetControllerId sets ControllerId field to given value.


### GetTransportType

`func (o *SubsystemsGETID) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *SubsystemsGETID) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *SubsystemsGETID) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetSubType

`func (o *SubsystemsGETID) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *SubsystemsGETID) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *SubsystemsGETID) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetTREQ

`func (o *SubsystemsGETID) GetTREQ() string`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *SubsystemsGETID) GetTREQOk() (*string, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *SubsystemsGETID) SetTREQ(v string)`

SetTREQ sets TREQ field to given value.


### GetASQZ

`func (o *SubsystemsGETID) GetASQZ() float32`

GetASQZ returns the ASQZ field if non-nil, zero value otherwise.

### GetASQZOk

`func (o *SubsystemsGETID) GetASQZOk() (*float32, bool)`

GetASQZOk returns a tuple with the ASQZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASQZ

`func (o *SubsystemsGETID) SetASQZ(v float32)`

SetASQZ sets ASQZ field to given value.


### GetTransportServiceId

`func (o *SubsystemsGETID) GetTransportServiceId() string`

GetTransportServiceId returns the TransportServiceId field if non-nil, zero value otherwise.

### GetTransportServiceIdOk

`func (o *SubsystemsGETID) GetTransportServiceIdOk() (*string, bool)`

GetTransportServiceIdOk returns a tuple with the TransportServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportServiceId

`func (o *SubsystemsGETID) SetTransportServiceId(v string)`

SetTransportServiceId sets TransportServiceId field to given value.


### GetTSAS

`func (o *SubsystemsGETID) GetTSAS() string`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *SubsystemsGETID) GetTSASOk() (*string, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *SubsystemsGETID) SetTSAS(v string)`

SetTSAS sets TSAS field to given value.


### GetRcvdGenCounter

`func (o *SubsystemsGETID) GetRcvdGenCounter() float32`

GetRcvdGenCounter returns the RcvdGenCounter field if non-nil, zero value otherwise.

### GetRcvdGenCounterOk

`func (o *SubsystemsGETID) GetRcvdGenCounterOk() (*float32, bool)`

GetRcvdGenCounterOk returns a tuple with the RcvdGenCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvdGenCounter

`func (o *SubsystemsGETID) SetRcvdGenCounter(v float32)`

SetRcvdGenCounter sets RcvdGenCounter field to given value.


### GetRegistrationType

`func (o *SubsystemsGETID) GetRegistrationType() string`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *SubsystemsGETID) GetRegistrationTypeOk() (*string, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *SubsystemsGETID) SetRegistrationType(v string)`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *SubsystemsGETID) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *SubsystemsGETID) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *SubsystemsGETID) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetEKType

`func (o *SubsystemsGETID) GetEKType() string`

GetEKType returns the EKType field if non-nil, zero value otherwise.

### GetEKTypeOk

`func (o *SubsystemsGETID) GetEKTypeOk() (*string, bool)`

GetEKTypeOk returns a tuple with the EKType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEKType

`func (o *SubsystemsGETID) SetEKType(v string)`

SetEKType sets EKType field to given value.


### GetOdataId

`func (o *SubsystemsGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *SubsystemsGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *SubsystemsGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *SubsystemsGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *SubsystemsGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *SubsystemsGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *SubsystemsGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *SubsystemsGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *SubsystemsGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


