# GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportAddressFamily** | **[]string** | IP address family; possible values include IPv4 and IPv6 | 
**EKType** | **[]string** | Entity Key Type; possible values are Unknown, Port, and TRADDR | 
**FailureReason** | **[]string** | Reason for the host being offline; possible values include NONE, KATO, and Peer Closure | 
**RegistrationType** | **[]string** | Type of endpoint registration; possible values include Implicit, Explicit, Pull, and Manual | 
**ConnectionStatus** | **[]string** | Status of the TCP connection between the host and the CDC instance; possible values include Online and Offline | 
**TREQ** | **[]string** | Transport Request Type (TREQ); possible values include Secure channel not specified, Secure channel required, and Secure channel not required | 
**TransportType** | **[]string** | Supported transport types that can be used for communication with the controller; possible values are TCP, RoCE, FC | 
**TSAS** | **[]string** | Possible values are No Security and TLS | 

## Methods

### NewGetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response

`func NewGetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response(transportAddressFamily []string, eKType []string, failureReason []string, registrationType []string, connectionStatus []string, tREQ []string, transportType []string, tSAS []string, ) *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response`

NewGetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response instantiates a new GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHttpsIPRedfishV1SFSSInstanceHostsEnums200ResponseWithDefaults

`func NewGetHttpsIPRedfishV1SFSSInstanceHostsEnums200ResponseWithDefaults() *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response`

NewGetHttpsIPRedfishV1SFSSInstanceHostsEnums200ResponseWithDefaults instantiates a new GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportAddressFamily

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTransportAddressFamily() []string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTransportAddressFamilyOk() (*[]string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetTransportAddressFamily(v []string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetEKType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetEKType() []string`

GetEKType returns the EKType field if non-nil, zero value otherwise.

### GetEKTypeOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetEKTypeOk() (*[]string, bool)`

GetEKTypeOk returns a tuple with the EKType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEKType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetEKType(v []string)`

SetEKType sets EKType field to given value.


### GetFailureReason

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetFailureReason() []string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetFailureReasonOk() (*[]string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetFailureReason(v []string)`

SetFailureReason sets FailureReason field to given value.


### GetRegistrationType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetRegistrationType() []string`

GetRegistrationType returns the RegistrationType field if non-nil, zero value otherwise.

### GetRegistrationTypeOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetRegistrationTypeOk() (*[]string, bool)`

GetRegistrationTypeOk returns a tuple with the RegistrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetRegistrationType(v []string)`

SetRegistrationType sets RegistrationType field to given value.


### GetConnectionStatus

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetConnectionStatus() []string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetConnectionStatusOk() (*[]string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetConnectionStatus(v []string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetTREQ

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTREQ() []string`

GetTREQ returns the TREQ field if non-nil, zero value otherwise.

### GetTREQOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTREQOk() (*[]string, bool)`

GetTREQOk returns a tuple with the TREQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTREQ

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetTREQ(v []string)`

SetTREQ sets TREQ field to given value.


### GetTransportType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTransportType() []string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTransportTypeOk() (*[]string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetTransportType(v []string)`

SetTransportType sets TransportType field to given value.


### GetTSAS

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTSAS() []string`

GetTSAS returns the TSAS field if non-nil, zero value otherwise.

### GetTSASOk

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) GetTSASOk() (*[]string, bool)`

GetTSASOk returns a tuple with the TSAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSAS

`func (o *GetHttpsIPRedfishV1SFSSInstanceHostsEnums200Response) SetTSAS(v []string)`

SetTSAS sets TSAS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


