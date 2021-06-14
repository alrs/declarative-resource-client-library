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
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Backup) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}

func backupGetURL(userBasePath string, r *Backup) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups/{{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil
}

func backupListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil

}

func backupCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups?backupId={{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil

}

func backupDeleteURL(userBasePath string, r *Backup) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups/{{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil
}

// backupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type backupApiOperation interface {
	do(context.Context, *Backup, *Client) error
}

// newUpdateBackupUpdateBackupRequest creates a request for an
// Backup resource's UpdateBackup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBackupUpdateBackupRequest(ctx context.Context, f *Backup, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateBackupUpdateBackupRequest converts the update into
// the final JSON request body.
func marshalUpdateBackupUpdateBackupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBackupUpdateBackupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBackupUpdateBackupOperation) do(ctx context.Context, r *Backup, c *Client) error {
	_, err := c.GetBackup(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateBackup")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.Diffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateBackupUpdateBackupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateBackupUpdateBackupRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://file.googleapis.com/v1beta1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listBackupRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := backupListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BackupMaxPage {
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

type listBackupOperation struct {
	Backups []map[string]interface{} `json:"backups"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listBackup(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Backup, string, error) {
	b, err := c.listBackupRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBackupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Backup
	for _, v := range m.Backups {
		res, err := unmarshalMapBackup(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBackup(ctx context.Context, f func(*Backup) bool, resources []*Backup) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBackup(ctx, res)
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

type deleteBackupOperation struct{}

func (op *deleteBackupOperation) do(ctx context.Context, r *Backup, c *Client) error {

	_, err := c.GetBackup(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Backup not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetBackup checking for existence. error: %v", err)
		return err
	}

	u, err := backupDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://file.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetBackup(ctx, r.urlNormalized())
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
type createBackupOperation struct {
	response map[string]interface{}
}

func (op *createBackupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBackupOperation) do(ctx context.Context, r *Backup, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := backupCreateURL(c.Config.BasePath, project, location, name)

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
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://file.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetBackup(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBackupRaw(ctx context.Context, r *Backup) ([]byte, error) {

	u, err := backupGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) backupDiffsForRawDesired(ctx context.Context, rawDesired *Backup, opts ...dcl.ApplyOption) (initial, desired *Backup, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Backup
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Backup); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Backup, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBackup(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Backup resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Backup resource: %v", err)
		}
		c.Config.Logger.Info("Found that Backup resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBackupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Backup: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Backup: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBackupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Backup: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBackupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Backup: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBackup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBackupInitialState(rawInitial, rawDesired *Backup) (*Backup, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBackupDesiredState(rawDesired, rawInitial *Backup, opts ...dcl.ApplyOption) (*Backup, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.NameToSelfLink(rawDesired.SourceInstance, rawInitial.SourceInstance) {
		rawDesired.SourceInstance = rawInitial.SourceInstance
	}
	if dcl.StringCanonicalize(rawDesired.SourceFileShare, rawInitial.SourceFileShare) {
		rawDesired.SourceFileShare = rawInitial.SourceFileShare
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeBackupNewState(c *Client, rawNew, rawDesired *Backup) (*Backup, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CapacityGb) && dcl.IsEmptyValueIndirect(rawDesired.CapacityGb) {
		rawNew.CapacityGb = rawDesired.CapacityGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageBytes) && dcl.IsEmptyValueIndirect(rawDesired.StorageBytes) {
		rawNew.StorageBytes = rawDesired.StorageBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceInstance) && dcl.IsEmptyValueIndirect(rawDesired.SourceInstance) {
		rawNew.SourceInstance = rawDesired.SourceInstance
	} else {
		if dcl.NameToSelfLink(rawDesired.SourceInstance, rawNew.SourceInstance) {
			rawNew.SourceInstance = rawDesired.SourceInstance
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceFileShare) && dcl.IsEmptyValueIndirect(rawDesired.SourceFileShare) {
		rawNew.SourceFileShare = rawDesired.SourceFileShare
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceFileShare, rawNew.SourceFileShare) {
			rawNew.SourceFileShare = rawDesired.SourceFileShare
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceInstanceTier) && dcl.IsEmptyValueIndirect(rawDesired.SourceInstanceTier) {
		rawNew.SourceInstanceTier = rawDesired.SourceInstanceTier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DownloadBytes) && dcl.IsEmptyValueIndirect(rawDesired.DownloadBytes) {
		rawNew.DownloadBytes = rawDesired.DownloadBytes
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffBackup(c *Client, desired, actual *Backup, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackupUpdateBackupOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackupUpdateBackupOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CapacityGb, actual.CapacityGb, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CapacityGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageBytes, actual.StorageBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StorageBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceInstance, actual.SourceInstance, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceFileShare, actual.SourceFileShare, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceFileShare")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceInstanceTier, actual.SourceInstanceTier, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceInstanceTier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DownloadBytes, actual.DownloadBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DownloadBytes")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Backup) urlNormalized() *Backup {
	normalized := dcl.Copy(*r).(Backup)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceInstance = dcl.SelfLinkToName(r.SourceInstance)
	normalized.SourceFileShare = dcl.SelfLinkToName(r.SourceFileShare)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Backup) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Backup) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Backup) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Backup) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateBackup" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/backups/{{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Backup resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Backup) marshal(c *Client) ([]byte, error) {
	m, err := expandBackup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Backup: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBackup decodes JSON responses into the Backup resource schema.
func unmarshalBackup(b []byte, c *Client) (*Backup, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBackup(m, c)
}

func unmarshalMapBackup(m map[string]interface{}, c *Client) (*Backup, error) {

	return flattenBackup(c, m), nil
}

// expandBackup expands Backup into a JSON request object.
func expandBackup(c *Client, f *Backup) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.CapacityGb; !dcl.IsEmptyValueIndirect(v) {
		m["capacityGb"] = v
	}
	if v := f.StorageBytes; !dcl.IsEmptyValueIndirect(v) {
		m["storageBytes"] = v
	}
	if v := f.SourceInstance; !dcl.IsEmptyValueIndirect(v) {
		m["sourceInstance"] = v
	}
	if v := f.SourceFileShare; !dcl.IsEmptyValueIndirect(v) {
		m["sourceFileShare"] = v
	}
	if v := f.SourceInstanceTier; !dcl.IsEmptyValueIndirect(v) {
		m["sourceInstanceTier"] = v
	}
	if v := f.DownloadBytes; !dcl.IsEmptyValueIndirect(v) {
		m["downloadBytes"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenBackup flattens Backup from a JSON request object into the
// Backup type.
func flattenBackup(c *Client, i interface{}) *Backup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Backup{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.State = flattenBackupStateEnum(m["state"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.CapacityGb = dcl.FlattenInteger(m["capacityGb"])
	res.StorageBytes = dcl.FlattenInteger(m["storageBytes"])
	res.SourceInstance = dcl.FlattenString(m["sourceInstance"])
	res.SourceFileShare = dcl.FlattenString(m["sourceFileShare"])
	res.SourceInstanceTier = flattenBackupSourceInstanceTierEnum(m["sourceInstanceTier"])
	res.DownloadBytes = dcl.FlattenInteger(m["downloadBytes"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// flattenBackupStateEnumSlice flattens the contents of BackupStateEnum from a JSON
// response object.
func flattenBackupStateEnumSlice(c *Client, i interface{}) []BackupStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackupStateEnum{}
	}

	if len(a) == 0 {
		return []BackupStateEnum{}
	}

	items := make([]BackupStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackupStateEnum(item.(interface{})))
	}

	return items
}

// flattenBackupStateEnum asserts that an interface is a string, and returns a
// pointer to a *BackupStateEnum with the same value as that string.
func flattenBackupStateEnum(i interface{}) *BackupStateEnum {
	s, ok := i.(string)
	if !ok {
		return BackupStateEnumRef("")
	}

	return BackupStateEnumRef(s)
}

// flattenBackupSourceInstanceTierEnumSlice flattens the contents of BackupSourceInstanceTierEnum from a JSON
// response object.
func flattenBackupSourceInstanceTierEnumSlice(c *Client, i interface{}) []BackupSourceInstanceTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackupSourceInstanceTierEnum{}
	}

	if len(a) == 0 {
		return []BackupSourceInstanceTierEnum{}
	}

	items := make([]BackupSourceInstanceTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackupSourceInstanceTierEnum(item.(interface{})))
	}

	return items
}

// flattenBackupSourceInstanceTierEnum asserts that an interface is a string, and returns a
// pointer to a *BackupSourceInstanceTierEnum with the same value as that string.
func flattenBackupSourceInstanceTierEnum(i interface{}) *BackupSourceInstanceTierEnum {
	s, ok := i.(string)
	if !ok {
		return BackupSourceInstanceTierEnumRef("")
	}

	return BackupSourceInstanceTierEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Backup) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBackup(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

type backupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         backupApiOperation
}

func convertFieldDiffToBackupOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]backupDiff, error) {
	var diffs []backupDiff
	for _, op := range ops {
		diff := backupDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTobackupApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTobackupApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (backupApiOperation, error) {
	switch op {

	case "updateBackupUpdateBackupOperation":
		return &updateBackupUpdateBackupOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
