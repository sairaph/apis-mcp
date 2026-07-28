---
title: Enable validation for a schema
page_id: operation-patch-zones-zone-id-api-gateway-user-schemas-schema-id-5293996f
path: operations/api-shield-schema-validation-2-0
description: Activates schema validation for an uploaded OpenAPI schema. Requests to matching endpoints will be validated against the schema definitions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas/{schema_id}
operation_ids:
    - api-shield-schema-validation-enable-validation-for-a-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable validation for a schema

`PATCH /zones/{zone_id}/api_gateway/user_schemas/{schema_id}`

Operation ID: `api-shield-schema-validation-enable-validation-for-a-schema`

Activates schema validation for an uploaded OpenAPI schema. Requests to matching endpoints will be validated against the schema definitions.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_old_schema_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-enable-validation-for-a-schema", "summary": "Enable validation for a schema", "description": "Activates schema validation for an uploaded OpenAPI schema. Requests to matching endpoints will be validated against the schema definitions.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"validation_enabled": {"allOf": [{"$ref": "#/components/schemas/api-shield_old_validation_enabled"}, {"enum": [true]}]}}}}}}, "responses": {"200": {"description": "Enable validation for a schema response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_old_public_schema"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Enable validation for a schema response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
