//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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
	"strconv"
	"strings"
)

// SessionHostsClient contains the methods for the SessionHosts group.
// Don't use this type directly, use NewSessionHostsClient() instead.
type SessionHostsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSessionHostsClient creates a new instance of SessionHostsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSessionHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SessionHostsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SessionHostsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Delete - Remove a SessionHost.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// hostPoolName - The name of the host pool within the specified resource group
// sessionHostName - The name of the session host within the specified host pool
// options - SessionHostsClientDeleteOptions contains the optional parameters for the SessionHostsClient.Delete method.
func (client *SessionHostsClient) Delete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *SessionHostsClientDeleteOptions) (SessionHostsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostPoolName, sessionHostName, options)
	if err != nil {
		return SessionHostsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SessionHostsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SessionHostsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SessionHostsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SessionHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *SessionHostsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if sessionHostName == "" {
		return nil, errors.New("parameter sessionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionHostName}", url.PathEscape(sessionHostName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a session host.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// hostPoolName - The name of the host pool within the specified resource group
// sessionHostName - The name of the session host within the specified host pool
// options - SessionHostsClientGetOptions contains the optional parameters for the SessionHostsClient.Get method.
func (client *SessionHostsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *SessionHostsClientGetOptions) (SessionHostsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, sessionHostName, options)
	if err != nil {
		return SessionHostsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SessionHostsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionHostsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SessionHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *SessionHostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if sessionHostName == "" {
		return nil, errors.New("parameter sessionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionHostName}", url.PathEscape(sessionHostName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SessionHostsClient) getHandleResponse(resp *http.Response) (SessionHostsClientGetResponse, error) {
	result := SessionHostsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHost); err != nil {
		return SessionHostsClientGetResponse{}, err
	}
	return result, nil
}

// List - List sessionHosts.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// hostPoolName - The name of the host pool within the specified resource group
// options - SessionHostsClientListOptions contains the optional parameters for the SessionHostsClient.List method.
func (client *SessionHostsClient) List(resourceGroupName string, hostPoolName string, options *SessionHostsClientListOptions) *SessionHostsClientListPager {
	return &SessionHostsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, hostPoolName, options)
		},
		advancer: func(ctx context.Context, resp SessionHostsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SessionHostList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SessionHostsClient) listCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SessionHostsClient) listHandleResponse(resp *http.Response) (SessionHostsClientListResponse, error) {
	result := SessionHostsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostList); err != nil {
		return SessionHostsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update a session host.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// hostPoolName - The name of the host pool within the specified resource group
// sessionHostName - The name of the session host within the specified host pool
// options - SessionHostsClientUpdateOptions contains the optional parameters for the SessionHostsClient.Update method.
func (client *SessionHostsClient) Update(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *SessionHostsClientUpdateOptions) (SessionHostsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostPoolName, sessionHostName, options)
	if err != nil {
		return SessionHostsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SessionHostsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionHostsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SessionHostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *SessionHostsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if sessionHostName == "" {
		return nil, errors.New("parameter sessionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionHostName}", url.PathEscape(sessionHostName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.SessionHost != nil {
		return req, runtime.MarshalAsJSON(req, *options.SessionHost)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SessionHostsClient) updateHandleResponse(resp *http.Response) (SessionHostsClientUpdateResponse, error) {
	result := SessionHostsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHost); err != nil {
		return SessionHostsClientUpdateResponse{}, err
	}
	return result, nil
}
