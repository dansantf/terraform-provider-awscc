---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_thing Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IoT::Thing
---

# awscc_iot_thing (Resource)

Resource Type definition for AWS::IoT::Thing



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `attribute_payload` (Attributes) (see [below for nested schema](#nestedatt--attribute_payload))
- `thing_name` (String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `thing_id` (String)

<a id="nestedatt--attribute_payload"></a>
### Nested Schema for `attribute_payload`

Optional:

- `attributes` (Map of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iot_thing.example <resource ID>
```
