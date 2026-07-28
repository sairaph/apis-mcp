---
title: Update configuration properties
page_id: operation-put-zones-zone-id-api-gateway-configuration-baf14eac
path: operations/api-shield-settings
description: Updates API Shield configuration settings for a zone. Can modify validation strictness, enforcement mode, and other global settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/api_gateway/configuration
operation_ids:
    - api-shield-settings-set-configuration-properties
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update configuration properties

`PUT /zones/{zone_id}/api_gateway/configuration`

Operation ID: `api-shield-settings-set-configuration-properties`

Updates API Shield configuration settings for a zone. Can modify validation strictness, enforcement mode, and other global settings.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"name": "normalize", "in": "query", "description": "Ensures that the configuration is written or retrieved in normalized fashion", "schema": {"type": "boolean"}}]
```

## Definition

```yaml
{"operationId": "api-shield-settings-set-configuration-properties", "summary": "Update configuration properties", "description": "Updates API Shield configuration settings for a zone. Can modify validation strictness, enforcement mode, and other global settings.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_config-update"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_config-update-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic-failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Settings"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.configurations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
