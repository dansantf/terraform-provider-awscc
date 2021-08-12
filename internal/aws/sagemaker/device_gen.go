// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

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
	registry.AddResourceTypeFactory("aws_sagemaker_device", deviceResourceType)
}

// deviceResourceType returns the Terraform aws_sagemaker_device resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::Device resource type.
func deviceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"device": {
			// Property: Device
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Edge device you want to create",
			//   "properties": {
			//     "Description": {
			//       "description": "Description of the device",
			//       "maxLength": 40,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "DeviceName": {
			//       "description": "The name of the device",
			//       "maxLength": 63,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "IotThingName": {
			//       "description": "AWS Internet of Things (IoT) object name.",
			//       "maxLength": 128,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "DeviceName"
			//   ],
			//   "type": "object"
			// }
			Description: "Edge device you want to create",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"description": {
						// Property: Description
						// CloudFormation resource type schema:
						// {
						//   "description": "Description of the device",
						//   "maxLength": 40,
						//   "minLength": 1,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "Description of the device",
						Type:        types.StringType,
						Optional:    true,
					},
					"device_name": {
						// Property: DeviceName
						// CloudFormation resource type schema:
						// {
						//   "description": "The name of the device",
						//   "maxLength": 63,
						//   "minLength": 1,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "The name of the device",
						Type:        types.StringType,
						Required:    true,
					},
					"iot_thing_name": {
						// Property: IotThingName
						// CloudFormation resource type schema:
						// {
						//   "description": "AWS Internet of Things (IoT) object name.",
						//   "maxLength": 128,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "AWS Internet of Things (IoT) object name.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// Device is a force-new attribute.
		},
		"device_fleet_name": {
			// Property: DeviceFleetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the edge device fleet",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the edge device fleet",
			Type:        types.StringType,
			Required:    true,
			// DeviceFleetName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Associate tags with the resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Associate tags with the resource",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						//   "maxLength": 128,
						//   "minLength": 1,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						// {
						//   "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						//   "maxLength": 256,
						//   "minLength": 0,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
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
		Description: "Resource schema for AWS::SageMaker::Device",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Device").WithTerraformTypeName("aws_sagemaker_device").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_sagemaker_device", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
