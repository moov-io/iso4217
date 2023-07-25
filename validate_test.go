// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package iso4217_test

import (
	"testing"

	"github.com/moov-io/iso4217"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	// alphanumeric codes
	_, exists := iso4217.Lookup("USD")
	assert.True(t, exists)

	_, exists = iso4217.Lookup("eur")
	assert.True(t, exists)

	_, exists = iso4217.Lookup("")
	assert.False(t, exists)

	_, exists = iso4217.Lookup("QZA")
	assert.False(t, exists)

	// numeric codes
	_, exists = iso4217.Lookup("0")
	assert.False(t, exists)

	_, exists = iso4217.Lookup("840")
	assert.True(t, exists)

	cc, exists := iso4217.Lookup("36")
	assert.True(t, exists)
	assert.Equal(t, "AUD", cc.Code)

	cc, exists = iso4217.Lookup("036")
	assert.True(t, exists)
	assert.Equal(t, "AUD", cc.Code)
}
