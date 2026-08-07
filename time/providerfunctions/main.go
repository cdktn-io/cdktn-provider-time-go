// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-time.providerFunctions.TimeProviderFunctions",
		reflect.TypeOf((*TimeProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "durationParse", GoMethod: "DurationParse"},
			_jsii_.MemberMethod{JsiiMethod: "rfc3339Parse", GoMethod: "Rfc3339Parse"},
			_jsii_.MemberMethod{JsiiMethod: "unixTimestampParse", GoMethod: "UnixTimestampParse"},
		},
		func() interface{} {
			return &jsiiProxy_TimeProviderFunctions{}
		},
	)
}
