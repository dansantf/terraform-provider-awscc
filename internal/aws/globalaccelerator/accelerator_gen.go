// Code generated by generators/resource/main.go; DO NOT EDIT.

package globalaccelerator

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
	registry.AddResourceTypeFactory("aws_globalaccelerator_accelerator", acceleratorResourceType)
}

// acceleratorResourceType returns the Terraform aws_globalaccelerator_accelerator resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GlobalAccelerator::Accelerator resource type.
func acceleratorResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"accelerator_arn": {
			// Property: AcceleratorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the accelerator.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the accelerator.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dns_name": {
			// Property: DnsName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.",
			//   "type": "string"
			// }
			Description: "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether an accelerator is enabled. The value is true or false.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether an accelerator is enabled. The value is true or false.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"ip_address_type": {
			// Property: IpAddressType
			// CloudFormation resource type schema:
			// {
			//   "description": "IP Address type.",
			//   "enum": [
			//     "IPV4",
			//     "IPV6"
			//   ],
			//   "type": "string"
			// }
			Description: "IP Address type.",
			Type:        types.StringType,
			Optional:    true,
		},
		"ip_addresses": {
			// Property: IpAddresses
			// CloudFormation resource type schema:
			// {
			//   "description": "The IP addresses from BYOIP Prefix pool.",
			//   "items": {
			//     "description": "The IP addresses from BYOIP Prefix pool.",
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The IP addresses from BYOIP Prefix pool.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of accelerator.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of accelerator.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "Tag is a key-value pair associated with accelerator.",
			//     "properties": {
			//       "Key": {
			//         "description": "Key of the tag. Value can be 1 to 127 characters.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Value for the tag. Value can be 1 to 255 characters.",
			//         "maxLength": 255,
			//         "minLength": 1,
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
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						// {
						//   "description": "Key of the tag. Value can be 1 to 127 characters.",
						//   "maxLength": 127,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "Key of the tag. Value can be 1 to 127 characters.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						// {
						//   "description": "Value for the tag. Value can be 1 to 255 characters.",
						//   "maxLength": 255,
						//   "minLength": 1,
						//   "type": "string"
						// }
						Description: "Value for the tag. Value can be 1 to 255 characters.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
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
		Description: "Resource Type definition for AWS::GlobalAccelerator::Accelerator",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::Accelerator").WithTerraformTypeName("aws_globalaccelerator_accelerator").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_globalaccelerator_accelerator", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
