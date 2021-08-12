// Code generated by generators/resource/main.go; DO NOT EDIT.

package athena

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
	registry.AddResourceTypeFactory("aws_athena_work_group", workGroupResourceType)
}

// workGroupResourceType returns the Terraform aws_athena_work_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Athena::WorkGroup resource type.
func workGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time the workgroup was created.",
			//   "type": "string"
			// }
			Description: "The date and time the workgroup was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The workgroup description.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The workgroup description.",
			Type:        types.StringType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The workGroup name.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The workGroup name.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"recursive_delete_option": {
			// Property: RecursiveDeleteOption
			// CloudFormation resource type schema:
			// {
			//   "description": "The option to delete the workgroup and its contents even if the workgroup contains any named queries.",
			//   "type": "boolean"
			// }
			Description: "The option to delete the workgroup and its contents even if the workgroup contains any named queries.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the workgroup: ENABLED or DISABLED.",
			//   "enum": [
			//     "ENABLED",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The state of the workgroup: ENABLED or DISABLED.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
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
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 128,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						// {
						//   "maxLength": 256,
						//   "minLength": 0,
						//   "type": "string"
						// }
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"work_group_configuration": {
			// Property: WorkGroupConfiguration
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "BytesScannedCutoffPerQuery": {
			//       "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
			//       "format": "int64",
			//       "type": "integer"
			//     },
			//     "EnforceWorkGroupConfiguration": {
			//       "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
			//       "type": "boolean"
			//     },
			//     "EngineVersion": {
			//       "description": "The Athena engine version for running queries.",
			//       "properties": {
			//         "EffectiveEngineVersion": {
			//           "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
			//           "type": "string"
			//         },
			//         "SelectedEngineVersion": {
			//           "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "PublishCloudWatchMetricsEnabled": {
			//       "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
			//       "type": "boolean"
			//     },
			//     "RequesterPaysEnabled": {
			//       "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
			//       "type": "boolean"
			//     },
			//     "ResultConfiguration": {
			//       "description": "The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as \"client-side settings\". If workgroup settings override client-side settings, then the query uses the workgroup settings.\n",
			//       "properties": {
			//         "EncryptionConfiguration": {
			//           "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
			//           "properties": {
			//             "EncryptionOption": {
			//               "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
			//               "enum": [
			//                 "SSE_S3",
			//                 "SSE_KMS",
			//                 "CSE_KMS"
			//               ],
			//               "type": "string"
			//             },
			//             "KmsKey": {
			//               "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "EncryptionOption"
			//           ],
			//           "type": "object"
			//         },
			//         "OutputLocation": {
			//           "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"bytes_scanned_cutoff_per_query": {
						// Property: BytesScannedCutoffPerQuery
						// CloudFormation resource type schema:
						// {
						//   "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
						//   "format": "int64",
						//   "type": "integer"
						// }
						Description: "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"enforce_work_group_configuration": {
						// Property: EnforceWorkGroupConfiguration
						// CloudFormation resource type schema:
						// {
						//   "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
						//   "type": "boolean"
						// }
						Description: "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
						Type:        types.BoolType,
						Optional:    true,
					},
					"engine_version": {
						// Property: EngineVersion
						// CloudFormation resource type schema:
						// {
						//   "description": "The Athena engine version for running queries.",
						//   "properties": {
						//     "EffectiveEngineVersion": {
						//       "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
						//       "type": "string"
						//     },
						//     "SelectedEngineVersion": {
						//       "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
						//       "type": "string"
						//     }
						//   },
						//   "type": "object"
						// }
						Description: "The Athena engine version for running queries.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"effective_engine_version": {
									// Property: EffectiveEngineVersion
									// CloudFormation resource type schema:
									// {
									//   "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
									//   "type": "string"
									// }
									Description: "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
									Type:        types.StringType,
									Computed:    true,
								},
								"selected_engine_version": {
									// Property: SelectedEngineVersion
									// CloudFormation resource type schema:
									// {
									//   "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
									//   "type": "string"
									// }
									Description: "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"publish_cloud_watch_metrics_enabled": {
						// Property: PublishCloudWatchMetricsEnabled
						// CloudFormation resource type schema:
						// {
						//   "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
						//   "type": "boolean"
						// }
						Description: "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"requester_pays_enabled": {
						// Property: RequesterPaysEnabled
						// CloudFormation resource type schema:
						// {
						//   "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
						//   "type": "boolean"
						// }
						Description: "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
						Type:        types.BoolType,
						Optional:    true,
					},
					"result_configuration": {
						// Property: ResultConfiguration
						// CloudFormation resource type schema:
						// {
						//   "description": "The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as \"client-side settings\". If workgroup settings override client-side settings, then the query uses the workgroup settings.\n",
						//   "properties": {
						//     "EncryptionConfiguration": {
						//       "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
						//       "properties": {
						//         "EncryptionOption": {
						//           "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
						//           "enum": [
						//             "SSE_S3",
						//             "SSE_KMS",
						//             "CSE_KMS"
						//           ],
						//           "type": "string"
						//         },
						//         "KmsKey": {
						//           "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
						//           "type": "string"
						//         }
						//       },
						//       "required": [
						//         "EncryptionOption"
						//       ],
						//       "type": "object"
						//     },
						//     "OutputLocation": {
						//       "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
						//       "type": "string"
						//     }
						//   },
						//   "type": "object"
						// }
						Description: "The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as \"client-side settings\". If workgroup settings override client-side settings, then the query uses the workgroup settings.\n",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"encryption_configuration": {
									// Property: EncryptionConfiguration
									// CloudFormation resource type schema:
									// {
									//   "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
									//   "properties": {
									//     "EncryptionOption": {
									//       "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
									//       "enum": [
									//         "SSE_S3",
									//         "SSE_KMS",
									//         "CSE_KMS"
									//       ],
									//       "type": "string"
									//     },
									//     "KmsKey": {
									//       "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
									//       "type": "string"
									//     }
									//   },
									//   "required": [
									//     "EncryptionOption"
									//   ],
									//   "type": "object"
									// }
									Description: "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"encryption_option": {
												// Property: EncryptionOption
												// CloudFormation resource type schema:
												// {
												//   "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
												//   "enum": [
												//     "SSE_S3",
												//     "SSE_KMS",
												//     "CSE_KMS"
												//   ],
												//   "type": "string"
												// }
												Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
												Type:        types.StringType,
												Required:    true,
											},
											"kms_key": {
												// Property: KmsKey
												// CloudFormation resource type schema:
												// {
												//   "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
												//   "type": "string"
												// }
												Description: "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
								"output_location": {
									// Property: OutputLocation
									// CloudFormation resource type schema:
									// {
									//   "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
									//   "type": "string"
									// }
									Description: "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"work_group_configuration_updates": {
			// Property: WorkGroupConfigurationUpdates
			// CloudFormation resource type schema:
			// {
			//   "description": "The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. ",
			//   "properties": {
			//     "BytesScannedCutoffPerQuery": {
			//       "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
			//       "format": "int64",
			//       "type": "integer"
			//     },
			//     "EnforceWorkGroupConfiguration": {
			//       "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
			//       "type": "boolean"
			//     },
			//     "EngineVersion": {
			//       "description": "The Athena engine version for running queries.",
			//       "properties": {
			//         "EffectiveEngineVersion": {
			//           "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
			//           "type": "string"
			//         },
			//         "SelectedEngineVersion": {
			//           "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "PublishCloudWatchMetricsEnabled": {
			//       "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
			//       "type": "boolean"
			//     },
			//     "RemoveBytesScannedCutoffPerQuery": {
			//       "description": "Indicates that the data usage control limit per query is removed.",
			//       "type": "boolean"
			//     },
			//     "RequesterPaysEnabled": {
			//       "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
			//       "type": "boolean"
			//     },
			//     "ResultConfigurationUpdates": {
			//       "description": "The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. ",
			//       "properties": {
			//         "EncryptionConfiguration": {
			//           "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
			//           "properties": {
			//             "EncryptionOption": {
			//               "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
			//               "enum": [
			//                 "SSE_S3",
			//                 "SSE_KMS",
			//                 "CSE_KMS"
			//               ],
			//               "type": "string"
			//             },
			//             "KmsKey": {
			//               "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "EncryptionOption"
			//           ],
			//           "type": "object"
			//         },
			//         "OutputLocation": {
			//           "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
			//           "type": "string"
			//         },
			//         "RemoveEncryptionConfiguration": {
			//           "type": "boolean"
			//         },
			//         "RemoveOutputLocation": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. ",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"bytes_scanned_cutoff_per_query": {
						// Property: BytesScannedCutoffPerQuery
						// CloudFormation resource type schema:
						// {
						//   "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
						//   "format": "int64",
						//   "type": "integer"
						// }
						Description: "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"enforce_work_group_configuration": {
						// Property: EnforceWorkGroupConfiguration
						// CloudFormation resource type schema:
						// {
						//   "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
						//   "type": "boolean"
						// }
						Description: "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
						Type:        types.BoolType,
						Optional:    true,
					},
					"engine_version": {
						// Property: EngineVersion
						// CloudFormation resource type schema:
						// {
						//   "description": "The Athena engine version for running queries.",
						//   "properties": {
						//     "EffectiveEngineVersion": {
						//       "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
						//       "type": "string"
						//     },
						//     "SelectedEngineVersion": {
						//       "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
						//       "type": "string"
						//     }
						//   },
						//   "type": "object"
						// }
						Description: "The Athena engine version for running queries.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"effective_engine_version": {
									// Property: EffectiveEngineVersion
									// CloudFormation resource type schema:
									// {
									//   "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
									//   "type": "string"
									// }
									Description: "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
									Type:        types.StringType,
									Computed:    true,
								},
								"selected_engine_version": {
									// Property: SelectedEngineVersion
									// CloudFormation resource type schema:
									// {
									//   "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
									//   "type": "string"
									// }
									Description: "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"publish_cloud_watch_metrics_enabled": {
						// Property: PublishCloudWatchMetricsEnabled
						// CloudFormation resource type schema:
						// {
						//   "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
						//   "type": "boolean"
						// }
						Description: "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"remove_bytes_scanned_cutoff_per_query": {
						// Property: RemoveBytesScannedCutoffPerQuery
						// CloudFormation resource type schema:
						// {
						//   "description": "Indicates that the data usage control limit per query is removed.",
						//   "type": "boolean"
						// }
						Description: "Indicates that the data usage control limit per query is removed.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"requester_pays_enabled": {
						// Property: RequesterPaysEnabled
						// CloudFormation resource type schema:
						// {
						//   "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
						//   "type": "boolean"
						// }
						Description: "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
						Type:        types.BoolType,
						Optional:    true,
					},
					"result_configuration_updates": {
						// Property: ResultConfigurationUpdates
						// CloudFormation resource type schema:
						// {
						//   "description": "The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. ",
						//   "properties": {
						//     "EncryptionConfiguration": {
						//       "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
						//       "properties": {
						//         "EncryptionOption": {
						//           "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
						//           "enum": [
						//             "SSE_S3",
						//             "SSE_KMS",
						//             "CSE_KMS"
						//           ],
						//           "type": "string"
						//         },
						//         "KmsKey": {
						//           "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
						//           "type": "string"
						//         }
						//       },
						//       "required": [
						//         "EncryptionOption"
						//       ],
						//       "type": "object"
						//     },
						//     "OutputLocation": {
						//       "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
						//       "type": "string"
						//     },
						//     "RemoveEncryptionConfiguration": {
						//       "type": "boolean"
						//     },
						//     "RemoveOutputLocation": {
						//       "type": "boolean"
						//     }
						//   },
						//   "type": "object"
						// }
						Description: "The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. ",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"encryption_configuration": {
									// Property: EncryptionConfiguration
									// CloudFormation resource type schema:
									// {
									//   "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
									//   "properties": {
									//     "EncryptionOption": {
									//       "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
									//       "enum": [
									//         "SSE_S3",
									//         "SSE_KMS",
									//         "CSE_KMS"
									//       ],
									//       "type": "string"
									//     },
									//     "KmsKey": {
									//       "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
									//       "type": "string"
									//     }
									//   },
									//   "required": [
									//     "EncryptionOption"
									//   ],
									//   "type": "object"
									// }
									Description: "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"encryption_option": {
												// Property: EncryptionOption
												// CloudFormation resource type schema:
												// {
												//   "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
												//   "enum": [
												//     "SSE_S3",
												//     "SSE_KMS",
												//     "CSE_KMS"
												//   ],
												//   "type": "string"
												// }
												Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
												Type:        types.StringType,
												Required:    true,
											},
											"kms_key": {
												// Property: KmsKey
												// CloudFormation resource type schema:
												// {
												//   "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
												//   "type": "string"
												// }
												Description: "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
								"output_location": {
									// Property: OutputLocation
									// CloudFormation resource type schema:
									// {
									//   "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
									//   "type": "string"
									// }
									Description: "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
									Type:        types.StringType,
									Optional:    true,
								},
								"remove_encryption_configuration": {
									// Property: RemoveEncryptionConfiguration
									// CloudFormation resource type schema:
									// {
									//   "type": "boolean"
									// }
									Type:     types.BoolType,
									Optional: true,
								},
								"remove_output_location": {
									// Property: RemoveOutputLocation
									// CloudFormation resource type schema:
									// {
									//   "type": "boolean"
									// }
									Type:     types.BoolType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
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
		Description: "Resource schema for AWS::Athena::WorkGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::WorkGroup").WithTerraformTypeName("aws_athena_work_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_athena_work_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
