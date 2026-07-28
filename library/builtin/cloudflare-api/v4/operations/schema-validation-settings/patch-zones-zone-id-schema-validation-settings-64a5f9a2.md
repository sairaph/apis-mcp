---
title: Edit global schema validation settings
page_id: operation-patch-zones-zone-id-schema-validation-settings-0ec1be46
path: operations/schema-validation-settings
description: Partially updates global schema validation settings for a zone using PATCH semantics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings
operation_ids:
    - schema-validation-edit-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit global schema validation settings

`PATCH /zones/{zone_id}/schema_validation/settings`

Operation ID: `schema-validation-edit-settings`

Partially updates global schema validation settings for a zone using PATCH semantics.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-edit-settings", "summary": "Edit global schema validation settings", "description": "Partially updates global schema validation settings for a zone using PATCH semantics.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_global_settings_edit"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_global_settings_edit_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
