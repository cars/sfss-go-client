# SubsystemsEnums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportAddressFamily** | **[]map[string]interface{}** | IP address family; possible values include IPv4 and IPv6 | 
**EKType** | **[]map[string]interface{}** | Entity Key Type; possible values are Unknown, Port, and TRADDR | 
**FailureReason** | **[]map[string]interface{}** | Connection Failure Reason - Possible Values are NONE, Peer Closure,Dial Failure, KATO | 
**RegistrationType** | **[]map[string]interface{}** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **[]map[string]interface{}** | Status of the TCP connection between the subsystem and the CDC instance; possible values include Online and Offline | 
**SubType** | **[]map[string]interface{}** | Possible values include Reserved, Discovery Service, and NVM Subsystem | 
**TREQ** | **[]map[string]interface{}** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**TransportType** | **[]map[string]interface{}** | Protocol used for communication with the host; possible values are TCP, RoCE, FC | 
**TSAS** | **[]map[string]interface{}** | NVMe Transport Service Address SubType; possible values are No Security and TLS | 

## Methods

### NewSubsystemsEnums

`func NewSubsystemsEnums(transportAddressFamily []map[string]interface{}, eKType []map[string]interface{}, failureReason []map[string]interface{}, registrationType []map[string]interface{}, connectionStatus []map[string]interface{}, subType []map[string]interface{}, tREQ []map[string]interface{}, transportType []map[string]interface{}, tSAS []map[string]interface{}, ) *SubsystemsEnums`

NewSubsystemsEnums instantiates a new SubsystemsEnums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsystemsEnumsWithDefaults

`func NewSubsystemsEnumsWithDefaults() *SubsystemsEnums`

NewSubsystemsEnumsWithDefaults instantiates a new SubsystemsEnums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportAddressFamily

`func (o *SubsystemsEnums) GetTransportAddressFamily() []map[string]interface{}`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *SubsystemsEnums) GetTransportAddressFamilyOk() (*[]map[string]interface{}, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *SubsystemsEnums) SetTransportAddressFamily(v []map[string]interface{})`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetEKType

`func (o *SubsystemsEnums) GetEKType() []map[string]interface{}`

GetEKType returns the EKType field if non-nil, zero value otherwise.

### GetEKTypeOk

`func (o *SubsystemsEnums) GetEKTypeOk() (*[]map[string]interface{}, bool)`

GetEKTypeOk returns a tuple with the EKType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEKType

`func (o *SubsystemsEnums) SetEKType(v []map[string]interface{})`

SetEKType sets EKType field to given value.


### GetFailureReason

`func (o *SubsystemsEnums) GetFailureReason() []map[string]interface{}`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *SubsystemsEnums) GetFailureReasonOk() (*[]map[string]interface{}, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *SubsystemsEnums) SetFailureReason(v []map[string]interface{})`

SetFailureReason sets FailureReason field to given value.


### GetRegistrationType

`func (o *SubsystemsEnums) GetRegistrationType() []map[string]interface{}`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *SubsystemsEnums) GetRegistrationTypeOk() (*[]map[string]interface{}, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *SubsystemsEnums) SetRegistrationType(v []map[string]interface{})`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *SubsystemsEnums) GetConnectionStatus() []map[string]interface{}`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *SubsystemsEnums) GetConnectionStatusOk() (*[]map[string]interface{}, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *SubsystemsEnums) SetConnectionStatus(v []map[string]interface{})`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetSubType

`func (o *SubsystemsEnums) GetSubType() []map[string]interface{}`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *SubsystemsEnums) GetSubTypeOk() (*[]map[string]interface{}, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *SubsystemsEnums) SetSubType(v []map[string]interface{})`

SetSubType sets SubType field to given value.


### GetTREQ

`func (o *SubsystemsEnums) GetTREQ() []map[string]interface{}`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *SubsystemsEnums) GetTREQOk() (*[]map[string]interface{}, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *SubsystemsEnums) SetTREQ(v []map[string]interface{})`

SetTREQ sets TREQ field to given value.


### GetTransportType

`func (o *SubsystemsEnums) GetTransportType() []map[string]interface{}`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *SubsystemsEnums) GetTransportTypeOk() (*[]map[string]interface{}, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *SubsystemsEnums) SetTransportType(v []map[string]interface{})`

SetTransportType sets TransportType field to given value.


### GetTSAS

`func (o *SubsystemsEnums) GetTSAS() []map[string]interface{}`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *SubsystemsEnums) GetTSASOk() (*[]map[string]interface{}, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *SubsystemsEnums) SetTSAS(v []map[string]interface{})`

SetTSAS sets TSAS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


