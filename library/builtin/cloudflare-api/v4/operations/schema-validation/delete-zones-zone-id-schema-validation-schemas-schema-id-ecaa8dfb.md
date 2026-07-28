---
title: Delete a schema
page_id: operation-delete-zones-zone-id-schema-validation-schemas-schema-id-be30e752
path: operations/schema-validation
description: Permanently removes an uploaded OpenAPI schema from API Shield. Operations using this schema will lose their validation rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas/{schema_id}
operation_ids:
    - schema-validation-delete-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a schema

`DELETE /zones/{zone_id}/schema_validation/schemas/{schema_id}`

Operation ID: `schema-validation-delete-schema`

Permanently removes an uploaded OpenAPI schema from API Shield. Operations using this schema will lose their validation rules.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_schema_id_path"}]
```

## Definition

```yaml
{"operationId": "schema-validation-delete-schema", "summary": "Delete a schema", "description": "Permanently removes an uploaded OpenAPI schema from API Shield. Operations using this schema will lose their validation rules.", "responses": {"200": {"$ref": "#/components/responses/api-shield_schema_delete_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
