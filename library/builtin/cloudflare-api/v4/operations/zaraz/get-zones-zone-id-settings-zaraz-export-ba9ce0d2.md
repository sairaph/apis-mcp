---
title: Export Zaraz configuration
page_id: operation-get-zones-zone-id-settings-zaraz-export-0178acfe
path: operations/zaraz
description: Exports full current published Zaraz configuration for a zone, secret variables included.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/export
operation_ids:
    - get-zones-zone_identifier-zaraz-export
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Export Zaraz configuration

`GET /zones/{zone_id}/settings/zaraz/export`

Operation ID: `get-zones-zone_identifier-zaraz-export`

Exports full current published Zaraz configuration for a zone, secret variables included.

## Definition

```yaml
{"operationId": "get-zones-zone_identifier-zaraz-export", "summary": "Export Zaraz configuration", "description": "Exports full current published Zaraz configuration for a zone, secret variables included.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "responses": {"200": {"description": "Get Zaraz configuration response.", "headers": {"Content-Disposition": {"example": "attachment; filename=zaraz-2023-11-10-23-00.json", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-return"}}}}, "4XX": {"description": "Get Zaraz configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Read", "Zaraz Admin"]}
```
