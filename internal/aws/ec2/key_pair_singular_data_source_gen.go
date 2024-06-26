// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_key_pair", keyPairDataSource)
}

// keyPairDataSource returns the Terraform awscc_ec2_key_pair data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::KeyPair resource.
func keyPairDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: KeyFingerprint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A short sequence of bytes used for public key verification",
		//	  "type": "string"
		//	}
		"key_fingerprint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A short sequence of bytes used for public key verification",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "pem",
		//	  "description": "The format of the private key",
		//	  "enum": [
		//	    "pem",
		//	    "ppk"
		//	  ],
		//	  "type": "string"
		//	}
		"key_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The format of the private key",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SSH key pair",
		//	  "type": "string"
		//	}
		"key_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SSH key pair",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyPairId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An AWS generated ID for the key pair",
		//	  "type": "string"
		//	}
		"key_pair_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An AWS generated ID for the key pair",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "rsa",
		//	  "description": "The crypto-system used to generate a key pair.",
		//	  "enum": [
		//	    "rsa",
		//	    "ed25519"
		//	  ],
		//	  "type": "string"
		//	}
		"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The crypto-system used to generate a key pair.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicKeyMaterial
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Plain text public key to import",
		//	  "type": "string"
		//	}
		"public_key_material": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Plain text public key to import",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::KeyPair",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::KeyPair").WithTerraformTypeName("awscc_ec2_key_pair")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                 "Key",
		"key_fingerprint":     "KeyFingerprint",
		"key_format":          "KeyFormat",
		"key_name":            "KeyName",
		"key_pair_id":         "KeyPairId",
		"key_type":            "KeyType",
		"public_key_material": "PublicKeyMaterial",
		"tags":                "Tags",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
