# GlobalPoliciesENUMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoningPolicy** | **[]map[string]interface{}** | Global zoning policy that allows all hosts to communicate with all subsystems; disabled by default | 
**NameServerEntityPurgeTOV** | **[]map[string]interface{}** | Timeout value to remove endpoints (hosts or subsystems) that have lost connection with the CDC  | 

## Methods

### NewGlobalPoliciesENUMS

`func NewGlobalPoliciesENUMS(zoningPolicy []map[string]interface{}, nameServerEntityPurgeTOV []map[string]interface{}, ) *GlobalPoliciesENUMS`

NewGlobalPoliciesENUMS instantiates a new GlobalPoliciesENUMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPoliciesENUMSWithDefaults

`func NewGlobalPoliciesENUMSWithDefaults() *GlobalPoliciesENUMS`

NewGlobalPoliciesENUMSWithDefaults instantiates a new GlobalPoliciesENUMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoningPolicy

`func (o *GlobalPoliciesENUMS) GetZoningPolicy() []map[string]interface{}`

GetZoningPolicy returns the ZoningPolicy field if non-nil, zero value otherwise.

### GetZoningPolicyOk

`func (o *GlobalPoliciesENUMS) GetZoningPolicyOk() (*[]map[string]interface{}, bool)`

GetZoningPolicyOk returns a tuple with the ZoningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoningPolicy

`func (o *GlobalPoliciesENUMS) SetZoningPolicy(v []map[string]interface{})`

SetZoningPolicy sets ZoningPolicy field to given value.


### GetNameServerEntityPurgeTOV

`func (o *GlobalPoliciesENUMS) GetNameServerEntityPurgeTOV() []map[string]interface{}`

GetNameServerEntityPurgeTOV returns the NameServerEntityPurgeTOV field if non-nil, zero value otherwise.

### GetNameServerEntityPurgeTOVOk

`func (o *GlobalPoliciesENUMS) GetNameServerEntityPurgeTOVOk() (*[]map[string]interface{}, bool)`

GetNameServerEntityPurgeTOVOk returns a tuple with the NameServerEntityPurgeTOV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServerEntityPurgeTOV

`func (o *GlobalPoliciesENUMS) SetNameServerEntityPurgeTOV(v []map[string]interface{})`

SetNameServerEntityPurgeTOV sets NameServerEntityPurgeTOV field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


