---
title: Retrieve information about a specific schema on a zone
page_id: operation-get-zones-zone-id-api-gateway-user-schemas-schema-id-c1705eed
path: operations/api-shield-schema-validation-2-0
description: Gets detailed information about a specific uploaded OpenAPI schema, including its contents and validation configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas/{schema_id}
operation_ids:
    - api-shield-schema-validation-retrieve-information-about-specific-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve information about a specific schema on a zone

`GET /zones/{zone_id}/api_gateway/user_schemas/{schema_id}`

Operation ID: `api-shield-schema-validation-retrieve-information-about-specific-schema`

Gets detailed information about a specific uploaded OpenAPI schema, including its contents and validation configuration.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_old_schema_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-retrieve-information-about-specific-schema", "summary": "Retrieve information about a specific schema on a zone", "description": "Gets detailed information about a specific uploaded OpenAPI schema, including its contents and validation configuration.", "parameters": [{"$ref": "#/components/parameters/api-shield_old_omit_source"}], "responses": {"200": {"description": "Retrieve information about a specific schema on a zone response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_old_public_schema"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve information about a specific schema zone response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
