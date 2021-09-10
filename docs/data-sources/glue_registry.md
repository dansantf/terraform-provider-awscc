---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_glue_registry Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Glue::Registry
---

# awscc_glue_registry (Data Source)

Data Source schema for AWS::Glue::Registry



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) Amazon Resource Name for the created Registry.
- **description** (String) A description of the registry. If description is not provided, there will not be any default value for this.
- **name** (String) Name of the registry to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.  No whitespace.
- **tags** (Attributes List) List of tags to tag the Registry (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) A key to identify the tag.
- **value** (String) Corresponding tag value for the key.

