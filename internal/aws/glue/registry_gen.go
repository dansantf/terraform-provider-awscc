// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue

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
	registry.AddResourceTypeFactory("aws_glue_registry", registryResourceType)
}

// registryResourceType returns the Terraform aws_glue_registry resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Glue::Registry resource type.
func registryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name for the created Registry.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name for the created Registry.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the registry. If description is not provided, there will not be any default value for this.",
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the registry. If description is not provided, there will not be any default value for this.",
			Type:        types.StringType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the registry to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.  No whitespace.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the registry to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.  No whitespace.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of tags to tag the Registry",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "A key to identify the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Corresponding tag value for the key.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 10,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "List of tags to tag the Registry",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "description": "A key to identify the tag.",
						//   "maxLength": 128,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "A key to identify the tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						// {
						//   "description": "Corresponding tag value for the key.",
						//   "maxLength": 256,
						//   "minLength": 0,
						//   "type": "string"
						// }
						Description: "Corresponding tag value for the key.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 10,
				},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
			// Tags is a write-only attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "This resource creates a Registry for authoring schemas as part of Glue Schema Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::Registry").WithTerraformTypeName("aws_glue_registry").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_glue_registry", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
