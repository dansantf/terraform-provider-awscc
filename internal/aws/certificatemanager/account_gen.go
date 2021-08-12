// Code generated by generators/resource/main.go; DO NOT EDIT.

package certificatemanager

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
	registry.AddResourceTypeFactory("aws_certificatemanager_account", accountResourceType)
}

// accountResourceType returns the Terraform aws_certificatemanager_account resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CertificateManager::Account resource type.
func accountResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"account_id": {
			// Property: AccountId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"expiry_events_configuration": {
			// Property: ExpiryEventsConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DaysBeforeExpiry": {
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"days_before_expiry": {
						// Property: DaysBeforeExpiry
						// CloudFormation resource type schema:
						// {
						//   "type": "integer"
						// }
						Type:     types.NumberType,
						Optional: true,
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
		Description: "Resource schema for AWS::CertificateManager::Account.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CertificateManager::Account").WithTerraformTypeName("aws_certificatemanager_account").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_certificatemanager_account", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
