//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package rbac

// CreateOrUpdateRoleDefinitionResponse contains the response from method Client.CreateOrUpdateRoleDefinition.
type CreateOrUpdateRoleDefinitionResponse struct {
	RoleDefinition
}

// CreateRoleAssignmentResponse contains the response from method Client.CreateRoleAssignment.
type CreateRoleAssignmentResponse struct {
	RoleAssignment
}

// DeleteRoleAssignmentResponse contains the response from method Client.DeleteRoleAssignment.
type DeleteRoleAssignmentResponse struct {
	RoleAssignment
}

// DeleteRoleDefinitionResponse contains the response from method Client.DeleteRoleDefinition.
type DeleteRoleDefinitionResponse struct {
	RoleDefinition
}

// GetRoleAssignmentResponse contains the response from method Client.GetRoleAssignment.
type GetRoleAssignmentResponse struct {
	RoleAssignment
}

// GetRoleDefinitionResponse contains the response from method Client.GetRoleDefinition.
type GetRoleDefinitionResponse struct {
	RoleDefinition
}

// ListRoleAssignmentsResponse contains the response from method Client.NewListRoleAssignmentsPager.
type ListRoleAssignmentsResponse struct {
	RoleAssignmentListResult
}

// ListRoleDefinitionsResponse contains the response from method Client.NewListRoleDefinitionsPager.
type ListRoleDefinitionsResponse struct {
	RoleDefinitionListResult
}