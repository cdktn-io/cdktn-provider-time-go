// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-time-go/time/v14/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the time provider.
type TimeProviderFunctions interface {
	// Given a [Go duration string](https://pkg.go.dev/time#ParseDuration), will parse and return an object representation of that duration.
	DurationParse(duration *string) cdktn.IResolvable
	// Given an RFC3339 timestamp string, will parse and return an object representation of that date and time.
	Rfc3339Parse(timestamp *string) cdktn.IResolvable
	// Given a unix timestamp integer, will parse and return an object representation of that date and time.
	//
	// A unix timestamp is the number of seconds elapsed since January 1, 1970 UTC.
	UnixTimestampParse(unixTimestamp *float64) cdktn.IResolvable
}

// The jsii proxy struct for TimeProviderFunctions
type jsiiProxy_TimeProviderFunctions struct {
	_ byte // padding
}

func NewTimeProviderFunctions(providerLocalName *string) TimeProviderFunctions {
	_init_.Initialize()

	if err := validateNewTimeProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimeProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-time.providerFunctions.TimeProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewTimeProviderFunctions_Override(t TimeProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-time.providerFunctions.TimeProviderFunctions",
		[]interface{}{providerLocalName},
		t,
	)
}

func (t *jsiiProxy_TimeProviderFunctions) DurationParse(duration *string) cdktn.IResolvable {
	if err := t.validateDurationParseParameters(duration); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"durationParse",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeProviderFunctions) Rfc3339Parse(timestamp *string) cdktn.IResolvable {
	if err := t.validateRfc3339ParseParameters(timestamp); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"rfc3339Parse",
		[]interface{}{timestamp},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeProviderFunctions) UnixTimestampParse(unixTimestamp *float64) cdktn.IResolvable {
	if err := t.validateUnixTimestampParseParameters(unixTimestamp); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"unixTimestampParse",
		[]interface{}{unixTimestamp},
		&returns,
	)

	return returns
}

