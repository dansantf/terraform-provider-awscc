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
	registry.AddDataSourceFactory("awscc_apigatewayv2_api", apiDataSource)
}

// apiDataSource returns the Terraform awscc_apigatewayv2_api data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGatewayV2::Api resource.
func apiDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ApiKeySelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).",
		//	  "type": "string"
		//	}
		"api_key_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BasePath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.",
		//	  "type": "string"
		//	}
		"base_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Body
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.",
		//	  "type": "object"
		//	}
		"body": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BodyS3Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.",
		//	  "properties": {
		//	    "Bucket": {
		//	      "description": "The S3 bucket that contains the OpenAPI definition to import. Required if you specify a ``BodyS3Location`` for an API.",
		//	      "type": "string"
		//	    },
		//	    "Etag": {
		//	      "description": "The Etag of the S3 object.",
		//	      "type": "string"
		//	    },
		//	    "Key": {
		//	      "description": "The key of the S3 object. Required if you specify a ``BodyS3Location`` for an API.",
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the S3 object.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"body_s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Bucket
				"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The S3 bucket that contains the OpenAPI definition to import. Required if you specify a ``BodyS3Location`` for an API.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Etag
				"etag": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Etag of the S3 object.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Key
				"key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The key of the S3 object. Required if you specify a ``BodyS3Location`` for an API.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the S3 object.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CorsConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
		//	  "properties": {
		//	    "AllowCredentials": {
		//	      "description": "Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.",
		//	      "type": "boolean"
		//	    },
		//	    "AllowHeaders": {
		//	      "description": "Represents a collection of allowed headers. Supported only for HTTP APIs.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "AllowMethods": {
		//	      "description": "Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "AllowOrigins": {
		//	      "description": "Represents a collection of allowed origins. Supported only for HTTP APIs.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "ExposeHeaders": {
		//	      "description": "Represents a collection of exposed headers. Supported only for HTTP APIs.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "MaxAge": {
		//	      "description": "The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cors_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowCredentials
				"allow_credentials": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowHeaders
				"allow_headers": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of allowed headers. Supported only for HTTP APIs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowMethods
				"allow_methods": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowOrigins
				"allow_origins": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of allowed origins. Supported only for HTTP APIs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExposeHeaders
				"expose_headers": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of exposed headers. Supported only for HTTP APIs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxAge
				"max_age": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CredentialsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.",
		//	  "type": "string"
		//	}
		"credentials_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the API.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the API.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisableExecuteApiEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
		//	  "type": "boolean"
		//	}
		"disable_execute_api_endpoint": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisableSchemaValidation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Avoid validating models when creating a deployment. Supported only for WebSocket APIs.",
		//	  "type": "boolean"
		//	}
		"disable_schema_validation": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Avoid validating models when creating a deployment. Supported only for WebSocket APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FailOnWarnings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.",
		//	  "type": "boolean"
		//	}
		"fail_on_warnings": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProtocolType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.",
		//	  "type": "string"
		//	}
		"protocol_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RouteKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.",
		//	  "type": "string"
		//	}
		"route_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RouteSelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.",
		//	  "type": "string"
		//	}
		"route_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The collection of tags. Each tag element is associated with a given resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The collection of tags. Each tag element is associated with a given resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Target
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.",
		//	  "type": "string"
		//	}
		"target": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A version identifier for the API.",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A version identifier for the API.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGatewayV2::Api",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::Api").WithTerraformTypeName("awscc_apigatewayv2_api")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_credentials":            "AllowCredentials",
		"allow_headers":                "AllowHeaders",
		"allow_methods":                "AllowMethods",
		"allow_origins":                "AllowOrigins",
		"api_endpoint":                 "ApiEndpoint",
		"api_id":                       "ApiId",
		"api_key_selection_expression": "ApiKeySelectionExpression",
		"base_path":                    "BasePath",
		"body":                         "Body",
		"body_s3_location":             "BodyS3Location",
		"bucket":                       "Bucket",
		"cors_configuration":           "CorsConfiguration",
		"credentials_arn":              "CredentialsArn",
		"description":                  "Description",
		"disable_execute_api_endpoint": "DisableExecuteApiEndpoint",
		"disable_schema_validation":    "DisableSchemaValidation",
		"etag":                         "Etag",
		"expose_headers":               "ExposeHeaders",
		"fail_on_warnings":             "FailOnWarnings",
		"key":                          "Key",
		"max_age":                      "MaxAge",
		"name":                         "Name",
		"protocol_type":                "ProtocolType",
		"route_key":                    "RouteKey",
		"route_selection_expression":   "RouteSelectionExpression",
		"tags":                         "Tags",
		"target":                       "Target",
		"version":                      "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
