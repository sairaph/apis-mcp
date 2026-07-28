---
title: Upload a schema to a zone
page_id: operation-post-zones-zone-id-api-gateway-user-schemas-03353d76
path: operations/api-shield-schema-validation-2-0
description: Uploads a new OpenAPI schema for API Shield schema validation. The schema defines expected request/response formats for API endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas
operation_ids:
    - api-shield-schema-validation-post-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload a schema to a zone

`POST /zones/{zone_id}/api_gateway/user_schemas`

Operation ID: `api-shield-schema-validation-post-schema`

Uploads a new OpenAPI schema for API Shield schema validation. The schema defines expected request/response formats for API endpoints.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-post-schema", "summary": "Upload a schema to a zone", "description": "Uploads a new OpenAPI schema for API Shield schema validation. The schema defines expected request/response formats for API endpoints.", "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"file": {"description": "Schema file bytes", "type": "string", "format": "binary"}, "kind": {"$ref": "#/components/schemas/api-shield_old_kind"}, "name": {"description": "Name of the schema", "type": "string", "example": "petstore schema"}, "validation_enabled": {"description": "Flag whether schema is enabled for validation.", "type": "string", "enum": ["true", "false"]}}, "required": ["file", "kind"]}}}}, "responses": {"200": {"description": "Upload a schema response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_old_schema_upload_response"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Upload a schema response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_old_schema_upload_failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
