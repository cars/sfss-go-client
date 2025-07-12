# PostRedfishV1SFSSAppAuthenticationSequenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationSequence** | **[]string** | Defines a set of authentication servers (local and remote) configured in SFSS and the order in which SFSS looks up user information in these servers | 

## Methods

### NewPostRedfishV1SFSSAppAuthenticationSequenceRequest

`func NewPostRedfishV1SFSSAppAuthenticationSequenceRequest(authenticationSequence []string, ) *PostRedfishV1SFSSAppAuthenticationSequenceRequest`

NewPostRedfishV1SFSSAppAuthenticationSequenceRequest instantiates a new PostRedfishV1SFSSAppAuthenticationSequenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSAppAuthenticationSequenceRequestWithDefaults

`func NewPostRedfishV1SFSSAppAuthenticationSequenceRequestWithDefaults() *PostRedfishV1SFSSAppAuthenticationSequenceRequest`

NewPostRedfishV1SFSSAppAuthenticationSequenceRequestWithDefaults instantiates a new PostRedfishV1SFSSAppAuthenticationSequenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationSequence

`func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) GetAuthenticationSequence() []string`

GetAuthenticationSequence returns the AuthenticationSequence field if non-nil, zero value otherwise.

### GetAuthenticationSequenceOk

`func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) GetAuthenticationSequenceOk() (*[]string, bool)`

GetAuthenticationSequenceOk returns a tuple with the AuthenticationSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSequence

`func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) SetAuthenticationSequence(v []string)`

SetAuthenticationSequence sets AuthenticationSequence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


