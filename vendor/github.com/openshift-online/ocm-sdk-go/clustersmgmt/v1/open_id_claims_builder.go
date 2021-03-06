/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// OpenIDClaimsBuilder contains the data and logic needed to build 'open_ID_claims' objects.
//
// _OpenID_ identity provider claims.
type OpenIDClaimsBuilder struct {
	email             []string
	name              []string
	preferredUsername []string
}

// NewOpenIDClaims creates a new builder of 'open_ID_claims' objects.
func NewOpenIDClaims() *OpenIDClaimsBuilder {
	return new(OpenIDClaimsBuilder)
}

// Email sets the value of the 'email' attribute
// to the given values.
//
//
func (b *OpenIDClaimsBuilder) Email(values ...string) *OpenIDClaimsBuilder {
	b.email = make([]string, len(values))
	copy(b.email, values)
	return b
}

// Name sets the value of the 'name' attribute
// to the given values.
//
//
func (b *OpenIDClaimsBuilder) Name(values ...string) *OpenIDClaimsBuilder {
	b.name = make([]string, len(values))
	copy(b.name, values)
	return b
}

// PreferredUsername sets the value of the 'preferred_username' attribute
// to the given values.
//
//
func (b *OpenIDClaimsBuilder) PreferredUsername(values ...string) *OpenIDClaimsBuilder {
	b.preferredUsername = make([]string, len(values))
	copy(b.preferredUsername, values)
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *OpenIDClaimsBuilder) Copy(object *OpenIDClaims) *OpenIDClaimsBuilder {
	if object == nil {
		return b
	}
	if len(object.email) > 0 {
		b.email = make([]string, len(object.email))
		copy(b.email, object.email)
	} else {
		b.email = nil
	}
	if len(object.name) > 0 {
		b.name = make([]string, len(object.name))
		copy(b.name, object.name)
	} else {
		b.name = nil
	}
	if len(object.preferredUsername) > 0 {
		b.preferredUsername = make([]string, len(object.preferredUsername))
		copy(b.preferredUsername, object.preferredUsername)
	} else {
		b.preferredUsername = nil
	}
	return b
}

// Build creates a 'open_ID_claims' object using the configuration stored in the builder.
func (b *OpenIDClaimsBuilder) Build() (object *OpenIDClaims, err error) {
	object = new(OpenIDClaims)
	if b.email != nil {
		object.email = make([]string, len(b.email))
		copy(object.email, b.email)
	}
	if b.name != nil {
		object.name = make([]string, len(b.name))
		copy(object.name, b.name)
	}
	if b.preferredUsername != nil {
		object.preferredUsername = make([]string, len(b.preferredUsername))
		copy(object.preferredUsername, b.preferredUsername)
	}
	return
}
