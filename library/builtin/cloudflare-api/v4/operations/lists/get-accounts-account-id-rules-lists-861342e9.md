---
title: Get lists
page_id: operation-get-accounts-account-id-rules-lists-01bce3dd
path: operations/lists
description: Fetches all lists in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rules/lists
operation_ids:
    - lists-get-lists
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get lists

`GET /accounts/{account_id}/rules/lists`

Operation ID: `lists-get-lists`

Fetches all lists in the account.

## Definition

```yaml
{"operationId": "lists-get-lists", "summary": "Get lists", "description": "Fetches all lists in the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "responses": {"200": {"description": "Get lists response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_lists-response-collection"}}}}, "4XX": {"description": "Get lists response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_lists-response-collection"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Read", "Account Filter Lists Edit"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
