// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_eks_addon", addonResourceType)
}

// addonResourceType returns the Terraform aws_eks_addon resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EKS::Addon resource type.
func addonResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"addon_name": {
			// Property: AddonName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Addon",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of Addon",
			Type:        types.StringType,
			Required:    true,
			// AddonName is a force-new attribute.
		},
		"addon_version": {
			// Property: AddonVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "Version of Addon",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Version of Addon",
			Type:        types.StringType,
			Optional:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the add-on",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the add-on",
			Type:        types.StringType,
			Computed:    true,
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Cluster",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of Cluster",
			Type:        types.StringType,
			Required:    true,
			// ClusterName is a force-new attribute.
		},
		"resolve_conflicts": {
			// Property: ResolveConflicts
			// CloudFormation resource type schema:
			// {
			//   "description": "Resolve parameter value conflicts",
			//   "enum": [
			//     "NONE",
			//     "OVERWRITE"
			//   ],
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Resolve parameter value conflicts",
			Type:        types.StringType,
			Optional:    true,
			// ResolveConflicts is a write-only attribute.
		},
		"service_account_role_arn": {
			// Property: ServiceAccountRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "IAM role to bind to the add-on's service account",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "IAM role to bind to the add-on's service account",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						//   "maxLength": 127,
						//   "minLength": 1,
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
						//   "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						//   "maxLength": 255,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
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
		Description: "Resource Schema for AWS::EKS::Addon",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Addon").WithTerraformTypeName("aws_eks_addon").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ResolveConflicts",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_eks_addon", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
