---
title: Get Spectrum application configuration
page_id: operation-get-zones-zone-id-spectrum-apps-app-id-795c1be6
path: operations/spectrum-applications
description: Gets the application configuration of a specific application inside a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/spectrum/apps/{app_id}
operation_ids:
    - spectrum-applications-get-spectrum-application-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Spectrum application configuration

`GET /zones/{zone_id}/spectrum/apps/{app_id}`

Operation ID: `spectrum-applications-get-spectrum-application-configuration`

Gets the application configuration of a specific application inside a zone.

## Definition

```yaml
{"operationId": "spectrum-applications-get-spectrum-application-configuration", "summary": "Get Spectrum application configuration", "description": "Gets the application configuration of a specific application inside a zone.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_app_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_zone_identifier"}}], "responses": {"200": {"description": "Get Spectrum application configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_app_config_single"}}}}, "4XX": {"description": "Get Spectrum application configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Applications"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "spectrum.apps", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
