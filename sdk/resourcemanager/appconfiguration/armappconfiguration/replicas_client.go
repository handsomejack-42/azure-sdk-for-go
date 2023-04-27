//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappconfiguration

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicasClient contains the methods for the Replicas group.
// Don't use this type directly, use NewReplicasClient() instead.
type ReplicasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicasClient creates a new instance of ReplicasClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicasClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicasClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a replica with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - configStoreName - The name of the configuration store.
//   - replicaName - The name of the replica.
//   - replicaCreationParameters - The parameters for creating a replica.
//   - options - ReplicasClientBeginCreateOptions contains the optional parameters for the ReplicasClient.BeginCreate method.
func (client *ReplicasClient) BeginCreate(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, replicaCreationParameters Replica, options *ReplicasClientBeginCreateOptions) (*runtime.Poller[ReplicasClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, configStoreName, replicaName, replicaCreationParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicasClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ReplicasClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a replica with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ReplicasClient) create(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, replicaCreationParameters Replica, options *ReplicasClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, configStoreName, replicaName, replicaCreationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicasClient) createCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, replicaCreationParameters Replica, options *ReplicasClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/replicas/{replicaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if replicaName == "" {
		return nil, errors.New("parameter replicaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicaName}", url.PathEscape(replicaName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, replicaCreationParameters)
}

// BeginDelete - Deletes a replica.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - configStoreName - The name of the configuration store.
//   - replicaName - The name of the replica.
//   - options - ReplicasClientBeginDeleteOptions contains the optional parameters for the ReplicasClient.BeginDelete method.
func (client *ReplicasClient) BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, options *ReplicasClientBeginDeleteOptions) (*runtime.Poller[ReplicasClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, configStoreName, replicaName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicasClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ReplicasClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a replica.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ReplicasClient) deleteOperation(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, options *ReplicasClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configStoreName, replicaName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicasClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, options *ReplicasClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/replicas/{replicaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if replicaName == "" {
		return nil, errors.New("parameter replicaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicaName}", url.PathEscape(replicaName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified replica.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - configStoreName - The name of the configuration store.
//   - replicaName - The name of the replica.
//   - options - ReplicasClientGetOptions contains the optional parameters for the ReplicasClient.Get method.
func (client *ReplicasClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, options *ReplicasClientGetOptions) (ReplicasClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configStoreName, replicaName, options)
	if err != nil {
		return ReplicasClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicasClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicasClient) getCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, replicaName string, options *ReplicasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/replicas/{replicaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if replicaName == "" {
		return nil, errors.New("parameter replicaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicaName}", url.PathEscape(replicaName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicasClient) getHandleResponse(resp *http.Response) (ReplicasClientGetResponse, error) {
	result := ReplicasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Replica); err != nil {
		return ReplicasClientGetResponse{}, err
	}
	return result, nil
}

// NewListByConfigurationStorePager - Lists the replicas for a given configuration store.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - configStoreName - The name of the configuration store.
//   - options - ReplicasClientListByConfigurationStoreOptions contains the optional parameters for the ReplicasClient.NewListByConfigurationStorePager
//     method.
func (client *ReplicasClient) NewListByConfigurationStorePager(resourceGroupName string, configStoreName string, options *ReplicasClientListByConfigurationStoreOptions) *runtime.Pager[ReplicasClientListByConfigurationStoreResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicasClientListByConfigurationStoreResponse]{
		More: func(page ReplicasClientListByConfigurationStoreResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicasClientListByConfigurationStoreResponse) (ReplicasClientListByConfigurationStoreResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByConfigurationStoreCreateRequest(ctx, resourceGroupName, configStoreName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicasClientListByConfigurationStoreResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicasClientListByConfigurationStoreResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicasClientListByConfigurationStoreResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConfigurationStoreHandleResponse(resp)
		},
	})
}

// listByConfigurationStoreCreateRequest creates the ListByConfigurationStore request.
func (client *ReplicasClient) listByConfigurationStoreCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ReplicasClientListByConfigurationStoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/replicas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByConfigurationStoreHandleResponse handles the ListByConfigurationStore response.
func (client *ReplicasClient) listByConfigurationStoreHandleResponse(resp *http.Response) (ReplicasClientListByConfigurationStoreResponse, error) {
	result := ReplicasClientListByConfigurationStoreResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicaListResult); err != nil {
		return ReplicasClientListByConfigurationStoreResponse{}, err
	}
	return result, nil
}