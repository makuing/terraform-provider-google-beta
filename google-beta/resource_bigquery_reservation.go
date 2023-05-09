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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceBigqueryReservationReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationReservationCreate,
		Read:   resourceBigqueryReservationReservationRead,
		Update: resourceBigqueryReservationReservationUpdate,
		Delete: resourceBigqueryReservationReservationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationReservationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the reservation. This field must only contain alphanumeric characters or dash.`,
			},
			"slot_capacity": {
				Type:     schema.TypeInt,
				Required: true,
				Description: `Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.`,
			},
			"autoscale": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The configuration parameters for the auto scaling feature.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_slots": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Number of slots to be scaled when needed.`,
						},
						"current_slots": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The slot capacity added to this reservation when autoscale happens. Will be between [0, max_slots].`,
						},
					},
				},
			},
			"concurrency": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.`,
				Default:     0,
			},
			"edition": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS`,
			},
			"ignore_idle_slots": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If false, any query using this reservation will use idle slots from other reservations within
the same admin project. If true, a query using this reservation will execute with the slot
capacity specified above at most.`,
				Default: false,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the transfer config should reside.
Examples: US, EU, asia-northeast1. The default value is US.`,
				Default: "US",
			},
			"multi_region_auxiliary": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBigqueryReservationReservationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	slotCapacityProp, err := expandBigqueryReservationReservationSlotCapacity(d.Get("slot_capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_capacity"); !tpgresource.IsEmptyValue(reflect.ValueOf(slotCapacityProp)) && (ok || !reflect.DeepEqual(v, slotCapacityProp)) {
		obj["slotCapacity"] = slotCapacityProp
	}
	ignoreIdleSlotsProp, err := expandBigqueryReservationReservationIgnoreIdleSlots(d.Get("ignore_idle_slots"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ignore_idle_slots"); !tpgresource.IsEmptyValue(reflect.ValueOf(ignoreIdleSlotsProp)) && (ok || !reflect.DeepEqual(v, ignoreIdleSlotsProp)) {
		obj["ignoreIdleSlots"] = ignoreIdleSlotsProp
	}
	concurrencyProp, err := expandBigqueryReservationReservationConcurrency(d.Get("concurrency"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("concurrency"); !tpgresource.IsEmptyValue(reflect.ValueOf(concurrencyProp)) && (ok || !reflect.DeepEqual(v, concurrencyProp)) {
		obj["concurrency"] = concurrencyProp
	}
	multiRegionAuxiliaryProp, err := expandBigqueryReservationReservationMultiRegionAuxiliary(d.Get("multi_region_auxiliary"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("multi_region_auxiliary"); !tpgresource.IsEmptyValue(reflect.ValueOf(multiRegionAuxiliaryProp)) && (ok || !reflect.DeepEqual(v, multiRegionAuxiliaryProp)) {
		obj["multiRegionAuxiliary"] = multiRegionAuxiliaryProp
	}
	editionProp, err := expandBigqueryReservationReservationEdition(d.Get("edition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("edition"); !tpgresource.IsEmptyValue(reflect.ValueOf(editionProp)) && (ok || !reflect.DeepEqual(v, editionProp)) {
		obj["edition"] = editionProp
	}
	autoscaleProp, err := expandBigqueryReservationReservationAutoscale(d.Get("autoscale"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscale"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscaleProp)) && (ok || !reflect.DeepEqual(v, autoscaleProp)) {
		obj["autoscale"] = autoscaleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations?reservationId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Reservation: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Reservation: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Reservation %q: %#v", d.Id(), res)

	return resourceBigqueryReservationReservationRead(d, meta)
}

func resourceBigqueryReservationReservationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationReservation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	if err := d.Set("slot_capacity", flattenBigqueryReservationReservationSlotCapacity(res["slotCapacity"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("ignore_idle_slots", flattenBigqueryReservationReservationIgnoreIdleSlots(res["ignoreIdleSlots"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("concurrency", flattenBigqueryReservationReservationConcurrency(res["concurrency"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("multi_region_auxiliary", flattenBigqueryReservationReservationMultiRegionAuxiliary(res["multiRegionAuxiliary"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("edition", flattenBigqueryReservationReservationEdition(res["edition"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("autoscale", flattenBigqueryReservationReservationAutoscale(res["autoscale"], d, config)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	return nil
}

func resourceBigqueryReservationReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	slotCapacityProp, err := expandBigqueryReservationReservationSlotCapacity(d.Get("slot_capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_capacity"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, slotCapacityProp)) {
		obj["slotCapacity"] = slotCapacityProp
	}
	ignoreIdleSlotsProp, err := expandBigqueryReservationReservationIgnoreIdleSlots(d.Get("ignore_idle_slots"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ignore_idle_slots"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ignoreIdleSlotsProp)) {
		obj["ignoreIdleSlots"] = ignoreIdleSlotsProp
	}
	concurrencyProp, err := expandBigqueryReservationReservationConcurrency(d.Get("concurrency"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("concurrency"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, concurrencyProp)) {
		obj["concurrency"] = concurrencyProp
	}
	multiRegionAuxiliaryProp, err := expandBigqueryReservationReservationMultiRegionAuxiliary(d.Get("multi_region_auxiliary"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("multi_region_auxiliary"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, multiRegionAuxiliaryProp)) {
		obj["multiRegionAuxiliary"] = multiRegionAuxiliaryProp
	}
	autoscaleProp, err := expandBigqueryReservationReservationAutoscale(d.Get("autoscale"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscale"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autoscaleProp)) {
		obj["autoscale"] = autoscaleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Reservation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("slot_capacity") {
		updateMask = append(updateMask, "slotCapacity")
	}

	if d.HasChange("ignore_idle_slots") {
		updateMask = append(updateMask, "ignoreIdleSlots")
	}

	if d.HasChange("concurrency") {
		updateMask = append(updateMask, "concurrency")
	}

	if d.HasChange("multi_region_auxiliary") {
		updateMask = append(updateMask, "multiRegionAuxiliary")
	}

	if d.HasChange("autoscale") {
		updateMask = append(updateMask, "autoscale")
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
		return fmt.Errorf("Error updating Reservation %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Reservation %q: %#v", d.Id(), res)
	}

	return resourceBigqueryReservationReservationRead(d, meta)
}

func resourceBigqueryReservationReservationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Reservation: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Reservation %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Reservation")
	}

	log.Printf("[DEBUG] Finished deleting Reservation %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryReservationReservationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/reservations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryReservationReservationSlotCapacity(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigqueryReservationReservationIgnoreIdleSlots(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationReservationConcurrency(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigqueryReservationReservationMultiRegionAuxiliary(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationReservationEdition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationReservationAutoscale(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["current_slots"] =
		flattenBigqueryReservationReservationAutoscaleCurrentSlots(original["currentSlots"], d, config)
	transformed["max_slots"] =
		flattenBigqueryReservationReservationAutoscaleMaxSlots(original["maxSlots"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryReservationReservationAutoscaleCurrentSlots(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigqueryReservationReservationAutoscaleMaxSlots(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandBigqueryReservationReservationSlotCapacity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationIgnoreIdleSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationConcurrency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationMultiRegionAuxiliary(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationEdition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAutoscale(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCurrentSlots, err := expandBigqueryReservationReservationAutoscaleCurrentSlots(original["current_slots"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentSlots); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentSlots"] = transformedCurrentSlots
	}

	transformedMaxSlots, err := expandBigqueryReservationReservationAutoscaleMaxSlots(original["max_slots"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxSlots); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxSlots"] = transformedMaxSlots
	}

	return transformed, nil
}

func expandBigqueryReservationReservationAutoscaleCurrentSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAutoscaleMaxSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
