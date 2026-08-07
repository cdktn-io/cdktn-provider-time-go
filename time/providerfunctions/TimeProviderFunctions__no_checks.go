// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimeProviderFunctions) validateDurationParseParameters(duration *string) error {
	return nil
}

func (t *jsiiProxy_TimeProviderFunctions) validateRfc3339ParseParameters(timestamp *string) error {
	return nil
}

func (t *jsiiProxy_TimeProviderFunctions) validateUnixTimestampParseParameters(unixTimestamp *float64) error {
	return nil
}

func validateNewTimeProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

