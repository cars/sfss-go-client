# GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identity of the Entity - Fully Qualified Name (FQN) of the host | 
**NQN** | **string** | NVMe Qualified Name (NQN) of the host | 
**TransportAddress** | **string** | IP address of the host | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**TransportType** | **string** | Supported transport types that can be used for communication with the host; possible value is TCP | 
**TREQ** | **string** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**TSAS** | **string** | Possible values are No Security and TLS | 
**RegistrationType** | **string** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **string** | Status of the TCP connection between the host and the CDC instance; possible values include Online and Offline | 
**FailureReason** | **string** | Reason for the host being offline; possible values include NONE, KATO, and Peer Closure | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner

`func NewGetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner(id string, nQN string, transportAddress string, transportAddressFamily string, transportType string, tREQ string, tSAS string, registrationType string, connectionStatus string, failureReason string, odataId string, odataType string, odataContext string, ) *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner`

NewGetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner instantiates a new GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInnerWithDefaults

`func NewGetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInnerWithDefaults() *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner`

NewGetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInnerWithDefaults instantiates a new GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetId(v string)`

SetId sets Id field to given value.


### GetNQN

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetNQN() string`

GetNQN returns the NQN field if non-nil, zero value otherwise.

### GetNQNOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetNQNOk() (*string, bool)`

GetNQNOk returns a tuple with the NQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNQN

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetNQN(v string)`

SetNQN sets NQN field to given value.


### GetTransportAddress

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetTransportType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetTREQ

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTREQ() string`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTREQOk() (*string, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetTREQ(v string)`

SetTREQ sets TREQ field to given value.


### GetTSAS

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTSAS() string`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetTSASOk() (*string, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetTSAS(v string)`

SetTSAS sets TSAS field to given value.


### GetRegistrationType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetRegistrationType() string`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetRegistrationTypeOk() (*string, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetRegistrationType(v string)`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetFailureReason

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetOdataId

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsExpandHosts200ResponseHostsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


