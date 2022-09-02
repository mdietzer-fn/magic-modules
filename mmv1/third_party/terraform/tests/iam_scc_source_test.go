package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecurityCenterSourceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterSourceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_scc_source_iam_binding.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecurityCenterSourceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_scc_source_iam_binding.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecurityCenterSourceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecurityCenterSourceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_scc_source_iam_member.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecurityCenterSourceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterSourceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_scc_source_iam_policy.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecurityCenterSourceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_scc_source_iam_policy.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecurityCenterSourceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "My Source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_member" "foo" {
  source = google_scc_source.custom_source.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecurityCenterSourceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "My Source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_scc_source_iam_policy" "foo" {
  source = google_scc_source.custom_source.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecurityCenterSourceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "My Source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

data "google_iam_policy" "foo" {
}

resource "google_scc_source_iam_policy" "foo" {
  source = google_scc_source.custom_source.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecurityCenterSourceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "My Source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_binding" "foo" {
  source = google_scc_source.custom_source.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecurityCenterSourceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "My Source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_binding" "foo" {
  source = google_scc_source.custom_source.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
