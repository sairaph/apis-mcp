---
title: Retrieves Issue Audit Log
page_id: operation-get-accounts-account-id-security-center-insights-issue-id-audit-log-83d0c2ea
path: operations/security-center-audit-log
description: Lists audit log entries for a specific Security Center insight, showing changes to its status and classification over time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/security-center/insights/{issue_id}/audit-log
operation_ids:
    - get-security-center-issue-audit-log
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Issue Audit Log

`GET /accounts/{account_id}/security-center/insights/{issue_id}/audit-log`

Operation ID: `get-security-center-issue-audit-log`

Lists audit log entries for a specific Security Center insight, showing changes to its status and classification over time.

## Definition

```yaml
{"operationId": "get-security-center-issue-audit-log", "summary": "Retrieves Issue Audit Log", "description": "Lists audit log entries for a specific Security Center insight, showing changes to its status and classification over time.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "issue_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"$ref": "#/components/parameters/security-center_auditLogPerPage"}, {"$ref": "#/components/parameters/security-center_auditLogCursor"}, {"$ref": "#/components/parameters/security-center_auditLogFieldChanged"}, {"$ref": "#/components/parameters/security-center_auditLogChangedBy"}, {"$ref": "#/components/parameters/security-center_auditLogSince"}, {"$ref": "#/components/parameters/security-center_auditLogBefore"}, {"$ref": "#/components/parameters/security-center_auditLogOrder"}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_auditLogResponse"}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Audit Log"]}
```
