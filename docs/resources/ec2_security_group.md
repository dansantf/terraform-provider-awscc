---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_security_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::SecurityGroup
---

# awscc_ec2_security_group (Resource)

Resource Type definition for AWS::EC2::SecurityGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_description` (String) A description for the security group.

### Optional

- `group_name` (String) The name of the security group.
- `security_group_egress` (Attributes List) [VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group. (see [below for nested schema](#nestedatt--security_group_egress))
- `security_group_ingress` (Attributes List) The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group. (see [below for nested schema](#nestedatt--security_group_ingress))
- `tags` (Attributes List) Any tags assigned to the security group. (see [below for nested schema](#nestedatt--tags))
- `vpc_id` (String) The ID of the VPC for the security group.

### Read-Only

- `group_id` (String) The group ID of the specified security group.
- `id` (String) Uniquely identifies the resource.
- `security_group_id` (String) The group name or group ID depending on whether the SG is created in default or specific VPC

<a id="nestedatt--security_group_egress"></a>
### Nested Schema for `security_group_egress`

Required:

- `ip_protocol` (String)

Optional:

- `cidr_ip` (String)
- `cidr_ipv_6` (String)
- `description` (String)
- `destination_prefix_list_id` (String)
- `destination_security_group_id` (String)
- `from_port` (Number)
- `to_port` (Number)


<a id="nestedatt--security_group_ingress"></a>
### Nested Schema for `security_group_ingress`

Required:

- `ip_protocol` (String)

Optional:

- `cidr_ip` (String)
- `cidr_ipv_6` (String)
- `description` (String)
- `from_port` (Number)
- `source_prefix_list_id` (String)
- `source_security_group_id` (String)
- `source_security_group_name` (String)
- `source_security_group_owner_id` (String)
- `to_port` (Number)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_security_group.example <resource ID>
```
