---
title: Retrieve operation-level schema validation settings
page_id: operation-get-zones-zone-id-api-gateway-operations-operation-id-schema-validation-e21e9ee6
path: operations/api-shield-schema-validation-2-0
description: Retrieves operation-level schema validation settings on the zone
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation
operation_ids:
    - api-shield-schema-validation-retrieve-operation-level-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve operation-level schema validation settings

`GET /zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation`

Operation ID: `api-shield-schema-validation-retrieve-operation-level-settings`

Retrieves operation-level schema validation settings on the zone

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_operation_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-retrieve-operation-level-settings", "summary": "Retrieve operation-level schema validation settings", "description": "Retrieves operation-level schema validation settings on the zone", "responses": {"200": {"description": "Operation-level schema validation settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_old_operation_schema_validation_settings"}}}}, "4XX": {"description": "Operation-level schema validation settings response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.operations.schema-validation", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
