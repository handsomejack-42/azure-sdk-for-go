//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuredata

const (
	moduleName    = "armazuredata"
	moduleVersion = "v0.2.0"
)

// IdentityType - The type of identity that creates/modifies resources
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "application"
	IdentityTypeKey             IdentityType = "key"
	IdentityTypeManagedIdentity IdentityType = "managedIdentity"
	IdentityTypeUser            IdentityType = "user"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// ToPtr returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginSystem OperationOrigin = "system"
	OperationOriginUser   OperationOrigin = "user"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginSystem,
		OperationOriginUser,
	}
}

// ToPtr returns a *OperationOrigin pointing to the current value.
func (c OperationOrigin) ToPtr() *OperationOrigin {
	return &c
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}
