---
title: Update a saved event query
page_id: operation-patch-accounts-account-id-cloudforce-one-events-queries-query-id-e56284cd
path: operations/event
description: Update an existing saved event query by its ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/queries/{query_id}
operation_ids:
    - patch_EventQueryUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a saved event query

`PATCH /accounts/{account_id}/cloudforce-one/events/queries/{query_id}`

Operation ID: `patch_EventQueryUpdate`

Update an existing saved event query by its ID

## Definition

```yaml
{"operationId": "patch_EventQueryUpdate", "summary": "Update a saved event query", "description": "Update an existing saved event query by its ID", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "query_id", "in": "path", "description": "Event query ID", "required": true, "schema": {"description": "Event query ID", "type": "integer"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"alert_enabled": {"description": "Enable alerts for this query", "type": "boolean"}, "alert_rollup_enabled": {"description": "Enable alert rollup for this query", "type": "boolean"}, "name": {"description": "Unique name for the saved query", "type": "string"}, "query_json": {"description": "JSON string containing the query parameters", "type": "string"}, "rule_enabled": {"description": "Enable rule for this query", "type": "boolean"}, "rule_scope": {"description": "Scope for the rule", "type": "string"}}}}}}, "responses": {"200": {"description": "Returns the updated event query.", "content": {"application/json": {"schema": {"type": "object", "properties": {"account_id": {"description": "Account ID", "type": "integer"}, "alert_enabled": {"description": "Whether alerts are enabled", "type": "boolean"}, "alert_rollup_enabled": {"description": "Whether alert rollup is enabled", "type": "boolean"}, "created_at": {"description": "Creation timestamp", "type": "string"}, "custom_threat_feed_id": {"description": "Intel Indicator Feed ID (numeric)", "type": "integer", "nullable": true}, "id": {"description": "Unique identifier for the saved query", "type": "integer"}, "name": {"description": "Name of the saved query", "type": "string"}, "query_json": {"description": "JSON string containing the query parameters", "type": "string"}, "rule_enabled": {"description": "Whether rule is enabled", "type": "boolean"}, "rule_list_id": {"description": "WAF rules list ID for blocking", "type": "string"}, "rule_scope": {"description": "Scope for the rule", "type": "string"}, "updated_at": {"description": "Last update timestamp", "type": "string"}, "user_email": {"description": "Email of the user who created the query", "type": "string"}}, "required": ["id", "account_id", "name", "user_email", "query_json", "alert_enabled", "alert_rollup_enabled", "rule_enabled", "created_at", "updated_at"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
