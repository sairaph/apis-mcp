---
title: Get per-operation schema validation setting
page_id: operation-get-zones-zone-id-schema-validation-settings-operations-operation-id-81eaed67
path: operations/schema-validation-settings
description: Retrieves the schema validation settings configured for a specific API operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/settings/operations/{operation_id}
operation_ids:
    - schema-validation-get-per-operation-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get per-operation schema validation setting

`GET /zones/{zone_id}/schema_validation/settings/operations/{operation_id}`

Operation ID: `schema-validation-get-per-operation-setting`

Retrieves the schema validation settings configured for a specific API operation.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-get-per-operation-setting", "summary": "Get per-operation schema validation setting", "description": "Retrieves the schema validation settings configured for a specific API operation.", "responses": {"200": {"$ref": "#/components/responses/api-shield_per_operation_setting_get_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation Settings"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.settings.operations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
