// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *LogMetric) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MetricDescriptor) {
		if err := r.MetricDescriptor.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BucketOptions) {
		if err := r.BucketOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *LogMetricMetricDescriptor) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Metadata) {
		if err := r.Metadata.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *LogMetricMetricDescriptorLabels) validate() error {
	return nil
}
func (r *LogMetricMetricDescriptorMetadata) validate() error {
	return nil
}
func (r *LogMetricBucketOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LinearBuckets) {
		if err := r.LinearBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExponentialBuckets) {
		if err := r.ExponentialBuckets.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExplicitBuckets) {
		if err := r.ExplicitBuckets.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *LogMetricBucketOptionsLinearBuckets) validate() error {
	return nil
}
func (r *LogMetricBucketOptionsExponentialBuckets) validate() error {
	return nil
}
func (r *LogMetricBucketOptionsExplicitBuckets) validate() error {
	return nil
}

func logMetricGetURL(userBasePath string, r *LogMetric) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/metrics/{{name}}", "https://logging.googleapis.com/v2/", userBasePath, params), nil
}

func logMetricListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/metrics", "https://logging.googleapis.com/v2/", userBasePath, params), nil

}

func logMetricCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/metrics", "https://logging.googleapis.com/v2/", userBasePath, params), nil

}

func logMetricDeleteURL(userBasePath string, r *LogMetric) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/metrics/{{name}}", "https://logging.googleapis.com/v2/", userBasePath, params), nil
}

// logMetricApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type logMetricApiOperation interface {
	do(context.Context, *LogMetric, *Client) error
}

// newUpdateLogMetricUpdateRequest creates a request for an
// LogMetric resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateLogMetricUpdateRequest(ctx context.Context, f *LogMetric, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		req["filter"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	if v, err := expandLogMetricMetricDescriptor(c, f.MetricDescriptor); err != nil {
		return nil, fmt.Errorf("error expanding MetricDescriptor into metricDescriptor: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["metricDescriptor"] = v
	}
	if v := f.ValueExtractor; !dcl.IsEmptyValueIndirect(v) {
		req["valueExtractor"] = v
	}
	if v := f.LabelExtractors; !dcl.IsEmptyValueIndirect(v) {
		req["labelExtractors"] = v
	}
	if v, err := expandLogMetricBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["bucketOptions"] = v
	}
	return req, nil
}

// marshalUpdateLogMetricUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateLogMetricUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateLogMetricUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateLogMetricUpdateOperation) do(ctx context.Context, r *LogMetric, c *Client) error {
	_, err := c.GetLogMetric(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateLogMetricUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateLogMetricUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listLogMetricRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := logMetricListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != LogMetricMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listLogMetricOperation struct {
	Metrics []map[string]interface{} `json:"metrics"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listLogMetric(ctx context.Context, project, pageToken string, pageSize int32) ([]*LogMetric, string, error) {
	b, err := c.listLogMetricRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listLogMetricOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*LogMetric
	for _, v := range m.Metrics {
		res, err := unmarshalMapLogMetric(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllLogMetric(ctx context.Context, f func(*LogMetric) bool, resources []*LogMetric) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteLogMetric(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteLogMetricOperation struct{}

func (op *deleteLogMetricOperation) do(ctx context.Context, r *LogMetric, c *Client) error {
	r, err := c.GetLogMetric(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("LogMetric not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetLogMetric checking for existence. error: %v", err)
		return err
	}

	u, err := logMetricDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete LogMetric: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetLogMetric(ctx, r.URLNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createLogMetricOperation struct {
	response map[string]interface{}
}

func (op *createLogMetricOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createLogMetricOperation) do(ctx context.Context, r *LogMetric, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := logMetricCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetLogMetric(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getLogMetricRaw(ctx context.Context, r *LogMetric) ([]byte, error) {

	u, err := logMetricGetURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) logMetricDiffsForRawDesired(ctx context.Context, rawDesired *LogMetric, opts ...dcl.ApplyOption) (initial, desired *LogMetric, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *LogMetric
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*LogMetric); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected LogMetric, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetLogMetric(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a LogMetric resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve LogMetric resource: %v", err)
		}
		c.Config.Logger.Info("Found that LogMetric resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeLogMetricDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for LogMetric: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for LogMetric: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeLogMetricInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for LogMetric: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeLogMetricDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for LogMetric: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffLogMetric(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeLogMetricInitialState(rawInitial, rawDesired *LogMetric) (*LogMetric, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeLogMetricDesiredState(rawDesired, rawInitial *LogMetric, opts ...dcl.ApplyOption) (*LogMetric, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MetricDescriptor = canonicalizeLogMetricMetricDescriptor(rawDesired.MetricDescriptor, nil, opts...)
		rawDesired.BucketOptions = canonicalizeLogMetricBucketOptions(rawDesired.BucketOptions, nil, opts...)

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.Filter, rawInitial.Filter) {
		rawDesired.Filter = rawInitial.Filter
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		rawDesired.Disabled = rawInitial.Disabled
	}
	rawDesired.MetricDescriptor = canonicalizeLogMetricMetricDescriptor(rawDesired.MetricDescriptor, rawInitial.MetricDescriptor, opts...)
	if dcl.StringCanonicalize(rawDesired.ValueExtractor, rawInitial.ValueExtractor) {
		rawDesired.ValueExtractor = rawInitial.ValueExtractor
	}
	if dcl.IsZeroValue(rawDesired.LabelExtractors) {
		rawDesired.LabelExtractors = rawInitial.LabelExtractors
	}
	rawDesired.BucketOptions = canonicalizeLogMetricBucketOptions(rawDesired.BucketOptions, rawInitial.BucketOptions, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeLogMetricNewState(c *Client, rawNew, rawDesired *LogMetric) (*LogMetric, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Filter) && dcl.IsEmptyValueIndirect(rawDesired.Filter) {
		rawNew.Filter = rawDesired.Filter
	} else {
		if dcl.StringCanonicalize(rawDesired.Filter, rawNew.Filter) {
			rawNew.Filter = rawDesired.Filter
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MetricDescriptor) && dcl.IsEmptyValueIndirect(rawDesired.MetricDescriptor) {
		rawNew.MetricDescriptor = rawDesired.MetricDescriptor
	} else {
		rawNew.MetricDescriptor = canonicalizeNewLogMetricMetricDescriptor(c, rawDesired.MetricDescriptor, rawNew.MetricDescriptor)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ValueExtractor) && dcl.IsEmptyValueIndirect(rawDesired.ValueExtractor) {
		rawNew.ValueExtractor = rawDesired.ValueExtractor
	} else {
		if dcl.StringCanonicalize(rawDesired.ValueExtractor, rawNew.ValueExtractor) {
			rawNew.ValueExtractor = rawDesired.ValueExtractor
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LabelExtractors) && dcl.IsEmptyValueIndirect(rawDesired.LabelExtractors) {
		rawNew.LabelExtractors = rawDesired.LabelExtractors
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BucketOptions) && dcl.IsEmptyValueIndirect(rawDesired.BucketOptions) {
		rawNew.BucketOptions = rawDesired.BucketOptions
	} else {
		rawNew.BucketOptions = canonicalizeNewLogMetricBucketOptions(c, rawDesired.BucketOptions, rawNew.BucketOptions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeLogMetricMetricDescriptor(des, initial *LogMetricMetricDescriptor, opts ...dcl.ApplyOption) *LogMetricMetricDescriptor {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.MetricKind) {
		des.MetricKind = initial.MetricKind
	}
	if canonicalizeLogMetricMetricDescriptorValueType(des.ValueType, initial.ValueType) || dcl.IsZeroValue(des.ValueType) {
		des.ValueType = initial.ValueType
	}
	if dcl.StringCanonicalize(des.Unit, initial.Unit) || dcl.IsZeroValue(des.Unit) {
		des.Unit = initial.Unit
	}
	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		des.DisplayName = initial.DisplayName
	}
	des.Metadata = canonicalizeLogMetricMetricDescriptorMetadata(des.Metadata, initial.Metadata, opts...)
	if dcl.IsZeroValue(des.LaunchStage) {
		des.LaunchStage = initial.LaunchStage
	}

	return des
}

func canonicalizeNewLogMetricMetricDescriptor(c *Client, des, nw *LogMetricMetricDescriptor) *LogMetricMetricDescriptor {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	nw.Labels = canonicalizeNewLogMetricMetricDescriptorLabelsSet(c, des.Labels, nw.Labels)
	if dcl.IsZeroValue(nw.MetricKind) {
		nw.MetricKind = des.MetricKind
	}
	if canonicalizeLogMetricMetricDescriptorValueType(des.ValueType, nw.ValueType) {
		nw.ValueType = des.ValueType
	}
	if dcl.StringCanonicalize(des.Unit, nw.Unit) {
		nw.Unit = des.Unit
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}
	nw.Metadata = canonicalizeNewLogMetricMetricDescriptorMetadata(c, des.Metadata, nw.Metadata)
	nw.LaunchStage = des.LaunchStage
	if dcl.IsZeroValue(nw.MonitoredResourceTypes) {
		nw.MonitoredResourceTypes = des.MonitoredResourceTypes
	}

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorSet(c *Client, des, nw []LogMetricMetricDescriptor) []LogMetricMetricDescriptor {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptor
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricMetricDescriptorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricMetricDescriptorSlice(c *Client, des, nw []LogMetricMetricDescriptor) []LogMetricMetricDescriptor {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricMetricDescriptor
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricMetricDescriptor(c, &d, &n))
	}

	return items
}

func canonicalizeLogMetricMetricDescriptorLabels(des, initial *LogMetricMetricDescriptorLabels, opts ...dcl.ApplyOption) *LogMetricMetricDescriptorLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if canonicalizeLogMetricMetricDescriptorLabelsValueType(des.ValueType, initial.ValueType) || dcl.IsZeroValue(des.ValueType) {
		des.ValueType = initial.ValueType
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewLogMetricMetricDescriptorLabels(c *Client, des, nw *LogMetricMetricDescriptorLabels) *LogMetricMetricDescriptorLabels {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if canonicalizeLogMetricMetricDescriptorLabelsValueType(des.ValueType, nw.ValueType) {
		nw.ValueType = des.ValueType
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorLabelsSet(c *Client, des, nw []LogMetricMetricDescriptorLabels) []LogMetricMetricDescriptorLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptorLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricMetricDescriptorLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricMetricDescriptorLabelsSlice(c *Client, des, nw []LogMetricMetricDescriptorLabels) []LogMetricMetricDescriptorLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricMetricDescriptorLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricMetricDescriptorLabels(c, &d, &n))
	}

	return items
}

func canonicalizeLogMetricMetricDescriptorMetadata(des, initial *LogMetricMetricDescriptorMetadata, opts ...dcl.ApplyOption) *LogMetricMetricDescriptorMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.SamplePeriod, initial.SamplePeriod) || dcl.IsZeroValue(des.SamplePeriod) {
		des.SamplePeriod = initial.SamplePeriod
	}
	if dcl.StringCanonicalize(des.IngestDelay, initial.IngestDelay) || dcl.IsZeroValue(des.IngestDelay) {
		des.IngestDelay = initial.IngestDelay
	}

	return des
}

func canonicalizeNewLogMetricMetricDescriptorMetadata(c *Client, des, nw *LogMetricMetricDescriptorMetadata) *LogMetricMetricDescriptorMetadata {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.SamplePeriod, nw.SamplePeriod) {
		nw.SamplePeriod = des.SamplePeriod
	}
	if dcl.StringCanonicalize(des.IngestDelay, nw.IngestDelay) {
		nw.IngestDelay = des.IngestDelay
	}

	return nw
}

func canonicalizeNewLogMetricMetricDescriptorMetadataSet(c *Client, des, nw []LogMetricMetricDescriptorMetadata) []LogMetricMetricDescriptorMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricMetricDescriptorMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricMetricDescriptorMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricMetricDescriptorMetadataSlice(c *Client, des, nw []LogMetricMetricDescriptorMetadata) []LogMetricMetricDescriptorMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricMetricDescriptorMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricMetricDescriptorMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeLogMetricBucketOptions(des, initial *LogMetricBucketOptions, opts ...dcl.ApplyOption) *LogMetricBucketOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.LinearBuckets = canonicalizeLogMetricBucketOptionsLinearBuckets(des.LinearBuckets, initial.LinearBuckets, opts...)
	des.ExponentialBuckets = canonicalizeLogMetricBucketOptionsExponentialBuckets(des.ExponentialBuckets, initial.ExponentialBuckets, opts...)
	des.ExplicitBuckets = canonicalizeLogMetricBucketOptionsExplicitBuckets(des.ExplicitBuckets, initial.ExplicitBuckets, opts...)

	return des
}

func canonicalizeNewLogMetricBucketOptions(c *Client, des, nw *LogMetricBucketOptions) *LogMetricBucketOptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.LinearBuckets = canonicalizeNewLogMetricBucketOptionsLinearBuckets(c, des.LinearBuckets, nw.LinearBuckets)
	nw.ExponentialBuckets = canonicalizeNewLogMetricBucketOptionsExponentialBuckets(c, des.ExponentialBuckets, nw.ExponentialBuckets)
	nw.ExplicitBuckets = canonicalizeNewLogMetricBucketOptionsExplicitBuckets(c, des.ExplicitBuckets, nw.ExplicitBuckets)

	return nw
}

func canonicalizeNewLogMetricBucketOptionsSet(c *Client, des, nw []LogMetricBucketOptions) []LogMetricBucketOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricBucketOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricBucketOptionsSlice(c *Client, des, nw []LogMetricBucketOptions) []LogMetricBucketOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricBucketOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricBucketOptions(c, &d, &n))
	}

	return items
}

func canonicalizeLogMetricBucketOptionsLinearBuckets(des, initial *LogMetricBucketOptionsLinearBuckets, opts ...dcl.ApplyOption) *LogMetricBucketOptionsLinearBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.Width) {
		des.Width = initial.Width
	}
	if dcl.IsZeroValue(des.Offset) {
		des.Offset = initial.Offset
	}

	return des
}

func canonicalizeNewLogMetricBucketOptionsLinearBuckets(c *Client, des, nw *LogMetricBucketOptionsLinearBuckets) *LogMetricBucketOptionsLinearBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.Width) {
		nw.Width = des.Width
	}
	if dcl.IsZeroValue(nw.Offset) {
		nw.Offset = des.Offset
	}

	return nw
}

func canonicalizeNewLogMetricBucketOptionsLinearBucketsSet(c *Client, des, nw []LogMetricBucketOptionsLinearBuckets) []LogMetricBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptionsLinearBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricBucketOptionsLinearBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricBucketOptionsLinearBucketsSlice(c *Client, des, nw []LogMetricBucketOptionsLinearBuckets) []LogMetricBucketOptionsLinearBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricBucketOptionsLinearBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricBucketOptionsLinearBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeLogMetricBucketOptionsExponentialBuckets(des, initial *LogMetricBucketOptionsExponentialBuckets, opts ...dcl.ApplyOption) *LogMetricBucketOptionsExponentialBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.NumFiniteBuckets) {
		des.NumFiniteBuckets = initial.NumFiniteBuckets
	}
	if dcl.IsZeroValue(des.GrowthFactor) {
		des.GrowthFactor = initial.GrowthFactor
	}
	if dcl.IsZeroValue(des.Scale) {
		des.Scale = initial.Scale
	}

	return des
}

func canonicalizeNewLogMetricBucketOptionsExponentialBuckets(c *Client, des, nw *LogMetricBucketOptionsExponentialBuckets) *LogMetricBucketOptionsExponentialBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.NumFiniteBuckets) {
		nw.NumFiniteBuckets = des.NumFiniteBuckets
	}
	if dcl.IsZeroValue(nw.GrowthFactor) {
		nw.GrowthFactor = des.GrowthFactor
	}
	if dcl.IsZeroValue(nw.Scale) {
		nw.Scale = des.Scale
	}

	return nw
}

func canonicalizeNewLogMetricBucketOptionsExponentialBucketsSet(c *Client, des, nw []LogMetricBucketOptionsExponentialBuckets) []LogMetricBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptionsExponentialBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricBucketOptionsExponentialBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricBucketOptionsExponentialBucketsSlice(c *Client, des, nw []LogMetricBucketOptionsExponentialBuckets) []LogMetricBucketOptionsExponentialBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricBucketOptionsExponentialBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricBucketOptionsExponentialBuckets(c, &d, &n))
	}

	return items
}

func canonicalizeLogMetricBucketOptionsExplicitBuckets(des, initial *LogMetricBucketOptionsExplicitBuckets, opts ...dcl.ApplyOption) *LogMetricBucketOptionsExplicitBuckets {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Bounds) {
		des.Bounds = initial.Bounds
	}

	return des
}

func canonicalizeNewLogMetricBucketOptionsExplicitBuckets(c *Client, des, nw *LogMetricBucketOptionsExplicitBuckets) *LogMetricBucketOptionsExplicitBuckets {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Bounds) {
		nw.Bounds = des.Bounds
	}

	return nw
}

func canonicalizeNewLogMetricBucketOptionsExplicitBucketsSet(c *Client, des, nw []LogMetricBucketOptionsExplicitBuckets) []LogMetricBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}
	var reorderedNew []LogMetricBucketOptionsExplicitBuckets
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareLogMetricBucketOptionsExplicitBucketsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewLogMetricBucketOptionsExplicitBucketsSlice(c *Client, des, nw []LogMetricBucketOptionsExplicitBuckets) []LogMetricBucketOptionsExplicitBuckets {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []LogMetricBucketOptionsExplicitBuckets
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewLogMetricBucketOptionsExplicitBuckets(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffLogMetric(c *Client, desired, actual *LogMetric, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MetricDescriptor, actual.MetricDescriptor, dcl.Info{ObjectFunction: compareLogMetricMetricDescriptorNewStyle, EmptyObject: EmptyLogMetricMetricDescriptor, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetricDescriptor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueExtractor, actual.ValueExtractor, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("ValueExtractor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LabelExtractors, actual.LabelExtractors, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("LabelExtractors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BucketOptions, actual.BucketOptions, dcl.Info{ObjectFunction: compareLogMetricBucketOptionsNewStyle, EmptyObject: EmptyLogMetricBucketOptions, OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("BucketOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareLogMetricMetricDescriptorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricMetricDescriptor)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricMetricDescriptor)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricMetricDescriptor or *LogMetricMetricDescriptor", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricMetricDescriptor)
	if !ok {
		actualNotPointer, ok := a.(LogMetricMetricDescriptor)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricMetricDescriptor", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{Type: "Set", ObjectFunction: compareLogMetricMetricDescriptorLabelsNewStyle, EmptyObject: EmptyLogMetricMetricDescriptorLabels, OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MetricKind, actual.MetricKind, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetricKind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueType, actual.ValueType, dcl.Info{Type: "EnumType", CustomDiff: canonicalizeLogMetricMetricDescriptorValueType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ValueType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Unit, actual.Unit, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Unit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{Ignore: true, ObjectFunction: compareLogMetricMetricDescriptorMetadataNewStyle, EmptyObject: EmptyLogMetricMetricDescriptorMetadata, OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LaunchStage, actual.LaunchStage, dcl.Info{Ignore: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("LaunchStage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MonitoredResourceTypes, actual.MonitoredResourceTypes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MonitoredResourceTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareLogMetricMetricDescriptorLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricMetricDescriptorLabels)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricMetricDescriptorLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricMetricDescriptorLabels or *LogMetricMetricDescriptorLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricMetricDescriptorLabels)
	if !ok {
		actualNotPointer, ok := a.(LogMetricMetricDescriptorLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricMetricDescriptorLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueType, actual.ValueType, dcl.Info{Type: "EnumType", CustomDiff: canonicalizeLogMetricMetricDescriptorLabelsValueType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ValueType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareLogMetricMetricDescriptorMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricMetricDescriptorMetadata)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricMetricDescriptorMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricMetricDescriptorMetadata or *LogMetricMetricDescriptorMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricMetricDescriptorMetadata)
	if !ok {
		actualNotPointer, ok := a.(LogMetricMetricDescriptorMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricMetricDescriptorMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SamplePeriod, actual.SamplePeriod, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("SamplePeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IngestDelay, actual.IngestDelay, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogMetricUpdateOperation")}, fn.AddNest("IngestDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareLogMetricBucketOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricBucketOptions)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptions or *LogMetricBucketOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricBucketOptions)
	if !ok {
		actualNotPointer, ok := a.(LogMetricBucketOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LinearBuckets, actual.LinearBuckets, dcl.Info{ObjectFunction: compareLogMetricBucketOptionsLinearBucketsNewStyle, EmptyObject: EmptyLogMetricBucketOptionsLinearBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LinearBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExponentialBuckets, actual.ExponentialBuckets, dcl.Info{ObjectFunction: compareLogMetricBucketOptionsExponentialBucketsNewStyle, EmptyObject: EmptyLogMetricBucketOptionsExponentialBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExponentialBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExplicitBuckets, actual.ExplicitBuckets, dcl.Info{ObjectFunction: compareLogMetricBucketOptionsExplicitBucketsNewStyle, EmptyObject: EmptyLogMetricBucketOptionsExplicitBuckets, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExplicitBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareLogMetricBucketOptionsLinearBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricBucketOptionsLinearBuckets)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptionsLinearBuckets or *LogMetricBucketOptionsLinearBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricBucketOptionsLinearBuckets)
	if !ok {
		actualNotPointer, ok := a.(LogMetricBucketOptionsLinearBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptionsLinearBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Width, actual.Width, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Width")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Offset, actual.Offset, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Offset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareLogMetricBucketOptionsExponentialBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricBucketOptionsExponentialBuckets)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptionsExponentialBuckets or *LogMetricBucketOptionsExponentialBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricBucketOptionsExponentialBuckets)
	if !ok {
		actualNotPointer, ok := a.(LogMetricBucketOptionsExponentialBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptionsExponentialBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NumFiniteBuckets, actual.NumFiniteBuckets, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumFiniteBuckets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrowthFactor, actual.GrowthFactor, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GrowthFactor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareLogMetricBucketOptionsExplicitBucketsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*LogMetricBucketOptionsExplicitBuckets)
	if !ok {
		desiredNotPointer, ok := d.(LogMetricBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptionsExplicitBuckets or *LogMetricBucketOptionsExplicitBuckets", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*LogMetricBucketOptionsExplicitBuckets)
	if !ok {
		actualNotPointer, ok := a.(LogMetricBucketOptionsExplicitBuckets)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a LogMetricBucketOptionsExplicitBuckets", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Bounds, actual.Bounds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *LogMetric) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *LogMetric) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *LogMetric) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *LogMetric) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/metrics/{{name}}", "https://logging.googleapis.com/v2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the LogMetric resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *LogMetric) marshal(c *Client) ([]byte, error) {
	m, err := expandLogMetric(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling LogMetric: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalLogMetric decodes JSON responses into the LogMetric resource schema.
func unmarshalLogMetric(b []byte, c *Client) (*LogMetric, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapLogMetric(m, c)
}

func unmarshalMapLogMetric(m map[string]interface{}, c *Client) (*LogMetric, error) {

	return flattenLogMetric(c, m), nil
}

// expandLogMetric expands LogMetric into a JSON request object.
func expandLogMetric(c *Client, f *LogMetric) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v, err := expandLogMetricMetricDescriptor(c, f.MetricDescriptor); err != nil {
		return nil, fmt.Errorf("error expanding MetricDescriptor into metricDescriptor: %w", err)
	} else if v != nil {
		m["metricDescriptor"] = v
	}
	if v := f.ValueExtractor; !dcl.IsEmptyValueIndirect(v) {
		m["valueExtractor"] = v
	}
	if v := f.LabelExtractors; !dcl.IsEmptyValueIndirect(v) {
		m["labelExtractors"] = v
	}
	if v, err := expandLogMetricBucketOptions(c, f.BucketOptions); err != nil {
		return nil, fmt.Errorf("error expanding BucketOptions into bucketOptions: %w", err)
	} else if v != nil {
		m["bucketOptions"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenLogMetric flattens LogMetric from a JSON request object into the
// LogMetric type.
func flattenLogMetric(c *Client, i interface{}) *LogMetric {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &LogMetric{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Filter = dcl.FlattenString(m["filter"])
	res.Disabled = dcl.FlattenBool(m["disabled"])
	res.MetricDescriptor = flattenLogMetricMetricDescriptor(c, m["metricDescriptor"])
	res.ValueExtractor = dcl.FlattenString(m["valueExtractor"])
	res.LabelExtractors = dcl.FlattenKeyValuePairs(m["labelExtractors"])
	res.BucketOptions = flattenLogMetricBucketOptions(c, m["bucketOptions"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandLogMetricMetricDescriptorMap expands the contents of LogMetricMetricDescriptor into a JSON
// request object.
func expandLogMetricMetricDescriptorMap(c *Client, f map[string]LogMetricMetricDescriptor) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptor(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorSlice expands the contents of LogMetricMetricDescriptor into a JSON
// request object.
func expandLogMetricMetricDescriptorSlice(c *Client, f []LogMetricMetricDescriptor) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptor(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorMap flattens the contents of LogMetricMetricDescriptor from a JSON
// response object.
func flattenLogMetricMetricDescriptorMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptor {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptor{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptor{}
	}

	items := make(map[string]LogMetricMetricDescriptor)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptor(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorSlice flattens the contents of LogMetricMetricDescriptor from a JSON
// response object.
func flattenLogMetricMetricDescriptorSlice(c *Client, i interface{}) []LogMetricMetricDescriptor {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptor{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptor{}
	}

	items := make([]LogMetricMetricDescriptor, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptor(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptor expands an instance of LogMetricMetricDescriptor into a JSON
// request object.
func expandLogMetricMetricDescriptor(c *Client, f *LogMetricMetricDescriptor) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v, err := expandLogMetricMetricDescriptorLabelsSlice(c, f.Labels); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.MetricKind; !dcl.IsEmptyValueIndirect(v) {
		m["metricKind"] = v
	}
	if v := f.ValueType; !dcl.IsEmptyValueIndirect(v) {
		m["valueType"] = v
	}
	if v := f.Unit; !dcl.IsEmptyValueIndirect(v) {
		m["unit"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandLogMetricMetricDescriptorMetadata(c, f.Metadata); err != nil {
		return nil, fmt.Errorf("error expanding Metadata into metadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		m["launchStage"] = v
	}
	if v := f.MonitoredResourceTypes; !dcl.IsEmptyValueIndirect(v) {
		m["monitoredResourceTypes"] = v
	}

	return m, nil
}

// flattenLogMetricMetricDescriptor flattens an instance of LogMetricMetricDescriptor from a JSON
// response object.
func flattenLogMetricMetricDescriptor(c *Client, i interface{}) *LogMetricMetricDescriptor {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptor{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricMetricDescriptor
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Type = dcl.FlattenString(m["type"])
	r.Labels = flattenLogMetricMetricDescriptorLabelsSlice(c, m["labels"])
	r.MetricKind = flattenLogMetricMetricDescriptorMetricKindEnum(m["metricKind"])
	r.ValueType = flattenLogMetricMetricDescriptorValueTypeEnum(m["valueType"])
	r.Unit = dcl.FlattenString(m["unit"])
	r.Description = dcl.FlattenString(m["description"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Metadata = flattenLogMetricMetricDescriptorMetadata(c, m["metadata"])
	r.LaunchStage = flattenLogMetricMetricDescriptorLaunchStageEnum(m["launchStage"])
	r.MonitoredResourceTypes = dcl.FlattenStringSlice(m["monitoredResourceTypes"])

	return r
}

// expandLogMetricMetricDescriptorLabelsMap expands the contents of LogMetricMetricDescriptorLabels into a JSON
// request object.
func expandLogMetricMetricDescriptorLabelsMap(c *Client, f map[string]LogMetricMetricDescriptorLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptorLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorLabelsSlice expands the contents of LogMetricMetricDescriptorLabels into a JSON
// request object.
func expandLogMetricMetricDescriptorLabelsSlice(c *Client, f []LogMetricMetricDescriptorLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptorLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorLabelsMap flattens the contents of LogMetricMetricDescriptorLabels from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabelsMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptorLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptorLabels{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptorLabels{}
	}

	items := make(map[string]LogMetricMetricDescriptorLabels)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptorLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorLabelsSlice flattens the contents of LogMetricMetricDescriptorLabels from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabelsSlice(c *Client, i interface{}) []LogMetricMetricDescriptorLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorLabels{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorLabels{}
	}

	items := make([]LogMetricMetricDescriptorLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptorLabels expands an instance of LogMetricMetricDescriptorLabels into a JSON
// request object.
func expandLogMetricMetricDescriptorLabels(c *Client, f *LogMetricMetricDescriptorLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.ValueType; !dcl.IsEmptyValueIndirect(v) {
		m["valueType"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenLogMetricMetricDescriptorLabels flattens an instance of LogMetricMetricDescriptorLabels from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabels(c *Client, i interface{}) *LogMetricMetricDescriptorLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptorLabels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricMetricDescriptorLabels
	}
	r.Key = dcl.FlattenString(m["key"])
	r.ValueType = flattenLogMetricMetricDescriptorLabelsValueTypeEnum(m["valueType"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandLogMetricMetricDescriptorMetadataMap expands the contents of LogMetricMetricDescriptorMetadata into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataMap(c *Client, f map[string]LogMetricMetricDescriptorMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricMetricDescriptorMetadataSlice expands the contents of LogMetricMetricDescriptorMetadata into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadataSlice(c *Client, f []LogMetricMetricDescriptorMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricMetricDescriptorMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricMetricDescriptorMetadataMap flattens the contents of LogMetricMetricDescriptorMetadata from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataMap(c *Client, i interface{}) map[string]LogMetricMetricDescriptorMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricMetricDescriptorMetadata{}
	}

	if len(a) == 0 {
		return map[string]LogMetricMetricDescriptorMetadata{}
	}

	items := make(map[string]LogMetricMetricDescriptorMetadata)
	for k, item := range a {
		items[k] = *flattenLogMetricMetricDescriptorMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetadataSlice flattens the contents of LogMetricMetricDescriptorMetadata from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadataSlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetadata{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetadata{}
	}

	items := make([]LogMetricMetricDescriptorMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricMetricDescriptorMetadata expands an instance of LogMetricMetricDescriptorMetadata into a JSON
// request object.
func expandLogMetricMetricDescriptorMetadata(c *Client, f *LogMetricMetricDescriptorMetadata) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SamplePeriod; !dcl.IsEmptyValueIndirect(v) {
		m["samplePeriod"] = v
	}
	if v := f.IngestDelay; !dcl.IsEmptyValueIndirect(v) {
		m["ingestDelay"] = v
	}

	return m, nil
}

// flattenLogMetricMetricDescriptorMetadata flattens an instance of LogMetricMetricDescriptorMetadata from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetadata(c *Client, i interface{}) *LogMetricMetricDescriptorMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricMetricDescriptorMetadata{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricMetricDescriptorMetadata
	}
	r.SamplePeriod = dcl.FlattenString(m["samplePeriod"])
	r.IngestDelay = dcl.FlattenString(m["ingestDelay"])

	return r
}

// expandLogMetricBucketOptionsMap expands the contents of LogMetricBucketOptions into a JSON
// request object.
func expandLogMetricBucketOptionsMap(c *Client, f map[string]LogMetricBucketOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsSlice expands the contents of LogMetricBucketOptions into a JSON
// request object.
func expandLogMetricBucketOptionsSlice(c *Client, f []LogMetricBucketOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsMap flattens the contents of LogMetricBucketOptions from a JSON
// response object.
func flattenLogMetricBucketOptionsMap(c *Client, i interface{}) map[string]LogMetricBucketOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptions{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptions{}
	}

	items := make(map[string]LogMetricBucketOptions)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsSlice flattens the contents of LogMetricBucketOptions from a JSON
// response object.
func flattenLogMetricBucketOptionsSlice(c *Client, i interface{}) []LogMetricBucketOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptions{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptions{}
	}

	items := make([]LogMetricBucketOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptions expands an instance of LogMetricBucketOptions into a JSON
// request object.
func expandLogMetricBucketOptions(c *Client, f *LogMetricBucketOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandLogMetricBucketOptionsLinearBuckets(c, f.LinearBuckets); err != nil {
		return nil, fmt.Errorf("error expanding LinearBuckets into linearBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["linearBuckets"] = v
	}
	if v, err := expandLogMetricBucketOptionsExponentialBuckets(c, f.ExponentialBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExponentialBuckets into exponentialBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exponentialBuckets"] = v
	}
	if v, err := expandLogMetricBucketOptionsExplicitBuckets(c, f.ExplicitBuckets); err != nil {
		return nil, fmt.Errorf("error expanding ExplicitBuckets into explicitBuckets: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["explicitBuckets"] = v
	}

	return m, nil
}

// flattenLogMetricBucketOptions flattens an instance of LogMetricBucketOptions from a JSON
// response object.
func flattenLogMetricBucketOptions(c *Client, i interface{}) *LogMetricBucketOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricBucketOptions
	}
	r.LinearBuckets = flattenLogMetricBucketOptionsLinearBuckets(c, m["linearBuckets"])
	r.ExponentialBuckets = flattenLogMetricBucketOptionsExponentialBuckets(c, m["exponentialBuckets"])
	r.ExplicitBuckets = flattenLogMetricBucketOptionsExplicitBuckets(c, m["explicitBuckets"])

	return r
}

// expandLogMetricBucketOptionsLinearBucketsMap expands the contents of LogMetricBucketOptionsLinearBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsLinearBucketsMap(c *Client, f map[string]LogMetricBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsLinearBucketsSlice expands the contents of LogMetricBucketOptionsLinearBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsLinearBucketsSlice(c *Client, f []LogMetricBucketOptionsLinearBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptionsLinearBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsLinearBucketsMap flattens the contents of LogMetricBucketOptionsLinearBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsLinearBucketsMap(c *Client, i interface{}) map[string]LogMetricBucketOptionsLinearBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptionsLinearBuckets{}
	}

	items := make(map[string]LogMetricBucketOptionsLinearBuckets)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptionsLinearBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsLinearBucketsSlice flattens the contents of LogMetricBucketOptionsLinearBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsLinearBucketsSlice(c *Client, i interface{}) []LogMetricBucketOptionsLinearBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptionsLinearBuckets{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptionsLinearBuckets{}
	}

	items := make([]LogMetricBucketOptionsLinearBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptionsLinearBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptionsLinearBuckets expands an instance of LogMetricBucketOptionsLinearBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsLinearBuckets(c *Client, f *LogMetricBucketOptionsLinearBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.Width; !dcl.IsEmptyValueIndirect(v) {
		m["width"] = v
	}
	if v := f.Offset; !dcl.IsEmptyValueIndirect(v) {
		m["offset"] = v
	}

	return m, nil
}

// flattenLogMetricBucketOptionsLinearBuckets flattens an instance of LogMetricBucketOptionsLinearBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsLinearBuckets(c *Client, i interface{}) *LogMetricBucketOptionsLinearBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptionsLinearBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricBucketOptionsLinearBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.Width = dcl.FlattenDouble(m["width"])
	r.Offset = dcl.FlattenDouble(m["offset"])

	return r
}

// expandLogMetricBucketOptionsExponentialBucketsMap expands the contents of LogMetricBucketOptionsExponentialBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExponentialBucketsMap(c *Client, f map[string]LogMetricBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsExponentialBucketsSlice expands the contents of LogMetricBucketOptionsExponentialBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExponentialBucketsSlice(c *Client, f []LogMetricBucketOptionsExponentialBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptionsExponentialBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsExponentialBucketsMap flattens the contents of LogMetricBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExponentialBucketsMap(c *Client, i interface{}) map[string]LogMetricBucketOptionsExponentialBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptionsExponentialBuckets{}
	}

	items := make(map[string]LogMetricBucketOptionsExponentialBuckets)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptionsExponentialBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsExponentialBucketsSlice flattens the contents of LogMetricBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExponentialBucketsSlice(c *Client, i interface{}) []LogMetricBucketOptionsExponentialBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptionsExponentialBuckets{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptionsExponentialBuckets{}
	}

	items := make([]LogMetricBucketOptionsExponentialBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptionsExponentialBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptionsExponentialBuckets expands an instance of LogMetricBucketOptionsExponentialBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExponentialBuckets(c *Client, f *LogMetricBucketOptionsExponentialBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NumFiniteBuckets; !dcl.IsEmptyValueIndirect(v) {
		m["numFiniteBuckets"] = v
	}
	if v := f.GrowthFactor; !dcl.IsEmptyValueIndirect(v) {
		m["growthFactor"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}

	return m, nil
}

// flattenLogMetricBucketOptionsExponentialBuckets flattens an instance of LogMetricBucketOptionsExponentialBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExponentialBuckets(c *Client, i interface{}) *LogMetricBucketOptionsExponentialBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptionsExponentialBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricBucketOptionsExponentialBuckets
	}
	r.NumFiniteBuckets = dcl.FlattenInteger(m["numFiniteBuckets"])
	r.GrowthFactor = dcl.FlattenDouble(m["growthFactor"])
	r.Scale = dcl.FlattenDouble(m["scale"])

	return r
}

// expandLogMetricBucketOptionsExplicitBucketsMap expands the contents of LogMetricBucketOptionsExplicitBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExplicitBucketsMap(c *Client, f map[string]LogMetricBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandLogMetricBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandLogMetricBucketOptionsExplicitBucketsSlice expands the contents of LogMetricBucketOptionsExplicitBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExplicitBucketsSlice(c *Client, f []LogMetricBucketOptionsExplicitBuckets) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandLogMetricBucketOptionsExplicitBuckets(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenLogMetricBucketOptionsExplicitBucketsMap flattens the contents of LogMetricBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExplicitBucketsMap(c *Client, i interface{}) map[string]LogMetricBucketOptionsExplicitBuckets {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]LogMetricBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return map[string]LogMetricBucketOptionsExplicitBuckets{}
	}

	items := make(map[string]LogMetricBucketOptionsExplicitBuckets)
	for k, item := range a {
		items[k] = *flattenLogMetricBucketOptionsExplicitBuckets(c, item.(map[string]interface{}))
	}

	return items
}

// flattenLogMetricBucketOptionsExplicitBucketsSlice flattens the contents of LogMetricBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExplicitBucketsSlice(c *Client, i interface{}) []LogMetricBucketOptionsExplicitBuckets {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricBucketOptionsExplicitBuckets{}
	}

	if len(a) == 0 {
		return []LogMetricBucketOptionsExplicitBuckets{}
	}

	items := make([]LogMetricBucketOptionsExplicitBuckets, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricBucketOptionsExplicitBuckets(c, item.(map[string]interface{})))
	}

	return items
}

// expandLogMetricBucketOptionsExplicitBuckets expands an instance of LogMetricBucketOptionsExplicitBuckets into a JSON
// request object.
func expandLogMetricBucketOptionsExplicitBuckets(c *Client, f *LogMetricBucketOptionsExplicitBuckets) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Bounds; !dcl.IsEmptyValueIndirect(v) {
		m["bounds"] = v
	}

	return m, nil
}

// flattenLogMetricBucketOptionsExplicitBuckets flattens an instance of LogMetricBucketOptionsExplicitBuckets from a JSON
// response object.
func flattenLogMetricBucketOptionsExplicitBuckets(c *Client, i interface{}) *LogMetricBucketOptionsExplicitBuckets {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &LogMetricBucketOptionsExplicitBuckets{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyLogMetricBucketOptionsExplicitBuckets
	}
	r.Bounds = dcl.FlattenFloatSlice(m["bounds"])

	return r
}

// flattenLogMetricMetricDescriptorLabelsValueTypeEnumSlice flattens the contents of LogMetricMetricDescriptorLabelsValueTypeEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorLabelsValueTypeEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorLabelsValueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorLabelsValueTypeEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorLabelsValueTypeEnum{}
	}

	items := make([]LogMetricMetricDescriptorLabelsValueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorLabelsValueTypeEnum(item.(interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorLabelsValueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorLabelsValueTypeEnum with the same value as that string.
func flattenLogMetricMetricDescriptorLabelsValueTypeEnum(i interface{}) *LogMetricMetricDescriptorLabelsValueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorLabelsValueTypeEnumRef("")
	}

	return LogMetricMetricDescriptorLabelsValueTypeEnumRef(s)
}

// flattenLogMetricMetricDescriptorMetricKindEnumSlice flattens the contents of LogMetricMetricDescriptorMetricKindEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorMetricKindEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorMetricKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorMetricKindEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorMetricKindEnum{}
	}

	items := make([]LogMetricMetricDescriptorMetricKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorMetricKindEnum(item.(interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorMetricKindEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorMetricKindEnum with the same value as that string.
func flattenLogMetricMetricDescriptorMetricKindEnum(i interface{}) *LogMetricMetricDescriptorMetricKindEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorMetricKindEnumRef("")
	}

	return LogMetricMetricDescriptorMetricKindEnumRef(s)
}

// flattenLogMetricMetricDescriptorValueTypeEnumSlice flattens the contents of LogMetricMetricDescriptorValueTypeEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorValueTypeEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorValueTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorValueTypeEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorValueTypeEnum{}
	}

	items := make([]LogMetricMetricDescriptorValueTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorValueTypeEnum(item.(interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorValueTypeEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorValueTypeEnum with the same value as that string.
func flattenLogMetricMetricDescriptorValueTypeEnum(i interface{}) *LogMetricMetricDescriptorValueTypeEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorValueTypeEnumRef("")
	}

	return LogMetricMetricDescriptorValueTypeEnumRef(s)
}

// flattenLogMetricMetricDescriptorLaunchStageEnumSlice flattens the contents of LogMetricMetricDescriptorLaunchStageEnum from a JSON
// response object.
func flattenLogMetricMetricDescriptorLaunchStageEnumSlice(c *Client, i interface{}) []LogMetricMetricDescriptorLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []LogMetricMetricDescriptorLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []LogMetricMetricDescriptorLaunchStageEnum{}
	}

	items := make([]LogMetricMetricDescriptorLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenLogMetricMetricDescriptorLaunchStageEnum(item.(interface{})))
	}

	return items
}

// flattenLogMetricMetricDescriptorLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *LogMetricMetricDescriptorLaunchStageEnum with the same value as that string.
func flattenLogMetricMetricDescriptorLaunchStageEnum(i interface{}) *LogMetricMetricDescriptorLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return LogMetricMetricDescriptorLaunchStageEnumRef("")
	}

	return LogMetricMetricDescriptorLaunchStageEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *LogMetric) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalLogMetric(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type logMetricDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         logMetricApiOperation
}

func convertFieldDiffToLogMetricOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]logMetricDiff, error) {
	var diffs []logMetricDiff
	for _, op := range ops {
		diff := logMetricDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTologMetricApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTologMetricApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (logMetricApiOperation, error) {
	switch op {

	case "updateLogMetricUpdateOperation":
		return &updateLogMetricUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
