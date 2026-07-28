---
title: List account routing rules
page_id: operation-get-accounts-account-id-email-routing-rules-c44fabb3
path: operations/email-routing-routing-rules
description: Lists existing routing rules across all zones in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/routing/rules
operation_ids:
    - email-routing-routing-rules-list-account-routing-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account routing rules

`GET /accounts/{account_id}/email/routing/rules`

Operation ID: `email-routing-routing-rules-list-account-routing-rules`

Lists existing routing rules across all zones in the account.

## Definition

```yaml
{"operationId": "email-routing-routing-rules-list-account-routing-rules", "summary": "List account routing rules", "description": "Lists existing routing rules across all zones in the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "enabled", "in": "query", "schema": {"description": "Filter by enabled routing rules.", "type": "boolean", "example": true, "enum": [true, false]}}], "responses": {"200": {"description": "List account routing rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_account_rules_response_collection"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing routing rules"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.rules", "x-fern-sdk-method-name": "list-account", "x-forge-hidden": true}
```
