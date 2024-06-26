---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cleanrooms_privacy_budget_template Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CleanRooms::PrivacyBudgetTemplate
---

# awscc_cleanrooms_privacy_budget_template (Data Source)

Data Source schema for AWS::CleanRooms::PrivacyBudgetTemplate



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `auto_refresh` (String)
- `collaboration_arn` (String)
- `collaboration_identifier` (String)
- `membership_arn` (String)
- `membership_identifier` (String)
- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))
- `privacy_budget_template_identifier` (String)
- `privacy_budget_type` (String)
- `tags` (Attributes List) An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Read-Only:

- `epsilon` (Number)
- `users_noise_per_query` (Number)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
