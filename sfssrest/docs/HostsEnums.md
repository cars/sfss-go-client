# HostsEnums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportAddressFamily** | **[]map[string]interface{}** | IP address family; possible values include IPv4 and IPv6 | 
**EKType** | **[]map[string]interface{}** | Entity Key Type; possible values are Unknown, Port, and TRADDR | 
**FailureReason** | **[]map[string]interface{}** | Reason for the host being offline; possible values include NONE, KATO, and Peer Closure | 
**RegistrationType** | **[]map[string]interface{}** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **[]map[string]interface{}** | Status of the TCP connection between the host and the CDC instance; possible values include Online and Offline | 
**TREQ** | **[]map[string]interface{}** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**TransportType** | **[]map[string]interface{}** | Protocol used for communication with the host; possible value is TCP | 
**TSAS** | **[]map[string]interface{}** | Possible values are No Security and TLS | 

## Methods

### NewHostsEnums

`func NewHostsEnums(transportAddressFamily []map[string]interface{}, eKType []map[string]interface{}, failureReason []map[string]interface{}, registrationType []map[string]interface{}, connectionStatus []map[string]interface{}, tREQ []map[string]interface{}, transportType []map[string]interface{}, tSAS []map[string]interface{}, ) *HostsEnums`

NewHostsEnums instantiates a new HostsEnums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsEnumsWithDefaults

`func NewHostsEnumsWithDefaults() *HostsEnums`

NewHostsEnumsWithDefaults instantiates a new HostsEnums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportAddressFamily

`func (o *HostsEnums) GetTransportAddressFamily() []map[string]interface{}`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *HostsEnums) GetTransportAddressFamilyOk() (*[]map[string]interface{}, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *HostsEnums) SetTransportAddressFamily(v []map[string]interface{})`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetEKType

`func (o *HostsEnums) GetEKType() []map[string]interface{}`

GetEKType returns the EKType field if non-nil, zero value otherwise.

### GetEKTypeOk

`func (o *HostsEnums) GetEKTypeOk() (*[]map[string]interface{}, bool)`

GetEKTypeOk returns a tuple with the EKType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEKType

`func (o *HostsEnums) SetEKType(v []map[string]interface{})`

SetEKType sets EKType field to given value.


### GetFailureReason

`func (o *HostsEnums) GetFailureReason() []map[string]interface{}`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HostsEnums) GetFailureReasonOk() (*[]map[string]interface{}, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HostsEnums) SetFailureReason(v []map[string]interface{})`

SetFailureReason sets FailureReason field to given value.


### GetRegistrationType

`func (o *HostsEnums) GetRegistrationType() []map[string]interface{}`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *HostsEnums) GetRegistrationTypeOk() (*[]map[string]interface{}, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *HostsEnums) SetRegistrationType(v []map[string]interface{})`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *HostsEnums) GetConnectionStatus() []map[string]interface{}`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *HostsEnums) GetConnectionStatusOk() (*[]map[string]interface{}, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *HostsEnums) SetConnectionStatus(v []map[string]interface{})`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetTREQ

`func (o *HostsEnums) GetTREQ() []map[string]interface{}`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *HostsEnums) GetTREQOk() (*[]map[string]interface{}, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *HostsEnums) SetTREQ(v []map[string]interface{})`

SetTREQ sets TREQ field to given value.


### GetTransportType

`func (o *HostsEnums) GetTransportType() []map[string]interface{}`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *HostsEnums) GetTransportTypeOk() (*[]map[string]interface{}, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *HostsEnums) SetTransportType(v []map[string]interface{})`

SetTransportType sets TransportType field to given value.


### GetTSAS

`func (o *HostsEnums) GetTSAS() []map[string]interface{}`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *HostsEnums) GetTSASOk() (*[]map[string]interface{}, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *HostsEnums) SetTSAS(v []map[string]interface{})`

SetTSAS sets TSAS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


