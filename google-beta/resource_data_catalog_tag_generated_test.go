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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDataCatalogTag_dataCatalogEntryTagBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"force_delete":  true,
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogTagDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTag_dataCatalogEntryTagBasicExample(context),
			},
			{
				ResourceName:            "google_data_catalog_tag.basic_tag",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataCatalogTag_dataCatalogEntryTagBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_entry" "entry" {
  entry_group = google_data_catalog_entry_group.entry_group.id
  entry_id = "tf_test_my_entry%{random_suffix}"

  user_specified_type = "my_custom_type"
  user_specified_system = "SomethingExternal"
}

resource "google_data_catalog_entry_group" "entry_group" {
  entry_group_id = "tf_test_my_entry_group%{random_suffix}"
}

resource "google_data_catalog_tag_template" "tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag" "basic_tag" {
  parent   = google_data_catalog_entry.entry.id
  template = google_data_catalog_tag_template.tag_template.id

  fields {
    field_name   = "source"
    string_value = "my-string"
  }
}
`, context)
}

func TestAccDataCatalogTag_dataCatalogEntryGroupTagExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"force_delete":  true,
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogTagDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTag_dataCatalogEntryGroupTagExample(context),
			},
			{
				ResourceName:            "google_data_catalog_tag.entry_group_tag",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataCatalogTag_dataCatalogEntryGroupTagExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_entry" "first_entry" {
  entry_group = google_data_catalog_entry_group.entry_group.id
  entry_id = "tf_test_first_entry%{random_suffix}"

  user_specified_type = "my_custom_type"
  user_specified_system = "SomethingExternal"
}

resource "google_data_catalog_entry" "second_entry" {
  entry_group = google_data_catalog_entry_group.entry_group.id
  entry_id = "tf_test_second_entry%{random_suffix}"

  user_specified_type = "another_custom_type"
  user_specified_system = "SomethingElseExternal"
}

resource "google_data_catalog_entry_group" "entry_group" {
  entry_group_id = "tf_test_my_entry_group%{random_suffix}"
}

resource "google_data_catalog_tag_template" "tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag" "entry_group_tag" {
  parent   = google_data_catalog_entry_group.entry_group.id
  template = google_data_catalog_tag_template.tag_template.id

  fields {
    field_name   = "source"
    string_value = "my-string"
  }
}
`, context)
}

func TestAccDataCatalogTag_dataCatalogEntryTagFullExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"force_delete":  true,
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogTagDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTag_dataCatalogEntryTagFullExample(context),
			},
			{
				ResourceName:            "google_data_catalog_tag.basic_tag",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataCatalogTag_dataCatalogEntryTagFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_entry" "entry" {
  entry_group = google_data_catalog_entry_group.entry_group.id
  entry_id = "tf_test_my_entry%{random_suffix}"

  user_specified_type = "my_custom_type"
  user_specified_system = "SomethingExternal"

  schema = <<EOF
{
  "columns": [
    {
      "column": "first_name",
      "description": "First name",
      "mode": "REQUIRED",
      "type": "STRING"
    },
    {
      "column": "last_name",
      "description": "Last name",
      "mode": "REQUIRED",
      "type": "STRING"
    },
    {
      "column": "address",
      "description": "Address",
      "mode": "REPEATED",
      "subcolumns": [
        {
          "column": "city",
          "description": "City",
          "mode": "NULLABLE",
          "type": "STRING"
        },
        {
          "column": "state",
          "description": "State",
          "mode": "NULLABLE",
          "type": "STRING"
        }
      ],
      "type": "RECORD"
    }
  ]
}
EOF
}

resource "google_data_catalog_entry_group" "entry_group" {
  entry_group_id = "tf_test_my_entry_group%{random_suffix}"
}

resource "google_data_catalog_tag_template" "tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag" "basic_tag" {
  parent   = google_data_catalog_entry.entry.id
  template = google_data_catalog_tag_template.tag_template.id

  fields {
    field_name   = "source"
    string_value = "my-string"
  }

  fields {
    field_name   = "num_rows"
    double_value = 5
  }

  fields {
    field_name = "pii_type"
    enum_value = "EMAIL"
  }

  column = "address"
}

resource "google_data_catalog_tag" "second-tag" {
  parent   = google_data_catalog_entry.entry.id
  template = google_data_catalog_tag_template.tag_template.id

  fields {
    field_name   = "source"
    string_value = "my-string"
  }

  fields {
    field_name = "pii_type"
    enum_value = "NONE"
  }

  column = "first_name"
}
`, context)
}

func TestAccDataCatalogTag_dataCatalogEntryTagFalseExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"force_delete":  true,
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogTagDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTag_dataCatalogEntryTagFalseExample(context),
			},
			{
				ResourceName:            "google_data_catalog_tag.basic_tag",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataCatalogTag_dataCatalogEntryTagFalseExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_entry" "entry" {
  entry_group = google_data_catalog_entry_group.entry_group.id
  entry_id = "tf_test_my_entry%{random_suffix}"

  user_specified_type = "my_custom_type"
  user_specified_system = "SomethingExternal"
}

resource "google_data_catalog_entry_group" "entry_group" {
  entry_group_id = "tf_test_my_entry_group%{random_suffix}"
}

resource "google_data_catalog_tag_template" "tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "test boolean value"
    type {
      primitive_type = "BOOL"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag" "basic_tag" {
  parent   = google_data_catalog_entry.entry.id
  template = google_data_catalog_tag_template.tag_template.id

  fields {
    field_name   = "source"
    bool_value = false
  }
}
`, context)
}

func testAccCheckDataCatalogTagDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_catalog_tag" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataCatalogBasePath}}{{parent}}/tags")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("DataCatalogTag still exists at %s", url)
			}
		}

		return nil
	}
}
