// Package operationalinsights implements the Azure ARM Operationalinsights service API version v1.
//
// Log Analytics Data Plane Client
package operationalinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Operationalinsights
	DefaultBaseURI = "https://api.loganalytics.io/v1"
)

// BaseClient is the base client for Operationalinsights.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

// Query executes an Analytics query for data. [Here](https://dev.loganalytics.io/documentation/Using-the-API) is an
// example for using POST with an Analytics query.
// Parameters:
// workspaceID - ID of the workspace. This is Workspace ID from the Properties blade in the Azure portal.
// body - the Analytics query. Learn more about the [Analytics query
// syntax](https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/)
func (client BaseClient) Query(ctx context.Context, workspaceID string, body QueryBody) (result QueryResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Query", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.BaseClient", "Query", err.Error())
	}

	req, err := client.QueryPreparer(ctx, workspaceID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.BaseClient", "Query", nil, "Failure preparing request")
		return
	}

	resp, err := client.QuerySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.BaseClient", "Query", resp, "Failure sending request")
		return
	}

	result, err = client.QueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.BaseClient", "Query", resp, "Failure responding to request")
	}

	return
}

// QueryPreparer prepares the Query request.
func (client BaseClient) QueryPreparer(ctx context.Context, workspaceID string, body QueryBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"workspaceId": autorest.Encode("path", workspaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/workspaces/{workspaceId}/query", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QuerySender sends the Query request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) QuerySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// QueryResponder handles the response to the Query request. The method always
// closes the http.Response Body.
func (client BaseClient) QueryResponder(resp *http.Response) (result QueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
