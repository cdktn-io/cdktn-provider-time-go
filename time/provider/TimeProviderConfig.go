// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type TimeProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.13.1/docs#alias TimeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

