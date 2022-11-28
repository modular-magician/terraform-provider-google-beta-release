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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccVertexAIFeaturestoreIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/viewer",
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),

		"kms_key_name": BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccVertexAIFeaturestoreIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeaturestoreIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/viewer",
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),

		"kms_key_name": BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccVertexAIFeaturestoreIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeaturestoreIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/viewer",
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),

		"kms_key_name": BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVertexAIFeaturestoreIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAIFeaturestoreIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

resource "google_vertex_ai_featurestore_iam_member" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccVertexAIFeaturestoreIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_vertex_ai_featurestore_iam_policy" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIFeaturestoreIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

data "google_iam_policy" "foo" {
}

resource "google_vertex_ai_featurestore_iam_policy" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIFeaturestoreIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

resource "google_vertex_ai_featurestore_iam_binding" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccVertexAIFeaturestoreIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

resource "google_vertex_ai_featurestore_iam_binding" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}