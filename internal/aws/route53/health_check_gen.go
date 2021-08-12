// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53

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
	registry.AddResourceTypeFactory("aws_route53_health_check", healthCheckResourceType)
}

// healthCheckResourceType returns the Terraform aws_route53_health_check resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53::HealthCheck resource type.
func healthCheckResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"health_check_config": {
			// Property: HealthCheckConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A complex type that contains information about the health check.",
			//   "properties": {
			//     "AlarmIdentifier": {
			//       "additionalProperties": false,
			//       "description": "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
			//       "properties": {
			//         "Name": {
			//           "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
			//           "maxLength": 256,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Region": {
			//           "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Name",
			//         "Region"
			//       ],
			//       "type": "object"
			//     },
			//     "ChildHealthChecks": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 256,
			//       "type": "array"
			//     },
			//     "EnableSNI": {
			//       "type": "boolean"
			//     },
			//     "FailureThreshold": {
			//       "type": "integer"
			//     },
			//     "FullyQualifiedDomainName": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "HealthThreshold": {
			//       "type": "integer"
			//     },
			//     "IPAddress": {
			//       "maxLength": 45,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "InsufficientDataHealthStatus": {
			//       "enum": [
			//         "Healthy",
			//         "LastKnownStatus",
			//         "Unhealthy"
			//       ],
			//       "type": "string"
			//     },
			//     "Inverted": {
			//       "type": "boolean"
			//     },
			//     "MeasureLatency": {
			//       "type": "boolean"
			//     },
			//     "Port": {
			//       "type": "integer"
			//     },
			//     "Regions": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 64,
			//       "type": "array"
			//     },
			//     "RequestInterval": {
			//       "type": "integer"
			//     },
			//     "ResourcePath": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "RoutingControlArn": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SearchString": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "Type": {
			//       "enum": [
			//         "CALCULATED",
			//         "CLOUDWATCH_METRIC",
			//         "HTTP",
			//         "HTTP_STR_MATCH",
			//         "HTTPS",
			//         "HTTPS_STR_MATCH",
			//         "TCP",
			//         "RECOVERY_CONTROL"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Description: "A complex type that contains information about the health check.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"alarm_identifier": {
						// Property: AlarmIdentifier
						// CloudFormation resource type schema:
						// {
						//   "additionalProperties": false,
						//   "description": "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
						//   "properties": {
						//     "Name": {
						//       "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
						//       "maxLength": 256,
						//       "minLength": 1,
						//       "type": "string"
						//     },
						//     "Region": {
						//       "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
						//       "type": "string"
						//     }
						//   },
						//   "required": [
						//     "Name",
						//     "Region"
						//   ],
						//   "type": "object"
						// }
						Description: "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"name": {
									// Property: Name
									// CloudFormation resource type schema:
									// {
									//   "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
									//   "maxLength": 256,
									//   "minLength": 1,
									//   "type": "string"
									// }
									Description: "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
									Type:        types.StringType,
									Required:    true,
								},
								"region": {
									// Property: Region
									// CloudFormation resource type schema:
									// {
									//   "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
									//   "type": "string"
									// }
									Description: "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
					"child_health_checks": {
						// Property: ChildHealthChecks
						// CloudFormation resource type schema:
						// {
						//   "insertionOrder": false,
						//   "items": {
						//     "type": "string"
						//   },
						//   "maxItems": 256,
						//   "type": "array"
						// }
						// Multiset.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"enable_sni": {
						// Property: EnableSNI
						// CloudFormation resource type schema:
						// {
						//   "type": "boolean"
						// }
						Type:     types.BoolType,
						Optional: true,
					},
					"failure_threshold": {
						// Property: FailureThreshold
						// CloudFormation resource type schema:
						// {
						//   "type": "integer"
						// }
						Type:     types.NumberType,
						Optional: true,
					},
					"fully_qualified_domain_name": {
						// Property: FullyQualifiedDomainName
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 255,
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"health_threshold": {
						// Property: HealthThreshold
						// CloudFormation resource type schema:
						// {
						//   "type": "integer"
						// }
						Type:     types.NumberType,
						Optional: true,
					},
					"ip_address": {
						// Property: IPAddress
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 45,
						//   "pattern": "",
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"insufficient_data_health_status": {
						// Property: InsufficientDataHealthStatus
						// CloudFormation resource type schema:
						// {
						//   "enum": [
						//     "Healthy",
						//     "LastKnownStatus",
						//     "Unhealthy"
						//   ],
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"inverted": {
						// Property: Inverted
						// CloudFormation resource type schema:
						// {
						//   "type": "boolean"
						// }
						Type:     types.BoolType,
						Optional: true,
					},
					"measure_latency": {
						// Property: MeasureLatency
						// CloudFormation resource type schema:
						// {
						//   "type": "boolean"
						// }
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						// MeasureLatency is a force-new attribute.
					},
					"port": {
						// Property: Port
						// CloudFormation resource type schema:
						// {
						//   "type": "integer"
						// }
						Type:     types.NumberType,
						Optional: true,
					},
					"regions": {
						// Property: Regions
						// CloudFormation resource type schema:
						// {
						//   "insertionOrder": false,
						//   "items": {
						//     "type": "string"
						//   },
						//   "maxItems": 64,
						//   "type": "array"
						// }
						// Multiset.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"request_interval": {
						// Property: RequestInterval
						// CloudFormation resource type schema:
						// {
						//   "type": "integer"
						// }
						Type:     types.NumberType,
						Optional: true,
						Computed: true,
						// RequestInterval is a force-new attribute.
					},
					"resource_path": {
						// Property: ResourcePath
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 255,
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"routing_control_arn": {
						// Property: RoutingControlArn
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 255,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"search_string": {
						// Property: SearchString
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 255,
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"type": {
						// Property: Type
						// CloudFormation resource type schema:
						// {
						//   "enum": [
						//     "CALCULATED",
						//     "CLOUDWATCH_METRIC",
						//     "HTTP",
						//     "HTTP_STR_MATCH",
						//     "HTTPS",
						//     "HTTPS_STR_MATCH",
						//     "TCP",
						//     "RECOVERY_CONTROL"
						//   ],
						//   "type": "string"
						// }
						Type:     types.StringType,
						Required: true,
						// Type is a force-new attribute.
					},
				},
			),
			Required: true,
		},
		"health_check_id": {
			// Property: HealthCheckId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"health_check_tags": {
			// Property: HealthCheckTags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag.",
			//         "maxLength": 128,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag.",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
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
						//   "description": "The key name of the tag.",
						//   "maxLength": 128,
						//   "type": "string"
						// }
						Description: "The key name of the tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						// {
						//   "description": "The value for the tag.",
						//   "maxLength": 256,
						//   "type": "string"
						// }
						Description: "The value for the tag.",
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
		Description: "Resource schema for AWS::Route53::HealthCheck.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::HealthCheck").WithTerraformTypeName("aws_route53_health_check").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_route53_health_check", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
