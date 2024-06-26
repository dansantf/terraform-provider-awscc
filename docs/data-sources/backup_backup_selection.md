---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_backup_backup_selection Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Backup::BackupSelection
---

# awscc_backup_backup_selection (Data Source)

Data Source schema for AWS::Backup::BackupSelection



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `backup_plan_id` (String)
- `backup_selection` (Attributes) (see [below for nested schema](#nestedatt--backup_selection))
- `backup_selection_id` (String)
- `selection_id` (String)

<a id="nestedatt--backup_selection"></a>
### Nested Schema for `backup_selection`

Read-Only:

- `conditions` (Attributes) (see [below for nested schema](#nestedatt--backup_selection--conditions))
- `iam_role_arn` (String)
- `list_of_tags` (Attributes List) (see [below for nested schema](#nestedatt--backup_selection--list_of_tags))
- `not_resources` (List of String)
- `resources` (List of String)
- `selection_name` (String)

<a id="nestedatt--backup_selection--conditions"></a>
### Nested Schema for `backup_selection.conditions`

Read-Only:

- `string_equals` (Attributes List) (see [below for nested schema](#nestedatt--backup_selection--conditions--string_equals))
- `string_like` (Attributes List) (see [below for nested schema](#nestedatt--backup_selection--conditions--string_like))
- `string_not_equals` (Attributes List) (see [below for nested schema](#nestedatt--backup_selection--conditions--string_not_equals))
- `string_not_like` (Attributes List) (see [below for nested schema](#nestedatt--backup_selection--conditions--string_not_like))

<a id="nestedatt--backup_selection--conditions--string_equals"></a>
### Nested Schema for `backup_selection.conditions.string_equals`

Read-Only:

- `condition_key` (String)
- `condition_value` (String)


<a id="nestedatt--backup_selection--conditions--string_like"></a>
### Nested Schema for `backup_selection.conditions.string_like`

Read-Only:

- `condition_key` (String)
- `condition_value` (String)


<a id="nestedatt--backup_selection--conditions--string_not_equals"></a>
### Nested Schema for `backup_selection.conditions.string_not_equals`

Read-Only:

- `condition_key` (String)
- `condition_value` (String)


<a id="nestedatt--backup_selection--conditions--string_not_like"></a>
### Nested Schema for `backup_selection.conditions.string_not_like`

Read-Only:

- `condition_key` (String)
- `condition_value` (String)



<a id="nestedatt--backup_selection--list_of_tags"></a>
### Nested Schema for `backup_selection.list_of_tags`

Read-Only:

- `condition_key` (String)
- `condition_type` (String)
- `condition_value` (String)
