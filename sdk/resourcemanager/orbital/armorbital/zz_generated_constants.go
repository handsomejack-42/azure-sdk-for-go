//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

const (
	moduleName    = "armorbital"
	moduleVersion = "v0.2.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// AuthorizationStatus - Authorization status of spacecraft.
type AuthorizationStatus string

const (
	AuthorizationStatusAllowed AuthorizationStatus = "Allowed"
	AuthorizationStatusPending AuthorizationStatus = "Pending"
	AuthorizationStatusDenied  AuthorizationStatus = "Denied"
)

// PossibleAuthorizationStatusValues returns the possible values for the AuthorizationStatus const type.
func PossibleAuthorizationStatusValues() []AuthorizationStatus {
	return []AuthorizationStatus{
		AuthorizationStatusAllowed,
		AuthorizationStatusPending,
		AuthorizationStatusDenied,
	}
}

// ToPtr returns a *AuthorizationStatus pointing to the current value.
func (c AuthorizationStatus) ToPtr() *AuthorizationStatus {
	return &c
}

// AutoTrackingConfiguration - Auto track configuration.
type AutoTrackingConfiguration string

const (
	AutoTrackingConfigurationDisabled AutoTrackingConfiguration = "disabled"
	AutoTrackingConfigurationXBand    AutoTrackingConfiguration = "xBand"
	AutoTrackingConfigurationSBand    AutoTrackingConfiguration = "sBand"
)

// PossibleAutoTrackingConfigurationValues returns the possible values for the AutoTrackingConfiguration const type.
func PossibleAutoTrackingConfigurationValues() []AutoTrackingConfiguration {
	return []AutoTrackingConfiguration{
		AutoTrackingConfigurationDisabled,
		AutoTrackingConfigurationXBand,
		AutoTrackingConfigurationSBand,
	}
}

// ToPtr returns a *AutoTrackingConfiguration pointing to the current value.
func (c AutoTrackingConfiguration) ToPtr() *AutoTrackingConfiguration {
	return &c
}

// Capability - Capability of the Ground Station.
type Capability string

const (
	CapabilityCommunication    Capability = "Communication"
	CapabilityEarthObservation Capability = "EarthObservation"
)

// PossibleCapabilityValues returns the possible values for the Capability const type.
func PossibleCapabilityValues() []Capability {
	return []Capability{
		CapabilityCommunication,
		CapabilityEarthObservation,
	}
}

// ToPtr returns a *Capability pointing to the current value.
func (c Capability) ToPtr() *Capability {
	return &c
}

type CapabilityType string

const (
	CapabilityTypeCommunication    CapabilityType = "Communication"
	CapabilityTypeEarthObservation CapabilityType = "EarthObservation"
)

// PossibleCapabilityTypeValues returns the possible values for the CapabilityType const type.
func PossibleCapabilityTypeValues() []CapabilityType {
	return []CapabilityType{
		CapabilityTypeCommunication,
		CapabilityTypeEarthObservation,
	}
}

// ToPtr returns a *CapabilityType pointing to the current value.
func (c CapabilityType) ToPtr() *CapabilityType {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// Direction - Direction (uplink or downlink)
type Direction string

const (
	DirectionDownlink Direction = "downlink"
	DirectionUplink   Direction = "uplink"
)

// PossibleDirectionValues returns the possible values for the Direction const type.
func PossibleDirectionValues() []Direction {
	return []Direction{
		DirectionDownlink,
		DirectionUplink,
	}
}

// ToPtr returns a *Direction pointing to the current value.
func (c Direction) ToPtr() *Direction {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// Polarization - polarization. eg (RHCP, LHCP)
type Polarization string

const (
	PolarizationDualRhcpLhcp     Polarization = "dualRhcpLhcp"
	PolarizationLHCP             Polarization = "LHCP"
	PolarizationLinearHorizontal Polarization = "linearHorizontal"
	PolarizationLinearVertical   Polarization = "linearVertical"
	PolarizationRHCP             Polarization = "RHCP"
)

// PossiblePolarizationValues returns the possible values for the Polarization const type.
func PossiblePolarizationValues() []Polarization {
	return []Polarization{
		PolarizationDualRhcpLhcp,
		PolarizationLHCP,
		PolarizationLinearHorizontal,
		PolarizationLinearVertical,
		PolarizationRHCP,
	}
}

// ToPtr returns a *Polarization pointing to the current value.
func (c Polarization) ToPtr() *Polarization {
	return &c
}

// Protocol - Protocol either UDP or TCP.
type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

// PossibleProtocolValues returns the possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{
		ProtocolTCP,
		ProtocolUDP,
	}
}

// ToPtr returns a *Protocol pointing to the current value.
func (c Protocol) ToPtr() *Protocol {
	return &c
}

// Status - Status of a contact.
type Status string

const (
	StatusScheduled         Status = "scheduled"
	StatusCancelled         Status = "cancelled"
	StatusSucceeded         Status = "succeeded"
	StatusFailed            Status = "failed"
	StatusProviderCancelled Status = "providerCancelled"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusScheduled,
		StatusCancelled,
		StatusSucceeded,
		StatusFailed,
		StatusProviderCancelled,
	}
}

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
}
