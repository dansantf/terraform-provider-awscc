// Code generated by generators/resource/main.go; DO NOT EDIT.

package codeartifact

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
	registry.AddResourceTypeFactory("aws_codeartifact_repository", repositoryResourceType)
}

// repositoryResourceType returns the Terraform aws_codeartifact_repository resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeArtifact::Repository resource type.
func repositoryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the repository.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The ARN of the repository.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A text description of the repository.",
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Description: "A text description of the repository.",
			Type:        types.StringType,
			Optional:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the domain that contains the repository.",
			//   "maxLength": 50,
			//   "minLength": 2,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the domain that contains the repository.",
			Type:        types.StringType,
			Required:    true,
			// DomainName is a force-new attribute.
		},
		"domain_owner": {
			// Property: DomainOwner
			// CloudFormation resource type schema:
			// {
			//   "description": "The 12-digit account ID of the AWS account that owns the domain.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The 12-digit account ID of the AWS account that owns the domain.",
			Type:        types.StringType,
			Computed:    true,
			// DomainOwner is a force-new attribute.
		},
		"external_connections": {
			// Property: ExternalConnections
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of external connections associated with the repository.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of external connections associated with the repository.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the repository. This is used for GetAtt",
			//   "maxLength": 100,
			//   "minLength": 2,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the repository. This is used for GetAtt",
			Type:        types.StringType,
			Computed:    true,
		},
		"permissions_policy_document": {
			// Property: PermissionsPolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "The access control resource policy on the provided repository.",
			//   "maxLength": 5120,
			//   "minLength": 2,
			//   "type": "object"
			// }
			Description: "The access control resource policy on the provided repository.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"repository_name": {
			// Property: RepositoryName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the repository.",
			//   "maxLength": 100,
			//   "minLength": 2,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the repository.",
			Type:        types.StringType,
			Required:    true,
			// RepositoryName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						//   "maxLength": 128,
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
						//   "maxLength": 256,
						//   "minLength": 0,
						//   "type": "string"
						// }
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"upstreams": {
			// Property: Upstreams
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of upstream repositories associated with the repository.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of upstream repositories associated with the repository.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The resource schema to create a CodeArtifact repository.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeArtifact::Repository").WithTerraformTypeName("aws_codeartifact_repository").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_codeartifact_repository", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
