// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigatewayv2_integration_response", integrationResponseDataSource)
}

// integrationResponseDataSource returns the Terraform awscc_apigatewayv2_integration_response data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGatewayV2::IntegrationResponse resource.
func integrationResponseDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API identifier.",
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContentHandlingStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are ``CONVERT_TO_BINARY`` and ``CONVERT_TO_TEXT``, with the following behaviors:\n  ``CONVERT_TO_BINARY``: Converts a response payload from a Base64-encoded string to the corresponding binary blob.\n  ``CONVERT_TO_TEXT``: Converts a response payload from a binary blob to a Base64-encoded string.\n If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.",
		//	  "type": "string"
		//	}
		"content_handling_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are ``CONVERT_TO_BINARY`` and ``CONVERT_TO_TEXT``, with the following behaviors:\n  ``CONVERT_TO_BINARY``: Converts a response payload from a Base64-encoded string to the corresponding binary blob.\n  ``CONVERT_TO_TEXT``: Converts a response payload from a binary blob to a Base64-encoded string.\n If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The integration ID.",
		//	  "type": "string"
		//	}
		"integration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The integration ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationResponseId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Integration Response ID generated by service",
		//	  "type": "string"
		//	}
		"integration_response_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Integration Response ID generated by service",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationResponseKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The integration response key.",
		//	  "type": "string"
		//	}
		"integration_response_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The integration response key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResponseParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ``method.response.header.{name}``, where name is a valid and unique header name. The mapped non-static value must match the pattern of ``integration.response.header.{name}`` or ``integration.response.body.{JSON-expression}``, where ``{name}`` is a valid and unique response header name and ``{JSON-expression}`` is a valid JSON expression without the ``$`` prefix.",
		//	  "type": "object"
		//	}
		"response_parameters": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ``method.response.header.{name}``, where name is a valid and unique header name. The mapped non-static value must match the pattern of ``integration.response.header.{name}`` or ``integration.response.body.{JSON-expression}``, where ``{name}`` is a valid and unique response header name and ``{JSON-expression}`` is a valid JSON expression without the ``$`` prefix.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResponseTemplates
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.",
		//	  "type": "object"
		//	}
		"response_templates": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateSelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The template selection expression for the integration response. Supported only for WebSocket APIs.",
		//	  "type": "string"
		//	}
		"template_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The template selection expression for the integration response. Supported only for WebSocket APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGatewayV2::IntegrationResponse",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::IntegrationResponse").WithTerraformTypeName("awscc_apigatewayv2_integration_response")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":                        "ApiId",
		"content_handling_strategy":     "ContentHandlingStrategy",
		"integration_id":                "IntegrationId",
		"integration_response_id":       "IntegrationResponseId",
		"integration_response_key":      "IntegrationResponseKey",
		"response_parameters":           "ResponseParameters",
		"response_templates":            "ResponseTemplates",
		"template_selection_expression": "TemplateSelectionExpression",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
