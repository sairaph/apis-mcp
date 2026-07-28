---
title: Get global schema validation settings
page_id: operation-get-zones-zone-id-schema-validation-settings-5181c07c
path: operations/schema-validation-settings
description: Retrieves the current global schema validation settings for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings
operation_ids:
    - schema-validation-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get global schema validation settings

`GET /zones/{zone_id}/schema_validation/settings`

Operation ID: `schema-validation-get-settings`

Retrieves the current global schema validation settings for a zone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-get-settings", "summary": "Get global schema validation settings", "description": "Retrieves the current global schema validation settings for a zone.", "responses": {"200": {"$ref": "#/components/responses/api-shield_global_settings_get_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
