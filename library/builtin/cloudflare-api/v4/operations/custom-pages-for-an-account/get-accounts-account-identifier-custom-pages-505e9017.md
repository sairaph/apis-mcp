---
title: List custom pages
page_id: operation-get-accounts-account-identifier-custom-pages-b644dcfb
path: operations/custom-pages-for-an-account
description: Fetches all the custom pages at the account level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_identifier}/custom_pages
operation_ids:
    - custom-pages-for-an-account-list-custom-pages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List custom pages

`GET /accounts/{account_identifier}/custom_pages`

Operation ID: `custom-pages-for-an-account-list-custom-pages`

Fetches all the custom pages at the account level.

## Definition

```yaml
{"operationId": "custom-pages-for-an-account-list-custom-pages", "summary": "List custom pages", "description": "Fetches all the custom pages at the account level.", "parameters": [{"name": "account_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "responses": {"200": {"description": "List custom pages response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_page_result_list"}}}}, "4XX": {"description": "List custom pages response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_page_result_list"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom pages for an account"], "x-api-token-group": ["Zero Trust: PII Read", "Account Custom Pages Write", "Account Custom Pages Read", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
