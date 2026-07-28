---
title: Edit details of a schema to enable validation
page_id: operation-patch-zones-zone-id-schema-validation-schemas-schema-id-d2d3434b
path: operations/schema-validation
description: Modifies an existing OpenAPI schema in API Shield, updating the validation rules for associated API operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas/{schema_id}
operation_ids:
    - schema-validation-edit-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit details of a schema to enable validation

`PATCH /zones/{zone_id}/schema_validation/schemas/{schema_id}`

Operation ID: `schema-validation-edit-schema`

Modifies an existing OpenAPI schema in API Shield, updating the validation rules for associated API operations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_schema_id_path"}]
```

## Definition

```yaml
{"operationId": "schema-validation-edit-schema", "summary": "Edit details of a schema to enable validation", "description": "Modifies an existing OpenAPI schema in API Shield, updating the validation rules for associated API operations.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_schema_edit"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_schema_edit_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
