---
title: Get Zaraz historical configurations by ID(s)
page_id: operation-get-zones-zone-id-settings-zaraz-history-configs-6e8dba9b
path: operations/zaraz
description: Gets a history of published Zaraz configurations by ID(s) for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/history/configs
operation_ids:
    - get-zones-zone_identifier-zaraz-config-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zaraz historical configurations by ID(s)

`GET /zones/{zone_id}/settings/zaraz/history/configs`

Operation ID: `get-zones-zone_identifier-zaraz-config-history`

Gets a history of published Zaraz configurations by ID(s) for a zone.

## Definition

```yaml
{"operationId": "get-zones-zone_identifier-zaraz-config-history", "summary": "Get Zaraz historical configurations by ID(s)", "description": "Gets a history of published Zaraz configurations by ID(s) for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}, {"name": "ids", "in": "query", "description": "Comma separated list of Zaraz configuration IDs.", "required": true, "schema": {"type": "array", "items": {"type": "integer"}}, "example": [12345, 23456], "explode": false, "style": "form"}], "responses": {"200": {"description": "Get Zaraz historical configurations by ID(s) response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-history-response"}}}}, "4XX": {"description": "Get Zaraz historical configurations by ID(s) failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"]}
```
