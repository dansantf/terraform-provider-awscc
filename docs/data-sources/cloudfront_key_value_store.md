---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_key_value_store Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFront::KeyValueStore
---

# awscc_cloudfront_key_value_store (Data Source)

Data Source schema for AWS::CloudFront::KeyValueStore



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `comment` (String)
- `import_source` (Attributes) (see [below for nested schema](#nestedatt--import_source))
- `key_value_store_id` (String)
- `name` (String)
- `status` (String)

<a id="nestedatt--import_source"></a>
### Nested Schema for `import_source`

Read-Only:

- `source_arn` (String)
- `source_type` (String)
