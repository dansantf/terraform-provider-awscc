---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_workspacesthinclient_environment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource type definition for AWS::WorkSpacesThinClient::Environment.
---

# awscc_workspacesthinclient_environment (Resource)

Resource type definition for AWS::WorkSpacesThinClient::Environment.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `desktop_arn` (String) The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0.

### Optional

- `desired_software_set_id` (String) The ID of the software set to apply.
- `desktop_endpoint` (String) The URL for the identity provider login (only for environments that use AppStream 2.0).
- `kms_key_arn` (String) The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment.
- `maintenance_window` (Attributes) A specification for a time window to apply software updates. (see [below for nested schema](#nestedatt--maintenance_window))
- `name` (String) The name of the environment.
- `software_set_update_mode` (String) An option to define which software updates to apply.
- `software_set_update_schedule` (String) An option to define if software updates should be applied within a maintenance window.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `activation_code` (String) Activation code for devices associated with environment.
- `arn` (String) The environment ARN.
- `created_at` (String) The timestamp in unix epoch format when environment was created.
- `desktop_type` (String) The type of VDI.
- `id` (String) Unique identifier of the environment.
- `pending_software_set_id` (String) The ID of the software set that is pending to be installed.
- `pending_software_set_version` (String) The version of the software set that is pending to be installed.
- `registered_devices_count` (Number) Number of devices registered to the environment.
- `software_set_compliance_status` (String) Describes if the software currently installed on all devices in the environment is a supported version.
- `updated_at` (String) The timestamp in unix epoch format when environment was last updated.

<a id="nestedatt--maintenance_window"></a>
### Nested Schema for `maintenance_window`

Required:

- `type` (String) The type of maintenance window.

Optional:

- `apply_time_of` (String) The desired time zone maintenance window.
- `days_of_the_week` (Set of String) The date of maintenance window.
- `end_time_hour` (Number) The hour end time of maintenance window.
- `end_time_minute` (Number) The minute end time of maintenance window.
- `start_time_hour` (Number) The hour start time of maintenance window.
- `start_time_minute` (Number) The minute start time of maintenance window.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_workspacesthinclient_environment.example <resource ID>
```