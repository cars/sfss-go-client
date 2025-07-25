/*
SmartFabric Storage Software Application REST APIs

REST APIs used for managing SFSS application are captured in this section.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sfss

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetRedfishV1SFSSAppIpAddressManagementsInterface200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRedfishV1SFSSAppIpAddressManagementsInterface200Response{}

// GetRedfishV1SFSSAppIpAddressManagementsInterface200Response 
type GetRedfishV1SFSSAppIpAddressManagementsInterface200Response struct {
	// Interface identifier
	Interface string `json:"Interface"`
	// Interface type; possible values include Ethernet and VLAN
	Type string `json:"Type"`
	// IPv4 configuration type; possible values include Manual, Automatic, and Disabled
	IPV4Config string `json:"IPV4Config"`
	// IPv4 static route configuration
	IPV4Route []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner `json:"IPV4Route,omitempty"`
	// IPv6 configuration type; possible values include Manual, Automatic, and Disabled
	IPV6Config string `json:"IPV6Config"`
	// IPv6 static route configuration
	IPV6Route []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner `json:"IPV6Route,omitempty"`
	// Friendly name for the interface
	Name *string `json:"Name,omitempty"`
	// Ethernet interface on which the VLAN interface is created
	ParentInterface string `json:"ParentInterface"`
	// VLAN ID
	VlanId float32 `json:"VlanId"`
	OdataId string `json:"@odata.id"`
	OdataType string `json:"@odata.type"`
	OdataContext string `json:"@odata.context"`
}

type _GetRedfishV1SFSSAppIpAddressManagementsInterface200Response GetRedfishV1SFSSAppIpAddressManagementsInterface200Response

// NewGetRedfishV1SFSSAppIpAddressManagementsInterface200Response instantiates a new GetRedfishV1SFSSAppIpAddressManagementsInterface200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedfishV1SFSSAppIpAddressManagementsInterface200Response(interface_ string, type_ string, iPV4Config string, iPV6Config string, parentInterface string, vlanId float32, odataId string, odataType string, odataContext string) *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response {
	this := GetRedfishV1SFSSAppIpAddressManagementsInterface200Response{}
	this.Interface = interface_
	this.Type = type_
	this.IPV4Config = iPV4Config
	this.IPV6Config = iPV6Config
	this.ParentInterface = parentInterface
	this.VlanId = vlanId
	this.OdataId = odataId
	this.OdataType = odataType
	this.OdataContext = odataContext
	return &this
}

// NewGetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppIpAddressManagementsInterface200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseWithDefaults() *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response {
	this := GetRedfishV1SFSSAppIpAddressManagementsInterface200Response{}
	return &this
}

// GetInterface returns the Interface field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetInterface() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetInterface(v string) {
	o.Interface = v
}

// GetType returns the Type field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetType(v string) {
	o.Type = v
}

// GetIPV4Config returns the IPV4Config field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4Config() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPV4Config
}

// GetIPV4ConfigOk returns a tuple with the IPV4Config field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4ConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV4Config, true
}

// SetIPV4Config sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV4Config(v string) {
	o.IPV4Config = v
}

// GetIPV4Route returns the IPV4Route field value if set, zero value otherwise.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4Route() []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner {
	if o == nil || IsNil(o.IPV4Route) {
		var ret []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner
		return ret
	}
	return o.IPV4Route
}

// GetIPV4RouteOk returns a tuple with the IPV4Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV4RouteOk() ([]GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner, bool) {
	if o == nil || IsNil(o.IPV4Route) {
		return nil, false
	}
	return o.IPV4Route, true
}

// HasIPV4Route returns a boolean if a field has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) HasIPV4Route() bool {
	if o != nil && !IsNil(o.IPV4Route) {
		return true
	}

	return false
}

// SetIPV4Route gets a reference to the given []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner and assigns it to the IPV4Route field.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV4Route(v []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner) {
	o.IPV4Route = v
}

// GetIPV6Config returns the IPV6Config field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6Config() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPV6Config
}

// GetIPV6ConfigOk returns a tuple with the IPV6Config field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6ConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV6Config, true
}

// SetIPV6Config sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV6Config(v string) {
	o.IPV6Config = v
}

// GetIPV6Route returns the IPV6Route field value if set, zero value otherwise.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6Route() []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner {
	if o == nil || IsNil(o.IPV6Route) {
		var ret []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner
		return ret
	}
	return o.IPV6Route
}

// GetIPV6RouteOk returns a tuple with the IPV6Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetIPV6RouteOk() ([]GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner, bool) {
	if o == nil || IsNil(o.IPV6Route) {
		return nil, false
	}
	return o.IPV6Route, true
}

// HasIPV6Route returns a boolean if a field has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) HasIPV6Route() bool {
	if o != nil && !IsNil(o.IPV6Route) {
		return true
	}

	return false
}

// SetIPV6Route gets a reference to the given []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner and assigns it to the IPV6Route field.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetIPV6Route(v []GetRedfishV1SFSSAppIpAddressManagementsInterface200ResponseIPV4RouteInner) {
	o.IPV6Route = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetName(v string) {
	o.Name = &v
}

// GetParentInterface returns the ParentInterface field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetParentInterface() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentInterface
}

// GetParentInterfaceOk returns a tuple with the ParentInterface field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetParentInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentInterface, true
}

// SetParentInterface sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetParentInterface(v string) {
	o.ParentInterface = v
}

// GetVlanId returns the VlanId field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetVlanId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetVlanIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetVlanId(v float32) {
	o.VlanId = v
}

// GetOdataId returns the OdataId field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataContext returns the OdataContext field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) SetOdataContext(v string) {
	o.OdataContext = v
}

func (o GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Interface"] = o.Interface
	toSerialize["Type"] = o.Type
	toSerialize["IPV4Config"] = o.IPV4Config
	if !IsNil(o.IPV4Route) {
		toSerialize["IPV4Route"] = o.IPV4Route
	}
	toSerialize["IPV6Config"] = o.IPV6Config
	if !IsNil(o.IPV6Route) {
		toSerialize["IPV6Route"] = o.IPV6Route
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	toSerialize["ParentInterface"] = o.ParentInterface
	toSerialize["VlanId"] = o.VlanId
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.type"] = o.OdataType
	toSerialize["@odata.context"] = o.OdataContext
	return toSerialize, nil
}

func (o *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Interface",
		"Type",
		"IPV4Config",
		"IPV6Config",
		"ParentInterface",
		"VlanId",
		"@odata.id",
		"@odata.type",
		"@odata.context",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetRedfishV1SFSSAppIpAddressManagementsInterface200Response := _GetRedfishV1SFSSAppIpAddressManagementsInterface200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRedfishV1SFSSAppIpAddressManagementsInterface200Response)

	if err != nil {
		return err
	}

	*o = GetRedfishV1SFSSAppIpAddressManagementsInterface200Response(varGetRedfishV1SFSSAppIpAddressManagementsInterface200Response)

	return err
}

type NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response struct {
	value *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response
	isSet bool
}

func (v NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response) Get() *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response {
	return v.value
}

func (v *NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response) Set(val *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response(val *GetRedfishV1SFSSAppIpAddressManagementsInterface200Response) *NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response {
	return &NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response{value: val, isSet: true}
}

func (v NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedfishV1SFSSAppIpAddressManagementsInterface200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


