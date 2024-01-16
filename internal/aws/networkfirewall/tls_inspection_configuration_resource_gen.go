// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_networkfirewall_tls_inspection_configuration", tLSInspectionConfigurationResource)
}

// tLSInspectionConfigurationResource returns the Terraform awscc_networkfirewall_tls_inspection_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::NetworkFirewall::TLSInspectionConfiguration resource.
func tLSInspectionConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 512,
		//	  "minLength": 1,
		//	  "pattern": "^.*$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 512),
				stringvalidator.RegexMatches(regexp.MustCompile("^.*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TLSInspectionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ServerCertificateConfigurations": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "CertificateAuthorityArn": {
		//	            "description": "A resource ARN.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "pattern": "^(arn:aws.*)$",
		//	            "type": "string"
		//	          },
		//	          "CheckCertificateRevocationStatus": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "RevokedStatusAction": {
		//	                "enum": [
		//	                  "PASS",
		//	                  "DROP",
		//	                  "REJECT"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "UnknownStatusAction": {
		//	                "enum": [
		//	                  "PASS",
		//	                  "DROP",
		//	                  "REJECT"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Scopes": {
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "DestinationPorts": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "FromPort": {
		//	                        "maximum": 65535,
		//	                        "minimum": 0,
		//	                        "type": "integer"
		//	                      },
		//	                      "ToPort": {
		//	                        "maximum": 65535,
		//	                        "minimum": 0,
		//	                        "type": "integer"
		//	                      }
		//	                    },
		//	                    "required": [
		//	                      "FromPort",
		//	                      "ToPort"
		//	                    ],
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array",
		//	                  "uniqueItems": false
		//	                },
		//	                "Destinations": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "AddressDefinition": {
		//	                        "maxLength": 255,
		//	                        "minLength": 1,
		//	                        "pattern": "^([a-fA-F\\d:\\.]+/\\d{1,3})$",
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "required": [
		//	                      "AddressDefinition"
		//	                    ],
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array",
		//	                  "uniqueItems": false
		//	                },
		//	                "Protocols": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "maximum": 255,
		//	                    "minimum": 0,
		//	                    "type": "integer"
		//	                  },
		//	                  "type": "array",
		//	                  "uniqueItems": false
		//	                },
		//	                "SourcePorts": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "FromPort": {
		//	                        "maximum": 65535,
		//	                        "minimum": 0,
		//	                        "type": "integer"
		//	                      },
		//	                      "ToPort": {
		//	                        "maximum": 65535,
		//	                        "minimum": 0,
		//	                        "type": "integer"
		//	                      }
		//	                    },
		//	                    "required": [
		//	                      "FromPort",
		//	                      "ToPort"
		//	                    ],
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array",
		//	                  "uniqueItems": false
		//	                },
		//	                "Sources": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "AddressDefinition": {
		//	                        "maxLength": 255,
		//	                        "minLength": 1,
		//	                        "pattern": "^([a-fA-F\\d:\\.]+/\\d{1,3})$",
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "required": [
		//	                      "AddressDefinition"
		//	                    ],
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array",
		//	                  "uniqueItems": false
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "ServerCertificates": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "ResourceArn": {
		//	                  "description": "A resource ARN.",
		//	                  "maxLength": 256,
		//	                  "minLength": 1,
		//	                  "pattern": "^(arn:aws.*)$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tls_inspection_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ServerCertificateConfigurations
				"server_certificate_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CertificateAuthorityArn
							"certificate_authority_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A resource ARN.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 256),
									stringvalidator.RegexMatches(regexp.MustCompile("^(arn:aws.*)$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: CheckCertificateRevocationStatus
							"check_certificate_revocation_status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: RevokedStatusAction
									"revoked_status_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"PASS",
												"DROP",
												"REJECT",
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: UnknownStatusAction
									"unknown_status_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"PASS",
												"DROP",
												"REJECT",
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Scopes
							"scopes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: DestinationPorts
										"destination_ports": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: FromPort
													"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Required: true,
														Validators: []validator.Int64{ /*START VALIDATORS*/
															int64validator.Between(0, 65535),
														}, /*END VALIDATORS*/
													}, /*END ATTRIBUTE*/
													// Property: ToPort
													"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Required: true,
														Validators: []validator.Int64{ /*START VALIDATORS*/
															int64validator.Between(0, 65535),
														}, /*END VALIDATORS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Destinations
										"destinations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: AddressDefinition
													"address_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
														Required: true,
														Validators: []validator.String{ /*START VALIDATORS*/
															stringvalidator.LengthBetween(1, 255),
															stringvalidator.RegexMatches(regexp.MustCompile("^([a-fA-F\\d:\\.]+/\\d{1,3})$"), ""),
														}, /*END VALIDATORS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Protocols
										"protocols": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.Int64Type,
											Optional:    true,
											Computed:    true,
											Validators: []validator.List{ /*START VALIDATORS*/
												listvalidator.ValueInt64sAre(
													int64validator.Between(0, 255),
												),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: SourcePorts
										"source_ports": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: FromPort
													"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Required: true,
														Validators: []validator.Int64{ /*START VALIDATORS*/
															int64validator.Between(0, 65535),
														}, /*END VALIDATORS*/
													}, /*END ATTRIBUTE*/
													// Property: ToPort
													"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Required: true,
														Validators: []validator.Int64{ /*START VALIDATORS*/
															int64validator.Between(0, 65535),
														}, /*END VALIDATORS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Sources
										"sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: AddressDefinition
													"address_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
														Required: true,
														Validators: []validator.String{ /*START VALIDATORS*/
															stringvalidator.LengthBetween(1, 255),
															stringvalidator.RegexMatches(regexp.MustCompile("^([a-fA-F\\d:\\.]+/\\d{1,3})$"), ""),
														}, /*END VALIDATORS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ServerCertificates
							"server_certificates": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ResourceArn
										"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "A resource ARN.",
											Optional:    true,
											Computed:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 256),
												stringvalidator.RegexMatches(regexp.MustCompile("^(arn:aws.*)$"), ""),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: TLSInspectionConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource ARN.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^(arn:aws.*)$",
		//	  "type": "string"
		//	}
		"tls_inspection_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource ARN.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TLSInspectionConfigurationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^([0-9a-f]{8})-([0-9a-f]{4}-){3}([0-9a-f]{12})$",
		//	  "type": "string"
		//	}
		"tls_inspection_configuration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TLSInspectionConfigurationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"tls_inspection_configuration_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^.*$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "pattern": "^.*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^.*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
							stringvalidator.RegexMatches(regexp.MustCompile("^.*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::TLSInspectionConfiguration").WithTerraformTypeName("awscc_networkfirewall_tls_inspection_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address_definition":                  "AddressDefinition",
		"certificate_authority_arn":           "CertificateAuthorityArn",
		"check_certificate_revocation_status": "CheckCertificateRevocationStatus",
		"description":                         "Description",
		"destination_ports":                   "DestinationPorts",
		"destinations":                        "Destinations",
		"from_port":                           "FromPort",
		"key":                                 "Key",
		"protocols":                           "Protocols",
		"resource_arn":                        "ResourceArn",
		"revoked_status_action":               "RevokedStatusAction",
		"scopes":                              "Scopes",
		"server_certificate_configurations":   "ServerCertificateConfigurations",
		"server_certificates":                 "ServerCertificates",
		"source_ports":                        "SourcePorts",
		"sources":                             "Sources",
		"tags":                                "Tags",
		"tls_inspection_configuration":        "TLSInspectionConfiguration",
		"tls_inspection_configuration_arn":    "TLSInspectionConfigurationArn",
		"tls_inspection_configuration_id":     "TLSInspectionConfigurationId",
		"tls_inspection_configuration_name":   "TLSInspectionConfigurationName",
		"to_port":                             "ToPort",
		"unknown_status_action":               "UnknownStatusAction",
		"value":                               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
