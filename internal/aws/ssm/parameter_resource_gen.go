// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ssm_parameter", parameterResource)
}

// parameterResource returns the Terraform awscc_ssm_parameter resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSM::Parameter resource.
func parameterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllowedPattern
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The regular expression used to validate the parameter value.",
		//	  "type": "string"
		//	}
		"allowed_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The regular expression used to validate the parameter value.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// AllowedPattern is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: DataType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The corresponding DataType of the parameter.",
		//	  "enum": [
		//	    "text",
		//	    "aws:ec2:image"
		//	  ],
		//	  "type": "string"
		//	}
		"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The corresponding DataType of the parameter.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"text",
					"aws:ec2:image",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The information about the parameter.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The information about the parameter.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Description is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the parameter.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the parameter.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Policies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policies attached to the parameter.",
		//	  "type": "string"
		//	}
		"policies": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The policies attached to the parameter.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Policies is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The corresponding tier of the parameter.",
		//	  "enum": [
		//	    "Standard",
		//	    "Advanced",
		//	    "Intelligent-Tiering"
		//	  ],
		//	  "type": "string"
		//	}
		"tier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The corresponding tier of the parameter.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"Standard",
					"Advanced",
					"Intelligent-Tiering",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Tier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the parameter.",
		//	  "enum": [
		//	    "String",
		//	    "StringList",
		//	    "SecureString"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the parameter.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"String",
					"StringList",
					"SecureString",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Value
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The value associated with the parameter.",
		//	  "type": "string"
		//	}
		"value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The value associated with the parameter.",
			Required:    true,
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
		Description: "Resource Type definition for AWS::SSM::Parameter",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::Parameter").WithTerraformTypeName("awscc_ssm_parameter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_pattern": "AllowedPattern",
		"data_type":       "DataType",
		"description":     "Description",
		"name":            "Name",
		"policies":        "Policies",
		"tags":            "Tags",
		"tier":            "Tier",
		"type":            "Type",
		"value":           "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
		"/properties/Description",
		"/properties/Tier",
		"/properties/AllowedPattern",
		"/properties/Policies",
	})
	opts = opts.WithCreateTimeoutInMinutes(5).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(5)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
