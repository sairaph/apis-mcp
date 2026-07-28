---
title: Get list items
page_id: operation-get-accounts-account-id-rules-lists-list-id-items-9bc0e0ea
path: operations/lists
description: Fetches all the items in the list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}/items
operation_ids:
    - lists-get-list-items
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get list items

`GET /accounts/{account_id}/rules/lists/{list_id}/items`

Operation ID: `lists-get-list-items`

Fetches all the items in the list.

## Definition

```yaml
{"operationId": "lists-get-list-items", "summary": "Get list items", "description": "Fetches all the items in the list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}, {"name": "cursor", "in": "query", "schema": {"description": "The pagination cursor. An opaque string token indicating the position from which to continue when requesting the next/previous set of records. Cursor values are provided under `result_info.cursors` in the response. You should make no assumptions about a cursor's content or length.", "type": "string", "example": "zzz"}}, {"name": "per_page", "in": "query", "schema": {"description": "Amount of results to include in each paginated response. A non-negative 32 bit integer.", "type": "integer", "maximum": 500, "minimum": 1}}, {"name": "search", "in": "query", "schema": {"description": "A search query to filter returned items. Its meaning depends on the list type: IP addresses must start with the provided string, hostnames and bulk redirects must contain the string, and ASNs must match the string exactly.", "type": "string", "example": "1.1.1."}}], "responses": {"200": {"description": "Get list items response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_items-list-response-collection"}}}}, "4XX": {"description": "Get list items response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_items-list-response-collection"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit", "Account Filter Lists Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
