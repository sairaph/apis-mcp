---
title: Publish Zaraz preview configuration
page_id: operation-post-zones-zone-id-settings-zaraz-publish-57b99dd5
path: operations/zaraz
description: Publish current Zaraz preview configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/settings/zaraz/publish
operation_ids:
    - post-zones-zone_identifier-zaraz-publish
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Publish Zaraz preview configuration

`POST /zones/{zone_id}/settings/zaraz/publish`

Operation ID: `post-zones-zone_identifier-zaraz-publish`

Publish current Zaraz preview configuration for a zone.

## Definition

```yaml
{"operationId": "post-zones-zone_identifier-zaraz-publish", "summary": "Publish Zaraz preview configuration", "description": "Publish current Zaraz preview configuration for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zaraz_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"description": "Zaraz configuration description.", "type": "string", "example": "Config with enabled ecommerce tracking"}}}}, "responses": {"200": {"description": "Update Zaraz workflow response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zaraz_api-response-common"}, {"properties": {"result": {"type": "string", "example": "Config has been published successfully"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Zaraz workflow response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zaraz_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zaraz"], "x-api-token-group": ["Zaraz Admin"], "x-cfPermissionsRequired": {"enum": ["#zaraz:publish"]}}
```
