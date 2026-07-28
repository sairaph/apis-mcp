---
title: Create a preview token
page_id: operation-post-zones-zone-identifier-custom-pages-preview-tokens-59889303
path: operations/custom-pages-for-a-zone
description: Creates a signed JWT token used to preview custom pages before they are published. The API gateway rewrites zone-scoped requests to the account-level service endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/preview_tokens
operation_ids:
    - custom-pages-for-a-zone-create-preview-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a preview token

`POST /zones/{zone_identifier}/custom_pages/preview_tokens`

Operation ID: `custom-pages-for-a-zone-create-preview-token`

Creates a signed JWT token used to preview custom pages before they are published. The API gateway rewrites zone-scoped requests to the account-level service endpoint.

## Definition

```yaml
{"operationId": "custom-pages-for-a-zone-create-preview-token", "summary": "Create a preview token", "description": "Creates a signed JWT token used to preview custom pages before they are published. The API gateway rewrites zone-scoped requests to the account-level service endpoint.", "parameters": [{"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_preview_request"}}}}, "responses": {"200": {"description": "Create preview token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_preview_token_result"}}}}, "4XX": {"description": "Create preview token response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_preview_token_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom pages for a zone"], "x-api-token-group": ["Custom Pages Write", "Custom Pages Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
