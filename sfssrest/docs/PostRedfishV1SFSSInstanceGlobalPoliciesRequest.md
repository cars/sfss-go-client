# PostRedfishV1SFSSInstanceGlobalPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoningPolicy** | **string** | Global zoning policy that allows all hosts to communicate with all subsystems; disabled by default | 
**NameServerEntityPurgeTOV** | **string** | Timeout value to remove endpoints (hosts and subsystems) that have lost connection with the CDC | 

## Methods

### NewPostRedfishV1SFSSInstanceGlobalPoliciesRequest

`func NewPostRedfishV1SFSSInstanceGlobalPoliciesRequest(zoningPolicy string, nameServerEntityPurgeTOV string, ) *PostRedfishV1SFSSInstanceGlobalPoliciesRequest`

NewPostRedfishV1SFSSInstanceGlobalPoliciesRequest instantiates a new PostRedfishV1SFSSInstanceGlobalPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedfishV1SFSSInstanceGlobalPoliciesRequestWithDefaults

`func NewPostRedfishV1SFSSInstanceGlobalPoliciesRequestWithDefaults() *PostRedfishV1SFSSInstanceGlobalPoliciesRequest`

NewPostRedfishV1SFSSInstanceGlobalPoliciesRequestWithDefaults instantiates a new PostRedfishV1SFSSInstanceGlobalPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoningPolicy

`func (o *PostRedfishV1SFSSInstanceGlobalPoliciesRequest) GetZoningPolicy() string`

GetZoningPolicy returns the ZoningPolicy field if non-nil, zero value otherwise.

### GetZoningPolicyOk

`func (o *PostRedfishV1SFSSInstanceGlobalPoliciesRequest) GetZoningPolicyOk() (*string, bool)`

GetZoningPolicyOk returns a tuple with the ZoningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoningPolicy

`func (o *PostRedfishV1SFSSInstanceGlobalPoliciesRequest) SetZoningPolicy(v string)`

SetZoningPolicy sets ZoningPolicy field to given value.


### GetNameServerEntityPurgeTOV

`func (o *PostRedfishV1SFSSInstanceGlobalPoliciesRequest) GetNameServerEntityPurgeTOV() string`

GetNameServerEntityPurgeTOV returns the NameServerEntityPurgeTOV field if non-nil, zero value otherwise.

### GetNameServerEntityPurgeTOVOk

`func (o *PostRedfishV1SFSSInstanceGlobalPoliciesRequest) GetNameServerEntityPurgeTOVOk() (*string, bool)`

GetNameServerEntityPurgeTOVOk returns a tuple with the NameServerEntityPurgeTOV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServerEntityPurgeTOV

`func (o *PostRedfishV1SFSSInstanceGlobalPoliciesRequest) SetNameServerEntityPurgeTOV(v string)`

SetNameServerEntityPurgeTOV sets NameServerEntityPurgeTOV field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


