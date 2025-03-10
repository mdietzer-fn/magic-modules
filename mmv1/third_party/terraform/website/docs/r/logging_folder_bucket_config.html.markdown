---
subcategory: "Cloud (Stackdriver) Logging"
page_title: "Google: google_logging_folder_bucket_config"
description: |-
  Manages a folder-level logging bucket config.
---

# google\_logging\_folder\_bucket\_config

Manages a folder-level logging bucket config. For more information see
[the official logging documentation](https://cloud.google.com/logging/docs/) and
[Storing Logs](https://cloud.google.com/logging/docs/storage).

~> **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your terraform state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

## Example Usage

```hcl
resource "google_folder" "default" {
	display_name = "some-folder-name"
	parent       = "organizations/123456789"
}

resource "google_logging_folder_bucket_config" "basic" {
	folder    = google_folder.default.name
	location  = "global"
	retention_days = 30
	bucket_id = "_Default"
}
```

## Argument Reference

The following arguments are supported:

* `folder` - (Required) The parent resource that contains the logging bucket.

* `location` - (Required) The location of the bucket.

* `bucket_id` - (Required) The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.

* `description` - (Optional) Describes this bucket.

* `retention_days` - (Optional) Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are
exported:

* `id` - an identifier for the resource with format `folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}`

* `name` -  The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"

* `lifecycle_state` -  The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).

## Import

This resource can be imported using the following format:

```
$ terraform import google_logging_folder_bucket_config.default folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}
```
