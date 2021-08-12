// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

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
	registry.AddResourceTypeFactory("aws_cloudfront_cache_policy", cachePolicyResourceType)
}

// cachePolicyResourceType returns the Terraform aws_cloudfront_cache_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFront::CachePolicy resource type.
func cachePolicyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"cache_policy_config": {
			// Property: CachePolicyConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Comment": {
			//       "type": "string"
			//     },
			//     "DefaultTTL": {
			//       "type": "number"
			//     },
			//     "MaxTTL": {
			//       "type": "number"
			//     },
			//     "MinTTL": {
			//       "type": "number"
			//     },
			//     "Name": {
			//       "type": "string"
			//     },
			//     "ParametersInCacheKeyAndForwardedToOrigin": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CookiesConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "CookieBehavior": {
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Cookies": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "CookieBehavior"
			//           ],
			//           "type": "object"
			//         },
			//         "EnableAcceptEncodingBrotli": {
			//           "type": "boolean"
			//         },
			//         "EnableAcceptEncodingGzip": {
			//           "type": "boolean"
			//         },
			//         "HeadersConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "HeaderBehavior": {
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Headers": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "HeaderBehavior"
			//           ],
			//           "type": "object"
			//         },
			//         "QueryStringsConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "QueryStringBehavior": {
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "QueryStrings": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "QueryStringBehavior"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "EnableAcceptEncodingGzip",
			//         "HeadersConfig",
			//         "CookiesConfig",
			//         "QueryStringsConfig"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "MinTTL",
			//     "MaxTTL",
			//     "DefaultTTL",
			//     "ParametersInCacheKeyAndForwardedToOrigin"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"comment": {
						// Property: Comment
						// CloudFormation resource type schema:
						// {
						//   "type": "string"
						// }
						Type:     types.StringType,
						Optional: true,
					},
					"default_ttl": {
						// Property: DefaultTTL
						// CloudFormation resource type schema:
						// {
						//   "type": "number"
						// }
						Type:     types.NumberType,
						Required: true,
					},
					"max_ttl": {
						// Property: MaxTTL
						// CloudFormation resource type schema:
						// {
						//   "type": "number"
						// }
						Type:     types.NumberType,
						Required: true,
					},
					"min_ttl": {
						// Property: MinTTL
						// CloudFormation resource type schema:
						// {
						//   "type": "number"
						// }
						Type:     types.NumberType,
						Required: true,
					},
					"name": {
						// Property: Name
						// CloudFormation resource type schema:
						// {
						//   "type": "string"
						// }
						Type:     types.StringType,
						Required: true,
					},
					"parameters_in_cache_key_and_forwarded_to_origin": {
						// Property: ParametersInCacheKeyAndForwardedToOrigin
						// CloudFormation resource type schema:
						// {
						//   "additionalProperties": false,
						//   "properties": {
						//     "CookiesConfig": {
						//       "additionalProperties": false,
						//       "properties": {
						//         "CookieBehavior": {
						//           "pattern": "",
						//           "type": "string"
						//         },
						//         "Cookies": {
						//           "items": {
						//             "type": "string"
						//           },
						//           "type": "array",
						//           "uniqueItems": false
						//         }
						//       },
						//       "required": [
						//         "CookieBehavior"
						//       ],
						//       "type": "object"
						//     },
						//     "EnableAcceptEncodingBrotli": {
						//       "type": "boolean"
						//     },
						//     "EnableAcceptEncodingGzip": {
						//       "type": "boolean"
						//     },
						//     "HeadersConfig": {
						//       "additionalProperties": false,
						//       "properties": {
						//         "HeaderBehavior": {
						//           "pattern": "",
						//           "type": "string"
						//         },
						//         "Headers": {
						//           "items": {
						//             "type": "string"
						//           },
						//           "type": "array",
						//           "uniqueItems": false
						//         }
						//       },
						//       "required": [
						//         "HeaderBehavior"
						//       ],
						//       "type": "object"
						//     },
						//     "QueryStringsConfig": {
						//       "additionalProperties": false,
						//       "properties": {
						//         "QueryStringBehavior": {
						//           "pattern": "",
						//           "type": "string"
						//         },
						//         "QueryStrings": {
						//           "items": {
						//             "type": "string"
						//           },
						//           "type": "array",
						//           "uniqueItems": false
						//         }
						//       },
						//       "required": [
						//         "QueryStringBehavior"
						//       ],
						//       "type": "object"
						//     }
						//   },
						//   "required": [
						//     "EnableAcceptEncodingGzip",
						//     "HeadersConfig",
						//     "CookiesConfig",
						//     "QueryStringsConfig"
						//   ],
						//   "type": "object"
						// }
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"cookies_config": {
									// Property: CookiesConfig
									// CloudFormation resource type schema:
									// {
									//   "additionalProperties": false,
									//   "properties": {
									//     "CookieBehavior": {
									//       "pattern": "",
									//       "type": "string"
									//     },
									//     "Cookies": {
									//       "items": {
									//         "type": "string"
									//       },
									//       "type": "array",
									//       "uniqueItems": false
									//     }
									//   },
									//   "required": [
									//     "CookieBehavior"
									//   ],
									//   "type": "object"
									// }
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"cookie_behavior": {
												// Property: CookieBehavior
												// CloudFormation resource type schema:
												// {
												//   "pattern": "",
												//   "type": "string"
												// }
												Type:     types.StringType,
												Required: true,
											},
											"cookies": {
												// Property: Cookies
												// CloudFormation resource type schema:
												// {
												//   "items": {
												//     "type": "string"
												//   },
												//   "type": "array",
												//   "uniqueItems": false
												// }
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
											},
										},
									),
									Required: true,
								},
								"enable_accept_encoding_brotli": {
									// Property: EnableAcceptEncodingBrotli
									// CloudFormation resource type schema:
									// {
									//   "type": "boolean"
									// }
									Type:     types.BoolType,
									Optional: true,
								},
								"enable_accept_encoding_gzip": {
									// Property: EnableAcceptEncodingGzip
									// CloudFormation resource type schema:
									// {
									//   "type": "boolean"
									// }
									Type:     types.BoolType,
									Required: true,
								},
								"headers_config": {
									// Property: HeadersConfig
									// CloudFormation resource type schema:
									// {
									//   "additionalProperties": false,
									//   "properties": {
									//     "HeaderBehavior": {
									//       "pattern": "",
									//       "type": "string"
									//     },
									//     "Headers": {
									//       "items": {
									//         "type": "string"
									//       },
									//       "type": "array",
									//       "uniqueItems": false
									//     }
									//   },
									//   "required": [
									//     "HeaderBehavior"
									//   ],
									//   "type": "object"
									// }
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"header_behavior": {
												// Property: HeaderBehavior
												// CloudFormation resource type schema:
												// {
												//   "pattern": "",
												//   "type": "string"
												// }
												Type:     types.StringType,
												Required: true,
											},
											"headers": {
												// Property: Headers
												// CloudFormation resource type schema:
												// {
												//   "items": {
												//     "type": "string"
												//   },
												//   "type": "array",
												//   "uniqueItems": false
												// }
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
											},
										},
									),
									Required: true,
								},
								"query_strings_config": {
									// Property: QueryStringsConfig
									// CloudFormation resource type schema:
									// {
									//   "additionalProperties": false,
									//   "properties": {
									//     "QueryStringBehavior": {
									//       "pattern": "",
									//       "type": "string"
									//     },
									//     "QueryStrings": {
									//       "items": {
									//         "type": "string"
									//       },
									//       "type": "array",
									//       "uniqueItems": false
									//     }
									//   },
									//   "required": [
									//     "QueryStringBehavior"
									//   ],
									//   "type": "object"
									// }
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"query_string_behavior": {
												// Property: QueryStringBehavior
												// CloudFormation resource type schema:
												// {
												//   "pattern": "",
												//   "type": "string"
												// }
												Type:     types.StringType,
												Required: true,
											},
											"query_strings": {
												// Property: QueryStrings
												// CloudFormation resource type schema:
												// {
												//   "items": {
												//     "type": "string"
												//   },
												//   "type": "array",
												//   "uniqueItems": false
												// }
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
											},
										},
									),
									Required: true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::CloudFront::CachePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::CachePolicy").WithTerraformTypeName("aws_cloudfront_cache_policy").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_cloudfront_cache_policy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
