---
title: Update global schema validation settings
page_id: operation-put-zones-zone-id-schema-validation-settings-63a20fa9
path: operations/schema-validation-settings
description: Fully updates global schema validation settings for a zone, replacing existing configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings
operation_ids:
    - schema-validation-update-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update global schema validation settings

`PUT /zones/{zone_id}/schema_validation/settings`

Operation ID: `schema-validation-update-settings`

Fully updates global schema validation settings for a zone, replacing existing configuration.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-update-settings", "summary": "Update global schema validation settings", "description": "Fully updates global schema validation settings for a zone, replacing existing configuration.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_global_settings_update"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_global_settings_update_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
