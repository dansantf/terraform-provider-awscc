// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package applicationautoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_applicationautoscaling_scalable_target", scalableTargetDataSource)
}

// scalableTargetDataSource returns the Terraform awscc_applicationautoscaling_scalable_target data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApplicationAutoScaling::ScalableTarget resource.
func scalableTargetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This value can be returned by using the Ref function. Ref returns the Cloudformation generated ID of the resource in format - ResourceId|ScalableDimension|ServiceNamespace",
		//	  "type": "string"
		//	}
		"scalable_target_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This value can be returned by using the Ref function. Ref returns the Cloudformation generated ID of the resource in format - ResourceId|ScalableDimension|ServiceNamespace",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand",
		//	  "type": "integer"
		//	}
		"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand",
		//	  "type": "integer"
		//	}
		"min_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the resource associated with the scalable target",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the resource associated with the scalable target",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. ",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScalableDimension
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property",
		//	  "type": "string"
		//	}
		"scalable_dimension": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScheduledActions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scheduled actions for the scalable target. Duplicates aren't allowed.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "specifies a scheduled action for a scalable target",
		//	    "properties": {
		//	      "EndTime": {
		//	        "type": "string"
		//	      },
		//	      "ScalableTargetAction": {
		//	        "additionalProperties": false,
		//	        "description": "specifies the minimum and maximum capacity",
		//	        "properties": {
		//	          "MaxCapacity": {
		//	            "type": "integer"
		//	          },
		//	          "MinCapacity": {
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Schedule": {
		//	        "type": "string"
		//	      },
		//	      "ScheduledActionName": {
		//	        "type": "string"
		//	      },
		//	      "StartTime": {
		//	        "type": "string"
		//	      },
		//	      "Timezone": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ScheduledActionName",
		//	      "Schedule"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"scheduled_actions": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EndTime
					"end_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ScalableTargetAction
					"scalable_target_action": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: MaxCapacity
							"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: MinCapacity
							"min_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "specifies the minimum and maximum capacity",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Schedule
					"schedule": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ScheduledActionName
					"scheduled_action_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: StartTime
					"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Timezone
					"timezone": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The scheduled actions for the scalable target. Duplicates aren't allowed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceNamespace
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The namespace of the AWS service that provides the resource, or a custom-resource",
		//	  "type": "string"
		//	}
		"service_namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The namespace of the AWS service that provides the resource, or a custom-resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SuspendedState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.",
		//	  "properties": {
		//	    "DynamicScalingInSuspended": {
		//	      "type": "boolean"
		//	    },
		//	    "DynamicScalingOutSuspended": {
		//	      "type": "boolean"
		//	    },
		//	    "ScheduledScalingSuspended": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"suspended_state": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DynamicScalingInSuspended
				"dynamic_scaling_in_suspended": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DynamicScalingOutSuspended
				"dynamic_scaling_out_suspended": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ScheduledScalingSuspended
				"scheduled_scaling_suspended": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApplicationAutoScaling::ScalableTarget",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApplicationAutoScaling::ScalableTarget").WithTerraformTypeName("awscc_applicationautoscaling_scalable_target")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dynamic_scaling_in_suspended":  "DynamicScalingInSuspended",
		"dynamic_scaling_out_suspended": "DynamicScalingOutSuspended",
		"end_time":                      "EndTime",
		"max_capacity":                  "MaxCapacity",
		"min_capacity":                  "MinCapacity",
		"resource_id":                   "ResourceId",
		"role_arn":                      "RoleARN",
		"scalable_dimension":            "ScalableDimension",
		"scalable_target_action":        "ScalableTargetAction",
		"scalable_target_id":            "Id",
		"schedule":                      "Schedule",
		"scheduled_action_name":         "ScheduledActionName",
		"scheduled_actions":             "ScheduledActions",
		"scheduled_scaling_suspended":   "ScheduledScalingSuspended",
		"service_namespace":             "ServiceNamespace",
		"start_time":                    "StartTime",
		"suspended_state":               "SuspendedState",
		"timezone":                      "Timezone",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
