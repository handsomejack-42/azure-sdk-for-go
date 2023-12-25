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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"net/http"
	"net/url"
	"regexp"
)

// SubscriptionsServer is a fake server for instances of the armstreamanalytics.SubscriptionsClient type.
type SubscriptionsServer struct {
	// ListQuotas is the fake for method SubscriptionsClient.ListQuotas
	// HTTP status codes to indicate success: http.StatusOK
	ListQuotas func(ctx context.Context, location string, options *armstreamanalytics.SubscriptionsClientListQuotasOptions) (resp azfake.Responder[armstreamanalytics.SubscriptionsClientListQuotasResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionsServerTransport creates a new instance of SubscriptionsServerTransport with the provided implementation.
// The returned SubscriptionsServerTransport instance is connected to an instance of armstreamanalytics.SubscriptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubscriptionsServerTransport(srv *SubscriptionsServer) *SubscriptionsServerTransport {
	return &SubscriptionsServerTransport{srv: srv}
}

// SubscriptionsServerTransport connects instances of armstreamanalytics.SubscriptionsClient to instances of SubscriptionsServer.
// Don't use this type directly, use NewSubscriptionsServerTransport instead.
type SubscriptionsServerTransport struct {
	srv *SubscriptionsServer
}

// Do implements the policy.Transporter interface for SubscriptionsServerTransport.
func (s *SubscriptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubscriptionsClient.ListQuotas":
		resp, err = s.dispatchListQuotas(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchListQuotas(req *http.Request) (*http.Response, error) {
	if s.srv.ListQuotas == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQuotas not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ListQuotas(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionQuotasListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}