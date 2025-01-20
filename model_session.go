/*
QRZ XML Logbook Data Specification

This document describes the interface specification for access to QRZ's XML subscription data service. This service provides real-time access to information from the QRZ.COM servers and databases. Access to this service requires user authentication through the use of a valid QRZ.COM username and password. While any QRZ user may login to the service, an active QRZ Logbook Data subscription is required to access most of its features. Non-subscriber access limits the data fields that are returned and is primarily intended for testing and troubleshooting purposes only. A description of subscription plans and rates is available on the [QRZ website](http://www.qrz.com/i/subscriptions.html).

API version: 1.34
Contact: va7stv@qrz.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qrz

import (
	"encoding/json"
)

// checks if the Session type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Session{}

// Session struct for Session
type Session struct {
	// a valid user session key. A user session is established whenever a session key is returned. Any response from the server that does not contain the Key element indicates that no valid session exists and that a re-login is required to continue.
	Key *string `json:"Key,omitempty" xml:"Key"`
	// Number of lookups performed by this user in the current 24 hour period
	Count *float32 `json:"Count,omitempty" xml:"Count"`
	// time and date that the users subscription will expire - or - \"non-subscriber\"
	SubExp *string `json:"SubExp,omitempty" xml:"SubExp"`
	// Time stamp for this message
	GMTime *string `json:"GMTime,omitempty" xml:"GMTime"`
	// An informational message for the user
	Message *string `json:"Message,omitempty" xml:"Message"`
	// XML system error message
	Error *string `json:"Error,omitempty" xml:"Error"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession() *Session {
	this := Session{}
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Session) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Session) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Session) SetKey(v string) {
	o.Key = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Session) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Session) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *Session) SetCount(v float32) {
	o.Count = &v
}

// GetSubExp returns the SubExp field value if set, zero value otherwise.
func (o *Session) GetSubExp() string {
	if o == nil || IsNil(o.SubExp) {
		var ret string
		return ret
	}
	return *o.SubExp
}

// GetSubExpOk returns a tuple with the SubExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetSubExpOk() (*string, bool) {
	if o == nil || IsNil(o.SubExp) {
		return nil, false
	}
	return o.SubExp, true
}

// HasSubExp returns a boolean if a field has been set.
func (o *Session) HasSubExp() bool {
	if o != nil && !IsNil(o.SubExp) {
		return true
	}

	return false
}

// SetSubExp gets a reference to the given string and assigns it to the SubExp field.
func (o *Session) SetSubExp(v string) {
	o.SubExp = &v
}

// GetGMTime returns the GMTime field value if set, zero value otherwise.
func (o *Session) GetGMTime() string {
	if o == nil || IsNil(o.GMTime) {
		var ret string
		return ret
	}
	return *o.GMTime
}

// GetGMTimeOk returns a tuple with the GMTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetGMTimeOk() (*string, bool) {
	if o == nil || IsNil(o.GMTime) {
		return nil, false
	}
	return o.GMTime, true
}

// HasGMTime returns a boolean if a field has been set.
func (o *Session) HasGMTime() bool {
	if o != nil && !IsNil(o.GMTime) {
		return true
	}

	return false
}

// SetGMTime gets a reference to the given string and assigns it to the GMTime field.
func (o *Session) SetGMTime(v string) {
	o.GMTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Session) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Session) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Session) SetMessage(v string) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Session) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Session) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *Session) SetError(v string) {
	o.Error = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Session) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if !IsNil(o.SubExp) {
		toSerialize["SubExp"] = o.SubExp
	}
	if !IsNil(o.GMTime) {
		toSerialize["GMTime"] = o.GMTime
	}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
