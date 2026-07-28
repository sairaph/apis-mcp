---
title: Restore Zaraz historical configuration by ID
page_id: operation-put-zones-zone-id-settings-zaraz-history-842e158d
path: operations/zaraz
description: Restores a historical published Zaraz configuration by ID for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/history
operation_ids:
    - put-zones-zone_identifier-zaraz-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Restore Zaraz historical configuration by ID

`PUT /zones/{zone_id}/settings/zaraz/history`

Operation ID: `put-zones-zone_identifier-zaraz-history`

Restores a historical published Zaraz configuration by ID for a zone.

## Definition

```yaml
{"operationId": "put-zones-zone_identifier-zaraz-history", "summary": "Restore Zaraz historical configuration by ID", "description": "Restores a historical published Zaraz configuration by ID for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"description": "ID of the Zaraz configuration to restore.", "type": "integer", "example": 12345, "minimum": 1}}}}, "responses": {"200": {"description": "Restore Zaraz historical configuration by ID response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-response"}}}}, "4XX": {"description": "Restore Zaraz historical configuration by ID failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Admin"]}
```
