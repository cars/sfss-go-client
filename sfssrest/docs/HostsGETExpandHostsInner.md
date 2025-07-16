# HostsGETExpandHostsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identity of the Entity - Fully Qualified Name (FQN) of the host | 
**NQN** | **string** | NVMe Qualified Name (NQN) of the host | 
**TransportAddress** | **string** | IP address of the host | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**TransportType** | **string** | Protocol used for communication with the host; possible value is TCP | 
**TREQ** | **string** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**TSAS** | **string** | Possible values are No Security and TLS | 
**RegistrationType** | **string** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **string** | Status of the TCP connection between the host and the CDC instance; possible values include Online and Offline | 
**FailureReason** | **string** | Reason for the host being offline; possible values include NONE, KATO, and Peer Closure | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewHostsGETExpandHostsInner

`func NewHostsGETExpandHostsInner(id string, nQN string, transportAddress string, transportAddressFamily string, transportType string, tREQ string, tSAS string, registrationType string, connectionStatus string, failureReason string, odataId string, odataType string, odataContext string, ) *HostsGETExpandHostsInner`

NewHostsGETExpandHostsInner instantiates a new HostsGETExpandHostsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsGETExpandHostsInnerWithDefaults

`func NewHostsGETExpandHostsInnerWithDefaults() *HostsGETExpandHostsInner`

NewHostsGETExpandHostsInnerWithDefaults instantiates a new HostsGETExpandHostsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostsGETExpandHostsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostsGETExpandHostsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostsGETExpandHostsInner) SetId(v string)`

SetId sets Id field to given value.


### GetNQN

`func (o *HostsGETExpandHostsInner) GetNQN() string`

GetNQN returns the NQN field if non-nil, zero value otherwise.

### GetNQNOk

`func (o *HostsGETExpandHostsInner) GetNQNOk() (*string, bool)`

GetNQNOk returns a tuple with the NQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNQN

`func (o *HostsGETExpandHostsInner) SetNQN(v string)`

SetNQN sets NQN field to given value.


### GetTransportAddress

`func (o *HostsGETExpandHostsInner) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *HostsGETExpandHostsInner) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *HostsGETExpandHostsInner) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *HostsGETExpandHostsInner) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *HostsGETExpandHostsInner) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *HostsGETExpandHostsInner) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetTransportType

`func (o *HostsGETExpandHostsInner) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *HostsGETExpandHostsInner) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *HostsGETExpandHostsInner) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetTREQ

`func (o *HostsGETExpandHostsInner) GetTREQ() string`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *HostsGETExpandHostsInner) GetTREQOk() (*string, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *HostsGETExpandHostsInner) SetTREQ(v string)`

SetTREQ sets TREQ field to given value.


### GetTSAS

`func (o *HostsGETExpandHostsInner) GetTSAS() string`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *HostsGETExpandHostsInner) GetTSASOk() (*string, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *HostsGETExpandHostsInner) SetTSAS(v string)`

SetTSAS sets TSAS field to given value.


### GetRegistrationType

`func (o *HostsGETExpandHostsInner) GetRegistrationType() string`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *HostsGETExpandHostsInner) GetRegistrationTypeOk() (*string, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *HostsGETExpandHostsInner) SetRegistrationType(v string)`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *HostsGETExpandHostsInner) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *HostsGETExpandHostsInner) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *HostsGETExpandHostsInner) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetFailureReason

`func (o *HostsGETExpandHostsInner) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HostsGETExpandHostsInner) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HostsGETExpandHostsInner) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetOdataId

`func (o *HostsGETExpandHostsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *HostsGETExpandHostsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *HostsGETExpandHostsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *HostsGETExpandHostsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *HostsGETExpandHostsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *HostsGETExpandHostsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *HostsGETExpandHostsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *HostsGETExpandHostsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *HostsGETExpandHostsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


