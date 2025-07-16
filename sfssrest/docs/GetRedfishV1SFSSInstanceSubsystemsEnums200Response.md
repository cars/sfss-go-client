# GetRedfishV1SFSSInstanceSubsystemsEnums200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportAddressFamily** | **[]string** | IP address family; possible values include IPv4 and IPv6 | 
**EKType** | **[]string** | Entity Key Type; possible values are Unknown, Port, and TRADDR | 
**FailureReason** | **[]string** | Connection Failure Reason - Possible Values are NONE, Peer Closure,Dial Failure, KATO | 
**RegistrationType** | **[]string** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **[]string** | Status of the TCP connection between the subsystem and the CDC instance; possible values include Online and Offline | 
**SubType** | **[]string** | Possible values include Reserved, Discovery Service, and NVM Subsystem | 
**TREQ** | **[]string** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**TransportType** | **[]string** | Supported transport types that can be used for communication with the controller; possible values are TCP, RoCE, FC | 
**TSAS** | **[]string** | NVMe Transport Service Address SubType | 

## Methods

### NewGetRedfishV1SFSSInstanceSubsystemsEnums200Response

`func NewGetRedfishV1SFSSInstanceSubsystemsEnums200Response(transportAddressFamily []string, eKType []string, failureReason []string, registrationType []string, connectionStatus []string, subType []string, tREQ []string, transportType []string, tSAS []string, ) *GetRedfishV1SFSSInstanceSubsystemsEnums200Response`

NewGetRedfishV1SFSSInstanceSubsystemsEnums200Response instantiates a new GetRedfishV1SFSSInstanceSubsystemsEnums200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSInstanceSubsystemsEnums200ResponseWithDefaults

`func NewGetRedfishV1SFSSInstanceSubsystemsEnums200ResponseWithDefaults() *GetRedfishV1SFSSInstanceSubsystemsEnums200Response`

NewGetRedfishV1SFSSInstanceSubsystemsEnums200ResponseWithDefaults instantiates a new GetRedfishV1SFSSInstanceSubsystemsEnums200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTransportAddressFamily() []string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTransportAddressFamilyOk() (*[]string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetTransportAddressFamily(v []string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetEKType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetEKType() []string`

GetEKType returns the EKType field if non-nil, zero value otherwise.

### GetEKTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetEKTypeOk() (*[]string, bool)`

GetEKTypeOk returns a tuple with the EKType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEKType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetEKType(v []string)`

SetEKType sets EKType field to given value.


### GetFailureReason

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetFailureReason() []string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetFailureReasonOk() (*[]string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetFailureReason(v []string)`

SetFailureReason sets FailureReason field to given value.


### GetRegistrationType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetRegistrationType() []string`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetRegistrationTypeOk() (*[]string, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetRegistrationType(v []string)`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetConnectionStatus() []string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetConnectionStatusOk() (*[]string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetConnectionStatus(v []string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetSubType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetSubType() []string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetSubTypeOk() (*[]string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetSubType(v []string)`

SetSubType sets SubType field to given value.


### GetTREQ

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTREQ() []string`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTREQOk() (*[]string, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetTREQ(v []string)`

SetTREQ sets TREQ field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTransportType() []string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTransportTypeOk() (*[]string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetTransportType(v []string)`

SetTransportType sets TransportType field to given value.


### GetTSAS

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTSAS() []string`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) GetTSASOk() (*[]string, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *GetRedfishV1SFSSInstanceSubsystemsEnums200Response) SetTSAS(v []string)`

SetTSAS sets TSAS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


