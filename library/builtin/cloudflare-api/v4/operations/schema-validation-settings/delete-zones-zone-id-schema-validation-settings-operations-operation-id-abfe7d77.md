---
title: Delete per-operation schema validation setting
page_id: operation-delete-zones-zone-id-schema-validation-settings-operations-operation-id-9068d7e8
path: operations/schema-validation-settings
description: Removes custom schema validation settings for a specific API operation, reverting to zone-level defaults.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings/operations/{operation_id}
operation_ids:
    - schema-validation-delete-per-operation-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete per-operation schema validation setting

`DELETE /zones/{zone_id}/schema_validation/settings/operations/{operation_id}`

Operation ID: `schema-validation-delete-per-operation-setting`

Removes custom schema validation settings for a specific API operation, reverting to zone-level defaults.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-delete-per-operation-setting", "summary": "Delete per-operation schema validation setting", "description": "Removes custom schema validation settings for a specific API operation, reverting to zone-level defaults.", "responses": {"200": {"$ref": "#/components/responses/api-shield_per_operation_settings_delete_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings.operations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
