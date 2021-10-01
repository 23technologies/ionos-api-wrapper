/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// UserPropertiesPost struct for UserPropertiesPost
type UserPropertiesPost struct {
	// first name of the user
	Firstname *string `json:"firstname,omitempty"`
	// last name of the user
	Lastname *string `json:"lastname,omitempty"`
	// email address of the user
	Email *string `json:"email,omitempty"`
	// indicates if the user has admin rights or not
	Administrator *bool `json:"administrator,omitempty"`
	// indicates if secure authentication should be forced on the user or not
	ForceSecAuth *bool `json:"forceSecAuth,omitempty"`
	// indicates if secure authentication is active for the user or not
	SecAuthActive *bool `json:"secAuthActive,omitempty"`
	// User password
	Password *string `json:"password,omitempty"`
	// indicates if the user is active
	Active *bool `json:"active,omitempty"`
}



// GetFirstname returns the Firstname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserPropertiesPost) GetFirstname() *string {
	if o == nil {
		return nil
	}


	return o.Firstname

}

// GetFirstnameOk returns a tuple with the Firstname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Firstname, true
}

// SetFirstname sets field value
func (o *UserPropertiesPost) SetFirstname(v string) {


	o.Firstname = &v

}

// HasFirstname returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}


// GetLastname returns the Lastname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserPropertiesPost) GetLastname() *string {
	if o == nil {
		return nil
	}


	return o.Lastname

}

// GetLastnameOk returns a tuple with the Lastname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Lastname, true
}

// SetLastname sets field value
func (o *UserPropertiesPost) SetLastname(v string) {


	o.Lastname = &v

}

// HasLastname returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}


// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserPropertiesPost) GetEmail() *string {
	if o == nil {
		return nil
	}


	return o.Email

}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Email, true
}

// SetEmail sets field value
func (o *UserPropertiesPost) SetEmail(v string) {


	o.Email = &v

}

// HasEmail returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}


// GetAdministrator returns the Administrator field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *UserPropertiesPost) GetAdministrator() *bool {
	if o == nil {
		return nil
	}


	return o.Administrator

}

// GetAdministratorOk returns a tuple with the Administrator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetAdministratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.Administrator, true
}

// SetAdministrator sets field value
func (o *UserPropertiesPost) SetAdministrator(v bool) {


	o.Administrator = &v

}

// HasAdministrator returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasAdministrator() bool {
	if o != nil && o.Administrator != nil {
		return true
	}

	return false
}


// GetForceSecAuth returns the ForceSecAuth field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *UserPropertiesPost) GetForceSecAuth() *bool {
	if o == nil {
		return nil
	}


	return o.ForceSecAuth

}

// GetForceSecAuthOk returns a tuple with the ForceSecAuth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetForceSecAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.ForceSecAuth, true
}

// SetForceSecAuth sets field value
func (o *UserPropertiesPost) SetForceSecAuth(v bool) {


	o.ForceSecAuth = &v

}

// HasForceSecAuth returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasForceSecAuth() bool {
	if o != nil && o.ForceSecAuth != nil {
		return true
	}

	return false
}


// GetSecAuthActive returns the SecAuthActive field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *UserPropertiesPost) GetSecAuthActive() *bool {
	if o == nil {
		return nil
	}


	return o.SecAuthActive

}

// GetSecAuthActiveOk returns a tuple with the SecAuthActive field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetSecAuthActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.SecAuthActive, true
}

// SetSecAuthActive sets field value
func (o *UserPropertiesPost) SetSecAuthActive(v bool) {


	o.SecAuthActive = &v

}

// HasSecAuthActive returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasSecAuthActive() bool {
	if o != nil && o.SecAuthActive != nil {
		return true
	}

	return false
}


// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserPropertiesPost) GetPassword() *string {
	if o == nil {
		return nil
	}


	return o.Password

}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Password, true
}

// SetPassword sets field value
func (o *UserPropertiesPost) SetPassword(v string) {


	o.Password = &v

}

// HasPassword returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}


// GetActive returns the Active field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *UserPropertiesPost) GetActive() *bool {
	if o == nil {
		return nil
	}


	return o.Active

}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPropertiesPost) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.Active, true
}

// SetActive sets field value
func (o *UserPropertiesPost) SetActive(v bool) {


	o.Active = &v

}

// HasActive returns a boolean if a field has been set.
func (o *UserPropertiesPost) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

func (o UserPropertiesPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}

	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}

	if o.Email != nil {
		toSerialize["email"] = o.Email
	}

	if o.Administrator != nil {
		toSerialize["administrator"] = o.Administrator
	}

	if o.ForceSecAuth != nil {
		toSerialize["forceSecAuth"] = o.ForceSecAuth
	}

	if o.SecAuthActive != nil {
		toSerialize["secAuthActive"] = o.SecAuthActive
	}

	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableUserPropertiesPost struct {
	value *UserPropertiesPost
	isSet bool
}

func (v NullableUserPropertiesPost) Get() *UserPropertiesPost {
	return v.value
}

func (v *NullableUserPropertiesPost) Set(val *UserPropertiesPost) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPropertiesPost) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPropertiesPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPropertiesPost(val *UserPropertiesPost) *NullableUserPropertiesPost {
	return &NullableUserPropertiesPost{value: val, isSet: true}
}

func (v NullableUserPropertiesPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPropertiesPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


