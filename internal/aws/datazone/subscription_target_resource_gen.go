// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datazone

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_datazone_subscription_target", subscriptionTargetResource)
}

// subscriptionTargetResource returns the Terraform awscc_datazone_subscription_target resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataZone::SubscriptionTarget resource.
func subscriptionTargetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicableAssetTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The asset types that can be included in the subscription target.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "minLength": 1,
		//	    "pattern": "^[^\\.]*",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"applicable_asset_types": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The asset types that can be included in the subscription target.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 256),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^\\.]*"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizedPrincipals
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorized principals of the subscription target.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^[a-zA-Z0-9:/_-]*$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 10,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"authorized_principals": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The authorized principals of the subscription target.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 10),
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9:/_-]*$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the subscription target was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the subscription target was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedBy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon DataZone user who created the subscription target.",
		//	  "type": "string"
		//	}
		"created_by": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon DataZone user who created the subscription target.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone domain in which subscription target is created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone domain in which subscription target is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone domain in which subscription target would be created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone domain in which subscription target would be created.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^dzd[-_][a-zA-Z0-9_-]{1,36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// DomainIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the environment in which subscription target is created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the environment in which subscription target is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the environment in which subscription target would be created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"environment_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the environment in which subscription target would be created.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{1,36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// EnvironmentIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subscription target.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"subscription_target_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subscription target.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManageAccessRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The manage access role that is used to create the subscription target.",
		//	  "type": "string"
		//	}
		"manage_access_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The manage access role that is used to create the subscription target.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subscription target.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subscription target.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the project specified in the subscription target.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the project specified in the subscription target.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Provider
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provider of the subscription target.",
		//	  "type": "string"
		//	}
		"provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provider of the subscription target.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionTargetConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The configuration of the subscription target.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The details of the subscription target configuration.",
		//	    "properties": {
		//	      "Content": {
		//	        "description": "The content of the subscription target configuration.",
		//	        "type": "string"
		//	      },
		//	      "FormName": {
		//	        "description": "The form name included in the subscription target configuration.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Content",
		//	      "FormName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"subscription_target_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Content
					"content": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The content of the subscription target configuration.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: FormName
					"form_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The form name included in the subscription target configuration.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The configuration of the subscription target.",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the subscription target.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the subscription target.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the subscription target was updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the subscription target was updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedBy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon DataZone user who updated the subscription target.",
		//	  "type": "string"
		//	}
		"updated_by": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon DataZone user who updated the subscription target.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Subscription targets enables one to access the data to which you have subscribed in your projects.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::SubscriptionTarget").WithTerraformTypeName("awscc_datazone_subscription_target")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"applicable_asset_types":     "ApplicableAssetTypes",
		"authorized_principals":      "AuthorizedPrincipals",
		"content":                    "Content",
		"created_at":                 "CreatedAt",
		"created_by":                 "CreatedBy",
		"domain_id":                  "DomainId",
		"domain_identifier":          "DomainIdentifier",
		"environment_id":             "EnvironmentId",
		"environment_identifier":     "EnvironmentIdentifier",
		"form_name":                  "FormName",
		"manage_access_role":         "ManageAccessRole",
		"name":                       "Name",
		"project_id":                 "ProjectId",
		"provider_name":              "Provider",
		"subscription_target_config": "SubscriptionTargetConfig",
		"subscription_target_id":     "Id",
		"type":                       "Type",
		"updated_at":                 "UpdatedAt",
		"updated_by":                 "UpdatedBy",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DomainIdentifier",
		"/properties/EnvironmentIdentifier",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
