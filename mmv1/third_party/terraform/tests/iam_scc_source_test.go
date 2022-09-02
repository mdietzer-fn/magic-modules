package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSecurityCenterSourceIamBinding(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/securitycenter.sourcesViewer",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterSourceIamBinding_basic(context),
			},
			{
				ResourceName: "google_scc_source_iam_binding.foo",
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					// This has to be a function because sources only use numeric IDs
					id := state.RootModule().Resources["google_scc_source.custom_source"].Primary.Attributes["id"]
					return fmt.Sprintf("%s %s",
						id,
						context["role"],
					), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecurityCenterSourceIamBinding_update(context),
			},
			{
				ResourceName: "google_scc_source_iam_binding.foo",
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					// This has to be a function because sources only use numeric IDs
					id := state.RootModule().Resources["google_scc_source.custom_source"].Primary.Attributes["id"]
					return fmt.Sprintf("%s %s",
						id,
						context["role"],
					), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// <% unless version == 'ga' -%>
func TestAccSecurityCenterSourceIamBinding_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/securitycenter.sourcesViewer",
		"org_id":        getTestOrgFromEnv(t),
		"condition":     "expires_after_2019_12_31",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterSourceIamBinding_basic_withCondition(context),
			},
			{
				ResourceName: "google_scc_source_iam_binding.foo",
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					// This has to be a function because sources only use numeric IDs
					id := state.RootModule().Resources["google_scc_source.custom_source"].Primary.Attributes["id"]
					return fmt.Sprintf("%s %s %s",
						id,
						context["role"],
						context["condition"],
					), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecurityCenterSourceIamBinding_update_withCondition(context),
			},
			{
				ResourceName: "google_scc_source_iam_binding.foo",
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					// This has to be a function because sources only use numeric IDs
					id := state.RootModule().Resources["google_scc_source.custom_source"].Primary.Attributes["id"]
					return fmt.Sprintf("%s %s %s",
						id,
						context["role"],
						context["condition"],
					), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// <% end -%>

func TestAccSecurityCenterSourceIamMember(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/securitycenter.sourcesViewer",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecurityCenterSourceIamMember_basic(context),
			},
			{
				ResourceName: "google_scc_source_iam_member.foo",
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					// This has to be a function because sources only use numeric IDs
					id := state.RootModule().Resources["google_scc_source.custom_source"].Primary.Attributes["id"]
					return fmt.Sprintf("%s %s user:admin@hashicorptest.com",
						id,
						context["role"],
					), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// <% unless version == 'ga' -%>
func TestAccSecurityCenterSourceIamMember_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/securitycenter.sourcesViewer",
		"org_id":        getTestOrgFromEnv(t),
		"condition":     "expires_after_2019_12_31",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecurityCenterSourceIamMember_basic_withCondition(context),
			},
			{
				ResourceName: "google_scc_source_iam_member.foo",
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					// This has to be a function because sources only use numeric IDs
					id := state.RootModule().Resources["google_scc_source.custom_source"].Primary.Attributes["id"]
					return fmt.Sprintf("%s %s user:admin@hashicorptest.com %s",
						id,
						context["role"],
						context["condition"],
					), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// <% end -%>

func TestAccSecurityCenterSourceIamPolicy(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/securitycenter.sourcesViewer",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterSourceIamPolicy_basic(context),
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

// <% unless version == 'ga' -%>
func TestAccSecurityCenterSourceIamPolicy_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/securitycenter.sourcesViewer",
		"org_id":        getTestOrgFromEnv(t),
		"condition":     "expires_after_2019_12_31",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterSourceIamPolicy_basic_withCondition(context),
			},
			{
				ResourceName:      "google_scc_source_iam_policy.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecurityCenterSourceIamPolicy_emptyBinding_withCondition(context),
			},
			{
				ResourceName:      "google_scc_source_iam_policy.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// <% end -%>

func testAccSecurityCenterSourceIamMember_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_member" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  role         = "%{role}"
  member       = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecurityCenterSourceIamPolicy_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

data "google_iam_policy" "foo" {
  binding {
    role    = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_scc_source_iam_policy" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  policy_data  = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecurityCenterSourceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

data "google_iam_policy" "foo" {
}

resource "google_scc_source_iam_policy" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  policy_data  = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecurityCenterSourceIamBinding_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_binding" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  role         = "%{role}"
  members      = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecurityCenterSourceIamBinding_update(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_binding" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  role         = "%{role}"
  members      = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

// <% unless version == 'ga' -%>
func testAccSecurityCenterSourceIamMember_basic_withCondition(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_member" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  role         = "%{role}"
  member       = "user:admin@hashicorptest.com"

  condition {
    title       = "%{condition}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, context)
}

func testAccSecurityCenterSourceIamPolicy_basic_withCondition(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

data "google_iam_policy" "foo" {
  binding {
    role    = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_scc_source_iam_policy" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  policy_data  = data.google_iam_policy.foo.policy_data

  condition {
    title       = "%{condition}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, context)
}

func testAccSecurityCenterSourceIamPolicy_emptyBinding_withCondition(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

data "google_iam_policy" "foo" {
}

resource "google_scc_source_iam_policy" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  policy_data  = data.google_iam_policy.foo.policy_data

  condition {
    title       = "%{condition}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, context)
}

func testAccSecurityCenterSourceIamBinding_basic_withCondition(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_binding" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  role         = "%{role}"
  members      = ["user:admin@hashicorptest.com"]

  condition {
    title       = "%{condition}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, context)
}

func testAccSecurityCenterSourceIamBinding_update_withCondition(context map[string]interface{}) string {
	return Nprintf(`
resource "google_scc_source" "custom_source" {
  display_name = "tf-test-source%{random_suffix}"
  organization = "%{org_id}"
  description  = "My custom Cloud Security Command Center Finding Source"
}

resource "google_scc_source_iam_binding" "foo" {
  source       = google_scc_source.custom_source.id
  organization = "%{org_id}"
  role         = "%{role}"
  members      = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]

  condition {
    title       = "%{condition}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}
`, context)
}

// <% end -%>
