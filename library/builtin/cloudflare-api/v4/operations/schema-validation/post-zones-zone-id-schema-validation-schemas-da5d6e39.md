---
title: Upload a schema
page_id: operation-post-zones-zone-id-schema-validation-schemas-37a2d494
path: operations/schema-validation
description: Uploads a new OpenAPI schema for API Shield schema validation. The schema defines expected request/response formats for API endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas
operation_ids:
    - schema-validation-create-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload a schema

`POST /zones/{zone_id}/schema_validation/schemas`

Operation ID: `schema-validation-create-schema`

Uploads a new OpenAPI schema for API Shield schema validation. The schema defines expected request/response formats for API endpoints.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-create-schema", "summary": "Upload a schema", "description": "Uploads a new OpenAPI schema for API Shield schema validation. The schema defines expected request/response formats for API endpoints.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_schema_create"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_schema_create_success"}, "4XX": {"$ref": "#/components/responses/api-shield_schema_create_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
