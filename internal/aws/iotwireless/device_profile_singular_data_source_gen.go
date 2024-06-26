// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_device_profile", deviceProfileDataSource)
}

// deviceProfileDataSource returns the Terraform awscc_iotwireless_device_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::DeviceProfile resource.
func deviceProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service profile Arn. Returned after successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service profile Arn. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service profile Id. Returned after successful create.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"device_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service profile Id. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoRaWAN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation",
		//	  "properties": {
		//	    "ClassBTimeout": {
		//	      "maximum": 1000,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "ClassCTimeout": {
		//	      "maximum": 1000,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "FactoryPresetFreqsList": {
		//	      "items": {
		//	        "maximum": 16700000,
		//	        "minimum": 1000000,
		//	        "type": "integer"
		//	      },
		//	      "maxItems": 20,
		//	      "type": "array"
		//	    },
		//	    "MacVersion": {
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "MaxDutyCycle": {
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "MaxEirp": {
		//	      "maximum": 15,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "PingSlotDr": {
		//	      "maximum": 15,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "PingSlotFreq": {
		//	      "maximum": 16700000,
		//	      "minimum": 1000000,
		//	      "type": "integer"
		//	    },
		//	    "PingSlotPeriod": {
		//	      "maximum": 4096,
		//	      "minimum": 128,
		//	      "type": "integer"
		//	    },
		//	    "RegParamsRevision": {
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "RfRegion": {
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "RxDataRate2": {
		//	      "maximum": 15,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "RxDelay1": {
		//	      "maximum": 15,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "RxDrOffset1": {
		//	      "maximum": 7,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "RxFreq2": {
		//	      "maximum": 16700000,
		//	      "minimum": 1000000,
		//	      "type": "integer"
		//	    },
		//	    "Supports32BitFCnt": {
		//	      "type": "boolean"
		//	    },
		//	    "SupportsClassB": {
		//	      "type": "boolean"
		//	    },
		//	    "SupportsClassC": {
		//	      "type": "boolean"
		//	    },
		//	    "SupportsJoin": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lo_ra_wan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ClassBTimeout
				"class_b_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ClassCTimeout
				"class_c_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: FactoryPresetFreqsList
				"factory_preset_freqs_list": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.Int64Type,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MacVersion
				"mac_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MaxDutyCycle
				"max_duty_cycle": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MaxEirp
				"max_eirp": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PingSlotDr
				"ping_slot_dr": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PingSlotFreq
				"ping_slot_freq": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PingSlotPeriod
				"ping_slot_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RegParamsRevision
				"reg_params_revision": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RfRegion
				"rf_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RxDataRate2
				"rx_data_rate_2": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RxDelay1
				"rx_delay_1": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RxDrOffset1
				"rx_dr_offset_1": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RxFreq2
				"rx_freq_2": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Supports32BitFCnt
				"supports_32_bit_f_cnt": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SupportsClassB
				"supports_class_b": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SupportsClassC
				"supports_class_c": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SupportsJoin
				"supports_join": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of service profile",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of service profile",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the device profile.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the device profile.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::DeviceProfile",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::DeviceProfile").WithTerraformTypeName("awscc_iotwireless_device_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"class_b_timeout":           "ClassBTimeout",
		"class_c_timeout":           "ClassCTimeout",
		"device_profile_id":         "Id",
		"factory_preset_freqs_list": "FactoryPresetFreqsList",
		"key":                       "Key",
		"lo_ra_wan":                 "LoRaWAN",
		"mac_version":               "MacVersion",
		"max_duty_cycle":            "MaxDutyCycle",
		"max_eirp":                  "MaxEirp",
		"name":                      "Name",
		"ping_slot_dr":              "PingSlotDr",
		"ping_slot_freq":            "PingSlotFreq",
		"ping_slot_period":          "PingSlotPeriod",
		"reg_params_revision":       "RegParamsRevision",
		"rf_region":                 "RfRegion",
		"rx_data_rate_2":            "RxDataRate2",
		"rx_delay_1":                "RxDelay1",
		"rx_dr_offset_1":            "RxDrOffset1",
		"rx_freq_2":                 "RxFreq2",
		"supports_32_bit_f_cnt":     "Supports32BitFCnt",
		"supports_class_b":          "SupportsClassB",
		"supports_class_c":          "SupportsClassC",
		"supports_join":             "SupportsJoin",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
