// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_security_group_ingress", securityGroupIngressResource)
}

// securityGroupIngressResource returns the Terraform awscc_ec2_security_group_ingress resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::SecurityGroupIngress resource.
func securityGroupIngressResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CidrIp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv4 ranges",
		//	  "type": "string"
		//	}
		"cidr_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv4 ranges",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CidrIpv6
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[VPC only] The IPv6 ranges",
		//	  "type": "string"
		//	}
		"cidr_ipv_6": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "[VPC only] The IPv6 ranges",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Updates the description of an ingress (inbound) security group rule. You can replace an existing description, or add a description to a rule that did not have one previously",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Updates the description of an ingress (inbound) security group rule. You can replace an existing description, or add a description to a rule that did not have one previously",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FromPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.\n\nUse this for ICMP and any protocol that uses ports.",
		//	  "type": "integer"
		//	}
		"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.\n\nUse this for ICMP and any protocol that uses ports.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.\n\nYou must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.",
		//	  "type": "string"
		//	}
		"group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.\n\nYou must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the security group.",
		//	  "type": "string"
		//	}
		"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the security group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Security Group Rule Id",
		//	  "type": "string"
		//	}
		"security_group_ingress_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Security Group Rule Id",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpProtocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP protocol name (tcp, udp, icmp, icmpv6) or number (see Protocol Numbers).\n\n[VPC only] Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.",
		//	  "type": "string"
		//	}
		"ip_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP protocol name (tcp, udp, icmp, icmpv6) or number (see Protocol Numbers).\n\n[VPC only] Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourcePrefixListId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[EC2-VPC only] The ID of a prefix list.\n\n",
		//	  "type": "string"
		//	}
		"source_prefix_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "[EC2-VPC only] The ID of a prefix list.\n\n",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceSecurityGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the security group. You must specify either the security group ID or the security group name. For security groups in a nondefault VPC, you must specify the security group ID.",
		//	  "type": "string"
		//	}
		"source_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the security group. You must specify either the security group ID or the security group name. For security groups in a nondefault VPC, you must specify the security group ID.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceSecurityGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[EC2-Classic, default VPC] The name of the source security group.\n\nYou must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.",
		//	  "type": "string"
		//	}
		"source_security_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "[EC2-Classic, default VPC] The name of the source security group.\n\nYou must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceSecurityGroupOwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[nondefault VPC] The AWS account ID that owns the source security group. You can't specify this property with an IP address range.\n\nIf you specify SourceSecurityGroupName or SourceSecurityGroupId and that security group is owned by a different account than the account creating the stack, you must specify the SourceSecurityGroupOwnerId; otherwise, this property is optional.",
		//	  "type": "string"
		//	}
		"source_security_group_owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "[nondefault VPC] The AWS account ID that owns the source security group. You can't specify this property with an IP address range.\n\nIf you specify SourceSecurityGroupName or SourceSecurityGroupId and that security group is owned by a different account than the account creating the stack, you must specify the SourceSecurityGroupOwnerId; otherwise, this property is optional.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ToPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes for the specified ICMP type. If you specify all ICMP/ICMPv6 types, you must specify all codes.\n\nUse this for ICMP and any protocol that uses ports.",
		//	  "type": "integer"
		//	}
		"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes for the specified ICMP type. If you specify all ICMP/ICMPv6 types, you must specify all codes.\n\nUse this for ICMP and any protocol that uses ports.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::EC2::SecurityGroupIngress",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::SecurityGroupIngress").WithTerraformTypeName("awscc_ec2_security_group_ingress")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr_ip":                        "CidrIp",
		"cidr_ipv_6":                     "CidrIpv6",
		"description":                    "Description",
		"from_port":                      "FromPort",
		"group_id":                       "GroupId",
		"group_name":                     "GroupName",
		"ip_protocol":                    "IpProtocol",
		"security_group_ingress_id":      "Id",
		"source_prefix_list_id":          "SourcePrefixListId",
		"source_security_group_id":       "SourceSecurityGroupId",
		"source_security_group_name":     "SourceSecurityGroupName",
		"source_security_group_owner_id": "SourceSecurityGroupOwnerId",
		"to_port":                        "ToPort",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
