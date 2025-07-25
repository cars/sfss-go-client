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

// checks if the IpAddressManagementGETID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpAddressManagementGETID{}

// IpAddressManagementGETID This model lists the interface information based on the specified interface ID.
type IpAddressManagementGETID struct {
	// Interface identifier
	Interface string `json:"Interface"`
	// IPv4 address configured on the interface
	IPV4Address []map[string]interface{} `json:"IPV4Address"`
	// Interface type; possible values include Ethernet and VLAN
	Type string `json:"Type"`
	// IPv4 configuration type; possible values include Manual, Automatic, and Disabled
	IPV4Config string `json:"IPV4Config"`
	// Friendly name for the configured interface
	Name string `json:"Name"`
	// IPv4 gateway address
	IPV4Gateway string `json:"IPV4Gateway"`
	// Subnet mask
	IPV4PrefixLength float32 `json:"IPV4PrefixLength"`
	// IPv4 static route configuration
	IPV4Route []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner `json:"IPV4Route,omitempty"`
	// IPv6 address configured on the interface
	IPV6Address []map[string]interface{} `json:"IPV6Address"`
	// IPv6 configuration type; possible values include Manual, Automatic, and Disabled
	IPV6Config string `json:"IPV6Config"`
	// IPv6 address subnet mask
	IPV6PrefixLength float32 `json:"IPV6PrefixLength"`
	// Maximum transmission unit for the interface; default value is 1500
	MTU float32 `json:"MTU"`
	// IPv6 static route configuration
	IPV6Route []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner `json:"IPV6Route,omitempty"`
	OdataId string `json:"@odata.id"`
	OdataType string `json:"@odata.type"`
	OdataContext string `json:"@odata.context"`
}

type _IpAddressManagementGETID IpAddressManagementGETID

// NewIpAddressManagementGETID instantiates a new IpAddressManagementGETID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAddressManagementGETID(interface_ string, iPV4Address []map[string]interface{}, type_ string, iPV4Config string, name string, iPV4Gateway string, iPV4PrefixLength float32, iPV6Address []map[string]interface{}, iPV6Config string, iPV6PrefixLength float32, mTU float32, odataId string, odataType string, odataContext string) *IpAddressManagementGETID {
	this := IpAddressManagementGETID{}
	this.Interface = interface_
	this.IPV4Address = iPV4Address
	this.Type = type_
	this.IPV4Config = iPV4Config
	this.Name = name
	this.IPV4Gateway = iPV4Gateway
	this.IPV4PrefixLength = iPV4PrefixLength
	this.IPV6Address = iPV6Address
	this.IPV6Config = iPV6Config
	this.IPV6PrefixLength = iPV6PrefixLength
	this.MTU = mTU
	this.OdataId = odataId
	this.OdataType = odataType
	this.OdataContext = odataContext
	return &this
}

// NewIpAddressManagementGETIDWithDefaults instantiates a new IpAddressManagementGETID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAddressManagementGETIDWithDefaults() *IpAddressManagementGETID {
	this := IpAddressManagementGETID{}
	return &this
}

// GetInterface returns the Interface field value
func (o *IpAddressManagementGETID) GetInterface() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *IpAddressManagementGETID) SetInterface(v string) {
	o.Interface = v
}

// GetIPV4Address returns the IPV4Address field value
func (o *IpAddressManagementGETID) GetIPV4Address() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.IPV4Address
}

// GetIPV4AddressOk returns a tuple with the IPV4Address field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV4AddressOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.IPV4Address, true
}

// SetIPV4Address sets field value
func (o *IpAddressManagementGETID) SetIPV4Address(v []map[string]interface{}) {
	o.IPV4Address = v
}

// GetType returns the Type field value
func (o *IpAddressManagementGETID) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IpAddressManagementGETID) SetType(v string) {
	o.Type = v
}

// GetIPV4Config returns the IPV4Config field value
func (o *IpAddressManagementGETID) GetIPV4Config() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPV4Config
}

// GetIPV4ConfigOk returns a tuple with the IPV4Config field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV4ConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV4Config, true
}

// SetIPV4Config sets field value
func (o *IpAddressManagementGETID) SetIPV4Config(v string) {
	o.IPV4Config = v
}

// GetName returns the Name field value
func (o *IpAddressManagementGETID) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpAddressManagementGETID) SetName(v string) {
	o.Name = v
}

// GetIPV4Gateway returns the IPV4Gateway field value
func (o *IpAddressManagementGETID) GetIPV4Gateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPV4Gateway
}

// GetIPV4GatewayOk returns a tuple with the IPV4Gateway field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV4GatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV4Gateway, true
}

// SetIPV4Gateway sets field value
func (o *IpAddressManagementGETID) SetIPV4Gateway(v string) {
	o.IPV4Gateway = v
}

// GetIPV4PrefixLength returns the IPV4PrefixLength field value
func (o *IpAddressManagementGETID) GetIPV4PrefixLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IPV4PrefixLength
}

// GetIPV4PrefixLengthOk returns a tuple with the IPV4PrefixLength field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV4PrefixLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV4PrefixLength, true
}

// SetIPV4PrefixLength sets field value
func (o *IpAddressManagementGETID) SetIPV4PrefixLength(v float32) {
	o.IPV4PrefixLength = v
}

// GetIPV4Route returns the IPV4Route field value if set, zero value otherwise.
func (o *IpAddressManagementGETID) GetIPV4Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner {
	if o == nil || IsNil(o.IPV4Route) {
		var ret []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner
		return ret
	}
	return o.IPV4Route
}

// GetIPV4RouteOk returns a tuple with the IPV4Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV4RouteOk() ([]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool) {
	if o == nil || IsNil(o.IPV4Route) {
		return nil, false
	}
	return o.IPV4Route, true
}

// HasIPV4Route returns a boolean if a field has been set.
func (o *IpAddressManagementGETID) HasIPV4Route() bool {
	if o != nil && !IsNil(o.IPV4Route) {
		return true
	}

	return false
}

// SetIPV4Route gets a reference to the given []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner and assigns it to the IPV4Route field.
func (o *IpAddressManagementGETID) SetIPV4Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) {
	o.IPV4Route = v
}

// GetIPV6Address returns the IPV6Address field value
func (o *IpAddressManagementGETID) GetIPV6Address() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.IPV6Address
}

// GetIPV6AddressOk returns a tuple with the IPV6Address field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV6AddressOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.IPV6Address, true
}

// SetIPV6Address sets field value
func (o *IpAddressManagementGETID) SetIPV6Address(v []map[string]interface{}) {
	o.IPV6Address = v
}

// GetIPV6Config returns the IPV6Config field value
func (o *IpAddressManagementGETID) GetIPV6Config() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPV6Config
}

// GetIPV6ConfigOk returns a tuple with the IPV6Config field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV6ConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV6Config, true
}

// SetIPV6Config sets field value
func (o *IpAddressManagementGETID) SetIPV6Config(v string) {
	o.IPV6Config = v
}

// GetIPV6PrefixLength returns the IPV6PrefixLength field value
func (o *IpAddressManagementGETID) GetIPV6PrefixLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IPV6PrefixLength
}

// GetIPV6PrefixLengthOk returns a tuple with the IPV6PrefixLength field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV6PrefixLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPV6PrefixLength, true
}

// SetIPV6PrefixLength sets field value
func (o *IpAddressManagementGETID) SetIPV6PrefixLength(v float32) {
	o.IPV6PrefixLength = v
}

// GetMTU returns the MTU field value
func (o *IpAddressManagementGETID) GetMTU() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MTU
}

// GetMTUOk returns a tuple with the MTU field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetMTUOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MTU, true
}

// SetMTU sets field value
func (o *IpAddressManagementGETID) SetMTU(v float32) {
	o.MTU = v
}

// GetIPV6Route returns the IPV6Route field value if set, zero value otherwise.
func (o *IpAddressManagementGETID) GetIPV6Route() []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner {
	if o == nil || IsNil(o.IPV6Route) {
		var ret []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner
		return ret
	}
	return o.IPV6Route
}

// GetIPV6RouteOk returns a tuple with the IPV6Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetIPV6RouteOk() ([]PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner, bool) {
	if o == nil || IsNil(o.IPV6Route) {
		return nil, false
	}
	return o.IPV6Route, true
}

// HasIPV6Route returns a boolean if a field has been set.
func (o *IpAddressManagementGETID) HasIPV6Route() bool {
	if o != nil && !IsNil(o.IPV6Route) {
		return true
	}

	return false
}

// SetIPV6Route gets a reference to the given []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner and assigns it to the IPV6Route field.
func (o *IpAddressManagementGETID) SetIPV6Route(v []PostRedfishV1SFSSAppIpAddressManagementsRequestIPV4RouteInner) {
	o.IPV6Route = v
}

// GetOdataId returns the OdataId field value
func (o *IpAddressManagementGETID) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *IpAddressManagementGETID) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *IpAddressManagementGETID) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *IpAddressManagementGETID) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataContext returns the OdataContext field value
func (o *IpAddressManagementGETID) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *IpAddressManagementGETID) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *IpAddressManagementGETID) SetOdataContext(v string) {
	o.OdataContext = v
}

func (o IpAddressManagementGETID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpAddressManagementGETID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Interface"] = o.Interface
	toSerialize["IPV4Address"] = o.IPV4Address
	toSerialize["Type"] = o.Type
	toSerialize["IPV4Config"] = o.IPV4Config
	toSerialize["Name"] = o.Name
	toSerialize["IPV4Gateway"] = o.IPV4Gateway
	toSerialize["IPV4PrefixLength"] = o.IPV4PrefixLength
	if !IsNil(o.IPV4Route) {
		toSerialize["IPV4Route"] = o.IPV4Route
	}
	toSerialize["IPV6Address"] = o.IPV6Address
	toSerialize["IPV6Config"] = o.IPV6Config
	toSerialize["IPV6PrefixLength"] = o.IPV6PrefixLength
	toSerialize["MTU"] = o.MTU
	if !IsNil(o.IPV6Route) {
		toSerialize["IPV6Route"] = o.IPV6Route
	}
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.type"] = o.OdataType
	toSerialize["@odata.context"] = o.OdataContext
	return toSerialize, nil
}

func (o *IpAddressManagementGETID) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Interface",
		"IPV4Address",
		"Type",
		"IPV4Config",
		"Name",
		"IPV4Gateway",
		"IPV4PrefixLength",
		"IPV6Address",
		"IPV6Config",
		"IPV6PrefixLength",
		"MTU",
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

	varIpAddressManagementGETID := _IpAddressManagementGETID{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpAddressManagementGETID)

	if err != nil {
		return err
	}

	*o = IpAddressManagementGETID(varIpAddressManagementGETID)

	return err
}

type NullableIpAddressManagementGETID struct {
	value *IpAddressManagementGETID
	isSet bool
}

func (v NullableIpAddressManagementGETID) Get() *IpAddressManagementGETID {
	return v.value
}

func (v *NullableIpAddressManagementGETID) Set(val *IpAddressManagementGETID) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddressManagementGETID) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddressManagementGETID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddressManagementGETID(val *IpAddressManagementGETID) *NullableIpAddressManagementGETID {
	return &NullableIpAddressManagementGETID{value: val, isSet: true}
}

func (v NullableIpAddressManagementGETID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddressManagementGETID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


