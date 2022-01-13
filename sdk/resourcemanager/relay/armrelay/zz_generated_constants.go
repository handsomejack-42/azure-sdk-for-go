//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

const (
	moduleName    = "armrelay"
	moduleVersion = "v0.3.0"
)

type AccessRights string

const (
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
	AccessRightsListen AccessRights = "Listen"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsManage,
		AccessRightsSend,
		AccessRightsListen,
	}
}

// ToPtr returns a *AccessRights pointing to the current value.
func (c AccessRights) ToPtr() *AccessRights {
	return &c
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// ToPtr returns a *KeyType pointing to the current value.
func (c KeyType) ToPtr() *KeyType {
	return &c
}

type ProvisioningStateEnum string

const (
	ProvisioningStateEnumCreated   ProvisioningStateEnum = "Created"
	ProvisioningStateEnumSucceeded ProvisioningStateEnum = "Succeeded"
	ProvisioningStateEnumDeleted   ProvisioningStateEnum = "Deleted"
	ProvisioningStateEnumFailed    ProvisioningStateEnum = "Failed"
	ProvisioningStateEnumUpdating  ProvisioningStateEnum = "Updating"
	ProvisioningStateEnumUnknown   ProvisioningStateEnum = "Unknown"
)

// PossibleProvisioningStateEnumValues returns the possible values for the ProvisioningStateEnum const type.
func PossibleProvisioningStateEnumValues() []ProvisioningStateEnum {
	return []ProvisioningStateEnum{
		ProvisioningStateEnumCreated,
		ProvisioningStateEnumSucceeded,
		ProvisioningStateEnumDeleted,
		ProvisioningStateEnumFailed,
		ProvisioningStateEnumUpdating,
		ProvisioningStateEnumUnknown,
	}
}

// ToPtr returns a *ProvisioningStateEnum pointing to the current value.
func (c ProvisioningStateEnum) ToPtr() *ProvisioningStateEnum {
	return &c
}

// Relaytype - WCF relay type.
type Relaytype string

const (
	RelaytypeNetTCP Relaytype = "NetTcp"
	RelaytypeHTTP   Relaytype = "Http"
)

// PossibleRelaytypeValues returns the possible values for the Relaytype const type.
func PossibleRelaytypeValues() []Relaytype {
	return []Relaytype{
		RelaytypeNetTCP,
		RelaytypeHTTP,
	}
}

// ToPtr returns a *Relaytype pointing to the current value.
func (c Relaytype) ToPtr() *Relaytype {
	return &c
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}

// ToPtr returns a *UnavailableReason pointing to the current value.
func (c UnavailableReason) ToPtr() *UnavailableReason {
	return &c
}
