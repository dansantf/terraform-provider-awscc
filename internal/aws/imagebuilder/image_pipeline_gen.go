// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_imagebuilder_image_pipeline", imagePipelineResourceType)
}

// imagePipelineResourceType returns the Terraform aws_imagebuilder_image_pipeline resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::ImagePipeline resource type.
func imagePipelineResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image pipeline.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"container_recipe_arn": {
			// Property: ContainerRecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			Type:        types.StringType,
			Optional:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the image pipeline.",
			//   "type": "string"
			// }
			Description: "The description of the image pipeline.",
			Type:        types.StringType,
			Optional:    true,
		},
		"distribution_configuration_arn": {
			// Property: DistributionConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
			Type:        types.StringType,
			Optional:    true,
		},
		"enhanced_image_metadata_enabled": {
			// Property: EnhancedImageMetadataEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			//   "type": "boolean"
			// }
			Description: "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"image_recipe_arn": {
			// Property: ImageRecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			Type:        types.StringType,
			Optional:    true,
		},
		"image_tests_configuration": {
			// Property: ImageTestsConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Image tests configuration.",
			//   "properties": {
			//     "ImageTestsEnabled": {
			//       "description": "Defines if tests should be executed when building this image.",
			//       "type": "boolean"
			//     },
			//     "TimeoutMinutes": {
			//       "description": "The maximum time in minutes that tests are permitted to run.",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Image tests configuration.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"image_tests_enabled": {
						// Property: ImageTestsEnabled
						// CloudFormation resource type schema:
						// {
						//   "description": "Defines if tests should be executed when building this image.",
						//   "type": "boolean"
						// }
						Description: "Defines if tests should be executed when building this image.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"timeout_minutes": {
						// Property: TimeoutMinutes
						// CloudFormation resource type schema:
						// {
						//   "description": "The maximum time in minutes that tests are permitted to run.",
						//   "type": "integer"
						// }
						Description: "The maximum time in minutes that tests are permitted to run.",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"infrastructure_configuration_arn": {
			// Property: InfrastructureConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
			Type:        types.StringType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image pipeline.",
			//   "type": "string"
			// }
			Description: "The name of the image pipeline.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"schedule": {
			// Property: Schedule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The schedule of the image pipeline.",
			//   "properties": {
			//     "PipelineExecutionStartCondition": {
			//       "description": "The condition configures when the pipeline should trigger a new image build.",
			//       "enum": [
			//         "EXPRESSION_MATCH_ONLY",
			//         "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
			//       ],
			//       "type": "string"
			//     },
			//     "ScheduleExpression": {
			//       "description": "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The schedule of the image pipeline.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"pipeline_execution_start_condition": {
						// Property: PipelineExecutionStartCondition
						// CloudFormation resource type schema:
						// {
						//   "description": "The condition configures when the pipeline should trigger a new image build.",
						//   "enum": [
						//     "EXPRESSION_MATCH_ONLY",
						//     "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
						//   ],
						//   "type": "string"
						// }
						Description: "The condition configures when the pipeline should trigger a new image build.",
						Type:        types.StringType,
						Optional:    true,
					},
					"schedule_expression": {
						// Property: ScheduleExpression
						// CloudFormation resource type schema:
						// {
						//   "description": "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
						//   "type": "string"
						// }
						Description: "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the image pipeline.",
			//   "enum": [
			//     "DISABLED",
			//     "ENABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the image pipeline.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags of this image pipeline.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags of this image pipeline.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::ImageBuilder::ImagePipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ImagePipeline").WithTerraformTypeName("aws_imagebuilder_image_pipeline").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_imagebuilder_image_pipeline", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
