---
title: Update per-operation schema validation setting
page_id: operation-put-zones-zone-id-schema-validation-settings-operations-operation-id-774ffe04
path: operations/schema-validation-settings
description: Fully updates schema validation settings for a specific API operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings/operations/{operation_id}
operation_ids:
    - schema-validation-update-per-operation-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update per-operation schema validation setting

`PUT /zones/{zone_id}/schema_validation/settings/operations/{operation_id}`

Operation ID: `schema-validation-update-per-operation-setting`

Fully updates schema validation settings for a specific API operation.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-update-per-operation-setting", "summary": "Update per-operation schema validation setting", "description": "Fully updates schema validation settings for a specific API operation.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_per_operation_setting_update"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_per_operation_setting_update_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings.operations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
