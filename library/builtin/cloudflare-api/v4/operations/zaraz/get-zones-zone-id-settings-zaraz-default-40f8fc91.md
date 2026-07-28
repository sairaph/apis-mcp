---
title: Get default Zaraz configuration
page_id: operation-get-zones-zone-id-settings-zaraz-default-156793fa
path: operations/zaraz
description: Gets default Zaraz configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/default
operation_ids:
    - get-zones-zone_identifier-zaraz-default
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get default Zaraz configuration

`GET /zones/{zone_id}/settings/zaraz/default`

Operation ID: `get-zones-zone_identifier-zaraz-default`

Gets default Zaraz configuration for a zone.

## Definition

```yaml
{"operationId": "get-zones-zone_identifier-zaraz-default", "summary": "Get default Zaraz configuration", "description": "Gets default Zaraz configuration for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "responses": {"200": {"description": "Get Zaraz default configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-response"}}}}, "4XX": {"description": "Get Zaraz default configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"]}
```
