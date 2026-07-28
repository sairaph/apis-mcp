---
title: Update operation-level schema validation settings
page_id: operation-put-zones-zone-id-api-gateway-operations-operation-id-schema-validation-e2a9fd9a
path: operations/api-shield-schema-validation-2-0
description: Updates operation-level schema validation settings on the zone
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation
operation_ids:
    - api-shield-schema-validation-update-operation-level-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update operation-level schema validation settings

`PUT /zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation`

Operation ID: `api-shield-schema-validation-update-operation-level-settings`

Updates operation-level schema validation settings on the zone

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-update-operation-level-settings", "summary": "Update operation-level schema validation settings", "description": "Updates operation-level schema validation settings on the zone", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_old_operation_schema_validation_settings_modify_request"}}}}, "responses": {"200": {"description": "Update operation-level schema validation settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_old_operation_schema_validation_settings"}}}}, "4XX": {"description": "Update operation-level schema validation settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.operations.schema-validation", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
