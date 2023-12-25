//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
	"net/http"
	"net/url"
	"regexp"
)

// CreateAndAssociateIPFilterServer is a fake server for instances of the armelastic.CreateAndAssociateIPFilterClient type.
type CreateAndAssociateIPFilterServer struct {
	// BeginCreate is the fake for method CreateAndAssociateIPFilterClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, monitorName string, options *armelastic.CreateAndAssociateIPFilterClientBeginCreateOptions) (resp azfake.PollerResponder[armelastic.CreateAndAssociateIPFilterClientCreateResponse], errResp azfake.ErrorResponder)
}

// NewCreateAndAssociateIPFilterServerTransport creates a new instance of CreateAndAssociateIPFilterServerTransport with the provided implementation.
// The returned CreateAndAssociateIPFilterServerTransport instance is connected to an instance of armelastic.CreateAndAssociateIPFilterClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCreateAndAssociateIPFilterServerTransport(srv *CreateAndAssociateIPFilterServer) *CreateAndAssociateIPFilterServerTransport {
	return &CreateAndAssociateIPFilterServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armelastic.CreateAndAssociateIPFilterClientCreateResponse]](),
	}
}

// CreateAndAssociateIPFilterServerTransport connects instances of armelastic.CreateAndAssociateIPFilterClient to instances of CreateAndAssociateIPFilterServer.
// Don't use this type directly, use NewCreateAndAssociateIPFilterServerTransport instead.
type CreateAndAssociateIPFilterServerTransport struct {
	srv         *CreateAndAssociateIPFilterServer
	beginCreate *tracker[azfake.PollerResponder[armelastic.CreateAndAssociateIPFilterClientCreateResponse]]
}

// Do implements the policy.Transporter interface for CreateAndAssociateIPFilterServerTransport.
func (c *CreateAndAssociateIPFilterServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CreateAndAssociateIPFilterClient.BeginCreate":
		resp, err = c.dispatchBeginCreate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreateAndAssociateIPFilterServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Elastic/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createAndAssociateIPFilter`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		iPsUnescaped, err := url.QueryUnescape(qp.Get("ips"))
		if err != nil {
			return nil, err
		}
		iPsParam := getOptional(iPsUnescaped)
		nameUnescaped, err := url.QueryUnescape(qp.Get("name"))
		if err != nil {
			return nil, err
		}
		nameParam := getOptional(nameUnescaped)
		var options *armelastic.CreateAndAssociateIPFilterClientBeginCreateOptions
		if iPsParam != nil || nameParam != nil {
			options = &armelastic.CreateAndAssociateIPFilterClientBeginCreateOptions{
				IPs:  iPsParam,
				Name: nameParam,
			}
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), resourceGroupNameParam, monitorNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}