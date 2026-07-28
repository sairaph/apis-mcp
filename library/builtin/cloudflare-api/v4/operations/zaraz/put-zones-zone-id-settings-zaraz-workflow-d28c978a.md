---
title: Update Zaraz workflow
page_id: operation-put-zones-zone-id-settings-zaraz-workflow-1e1f02a8
path: operations/zaraz
description: Updates Zaraz workflow for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/workflow
operation_ids:
    - put-zones-zone_identifier-zaraz-workflow
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zaraz workflow

`PUT /zones/{zone_id}/settings/zaraz/workflow`

Operation ID: `put-zones-zone_identifier-zaraz-workflow`

Updates Zaraz workflow for a zone.

## Definition

```yaml
{"operationId": "put-zones-zone_identifier-zaraz-workflow", "summary": "Update Zaraz workflow", "description": "Updates Zaraz workflow for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-workflow"}}}}, "responses": {"200": {"description": "Update Zaraz workflow response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_zaraz-workflow-response"}}}}, "4XX": {"description": "Update Zaraz workflow response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Edit", "Zaraz Admin"]}
```
