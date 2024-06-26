---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotanalytics_dataset Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IoTAnalytics::Dataset
---

# awscc_iotanalytics_dataset (Data Source)

Data Source schema for AWS::IoTAnalytics::Dataset



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `actions` (Attributes List) (see [below for nested schema](#nestedatt--actions))
- `content_delivery_rules` (Attributes List) (see [below for nested schema](#nestedatt--content_delivery_rules))
- `dataset_id` (String)
- `dataset_name` (String)
- `late_data_rules` (Attributes List) (see [below for nested schema](#nestedatt--late_data_rules))
- `retention_period` (Attributes) (see [below for nested schema](#nestedatt--retention_period))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `triggers` (Attributes List) (see [below for nested schema](#nestedatt--triggers))
- `versioning_configuration` (Attributes) (see [below for nested schema](#nestedatt--versioning_configuration))

<a id="nestedatt--actions"></a>
### Nested Schema for `actions`

Read-Only:

- `action_name` (String)
- `container_action` (Attributes) (see [below for nested schema](#nestedatt--actions--container_action))
- `query_action` (Attributes) (see [below for nested schema](#nestedatt--actions--query_action))

<a id="nestedatt--actions--container_action"></a>
### Nested Schema for `actions.container_action`

Read-Only:

- `execution_role_arn` (String)
- `image` (String)
- `resource_configuration` (Attributes) (see [below for nested schema](#nestedatt--actions--container_action--resource_configuration))
- `variables` (Attributes List) (see [below for nested schema](#nestedatt--actions--container_action--variables))

<a id="nestedatt--actions--container_action--resource_configuration"></a>
### Nested Schema for `actions.container_action.resource_configuration`

Read-Only:

- `compute_type` (String)
- `volume_size_in_gb` (Number)


<a id="nestedatt--actions--container_action--variables"></a>
### Nested Schema for `actions.container_action.variables`

Read-Only:

- `dataset_content_version_value` (Attributes) (see [below for nested schema](#nestedatt--actions--container_action--variables--dataset_content_version_value))
- `double_value` (Number)
- `output_file_uri_value` (Attributes) (see [below for nested schema](#nestedatt--actions--container_action--variables--output_file_uri_value))
- `string_value` (String)
- `variable_name` (String)

<a id="nestedatt--actions--container_action--variables--dataset_content_version_value"></a>
### Nested Schema for `actions.container_action.variables.variable_name`

Read-Only:

- `dataset_name` (String)


<a id="nestedatt--actions--container_action--variables--output_file_uri_value"></a>
### Nested Schema for `actions.container_action.variables.variable_name`

Read-Only:

- `file_name` (String)




<a id="nestedatt--actions--query_action"></a>
### Nested Schema for `actions.query_action`

Read-Only:

- `filters` (Attributes List) (see [below for nested schema](#nestedatt--actions--query_action--filters))
- `sql_query` (String)

<a id="nestedatt--actions--query_action--filters"></a>
### Nested Schema for `actions.query_action.filters`

Read-Only:

- `delta_time` (Attributes) (see [below for nested schema](#nestedatt--actions--query_action--filters--delta_time))

<a id="nestedatt--actions--query_action--filters--delta_time"></a>
### Nested Schema for `actions.query_action.filters.delta_time`

Read-Only:

- `offset_seconds` (Number)
- `time_expression` (String)





<a id="nestedatt--content_delivery_rules"></a>
### Nested Schema for `content_delivery_rules`

Read-Only:

- `destination` (Attributes) (see [below for nested schema](#nestedatt--content_delivery_rules--destination))
- `entry_name` (String)

<a id="nestedatt--content_delivery_rules--destination"></a>
### Nested Schema for `content_delivery_rules.destination`

Read-Only:

- `iot_events_destination_configuration` (Attributes) (see [below for nested schema](#nestedatt--content_delivery_rules--destination--iot_events_destination_configuration))
- `s3_destination_configuration` (Attributes) (see [below for nested schema](#nestedatt--content_delivery_rules--destination--s3_destination_configuration))

<a id="nestedatt--content_delivery_rules--destination--iot_events_destination_configuration"></a>
### Nested Schema for `content_delivery_rules.destination.iot_events_destination_configuration`

Read-Only:

- `input_name` (String)
- `role_arn` (String)


<a id="nestedatt--content_delivery_rules--destination--s3_destination_configuration"></a>
### Nested Schema for `content_delivery_rules.destination.s3_destination_configuration`

Read-Only:

- `bucket` (String)
- `glue_configuration` (Attributes) (see [below for nested schema](#nestedatt--content_delivery_rules--destination--s3_destination_configuration--glue_configuration))
- `key` (String)
- `role_arn` (String)

<a id="nestedatt--content_delivery_rules--destination--s3_destination_configuration--glue_configuration"></a>
### Nested Schema for `content_delivery_rules.destination.s3_destination_configuration.role_arn`

Read-Only:

- `database_name` (String)
- `table_name` (String)





<a id="nestedatt--late_data_rules"></a>
### Nested Schema for `late_data_rules`

Read-Only:

- `rule_configuration` (Attributes) (see [below for nested schema](#nestedatt--late_data_rules--rule_configuration))
- `rule_name` (String)

<a id="nestedatt--late_data_rules--rule_configuration"></a>
### Nested Schema for `late_data_rules.rule_configuration`

Read-Only:

- `delta_time_session_window_configuration` (Attributes) (see [below for nested schema](#nestedatt--late_data_rules--rule_configuration--delta_time_session_window_configuration))

<a id="nestedatt--late_data_rules--rule_configuration--delta_time_session_window_configuration"></a>
### Nested Schema for `late_data_rules.rule_configuration.delta_time_session_window_configuration`

Read-Only:

- `timeout_in_minutes` (Number)




<a id="nestedatt--retention_period"></a>
### Nested Schema for `retention_period`

Read-Only:

- `number_of_days` (Number)
- `unlimited` (Boolean)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


<a id="nestedatt--triggers"></a>
### Nested Schema for `triggers`

Read-Only:

- `schedule` (Attributes) (see [below for nested schema](#nestedatt--triggers--schedule))
- `triggering_dataset` (Attributes) (see [below for nested schema](#nestedatt--triggers--triggering_dataset))

<a id="nestedatt--triggers--schedule"></a>
### Nested Schema for `triggers.schedule`

Read-Only:

- `schedule_expression` (String)


<a id="nestedatt--triggers--triggering_dataset"></a>
### Nested Schema for `triggers.triggering_dataset`

Read-Only:

- `dataset_name` (String)



<a id="nestedatt--versioning_configuration"></a>
### Nested Schema for `versioning_configuration`

Read-Only:

- `max_versions` (Number)
- `unlimited` (Boolean)
