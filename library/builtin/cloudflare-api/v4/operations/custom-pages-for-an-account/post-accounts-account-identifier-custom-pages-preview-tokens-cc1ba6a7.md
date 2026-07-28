---
title: Create a preview token
page_id: operation-post-accounts-account-identifier-custom-pages-preview-tokens-81793973
path: operations/custom-pages-for-an-account
description: Creates a signed JWT token used to preview custom pages before they are published.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_identifier}/custom_pages/preview_tokens
operation_ids:
    - custom-pages-for-an-account-create-preview-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a preview token

`POST /accounts/{account_identifier}/custom_pages/preview_tokens`

Operation ID: `custom-pages-for-an-account-create-preview-token`

Creates a signed JWT token used to preview custom pages before they are published.

## Definition

```yaml
{"operationId": "custom-pages-for-an-account-create-preview-token", "summary": "Create a preview token", "description": "Creates a signed JWT token used to preview custom pages before they are published.", "parameters": [{"name": "account_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_preview_request"}}}}, "responses": {"200": {"description": "Create preview token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_preview_token_result"}}}}, "4XX": {"description": "Create preview token response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_preview_token_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom pages for an account"], "x-api-token-group": ["Zero Trust: PII Read", "Account Custom Pages Write", "Account Custom Pages Read", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
