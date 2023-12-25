//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v6"
	"net/http"
	"net/url"
	"regexp"
)

// SnapshotPoliciesServer is a fake server for instances of the armnetapp.SnapshotPoliciesClient type.
type SnapshotPoliciesServer struct {
	// Create is the fake for method SnapshotPoliciesClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body armnetapp.SnapshotPolicy, options *armnetapp.SnapshotPoliciesClientCreateOptions) (resp azfake.Responder[armnetapp.SnapshotPoliciesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SnapshotPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *armnetapp.SnapshotPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetapp.SnapshotPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SnapshotPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *armnetapp.SnapshotPoliciesClientGetOptions) (resp azfake.Responder[armnetapp.SnapshotPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SnapshotPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armnetapp.SnapshotPoliciesClientListOptions) (resp azfake.PagerResponder[armnetapp.SnapshotPoliciesClientListResponse])

	// ListVolumes is the fake for method SnapshotPoliciesClient.ListVolumes
	// HTTP status codes to indicate success: http.StatusOK
	ListVolumes func(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, options *armnetapp.SnapshotPoliciesClientListVolumesOptions) (resp azfake.Responder[armnetapp.SnapshotPoliciesClientListVolumesResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method SnapshotPoliciesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string, body armnetapp.SnapshotPolicyPatch, options *armnetapp.SnapshotPoliciesClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetapp.SnapshotPoliciesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSnapshotPoliciesServerTransport creates a new instance of SnapshotPoliciesServerTransport with the provided implementation.
// The returned SnapshotPoliciesServerTransport instance is connected to an instance of armnetapp.SnapshotPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSnapshotPoliciesServerTransport(srv *SnapshotPoliciesServer) *SnapshotPoliciesServerTransport {
	return &SnapshotPoliciesServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetapp.SnapshotPoliciesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetapp.SnapshotPoliciesClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armnetapp.SnapshotPoliciesClientUpdateResponse]](),
	}
}

// SnapshotPoliciesServerTransport connects instances of armnetapp.SnapshotPoliciesClient to instances of SnapshotPoliciesServer.
// Don't use this type directly, use NewSnapshotPoliciesServerTransport instead.
type SnapshotPoliciesServerTransport struct {
	srv          *SnapshotPoliciesServer
	beginDelete  *tracker[azfake.PollerResponder[armnetapp.SnapshotPoliciesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetapp.SnapshotPoliciesClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armnetapp.SnapshotPoliciesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for SnapshotPoliciesServerTransport.
func (s *SnapshotPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SnapshotPoliciesClient.Create":
		resp, err = s.dispatchCreate(req)
	case "SnapshotPoliciesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SnapshotPoliciesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SnapshotPoliciesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SnapshotPoliciesClient.ListVolumes":
		resp, err = s.dispatchListVolumes(req)
	case "SnapshotPoliciesClient.BeginUpdate":
		resp, err = s.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SnapshotPoliciesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshotPolicies/(?P<snapshotPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetapp.SnapshotPolicy](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	snapshotPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, snapshotPolicyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnapshotPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SnapshotPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshotPolicies/(?P<snapshotPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		snapshotPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, snapshotPolicyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SnapshotPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshotPolicies/(?P<snapshotPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	snapshotPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, snapshotPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnapshotPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SnapshotPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshotPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, accountNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *SnapshotPoliciesServerTransport) dispatchListVolumes(req *http.Request) (*http.Response, error) {
	if s.srv.ListVolumes == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListVolumes not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshotPolicies/(?P<snapshotPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	snapshotPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ListVolumes(req.Context(), resourceGroupNameParam, accountNameParam, snapshotPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnapshotPolicyVolumeList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SnapshotPoliciesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshotPolicies/(?P<snapshotPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.SnapshotPolicyPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		snapshotPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, accountNameParam, snapshotPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}