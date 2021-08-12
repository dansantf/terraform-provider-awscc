// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

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
	registry.AddResourceTypeFactory("aws_lambda_function", functionResourceType)
}

// functionResourceType returns the Terraform aws_lambda_function resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lambda::Function resource type.
func functionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier for function resources",
			//   "type": "string"
			// }
			Description: "Unique identifier for function resources",
			Type:        types.StringType,
			Computed:    true,
		},
		"code": {
			// Property: Code
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ImageUri": {
			//       "description": "ImageUri.",
			//       "type": "string"
			//     },
			//     "S3Bucket": {
			//       "description": "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
			//       "maxLength": 63,
			//       "minLength": 3,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "S3Key": {
			//       "description": "The Amazon S3 key of the deployment package.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "S3ObjectVersion": {
			//       "description": "For versioned objects, the version of the deployment package object to use.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "ZipFile": {
			//       "description": "The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"image_uri": {
						// Property: ImageUri
						// CloudFormation resource type schema:
						// {
						//   "description": "ImageUri.",
						//   "type": "string"
						// }
						Description: "ImageUri.",
						Type:        types.StringType,
						Optional:    true,
					},
					"s3_bucket": {
						// Property: S3Bucket
						// CloudFormation resource type schema:
						// {
						//   "description": "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
						//   "maxLength": 63,
						//   "minLength": 3,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
						Type:        types.StringType,
						Optional:    true,
					},
					"s3_key": {
						// Property: S3Key
						// CloudFormation resource type schema:
						// {
						//   "description": "The Amazon S3 key of the deployment package.",
						//   "maxLength": 1024,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "The Amazon S3 key of the deployment package.",
						Type:        types.StringType,
						Optional:    true,
					},
					"s3_object_version": {
						// Property: S3ObjectVersion
						// CloudFormation resource type schema:
						// {
						//   "description": "For versioned objects, the version of the deployment package object to use.",
						//   "maxLength": 1024,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "For versioned objects, the version of the deployment package object to use.",
						Type:        types.StringType,
						Optional:    true,
					},
					"zip_file": {
						// Property: ZipFile
						// CloudFormation resource type schema:
						// {
						//   "description": "The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..",
						//   "type": "string"
						// }
						Description: "The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
			// Code is a write-only attribute.
		},
		"code_signing_config_arn": {
			// Property: CodeSigningConfigArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique Arn for CodeSigningConfig resource",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique Arn for CodeSigningConfig resource",
			Type:        types.StringType,
			Optional:    true,
		},
		"dead_letter_config": {
			// Property: DeadLetterConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The dead-letter queue for failed asynchronous invocations.",
			//   "properties": {
			//     "TargetArn": {
			//       "description": "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The dead-letter queue for failed asynchronous invocations.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"target_arn": {
						// Property: TargetArn
						// CloudFormation resource type schema:
						// {
						//   "description": "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the function.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "A description of the function.",
			Type:        types.StringType,
			Optional:    true,
		},
		"environment": {
			// Property: Environment
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A function's environment variable settings.",
			//   "properties": {
			//     "Variables": {
			//       "additionalProperties": false,
			//       "description": "Environment variable key-value pairs.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A function's environment variable settings.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"variables": {
						// Property: Variables
						// CloudFormation resource type schema:
						// {
						//   "additionalProperties": false,
						//   "description": "Environment variable key-value pairs.",
						//   "patternProperties": {
						//     "": {
						//       "type": "string"
						//     }
						//   },
						//   "type": "object"
						// }
						Description: "Environment variable key-value pairs.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"file_system_configs": {
			// Property: FileSystemConfigs
			// CloudFormation resource type schema:
			// {
			//   "description": "Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Arn": {
			//         "description": "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
			//         "maxLength": 200,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "LocalMountPath": {
			//         "description": "The path where the function can access the file system, starting with /mnt/.",
			//         "maxLength": 160,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Arn",
			//       "LocalMountPath"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "type": "array"
			// }
			Description: "Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"arn": {
						// Property: Arn
						// CloudFormation resource type schema:
						// {
						//   "description": "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
						//   "maxLength": 200,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
						Type:        types.StringType,
						Required:    true,
					},
					"local_mount_path": {
						// Property: LocalMountPath
						// CloudFormation resource type schema:
						// {
						//   "description": "The path where the function can access the file system, starting with /mnt/.",
						//   "maxLength": 160,
						//   "pattern": "",
						//   "type": "string"
						// }
						Description: "The path where the function can access the file system, starting with /mnt/.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 1,
				},
			),
			Optional: true,
		},
		"function_name": {
			// Property: FunctionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// FunctionName is a force-new attribute.
		},
		"handler": {
			// Property: Handler
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime",
			Type:        types.StringType,
			Optional:    true,
		},
		"image_config": {
			// Property: ImageConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Command": {
			//       "description": "Command.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 1500,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "EntryPoint": {
			//       "description": "EntryPoint.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 1500,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "WorkingDirectory": {
			//       "description": "WorkingDirectory.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"command": {
						// Property: Command
						// CloudFormation resource type schema:
						// {
						//   "description": "Command.",
						//   "items": {
						//     "type": "string"
						//   },
						//   "maxItems": 1500,
						//   "type": "array",
						//   "uniqueItems": true
						// }
						Description: "Command.",
						// Ordered set.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"entry_point": {
						// Property: EntryPoint
						// CloudFormation resource type schema:
						// {
						//   "description": "EntryPoint.",
						//   "items": {
						//     "type": "string"
						//   },
						//   "maxItems": 1500,
						//   "type": "array",
						//   "uniqueItems": true
						// }
						Description: "EntryPoint.",
						// Ordered set.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"working_directory": {
						// Property: WorkingDirectory
						// CloudFormation resource type schema:
						// {
						//   "description": "WorkingDirectory.",
						//   "type": "string"
						// }
						Description: "WorkingDirectory.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"kms_key_arn": {
			// Property: KmsKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
			Type:        types.StringType,
			Optional:    true,
		},
		"layers": {
			// Property: Layers
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"memory_size": {
			// Property: MemorySize
			// CloudFormation resource type schema:
			// {
			//   "description": "The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.",
			//   "type": "integer"
			// }
			Description: "The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"package_type": {
			// Property: PackageType
			// CloudFormation resource type schema:
			// {
			//   "description": "PackageType.",
			//   "enum": [
			//     "Image",
			//     "Zip"
			//   ],
			//   "type": "string"
			// }
			Description: "PackageType.",
			Type:        types.StringType,
			Optional:    true,
		},
		"reserved_concurrent_executions": {
			// Property: ReservedConcurrentExecutions
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of simultaneous executions to reserve for the function.",
			//   "type": "integer"
			// }
			Description: "The number of simultaneous executions to reserve for the function.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"role": {
			// Property: Role
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the function's execution role.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the function's execution role.",
			Type:        types.StringType,
			Required:    true,
		},
		"runtime": {
			// Property: Runtime
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the function's runtime.",
			//   "type": "string"
			// }
			Description: "The identifier of the function's runtime.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the function.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of tags to apply to the function.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						//   "maxLength": 128,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						// {
						//   "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						//   "maxLength": 256,
						//   "minLength": 0,
						//   "type": "string"
						// }
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"timeout": {
			// Property: Timeout
			// CloudFormation resource type schema:
			// {
			//   "description": "The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.",
			//   "type": "integer"
			// }
			Description: "The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"tracing_config": {
			// Property: TracingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The function's AWS X-Ray tracing configuration. To sample and record incoming requests, set Mode to Active.",
			//   "properties": {
			//     "Mode": {
			//       "description": "The tracing mode.",
			//       "enum": [
			//         "Active",
			//         "PassThrough"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The function's AWS X-Ray tracing configuration. To sample and record incoming requests, set Mode to Active.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"mode": {
						// Property: Mode
						// CloudFormation resource type schema:
						// {
						//   "description": "The tracing mode.",
						//   "enum": [
						//     "Active",
						//     "PassThrough"
						//   ],
						//   "type": "string"
						// }
						Description: "The tracing mode.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"vpc_config": {
			// Property: VpcConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC.",
			//   "properties": {
			//     "SecurityGroupIds": {
			//       "description": "A list of VPC security groups IDs.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 5,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SubnetIds": {
			//       "description": "A list of VPC subnet IDs.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 16,
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"security_group_ids": {
						// Property: SecurityGroupIds
						// CloudFormation resource type schema:
						// {
						//   "description": "A list of VPC security groups IDs.",
						//   "items": {
						//     "type": "string"
						//   },
						//   "maxItems": 5,
						//   "type": "array",
						//   "uniqueItems": false
						// }
						Description: "A list of VPC security groups IDs.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
					"subnet_ids": {
						// Property: SubnetIds
						// CloudFormation resource type schema:
						// {
						//   "description": "A list of VPC subnet IDs.",
						//   "items": {
						//     "type": "string"
						//   },
						//   "maxItems": 16,
						//   "type": "array",
						//   "uniqueItems": false
						// }
						Description: "A list of VPC subnet IDs.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
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
		Description: "Resource Type definition for AWS::Lambda::Function",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::Function").WithTerraformTypeName("aws_lambda_function").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Code",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_lambda_function", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
