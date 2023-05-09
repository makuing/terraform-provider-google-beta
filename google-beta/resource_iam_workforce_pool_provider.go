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

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

const workforcePoolProviderIdRegexp = `^[a-z0-9-]{4,32}$`

func ValidateWorkforcePoolProviderId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "gcp-") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"gcp-\". "+
				"The prefix `gcp-` is reserved for use by Google, and may not be specified.", k, value))
	}

	if !regexp.MustCompile(workforcePoolProviderIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be 4-32 characters, and may contain the characters [a-z0-9-].", k, value))
	}

	return
}

func ResourceIAMWorkforcePoolWorkforcePoolProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMWorkforcePoolWorkforcePoolProviderCreate,
		Read:   resourceIAMWorkforcePoolWorkforcePoolProviderRead,
		Update: resourceIAMWorkforcePoolWorkforcePoolProviderUpdate,
		Delete: resourceIAMWorkforcePoolWorkforcePoolProviderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAMWorkforcePoolWorkforcePoolProviderImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource.`,
			},
			"provider_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateWorkforcePoolProviderId,
				Description: `The ID for the provider, which becomes the final component of the resource name.
This value must be 4-32 characters, and may contain the characters [a-z0-9-].
The prefix 'gcp-' is reserved for use by Google, and may not be specified.`,
			},
			"workforce_pool_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the pool, which becomes the final component of the resource name.
The IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.
It must start with a letter, and cannot have a trailing hyphen.
The prefix 'gcp-' is reserved for use by Google, and may not be specified.`,
			},
			"attribute_condition": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A [Common Expression Language](https://opensource.google/projects/cel) expression, in
plain text, to restrict what otherwise valid authentication credentials issued by the
provider should not be accepted.

The expression must output a boolean representing whether to allow the federation.

The following keywords may be referenced in the expressions:
  * 'assertion': JSON representing the authentication credential issued by the provider.
  * 'google': The Google attributes mapped from the assertion in the 'attribute_mappings'.
    'google.profile_photo' and 'google.display_name' are not supported.
  * 'attribute': The custom attributes mapped from the assertion in the 'attribute_mappings'.

The maximum length of the attribute condition expression is 4096 characters.
If unspecified, all valid authentication credentials will be accepted.

The following example shows how to only allow credentials with a mapped 'google.groups' value of 'admins':
'''
"'admins' in google.groups"
'''`,
			},
			"attribute_mapping": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Maps attributes from the authentication credentials issued by an external identity provider
to Google Cloud attributes, such as 'subject' and 'segment'.

Each key must be a string specifying the Google Cloud IAM attribute to map to.

The following keys are supported:
  * 'google.subject': The principal IAM is authenticating. You can reference this value in IAM bindings.
    This is also the subject that appears in Cloud Logging logs. This is a required field and
    the mapped subject cannot exceed 127 bytes.
  * 'google.groups': Groups the authenticating user belongs to. You can grant groups access to
    resources using an IAM 'principalSet' binding; access applies to all members of the group.
  * 'google.display_name': The name of the authenticated user. This is an optional field and
    the mapped display name cannot exceed 100 bytes. If not set, 'google.subject' will be displayed instead.
    This attribute cannot be referenced in IAM bindings.
  * 'google.profile_photo': The URL that specifies the authenticated user's thumbnail photo.
    This is an optional field. When set, the image will be visible as the user's profile picture.
    If not set, a generic user icon will be displayed instead.
    This attribute cannot be referenced in IAM bindings.

You can also provide custom attributes by specifying 'attribute.{custom_attribute}', where {custom_attribute}
is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.
The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].

You can reference these attributes in IAM policies to define fine-grained access for a workforce pool
to Google Cloud resources. For example:
  * 'google.subject':
    'principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}'
  * 'google.groups':
    'principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}'
  * 'attribute.{custom_attribute}':
    'principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}'

Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
function that maps an identity provider credential to the normalized attribute specified
by the corresponding map key.

You can use the 'assertion' keyword in the expression to access a JSON representation of
the authentication credential issued by the provider.

The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
the total size of all mapped attributes must not exceed 8KB.

For OIDC providers, you must supply a custom mapping that includes the 'google.subject' attribute.
For example, the following maps the sub claim of the incoming credential to the 'subject' attribute
on a Google token:
'''
{"google.subject": "assertion.sub"}
'''

An object containing a list of '"key": value' pairs.
Example: '{ "name": "wrench", "mass": "1.3kg", "count": "3" }'.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A user-specified description of the provider. Cannot exceed 256 characters.`,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
However, existing tokens still grant access.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A user-specified display name for the provider. Cannot exceed 32 characters.`,
			},
			"oidc": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Represents an OpenId Connect 1.0 identity provider.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The client ID. Must match the audience claim of the JWT issued by the identity provider.`,
						},
						"issuer_uri": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The OIDC issuer URI. Must be a valid URI using the 'https' scheme.`,
						},
						"web_sso_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: `Configuration for web single sign-on for the OIDC provider. Here, web sign-in refers to console sign-in and gcloud sign-in through the browser.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"assertion_claims_behavior": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"ONLY_ID_TOKEN_CLAIMS"}),
										Description: `The behavior for how OIDC Claims are included in the 'assertion' object used for attribute mapping and attribute condition.
* ONLY_ID_TOKEN_CLAIMS: Only include ID Token Claims. Possible values: ["ONLY_ID_TOKEN_CLAIMS"]`,
									},
									"response_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"ID_TOKEN"}),
										Description: `The Response Type to request for in the OIDC Authorization Request for web sign-in.
* ID_TOKEN: The 'response_type=id_token' selection uses the Implicit Flow for web sign-in. Possible values: ["ID_TOKEN"]`,
									},
								},
							},
						},
					},
				},
				ExactlyOneOf: []string{"saml", "oidc"},
			},
			"saml": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Represents a SAML identity provider.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp_metadata_xml": {
							Type:     schema.TypeString,
							Required: true,
							Description: `SAML Identity provider configuration metadata xml doc.
The xml document should comply with [SAML 2.0 specification](https://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf).
The max size of the acceptable xml document will be bounded to 128k characters.

The metadata xml document should satisfy the following constraints:
1) Must contain an Identity Provider Entity ID.
2) Must contain at least one non-expired signing key certificate.
3) For each signing key:
  a) Valid from should be no more than 7 days from now.
  b) Valid to should be no more than 10 years in the future.
4) Up to 3 IdP signing keys are allowed in the metadata xml.

When updating the provider's metadata xml, at least one non-expired signing key
must overlap with the existing metadata. This requirement is skipped if there are
no non-expired signing keys present in the existing metadata.`,
						},
					},
				},
				ExactlyOneOf: []string{"saml", "oidc"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The resource name of the provider.
Format: 'locations/{location}/workforcePools/{workforcePoolId}/providers/{providerId}'`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the provider.
* STATE_UNSPECIFIED: State unspecified.
* ACTIVE: The provider is active and may be used to validate authentication credentials.
* DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
  deleted after approximately 30 days. You can restore a soft-deleted provider using
  [providers.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools.providers/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePoolProvider).`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIAMWorkforcePoolWorkforcePoolProviderCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	attributeMappingProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(d.Get("attribute_mapping"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attribute_mapping"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributeMappingProp)) && (ok || !reflect.DeepEqual(v, attributeMappingProp)) {
		obj["attributeMapping"] = attributeMappingProp
	}
	attributeConditionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(d.Get("attribute_condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attribute_condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributeConditionProp)) && (ok || !reflect.DeepEqual(v, attributeConditionProp)) {
		obj["attributeCondition"] = attributeConditionProp
	}
	samlProp, err := expandIAMWorkforcePoolWorkforcePoolProviderSaml(d.Get("saml"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("saml"); !tpgresource.IsEmptyValue(reflect.ValueOf(samlProp)) && (ok || !reflect.DeepEqual(v, samlProp)) {
		obj["saml"] = samlProp
	}
	oidcProp, err := expandIAMWorkforcePoolWorkforcePoolProviderOidc(d.Get("oidc"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("oidc"); !tpgresource.IsEmptyValue(reflect.ValueOf(oidcProp)) && (ok || !reflect.DeepEqual(v, oidcProp)) {
		obj["oidc"] = oidcProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers?workforcePoolProviderId={{provider_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WorkforcePoolProvider: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating WorkforcePoolProvider: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = IAMWorkforcePoolOperationWaitTime(
		config, res, "Creating WorkforcePoolProvider", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create WorkforcePoolProvider: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkforcePoolProvider %q: %#v", d.Id(), res)

	return resourceIAMWorkforcePoolWorkforcePoolProviderRead(d, meta)
}

func resourceIAMWorkforcePoolWorkforcePoolProviderRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IAMWorkforcePoolWorkforcePoolProvider %q", d.Id()))
	}

	res, err = resourceIAMWorkforcePoolWorkforcePoolProviderDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing IAMWorkforcePoolWorkforcePoolProvider because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenIAMWorkforcePoolWorkforcePoolProviderName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("display_name", flattenIAMWorkforcePoolWorkforcePoolProviderDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("description", flattenIAMWorkforcePoolWorkforcePoolProviderDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("state", flattenIAMWorkforcePoolWorkforcePoolProviderState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("disabled", flattenIAMWorkforcePoolWorkforcePoolProviderDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("attribute_mapping", flattenIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(res["attributeMapping"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("attribute_condition", flattenIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(res["attributeCondition"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("saml", flattenIAMWorkforcePoolWorkforcePoolProviderSaml(res["saml"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}
	if err := d.Set("oidc", flattenIAMWorkforcePoolWorkforcePoolProviderOidc(res["oidc"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePoolProvider: %s", err)
	}

	return nil
}

func resourceIAMWorkforcePoolWorkforcePoolProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	attributeMappingProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(d.Get("attribute_mapping"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attribute_mapping"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attributeMappingProp)) {
		obj["attributeMapping"] = attributeMappingProp
	}
	attributeConditionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(d.Get("attribute_condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attribute_condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attributeConditionProp)) {
		obj["attributeCondition"] = attributeConditionProp
	}
	samlProp, err := expandIAMWorkforcePoolWorkforcePoolProviderSaml(d.Get("saml"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("saml"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, samlProp)) {
		obj["saml"] = samlProp
	}
	oidcProp, err := expandIAMWorkforcePoolWorkforcePoolProviderOidc(d.Get("oidc"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("oidc"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, oidcProp)) {
		obj["oidc"] = oidcProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WorkforcePoolProvider %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}

	if d.HasChange("attribute_mapping") {
		updateMask = append(updateMask, "attributeMapping")
	}

	if d.HasChange("attribute_condition") {
		updateMask = append(updateMask, "attributeCondition")
	}

	if d.HasChange("saml") {
		updateMask = append(updateMask, "saml")
	}

	if d.HasChange("oidc") {
		updateMask = append(updateMask, "oidc")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating WorkforcePoolProvider %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating WorkforcePoolProvider %q: %#v", d.Id(), res)
	}

	err = IAMWorkforcePoolOperationWaitTime(
		config, res, "Updating WorkforcePoolProvider", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceIAMWorkforcePoolWorkforcePoolProviderRead(d, meta)
}

func resourceIAMWorkforcePoolWorkforcePoolProviderDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting WorkforcePoolProvider %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "WorkforcePoolProvider")
	}

	err = IAMWorkforcePoolOperationWaitTime(
		config, res, "Deleting WorkforcePoolProvider", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting WorkforcePoolProvider %q: %#v", d.Id(), res)
	return nil
}

func resourceIAMWorkforcePoolWorkforcePoolProviderImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"locations/(?P<location>[^/]+)/workforcePools/(?P<workforce_pool_id>[^/]+)/providers/(?P<provider_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<workforce_pool_id>[^/]+)/(?P<provider_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAMWorkforcePoolWorkforcePoolProviderName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderSaml(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["idp_metadata_xml"] =
		flattenIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(original["idpMetadataXml"], d, config)
	return []interface{}{transformed}
}
func flattenIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderOidc(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["issuer_uri"] =
		flattenIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(original["issuerUri"], d, config)
	transformed["client_id"] =
		flattenIAMWorkforcePoolWorkforcePoolProviderOidcClientId(original["clientId"], d, config)
	transformed["web_sso_config"] =
		flattenIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfig(original["webSsoConfig"], d, config)
	return []interface{}{transformed}
}
func flattenIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderOidcClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["response_type"] =
		flattenIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigResponseType(original["responseType"], d, config)
	transformed["assertion_claims_behavior"] =
		flattenIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehavior(original["assertionClaimsBehavior"], d, config)
	return []interface{}{transformed}
}
func flattenIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigResponseType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehavior(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderSaml(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdpMetadataXml, err := expandIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(original["idp_metadata_xml"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpMetadataXml); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idpMetadataXml"] = transformedIdpMetadataXml
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIssuerUri, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(original["issuer_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIssuerUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["issuerUri"] = transformedIssuerUri
	}

	transformedClientId, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientId(original["client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientId"] = transformedClientId
	}

	transformedWebSsoConfig, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfig(original["web_sso_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebSsoConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["webSsoConfig"] = transformedWebSsoConfig
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResponseType, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigResponseType(original["response_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponseType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["responseType"] = transformedResponseType
	}

	transformedAssertionClaimsBehavior, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehavior(original["assertion_claims_behavior"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAssertionClaimsBehavior); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["assertionClaimsBehavior"] = transformedAssertionClaimsBehavior
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigResponseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehavior(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceIAMWorkforcePoolWorkforcePoolProviderDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DELETED" {
		return nil, nil
	}

	return res, nil
}
