// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_ipam", iPAMResource)
}

// iPAMResource returns the Terraform awscc_ec2_ipam resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::IPAM resource.
func iPAMResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultResourceDiscoveryAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the default association to the default resource discovery, created with this IPAM.",
		//	  "type": "string"
		//	}
		"default_resource_discovery_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the default association to the default resource discovery, created with this IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultResourceDiscoveryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the default resource discovery, created with this IPAM.",
		//	  "type": "string"
		//	}
		"default_resource_discovery_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the default resource discovery, created with this IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the IPAM.",
		//	  "type": "string"
		//	}
		"ipam_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OperatingRegions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
		//	    "properties": {
		//	      "RegionName": {
		//	        "description": "The name of the region.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "RegionName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"operating_regions": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RegionName
					"region_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the region.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrivateDefaultScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the default scope for publicly routable IP space, created with this IPAM.",
		//	  "type": "string"
		//	}
		"private_default_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the default scope for publicly routable IP space, created with this IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PublicDefaultScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the default scope for publicly routable IP space, created with this IPAM.",
		//	  "maxLength": 255,
		//	  "type": "string"
		//	}
		"public_default_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the default scope for publicly routable IP space, created with this IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceDiscoveryAssociationCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The count of resource discoveries associated with this IPAM.",
		//	  "type": "integer"
		//	}
		"resource_discovery_association_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The count of resource discoveries associated with this IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ScopeCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of scopes that currently exist in this IPAM.",
		//	  "type": "integer"
		//	}
		"scope_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of scopes that currently exist in this IPAM.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Schema of AWS::EC2::IPAM Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAM").WithTerraformTypeName("awscc_ec2_ipam")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn": "Arn",
		"default_resource_discovery_association_id": "DefaultResourceDiscoveryAssociationId",
		"default_resource_discovery_id":             "DefaultResourceDiscoveryId",
		"description":                               "Description",
		"ipam_id":                                   "IpamId",
		"key":                                       "Key",
		"operating_regions":                         "OperatingRegions",
		"private_default_scope_id":                  "PrivateDefaultScopeId",
		"public_default_scope_id":                   "PublicDefaultScopeId",
		"region_name":                               "RegionName",
		"resource_discovery_association_count":      "ResourceDiscoveryAssociationCount",
		"scope_count":                               "ScopeCount",
		"tags":                                      "Tags",
		"value":                                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
