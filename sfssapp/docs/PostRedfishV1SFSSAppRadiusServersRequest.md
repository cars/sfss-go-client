# PostRedfishV1SFSSAppRadiusServersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIp** | **string** | IP address of the RADIUS server; SFSS supports only IPv4 communication with remote servers | 
**ServerPass** | **string** | Password to access the RADIUS server | 

## Methods

### NewPostRedfishV1SFSSAppRadiusServersRequest

`func NewPostRedfishV1SFSSAppRadiusServersRequest(serverIp string, serverPass string, ) *PostRedfishV1SFSSAppRadiusServersRequest`

NewPostRedfishV1SFSSAppRadiusServersRequest instantiates a new PostRedfishV1SFSSAppRadiusServersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppRadiusServersRequestWithDefaults

`func NewPostRedfishV1SFSSAppRadiusServersRequestWithDefaults() *PostRedfishV1SFSSAppRadiusServersRequest`

NewPostRedfishV1SFSSAppRadiusServersRequestWithDefaults instantiates a new PostRedfishV1SFSSAppRadiusServersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIp

`func (o *PostRedfishV1SFSSAppRadiusServersRequest) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *PostRedfishV1SFSSAppRadiusServersRequest) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *PostRedfishV1SFSSAppRadiusServersRequest) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.


### GetServerPass

`func (o *PostRedfishV1SFSSAppRadiusServersRequest) GetServerPass() string`

GetServerPass returns the ServerPass field if non-nil, zero value otherwise.

### GetServerPassOk

`func (o *PostRedfishV1SFSSAppRadiusServersRequest) GetServerPassOk() (*string, bool)`

GetServerPassOk returns a tuple with the ServerPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPass

`func (o *PostRedfishV1SFSSAppRadiusServersRequest) SetServerPass(v string)`

SetServerPass sets ServerPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


