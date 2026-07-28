---
title: Get account audit logs
page_id: operation-get-accounts-account-id-audit-logs-6d6d9bef
path: operations/audit-logs
description: Gets a list of audit logs for an account. Can be filtered by who made the change, on which zone, and the timeframe of the change.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/audit_logs
operation_ids:
    - audit-logs-get-account-audit-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account audit logs

`GET /accounts/{account_id}/audit_logs`

Operation ID: `audit-logs-get-account-audit-logs`

Gets a list of audit logs for an account. Can be filtered by who made the change, on which zone, and the timeframe of the change.

## Definition

```yaml
{"operationId": "audit-logs-get-account-audit-logs", "summary": "Get account audit logs", "description": "Gets a list of audit logs for an account. Can be filtered by who made the change, on which zone, and the timeframe of the change.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_identifier"}}, {"name": "id", "in": "query", "schema": {"description": "Finds a specific log by its ID.", "type": "string", "example": "f174be97-19b1-40d6-954d-70cd5fbd52db"}}, {"name": "export", "in": "query", "schema": {"description": "Indicates that this request is an export of logs in CSV format.", "type": "boolean", "example": true}}, {"name": "action.type", "in": "query", "schema": {"description": "Filters by the action type.", "type": "string", "example": "add"}}, {"name": "actor.ip", "in": "query", "schema": {"description": "Filters by the IP address of the request that made the change by specific IP address or valid CIDR Range.", "type": "string", "example": "17.168.228.63"}}, {"name": "actor.email", "in": "query", "schema": {"description": "Filters by the email address of the actor that made the change.", "type": "string", "format": "email", "example": "alice@example.com"}}, {"name": "since", "in": "query", "schema": {"oneOf": [{"description": "Limits the returned results to logs newer than the specified date. A `full-date` that conforms to RFC3339.", "example": "2019-04-30", "format": "date", "type": "string"}, {"description": "Limits the returned results to logs newer than the specified date. A `date-time` that conforms to RFC3339.", "example": "2019-04-30T01:12:20Z", "format": "date-time", "type": "string"}]}}, {"name": "before", "in": "query", "schema": {"oneOf": [{"description": "Limits the returned results to logs older than the specified date. A `full-date` that conforms to RFC3339.", "example": "2019-04-30", "format": "date", "type": "string"}, {"description": "Limits the returned results to logs older than the specified date. A `date-time` that conforms to RFC3339.", "example": "2019-04-30T01:12:20Z", "format": "date-time", "type": "string"}]}}, {"name": "zone.name", "in": "query", "schema": {"description": "Filters by the name of the zone associated to the change.", "type": "string", "example": "example.com"}}, {"name": "direction", "in": "query", "schema": {"description": "Changes the direction of the chronological sorting.", "type": "string", "example": "desc", "default": "desc", "enum": ["desc", "asc"]}}, {"name": "per_page", "in": "query", "schema": {"description": "Sets the number of results to return per page.", "type": "number", "example": 25, "default": 100, "maximum": 1000, "minimum": 1}}, {"name": "page", "in": "query", "schema": {"description": "Defines which page of results to return.", "type": "number", "example": 50, "default": 1, "minimum": 1}}, {"name": "hide_user_logs", "in": "query", "schema": {"description": "Indicates whether or not to hide user level audit logs.", "type": "boolean", "default": false}}], "responses": {"200": {"description": "Get account audit logs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_audit_logs_response_collection"}}}}, "4XX": {"description": "Get account audit logs response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_audit_logs_response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Audit Logs"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
