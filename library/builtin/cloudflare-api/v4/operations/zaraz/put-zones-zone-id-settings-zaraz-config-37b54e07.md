---
title: Update Zaraz configuration
page_id: operation-put-zones-zone-id-settings-zaraz-config-f17153b1
path: operations/zaraz
description: Updates Zaraz configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/config
operation_ids:
    - put-zones-zone_identifier-zaraz-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zaraz configuration

`PUT /zones/{zone_id}/settings/zaraz/config`

Operation ID: `put-zones-zone_identifier-zaraz-config`

Updates Zaraz configuration for a zone.

## Definition

```yaml
{"operationId": "put-zones-zone_identifier-zaraz-config", "summary": "Update Zaraz configuration", "description": "Updates Zaraz configuration for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-body"}}}}, "responses": {"200": {"description": "Update Zaraz configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-config-response"}}}}, "4XX": {"description": "Update Zaraz configuration response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Admin"]}
```
