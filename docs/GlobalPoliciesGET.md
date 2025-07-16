# GlobalPoliciesGET

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoningPolicy** | **string** | Global zoning policy that allows all hosts to communicate with all subsystems; disabled by default | 
**NameServerEntityPurgeTOV** | **string** | Timeout value to remove endpoints (hosts or subsystems) that have lost connection with the CDC  | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGlobalPoliciesGET

`func NewGlobalPoliciesGET(zoningPolicy string, nameServerEntityPurgeTOV string, odataId string, odataType string, odataContext string, ) *GlobalPoliciesGET`

NewGlobalPoliciesGET instantiates a new GlobalPoliciesGET object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPoliciesGETWithDefaults

`func NewGlobalPoliciesGETWithDefaults() *GlobalPoliciesGET`

NewGlobalPoliciesGETWithDefaults instantiates a new GlobalPoliciesGET object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoningPolicy

`func (o *GlobalPoliciesGET) GetZoningPolicy() string`

GetZoningPolicy returns the ZoningPolicy field if non-nil, zero value otherwise.

### GetZoningPolicyOk

`func (o *GlobalPoliciesGET) GetZoningPolicyOk() (*string, bool)`

GetZoningPolicyOk returns a tuple with the ZoningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoningPolicy

`func (o *GlobalPoliciesGET) SetZoningPolicy(v string)`

SetZoningPolicy sets ZoningPolicy field to given value.


### GetNameServerEntityPurgeTOV

`func (o *GlobalPoliciesGET) GetNameServerEntityPurgeTOV() string`

GetNameServerEntityPurgeTOV returns the NameServerEntityPurgeTOV field if non-nil, zero value otherwise.

### GetNameServerEntityPurgeTOVOk

`func (o *GlobalPoliciesGET) GetNameServerEntityPurgeTOVOk() (*string, bool)`

GetNameServerEntityPurgeTOVOk returns a tuple with the NameServerEntityPurgeTOV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServerEntityPurgeTOV

`func (o *GlobalPoliciesGET) SetNameServerEntityPurgeTOV(v string)`

SetNameServerEntityPurgeTOV sets NameServerEntityPurgeTOV field to given value.


### GetOdataId

`func (o *GlobalPoliciesGET) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GlobalPoliciesGET) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GlobalPoliciesGET) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GlobalPoliciesGET) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GlobalPoliciesGET) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GlobalPoliciesGET) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GlobalPoliciesGET) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GlobalPoliciesGET) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GlobalPoliciesGET) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


