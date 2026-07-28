---
title: Get Zaraz workflow
page_id: operation-get-zones-zone-id-settings-zaraz-workflow-445bbf76
path: operations/zaraz
description: Gets Zaraz workflow for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/workflow
operation_ids:
    - get-zones-zone_identifier-zaraz-workflow
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zaraz workflow

`GET /zones/{zone_id}/settings/zaraz/workflow`

Operation ID: `get-zones-zone_identifier-zaraz-workflow`

Gets Zaraz workflow for a zone.

## Definition

```yaml
{"operationId": "get-zones-zone_identifier-zaraz-workflow", "summary": "Get Zaraz workflow", "description": "Gets Zaraz workflow for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "responses": {"200": {"description": "Get Zaraz workflow response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-workflow-response"}}}}, "4XX": {"description": "Get Zaraz workflow response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"]}
```
