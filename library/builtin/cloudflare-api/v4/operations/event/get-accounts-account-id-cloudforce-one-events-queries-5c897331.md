---
title: List all saved event queries
page_id: operation-get-accounts-account-id-cloudforce-one-events-queries-51d96a0a
path: operations/event
description: Retrieve all saved event queries for the account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/queries
operation_ids:
    - get_EventQueryList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all saved event queries

`GET /accounts/{account_id}/cloudforce-one/events/queries`

Operation ID: `get_EventQueryList`

Retrieve all saved event queries for the account

## Definition

```yaml
{"operationId": "get_EventQueryList", "summary": "List all saved event queries", "description": "Retrieve all saved event queries for the account", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns a list of event queries.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"account_id": {"description": "Account ID", "type": "integer"}, "alert_enabled": {"description": "Whether alerts are enabled", "type": "boolean"}, "alert_rollup_enabled": {"description": "Whether alert rollup is enabled", "type": "boolean"}, "created_at": {"description": "Creation timestamp", "type": "string"}, "custom_threat_feed_id": {"description": "Intel Indicator Feed ID (numeric)", "type": "integer", "nullable": true}, "id": {"description": "Unique identifier for the saved query", "type": "integer"}, "name": {"description": "Name of the saved query", "type": "string"}, "query_json": {"description": "JSON string containing the query parameters", "type": "string"}, "rule_enabled": {"description": "Whether rule is enabled", "type": "boolean"}, "rule_list_id": {"description": "WAF rules list ID for blocking", "type": "string"}, "rule_scope": {"description": "Scope for the rule", "type": "string"}, "updated_at": {"description": "Last update timestamp", "type": "string"}, "user_email": {"description": "Email of the user who created the query", "type": "string"}}, "required": ["id", "account_id", "name", "user_email", "query_json", "alert_enabled", "alert_rollup_enabled", "rule_enabled", "created_at", "updated_at"], "type": "object"}}}}}, "500": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
