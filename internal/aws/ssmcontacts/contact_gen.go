// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmcontacts

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
	registry.AddResourceTypeFactory("aws_ssmcontacts_contact", contactResourceType)
}

// contactResourceType returns the Terraform aws_ssmcontacts_contact resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSMContacts::Contact resource type.
func contactResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"alias": {
			// Property: Alias
			// CloudFormation resource type schema:
			// {
			//   "description": "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
			Type:        types.StringType,
			Required:    true,
			// Alias is a force-new attribute.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the contact.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the contact.",
			Type:        types.StringType,
			Computed:    true,
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
			Type:        types.StringType,
			Required:    true,
		},
		"plan": {
			// Property: Plan
			// CloudFormation resource type schema:
			// {
			//   "description": "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A set amount of time that an escalation plan or engagement plan engages the specified contacts or contact methods.",
			//     "properties": {
			//       "DurationInMinutes": {
			//         "description": "The time to wait until beginning the next stage.",
			//         "type": "integer"
			//       },
			//       "Targets": {
			//         "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
			//           "properties": {
			//             "ChannelTargetInfo": {
			//               "additionalProperties": false,
			//               "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
			//               "properties": {
			//                 "ChannelId": {
			//                   "description": "The Amazon Resource Name (ARN) of the contact channel.",
			//                   "type": "string"
			//                 },
			//                 "RetryIntervalInMinutes": {
			//                   "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
			//                   "type": "integer"
			//                 }
			//               },
			//               "required": [
			//                 "ChannelId",
			//                 "RetryIntervalInMinutes"
			//               ],
			//               "type": "object"
			//             },
			//             "ContactTargetInfo": {
			//               "additionalProperties": false,
			//               "description": "The contact that SSM Incident Manager is engaging during an incident.",
			//               "properties": {
			//                 "ContactId": {
			//                   "description": "The Amazon Resource Name (ARN) of the contact.",
			//                   "type": "string"
			//                 },
			//                 "IsEssential": {
			//                   "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
			//                   "type": "boolean"
			//                 }
			//               },
			//               "required": [
			//                 "ContactId",
			//                 "IsEssential"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "DurationInMinutes"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"duration_in_minutes": {
						// Property: DurationInMinutes
						// CloudFormation resource type schema:
						// {
						//   "description": "The time to wait until beginning the next stage.",
						//   "type": "integer"
						// }
						Description: "The time to wait until beginning the next stage.",
						Type:        types.NumberType,
						Required:    true,
					},
					"targets": {
						// Property: Targets
						// CloudFormation resource type schema:
						// {
						//   "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						//   "items": {
						//     "additionalProperties": false,
						//     "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						//     "properties": {
						//       "ChannelTargetInfo": {
						//         "additionalProperties": false,
						//         "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
						//         "properties": {
						//           "ChannelId": {
						//             "description": "The Amazon Resource Name (ARN) of the contact channel.",
						//             "type": "string"
						//           },
						//           "RetryIntervalInMinutes": {
						//             "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
						//             "type": "integer"
						//           }
						//         },
						//         "required": [
						//           "ChannelId",
						//           "RetryIntervalInMinutes"
						//         ],
						//         "type": "object"
						//       },
						//       "ContactTargetInfo": {
						//         "additionalProperties": false,
						//         "description": "The contact that SSM Incident Manager is engaging during an incident.",
						//         "properties": {
						//           "ContactId": {
						//             "description": "The Amazon Resource Name (ARN) of the contact.",
						//             "type": "string"
						//           },
						//           "IsEssential": {
						//             "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
						//             "type": "boolean"
						//           }
						//         },
						//         "required": [
						//           "ContactId",
						//           "IsEssential"
						//         ],
						//         "type": "object"
						//       }
						//     },
						//     "type": "object"
						//   },
						//   "type": "array"
						// }
						Description: "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"channel_target_info": {
									// Property: ChannelTargetInfo
									// CloudFormation resource type schema:
									// {
									//   "additionalProperties": false,
									//   "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
									//   "properties": {
									//     "ChannelId": {
									//       "description": "The Amazon Resource Name (ARN) of the contact channel.",
									//       "type": "string"
									//     },
									//     "RetryIntervalInMinutes": {
									//       "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
									//       "type": "integer"
									//     }
									//   },
									//   "required": [
									//     "ChannelId",
									//     "RetryIntervalInMinutes"
									//   ],
									//   "type": "object"
									// }
									Description: "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"channel_id": {
												// Property: ChannelId
												// CloudFormation resource type schema:
												// {
												//   "description": "The Amazon Resource Name (ARN) of the contact channel.",
												//   "type": "string"
												// }
												Description: "The Amazon Resource Name (ARN) of the contact channel.",
												Type:        types.StringType,
												Required:    true,
											},
											"retry_interval_in_minutes": {
												// Property: RetryIntervalInMinutes
												// CloudFormation resource type schema:
												// {
												//   "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
												//   "type": "integer"
												// }
												Description: "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
												Type:        types.NumberType,
												Required:    true,
											},
										},
									),
									Optional: true,
								},
								"contact_target_info": {
									// Property: ContactTargetInfo
									// CloudFormation resource type schema:
									// {
									//   "additionalProperties": false,
									//   "description": "The contact that SSM Incident Manager is engaging during an incident.",
									//   "properties": {
									//     "ContactId": {
									//       "description": "The Amazon Resource Name (ARN) of the contact.",
									//       "type": "string"
									//     },
									//     "IsEssential": {
									//       "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
									//       "type": "boolean"
									//     }
									//   },
									//   "required": [
									//     "ContactId",
									//     "IsEssential"
									//   ],
									//   "type": "object"
									// }
									Description: "The contact that SSM Incident Manager is engaging during an incident.",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"contact_id": {
												// Property: ContactId
												// CloudFormation resource type schema:
												// {
												//   "description": "The Amazon Resource Name (ARN) of the contact.",
												//   "type": "string"
												// }
												Description: "The Amazon Resource Name (ARN) of the contact.",
												Type:        types.StringType,
												Required:    true,
											},
											"is_essential": {
												// Property: IsEssential
												// CloudFormation resource type schema:
												// {
												//   "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
												//   "type": "boolean"
												// }
												Description: "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
												Type:        types.BoolType,
												Required:    true,
											},
										},
									),
									Optional: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
			// Plan is a write-only attribute.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
			//   "enum": [
			//     "PERSONAL",
			//     "CUSTOM",
			//     "SERVICE",
			//     "ESCALATION"
			//   ],
			//   "type": "string"
			// }
			Description: "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
			Type:        types.StringType,
			Required:    true,
			// Type is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SSMContacts::Contact",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::Contact").WithTerraformTypeName("aws_ssmcontacts_contact").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Plan",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ssmcontacts_contact", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
