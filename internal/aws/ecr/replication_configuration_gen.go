// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

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
	registry.AddResourceTypeFactory("aws_ecr_replication_configuration", replicationConfigurationResourceType)
}

// replicationConfigurationResourceType returns the Terraform aws_ecr_replication_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECR::ReplicationConfiguration resource type.
func replicationConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"registry_id": {
			// Property: RegistryId
			// CloudFormation resource type schema:
			// {
			//   "description": "The RegistryId associated with the aws account.",
			//   "type": "string"
			// }
			Description: "The RegistryId associated with the aws account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"replication_configuration": {
			// Property: ReplicationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object representing the replication configuration for a registry.",
			//   "properties": {
			//     "Rules": {
			//       "description": "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain only one replication rule but the rule may contain one or more replication destinations.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "An array of objects representing the details of a replication destination.",
			//         "properties": {
			//           "Destinations": {
			//             "description": "An array of objects representing the details of a replication destination.",
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "An array of objects representing the details of a replication destination.",
			//               "properties": {
			//                 "Region": {
			//                   "description": "A Region to replicate to.",
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "RegistryId": {
			//                   "description": "The account ID of the destination registry to replicate to.",
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Region",
			//                 "RegistryId"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 25,
			//             "minItems": 1,
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Destinations"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 1,
			//       "minItems": 0,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "Rules"
			//   ],
			//   "type": "object"
			// }
			Description: "An object representing the replication configuration for a registry.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"rules": {
						// Property: Rules
						// CloudFormation resource type schema:
						// {
						//   "description": "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain only one replication rule but the rule may contain one or more replication destinations.",
						//   "items": {
						//     "additionalProperties": false,
						//     "description": "An array of objects representing the details of a replication destination.",
						//     "properties": {
						//       "Destinations": {
						//         "description": "An array of objects representing the details of a replication destination.",
						//         "items": {
						//           "additionalProperties": false,
						//           "description": "An array of objects representing the details of a replication destination.",
						//           "properties": {
						//             "Region": {
						//               "description": "A Region to replicate to.",
						//               "pattern": "",
						//               "type": "string"
						//             },
						//             "RegistryId": {
						//               "description": "The account ID of the destination registry to replicate to.",
						//               "pattern": "",
						//               "type": "string"
						//             }
						//           },
						//           "required": [
						//             "Region",
						//             "RegistryId"
						//           ],
						//           "type": "object"
						//         },
						//         "maxItems": 25,
						//         "minItems": 1,
						//         "type": "array"
						//       }
						//     },
						//     "required": [
						//       "Destinations"
						//     ],
						//     "type": "object"
						//   },
						//   "maxItems": 1,
						//   "minItems": 0,
						//   "type": "array"
						// }
						Description: "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain only one replication rule but the rule may contain one or more replication destinations.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"destinations": {
									// Property: Destinations
									// CloudFormation resource type schema:
									// {
									//   "description": "An array of objects representing the details of a replication destination.",
									//   "items": {
									//     "additionalProperties": false,
									//     "description": "An array of objects representing the details of a replication destination.",
									//     "properties": {
									//       "Region": {
									//         "description": "A Region to replicate to.",
									//         "pattern": "",
									//         "type": "string"
									//       },
									//       "RegistryId": {
									//         "description": "The account ID of the destination registry to replicate to.",
									//         "pattern": "",
									//         "type": "string"
									//       }
									//     },
									//     "required": [
									//       "Region",
									//       "RegistryId"
									//     ],
									//     "type": "object"
									//   },
									//   "maxItems": 25,
									//   "minItems": 1,
									//   "type": "array"
									// }
									Description: "An array of objects representing the details of a replication destination.",
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"region": {
												// Property: Region
												// CloudFormation resource type schema:
												// {
												//   "description": "A Region to replicate to.",
												//   "pattern": "",
												//   "type": "string"
												// }
												Description: "A Region to replicate to.",
												Type:        types.StringType,
												Required:    true,
											},
											"registry_id": {
												// Property: RegistryId
												// CloudFormation resource type schema:
												// {
												//   "description": "The account ID of the destination registry to replicate to.",
												//   "pattern": "",
												//   "type": "string"
												// }
												Description: "The account ID of the destination registry to replicate to.",
												Type:        types.StringType,
												Required:    true,
											},
										},
										schema.ListNestedAttributesOptions{
											MinItems: 1,
											MaxItems: 25,
										},
									),
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 0,
								MaxItems: 1,
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::ECR::ReplicationConfiguration resource configures the replication destinations for an Amazon Elastic Container Registry (Amazon Private ECR). For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication.html",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::ReplicationConfiguration").WithTerraformTypeName("aws_ecr_replication_configuration").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ecr_replication_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
