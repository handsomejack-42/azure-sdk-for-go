//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtestbase

// AccountsClientCheckPackageNameAvailabilityResponse contains the response from method AccountsClient.CheckPackageNameAvailability.
type AccountsClientCheckPackageNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// AccountsClientCreateResponse contains the response from method AccountsClient.BeginCreate.
type AccountsClientCreateResponse struct {
	AccountResource
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetFileUploadURLResponse contains the response from method AccountsClient.GetFileUploadURL.
type AccountsClientGetFileUploadURLResponse struct {
	FileUploadURLResponse
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	AccountResource
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	AccountListResult
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	AccountListResult
}

// AccountsClientOffboardResponse contains the response from method AccountsClient.BeginOffboard.
type AccountsClientOffboardResponse struct {
	// placeholder for future response values
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	AccountResource
}

// AnalysisResultsClientGetResponse contains the response from method AnalysisResultsClient.Get.
type AnalysisResultsClientGetResponse struct {
	AnalysisResultSingletonResource
}

// AnalysisResultsClientListResponse contains the response from method AnalysisResultsClient.NewListPager.
type AnalysisResultsClientListResponse struct {
	AnalysisResultListResult
}

// AvailableOSClientGetResponse contains the response from method AvailableOSClient.Get.
type AvailableOSClientGetResponse struct {
	AvailableOSResource
}

// AvailableOSClientListResponse contains the response from method AvailableOSClient.NewListPager.
type AvailableOSClientListResponse struct {
	AvailableOSListResult
}

// CustomerEventsClientCreateResponse contains the response from method CustomerEventsClient.BeginCreate.
type CustomerEventsClientCreateResponse struct {
	CustomerEventResource
}

// CustomerEventsClientDeleteResponse contains the response from method CustomerEventsClient.BeginDelete.
type CustomerEventsClientDeleteResponse struct {
	// placeholder for future response values
}

// CustomerEventsClientGetResponse contains the response from method CustomerEventsClient.Get.
type CustomerEventsClientGetResponse struct {
	CustomerEventResource
}

// CustomerEventsClientListByTestBaseAccountResponse contains the response from method CustomerEventsClient.NewListByTestBaseAccountPager.
type CustomerEventsClientListByTestBaseAccountResponse struct {
	CustomerEventListResult
}

// EmailEventsClientGetResponse contains the response from method EmailEventsClient.Get.
type EmailEventsClientGetResponse struct {
	EmailEventResource
}

// EmailEventsClientListResponse contains the response from method EmailEventsClient.NewListPager.
type EmailEventsClientListResponse struct {
	EmailEventListResult
}

// FavoriteProcessesClientCreateResponse contains the response from method FavoriteProcessesClient.Create.
type FavoriteProcessesClientCreateResponse struct {
	FavoriteProcessResource
}

// FavoriteProcessesClientDeleteResponse contains the response from method FavoriteProcessesClient.Delete.
type FavoriteProcessesClientDeleteResponse struct {
	// placeholder for future response values
}

// FavoriteProcessesClientGetResponse contains the response from method FavoriteProcessesClient.Get.
type FavoriteProcessesClientGetResponse struct {
	FavoriteProcessResource
}

// FavoriteProcessesClientListResponse contains the response from method FavoriteProcessesClient.NewListPager.
type FavoriteProcessesClientListResponse struct {
	FavoriteProcessListResult
}

// FlightingRingsClientGetResponse contains the response from method FlightingRingsClient.Get.
type FlightingRingsClientGetResponse struct {
	FlightingRingResource
}

// FlightingRingsClientListResponse contains the response from method FlightingRingsClient.NewListPager.
type FlightingRingsClientListResponse struct {
	FlightingRingListResult
}

// OSUpdatesClientGetResponse contains the response from method OSUpdatesClient.Get.
type OSUpdatesClientGetResponse struct {
	OSUpdateResource
}

// OSUpdatesClientListResponse contains the response from method OSUpdatesClient.NewListPager.
type OSUpdatesClientListResponse struct {
	OSUpdateListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PackagesClientCreateResponse contains the response from method PackagesClient.BeginCreate.
type PackagesClientCreateResponse struct {
	PackageResource
}

// PackagesClientDeleteResponse contains the response from method PackagesClient.BeginDelete.
type PackagesClientDeleteResponse struct {
	// placeholder for future response values
}

// PackagesClientGetDownloadURLResponse contains the response from method PackagesClient.GetDownloadURL.
type PackagesClientGetDownloadURLResponse struct {
	DownloadURLResponse
}

// PackagesClientGetResponse contains the response from method PackagesClient.Get.
type PackagesClientGetResponse struct {
	PackageResource
}

// PackagesClientHardDeleteResponse contains the response from method PackagesClient.BeginHardDelete.
type PackagesClientHardDeleteResponse struct {
	// placeholder for future response values
}

// PackagesClientListByTestBaseAccountResponse contains the response from method PackagesClient.NewListByTestBaseAccountPager.
type PackagesClientListByTestBaseAccountResponse struct {
	PackageListResult
}

// PackagesClientUpdateResponse contains the response from method PackagesClient.BeginUpdate.
type PackagesClientUpdateResponse struct {
	PackageResource
}

// SKUsClientListResponse contains the response from method SKUsClient.NewListPager.
type SKUsClientListResponse struct {
	AccountSKUListResult
}

// TestResultsClientGetDownloadURLResponse contains the response from method TestResultsClient.GetDownloadURL.
type TestResultsClientGetDownloadURLResponse struct {
	DownloadURLResponse
}

// TestResultsClientGetResponse contains the response from method TestResultsClient.Get.
type TestResultsClientGetResponse struct {
	TestResultResource
}

// TestResultsClientGetVideoDownloadURLResponse contains the response from method TestResultsClient.GetVideoDownloadURL.
type TestResultsClientGetVideoDownloadURLResponse struct {
	DownloadURLResponse
}

// TestResultsClientListResponse contains the response from method TestResultsClient.NewListPager.
type TestResultsClientListResponse struct {
	TestResultListResult
}

// TestSummariesClientGetResponse contains the response from method TestSummariesClient.Get.
type TestSummariesClientGetResponse struct {
	TestSummaryResource
}

// TestSummariesClientListResponse contains the response from method TestSummariesClient.NewListPager.
type TestSummariesClientListResponse struct {
	TestSummaryListResult
}

// TestTypesClientGetResponse contains the response from method TestTypesClient.Get.
type TestTypesClientGetResponse struct {
	TestTypeResource
}

// TestTypesClientListResponse contains the response from method TestTypesClient.NewListPager.
type TestTypesClientListResponse struct {
	TestTypeListResult
}

// UsageClientListResponse contains the response from method UsageClient.NewListPager.
type UsageClientListResponse struct {
	AccountUsageDataList
}