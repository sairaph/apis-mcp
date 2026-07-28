---
title: Bulk edit per-operation schema validation settings
page_id: operation-patch-zones-zone-id-schema-validation-settings-operations-6844806e
path: operations/schema-validation-settings
description: Updates schema validation settings for multiple API operations in a single request. Efficient for applying consistent validation rules across endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings/operations
operation_ids:
    - schema-validation-bulk-edit-per-operation-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk edit per-operation schema validation settings

`PATCH /zones/{zone_id}/schema_validation/settings/operations`

Operation ID: `schema-validation-bulk-edit-per-operation-settings`

Updates schema validation settings for multiple API operations in a single request. Efficient for applying consistent validation rules across endpoints.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-bulk-edit-per-operation-settings", "summary": "Bulk edit per-operation schema validation settings", "description": "Updates schema validation settings for multiple API operations in a single request. Efficient for applying consistent validation rules across endpoints.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_per_operation_settings_bulk_edit"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_per_operation_settings_bulk_edit_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings.operations", "x-fern-sdk-method-name": "bulk-edit", "x-forge-hidden": true}
```
