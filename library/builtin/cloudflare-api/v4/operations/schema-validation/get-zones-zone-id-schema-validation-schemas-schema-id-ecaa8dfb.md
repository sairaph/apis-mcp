---
title: Get details of a schema
page_id: operation-get-zones-zone-id-schema-validation-schemas-schema-id-61f6631f
path: operations/schema-validation
description: Gets the contents and metadata of a specific OpenAPI schema uploaded to API Shield.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas/{schema_id}
operation_ids:
    - schema-validation-get-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get details of a schema

`GET /zones/{zone_id}/schema_validation/schemas/{schema_id}`

Operation ID: `schema-validation-get-schema`

Gets the contents and metadata of a specific OpenAPI schema uploaded to API Shield.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_schema_id_path"}]
```

## Definition

```yaml
{"operationId": "schema-validation-get-schema", "summary": "Get details of a schema", "description": "Gets the contents and metadata of a specific OpenAPI schema uploaded to API Shield.", "parameters": [{"$ref": "#/components/parameters/api-shield_omit_source_query"}], "responses": {"200": {"$ref": "#/components/responses/api-shield_schema_get_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
