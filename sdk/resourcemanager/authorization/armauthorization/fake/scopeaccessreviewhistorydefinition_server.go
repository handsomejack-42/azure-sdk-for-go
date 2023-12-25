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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ScopeAccessReviewHistoryDefinitionServer is a fake server for instances of the armauthorization.ScopeAccessReviewHistoryDefinitionClient type.
type ScopeAccessReviewHistoryDefinitionServer struct {
	// Create is the fake for method ScopeAccessReviewHistoryDefinitionClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, scope string, historyDefinitionID string, properties armauthorization.AccessReviewHistoryDefinitionProperties, options *armauthorization.ScopeAccessReviewHistoryDefinitionClientCreateOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewHistoryDefinitionClientCreateResponse], errResp azfake.ErrorResponder)

	// DeleteByID is the fake for method ScopeAccessReviewHistoryDefinitionClient.DeleteByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteByID func(ctx context.Context, scope string, historyDefinitionID string, options *armauthorization.ScopeAccessReviewHistoryDefinitionClientDeleteByIDOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewHistoryDefinitionClientDeleteByIDResponse], errResp azfake.ErrorResponder)
}

// NewScopeAccessReviewHistoryDefinitionServerTransport creates a new instance of ScopeAccessReviewHistoryDefinitionServerTransport with the provided implementation.
// The returned ScopeAccessReviewHistoryDefinitionServerTransport instance is connected to an instance of armauthorization.ScopeAccessReviewHistoryDefinitionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAccessReviewHistoryDefinitionServerTransport(srv *ScopeAccessReviewHistoryDefinitionServer) *ScopeAccessReviewHistoryDefinitionServerTransport {
	return &ScopeAccessReviewHistoryDefinitionServerTransport{srv: srv}
}

// ScopeAccessReviewHistoryDefinitionServerTransport connects instances of armauthorization.ScopeAccessReviewHistoryDefinitionClient to instances of ScopeAccessReviewHistoryDefinitionServer.
// Don't use this type directly, use NewScopeAccessReviewHistoryDefinitionServerTransport instead.
type ScopeAccessReviewHistoryDefinitionServerTransport struct {
	srv *ScopeAccessReviewHistoryDefinitionServer
}

// Do implements the policy.Transporter interface for ScopeAccessReviewHistoryDefinitionServerTransport.
func (s *ScopeAccessReviewHistoryDefinitionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ScopeAccessReviewHistoryDefinitionClient.Create":
		resp, err = s.dispatchCreate(req)
	case "ScopeAccessReviewHistoryDefinitionClient.DeleteByID":
		resp, err = s.dispatchDeleteByID(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ScopeAccessReviewHistoryDefinitionServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions/(?P<historyDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.AccessReviewHistoryDefinitionProperties](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	historyDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Create(req.Context(), scopeParam, historyDefinitionIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewHistoryDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewHistoryDefinitionServerTransport) dispatchDeleteByID(req *http.Request) (*http.Response, error) {
	if s.srv.DeleteByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteByID not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions/(?P<historyDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	historyDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.DeleteByID(req.Context(), scopeParam, historyDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}