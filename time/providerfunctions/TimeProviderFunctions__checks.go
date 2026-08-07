// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"
)

func (t *jsiiProxy_TimeProviderFunctions) validateDurationParseParameters(duration *string) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeProviderFunctions) validateRfc3339ParseParameters(timestamp *string) error {
	if timestamp == nil {
		return fmt.Errorf("parameter timestamp is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeProviderFunctions) validateUnixTimestampParseParameters(unixTimestamp *float64) error {
	if unixTimestamp == nil {
		return fmt.Errorf("parameter unixTimestamp is required, but nil was provided")
	}

	return nil
}

func validateNewTimeProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

