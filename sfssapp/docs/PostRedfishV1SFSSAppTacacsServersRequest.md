# PostRedfishV1SFSSAppTacacsServersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIp** | **string** | IP address of the TACACS+ server; SFSS supports only IPv4 communication with remote servers | 
**ServerPass** | **string** | TACACS+ server password | 

## Methods

### NewPostRedfishV1SFSSAppTacacsServersRequest

`func NewPostRedfishV1SFSSAppTacacsServersRequest(serverIp string, serverPass string, ) *PostRedfishV1SFSSAppTacacsServersRequest`

NewPostRedfishV1SFSSAppTacacsServersRequest instantiates a new PostRedfishV1SFSSAppTacacsServersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppTacacsServersRequestWithDefaults

`func NewPostRedfishV1SFSSAppTacacsServersRequestWithDefaults() *PostRedfishV1SFSSAppTacacsServersRequest`

NewPostRedfishV1SFSSAppTacacsServersRequestWithDefaults instantiates a new PostRedfishV1SFSSAppTacacsServersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIp

`func (o *PostRedfishV1SFSSAppTacacsServersRequest) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *PostRedfishV1SFSSAppTacacsServersRequest) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *PostRedfishV1SFSSAppTacacsServersRequest) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.


### GetServerPass

`func (o *PostRedfishV1SFSSAppTacacsServersRequest) GetServerPass() string`

GetServerPass returns the ServerPass field if non-nil, zero value otherwise.

### GetServerPassOk

`func (o *PostRedfishV1SFSSAppTacacsServersRequest) GetServerPassOk() (*string, bool)`

GetServerPassOk returns a tuple with the ServerPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPass

`func (o *PostRedfishV1SFSSAppTacacsServersRequest) SetServerPass(v string)`

SetServerPass sets ServerPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


