---
title: List per-operation schema validation settings
page_id: operation-get-zones-zone-id-schema-validation-settings-operations-ba198fc8
path: operations/schema-validation-settings
description: Lists all per-operation schema validation settings configured for the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings/operations
operation_ids:
    - schema-validation-list-per-operation-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List per-operation schema validation settings

`GET /zones/{zone_id}/schema_validation/settings/operations`

Operation ID: `schema-validation-list-per-operation-settings`

Lists all per-operation schema validation settings configured for the zone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-list-per-operation-settings", "summary": "List per-operation schema validation settings", "description": "Lists all per-operation schema validation settings configured for the zone.", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}], "responses": {"200": {"$ref": "#/components/responses/api-shield_per_operation_settings_list_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings.operations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
