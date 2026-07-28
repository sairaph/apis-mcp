---
title: Retrieve information about specific configuration properties
page_id: operation-get-zones-zone-id-api-gateway-configuration-de47ad18
path: operations/api-shield-settings
description: Gets the current API Shield configuration settings for a zone, including validation behavior and enforcement mode.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/configuration
operation_ids:
    - api-shield-settings-retrieve-information-about-specific-configuration-properties
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve information about specific configuration properties

`GET /zones/{zone_id}/api_gateway/configuration`

Operation ID: `api-shield-settings-retrieve-information-about-specific-configuration-properties`

Gets the current API Shield configuration settings for a zone, including validation behavior and enforcement mode.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"name": "normalize", "in": "query", "description": "Ensures that the configuration is written or retrieved in normalized fashion", "schema": {"type": "boolean"}}]
```

## Definition

```yaml
{"operationId": "api-shield-settings-retrieve-information-about-specific-configuration-properties", "summary": "Retrieve information about specific configuration properties", "description": "Gets the current API Shield configuration settings for a zone, including validation behavior and enforcement mode.", "responses": {"200": {"$ref": "#/components/responses/api-shield_config-get-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic-failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Settings"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.configurations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
