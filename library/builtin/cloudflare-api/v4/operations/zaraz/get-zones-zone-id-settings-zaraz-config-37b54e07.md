---
title: Get Zaraz configuration
page_id: operation-get-zones-zone-id-settings-zaraz-config-47b96850
path: operations/zaraz
description: Gets latest Zaraz configuration for a zone. It can be preview or published configuration, whichever was the last updated. Secret variables values will not be included.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/config
operation_ids:
    - get-zones-zone_identifier-zaraz-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zaraz configuration

`GET /zones/{zone_id}/settings/zaraz/config`

Operation ID: `get-zones-zone_identifier-zaraz-config`

Gets latest Zaraz configuration for a zone. It can be preview or published configuration, whichever was the last updated. Secret variables values will not be included.

## Definition

```yaml
{"operationId": "get-zones-zone_identifier-zaraz-config", "summary": "Get Zaraz configuration", "description": "Gets latest Zaraz configuration for a zone. It can be preview or published configuration, whichever was the last updated. Secret variables values will not be included.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "responses": {"200": {"description": "Get Zaraz configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-response"}}}}, "4XX": {"description": "Get Zaraz configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"]}
```
