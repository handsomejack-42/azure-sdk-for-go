//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SavedSearchesClient contains the methods for the SavedSearches group.
// Don't use this type directly, use NewSavedSearchesClient() instead.
type SavedSearchesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSavedSearchesClient creates a new instance of SavedSearchesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSavedSearchesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SavedSearchesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SavedSearchesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a saved search for a given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// savedSearchID - The id of the saved search.
// parameters - The parameters required to save a search.
// options - SavedSearchesClientCreateOrUpdateOptions contains the optional parameters for the SavedSearchesClient.CreateOrUpdate
// method.
func (client *SavedSearchesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters SavedSearch, options *SavedSearchesClientCreateOrUpdateOptions) (SavedSearchesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, savedSearchID, parameters, options)
	if err != nil {
		return SavedSearchesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavedSearchesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavedSearchesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SavedSearchesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters SavedSearch, options *SavedSearchesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if savedSearchID == "" {
		return nil, errors.New("parameter savedSearchID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savedSearchId}", url.PathEscape(savedSearchID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SavedSearchesClient) createOrUpdateHandleResponse(resp *http.Response) (SavedSearchesClientCreateOrUpdateResponse, error) {
	result := SavedSearchesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavedSearch); err != nil {
		return SavedSearchesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified saved search in a given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// savedSearchID - The id of the saved search.
// options - SavedSearchesClientDeleteOptions contains the optional parameters for the SavedSearchesClient.Delete method.
func (client *SavedSearchesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesClientDeleteOptions) (SavedSearchesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, savedSearchID, options)
	if err != nil {
		return SavedSearchesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavedSearchesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavedSearchesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SavedSearchesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SavedSearchesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if savedSearchID == "" {
		return nil, errors.New("parameter savedSearchID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savedSearchId}", url.PathEscape(savedSearchID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified saved search for a given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// savedSearchID - The id of the saved search.
// options - SavedSearchesClientGetOptions contains the optional parameters for the SavedSearchesClient.Get method.
func (client *SavedSearchesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesClientGetOptions) (SavedSearchesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, savedSearchID, options)
	if err != nil {
		return SavedSearchesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavedSearchesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavedSearchesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SavedSearchesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if savedSearchID == "" {
		return nil, errors.New("parameter savedSearchID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savedSearchId}", url.PathEscape(savedSearchID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SavedSearchesClient) getHandleResponse(resp *http.Response) (SavedSearchesClientGetResponse, error) {
	result := SavedSearchesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavedSearch); err != nil {
		return SavedSearchesClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Gets the saved searches for a given Log Analytics Workspace
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - SavedSearchesClientListByWorkspaceOptions contains the optional parameters for the SavedSearchesClient.ListByWorkspace
// method.
func (client *SavedSearchesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *SavedSearchesClientListByWorkspaceOptions) (SavedSearchesClientListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SavedSearchesClientListByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavedSearchesClientListByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavedSearchesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *SavedSearchesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SavedSearchesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *SavedSearchesClient) listByWorkspaceHandleResponse(resp *http.Response) (SavedSearchesClientListByWorkspaceResponse, error) {
	result := SavedSearchesClientListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavedSearchesListResult); err != nil {
		return SavedSearchesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
