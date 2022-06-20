// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_connect_task_template", taskTemplateResourceType)
}

// taskTemplateResourceType returns the Terraform awscc_connect_task_template resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Connect::TaskTemplate resource type.
func taskTemplateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier (arn) of the task template.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/task-template/[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$",
			//   "type": "string"
			// }
			Description: "The identifier (arn) of the task template.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"client_token": {
			// Property: ClientToken
			// CloudFormation resource type schema:
			// {
			//   "description": "the client token string in uuid format",
			//   "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$",
			//   "type": "string"
			// }
			Description: "the client token string in uuid format",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$"), ""),
			},
		},
		"constraints": {
			// Property: Constraints
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The constraints for the task template",
			//   "properties": {
			//     "InvisibleFields": {
			//       "description": "The list of the task template's invisible fields",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Invisible field info",
			//         "properties": {
			//           "Id": {
			//             "additionalProperties": false,
			//             "description": "the identifier (name) for the task template field",
			//             "properties": {
			//               "Name": {
			//                 "description": "The name of the task template field",
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Name"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Id"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 50,
			//       "type": "array"
			//     },
			//     "ReadOnlyFields": {
			//       "description": "The list of the task template's read only fields",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "ReadOnly field info",
			//         "properties": {
			//           "Id": {
			//             "additionalProperties": false,
			//             "description": "the identifier (name) for the task template field",
			//             "properties": {
			//               "Name": {
			//                 "description": "The name of the task template field",
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Name"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Id"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 50,
			//       "type": "array"
			//     },
			//     "RequiredFields": {
			//       "description": "The list of the task template's required fields",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Required field info",
			//         "properties": {
			//           "Id": {
			//             "additionalProperties": false,
			//             "description": "the identifier (name) for the task template field",
			//             "properties": {
			//               "Name": {
			//                 "description": "The name of the task template field",
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Name"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Id"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 50,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The constraints for the task template",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"invisible_fields": {
						// Property: InvisibleFields
						Description: "The list of the task template's invisible fields",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"id": {
									// Property: Id
									Description: "the identifier (name) for the task template field",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Description: "The name of the task template field",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 100),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
						},
					},
					"read_only_fields": {
						// Property: ReadOnlyFields
						Description: "The list of the task template's read only fields",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"id": {
									// Property: Id
									Description: "the identifier (name) for the task template field",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Description: "The name of the task template field",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 100),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
						},
					},
					"required_fields": {
						// Property: RequiredFields
						Description: "The list of the task template's required fields",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"id": {
									// Property: Id
									Description: "the identifier (name) for the task template field",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Description: "The name of the task template field",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 100),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
						},
					},
				},
			),
			Optional: true,
		},
		"contact_flow_arn": {
			// Property: ContactFlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the contact flow.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the contact flow.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$"), ""),
			},
		},
		"defaults": {
			// Property: Defaults
			// CloudFormation resource type schema:
			// {
			//   "description": "",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "the default value for the task template's field",
			//     "properties": {
			//       "DefaultValue": {
			//         "description": "the default value for the task template's field",
			//         "maxLength": 4096,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Id": {
			//         "additionalProperties": false,
			//         "description": "the identifier (name) for the task template field",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the task template field",
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Name"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "Id",
			//       "DefaultValue"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Description: "",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"default_value": {
						// Property: DefaultValue
						Description: "the default value for the task template's field",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 4096),
						},
					},
					"id": {
						// Property: Id
						Description: "the identifier (name) for the task template field",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the task template field",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the task template.",
			//   "maxLength": 255,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the task template.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 255),
			},
		},
		"fields": {
			// Property: Fields
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of task template's fields",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A task template field object.",
			//     "properties": {
			//       "Description": {
			//         "description": "The description of the task template's field",
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "Id": {
			//         "additionalProperties": false,
			//         "description": "the identifier (name) for the task template field",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the task template field",
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Name"
			//         ],
			//         "type": "object"
			//       },
			//       "SingleSelectOptions": {
			//         "description": "list of field options to be used with single select",
			//         "items": {
			//           "description": "Single select field identifier",
			//           "maxLength": 100,
			//           "minLength": 1,
			//           "pattern": "^[A-Za-z0-9](?:[A-Za-z0-9_.,\\s-]*[A-Za-z0-9_.,-])?$",
			//           "type": "string"
			//         },
			//         "maxItems": 50,
			//         "type": "array"
			//       },
			//       "Type": {
			//         "description": "The type of the task template's field",
			//         "enum": [
			//           "NAME",
			//           "DESCRIPTION",
			//           "SCHEDULED_TIME",
			//           "QUICK_CONNECT",
			//           "URL",
			//           "NUMBER",
			//           "TEXT",
			//           "TEXT_AREA",
			//           "DATE_TIME",
			//           "BOOLEAN",
			//           "SINGLE_SELECT",
			//           "EMAIL"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Id",
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Description: "The list of task template's fields",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Description: "The description of the task template's field",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 255),
						},
					},
					"id": {
						// Property: Id
						Description: "the identifier (name) for the task template field",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the task template field",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
								},
							},
						),
						Required: true,
					},
					"single_select_options": {
						// Property: SingleSelectOptions
						Description: "list of field options to be used with single select",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
							validate.ArrayForEach(validate.StringLenBetween(1, 100)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9](?:[A-Za-z0-9_.,\\s-]*[A-Za-z0-9_.,-])?$"), "")),
						},
					},
					"type": {
						// Property: Type
						Description: "The type of the task template's field",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NAME",
								"DESCRIPTION",
								"SCHEDULED_TIME",
								"QUICK_CONNECT",
								"URL",
								"NUMBER",
								"TEXT",
								"TEXT_AREA",
								"DATE_TIME",
								"BOOLEAN",
								"SINGLE_SELECT",
								"EMAIL",
							}),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier (arn) of the instance.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier (arn) of the instance.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$"), ""),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the task template.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the task template.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 100),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the task template",
			//   "enum": [
			//     "ACTIVE",
			//     "INACTIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the task template",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTIVE",
					"INACTIVE",
				}),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Connect::TaskTemplate.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::TaskTemplate").WithTerraformTypeName("awscc_connect_task_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"client_token":          "ClientToken",
		"constraints":           "Constraints",
		"contact_flow_arn":      "ContactFlowArn",
		"default_value":         "DefaultValue",
		"defaults":              "Defaults",
		"description":           "Description",
		"fields":                "Fields",
		"id":                    "Id",
		"instance_arn":          "InstanceArn",
		"invisible_fields":      "InvisibleFields",
		"key":                   "Key",
		"name":                  "Name",
		"read_only_fields":      "ReadOnlyFields",
		"required_fields":       "RequiredFields",
		"single_select_options": "SingleSelectOptions",
		"status":                "Status",
		"tags":                  "Tags",
		"type":                  "Type",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}