---
title: List all uploaded schemas
page_id: operation-get-zones-zone-id-schema-validation-schemas-934a4d67
path: operations/schema-validation
description: Lists all OpenAPI schemas uploaded to API Shield with pagination support.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas
operation_ids:
    - schema-validation-list-schemas-paginated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all uploaded schemas

`GET /zones/{zone_id}/schema_validation/schemas`

Operation ID: `schema-validation-list-schemas-paginated`

Lists all OpenAPI schemas uploaded to API Shield with pagination support.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-list-schemas-paginated", "summary": "List all uploaded schemas", "description": "Lists all OpenAPI schemas uploaded to API Shield with pagination support.", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"$ref": "#/components/parameters/api-shield_omit_source_query"}, {"name": "validation_enabled", "in": "query", "description": "Filter for enabled schemas", "schema": {"description": "Flag whether schema is enabled for validation.", "type": "boolean"}}], "responses": {"200": {"$ref": "#/components/responses/api-shield_schemas_list_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
