// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dataplex

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var DataplexAspectTypeIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"aspect_type_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type DataplexAspectTypeIamUpdater struct {
	project      string
	location     string
	aspectTypeId string
	d            tpgresource.TerraformResourceData
	Config       *transport_tpg.Config
}

func DataplexAspectTypeIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("aspect_type_id"); ok {
		values["aspect_type_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/aspectTypes/(?P<aspect_type_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<aspect_type_id>[^/]+)", "(?P<location>[^/]+)/(?P<aspect_type_id>[^/]+)", "(?P<aspect_type_id>[^/]+)"}, d, config, d.Get("aspect_type_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataplexAspectTypeIamUpdater{
		project:      values["project"],
		location:     values["location"],
		aspectTypeId: values["aspect_type_id"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("aspect_type_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting aspect_type_id: %s", err)
	}

	return u, nil
}

func DataplexAspectTypeIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/aspectTypes/(?P<aspect_type_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<aspect_type_id>[^/]+)", "(?P<location>[^/]+)/(?P<aspect_type_id>[^/]+)", "(?P<aspect_type_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataplexAspectTypeIamUpdater{
		project:      values["project"],
		location:     values["location"],
		aspectTypeId: values["aspect_type_id"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("aspect_type_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting aspect_type_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataplexAspectTypeIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyAspectTypeUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataplexAspectTypeIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyAspectTypeUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataplexAspectTypeIamUpdater) qualifyAspectTypeUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataplexBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", u.project, u.location, u.aspectTypeId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataplexAspectTypeIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", u.project, u.location, u.aspectTypeId)
}

func (u *DataplexAspectTypeIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dataplex-aspecttype-%s", u.GetResourceId())
}

func (u *DataplexAspectTypeIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dataplex aspecttype %q", u.GetResourceId())
}
