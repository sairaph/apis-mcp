---
title: Retrieves Zone Audit Log
page_id: operation-get-zones-zone-id-security-center-insights-audit-log-fb9bb3be
path: operations/security-center-audit-log
description: Lists audit log entries for all Security Center insights in a specific zone, showing changes to insight status and classification.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/security-center/insights/audit-log
operation_ids:
    - get-zone-security-center-audit-log
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Zone Audit Log

`GET /zones/{zone_id}/security-center/insights/audit-log`

Operation ID: `get-zone-security-center-audit-log`

Lists audit log entries for all Security Center insights in a specific zone, showing changes to insight status and classification.

## Definition

```yaml
{"operationId": "get-zone-security-center-audit-log", "summary": "Retrieves Zone Audit Log", "description": "Lists audit log entries for all Security Center insights in a specific zone, showing changes to insight status and classification.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"$ref": "#/components/parameters/security-center_auditLogPerPage"}, {"$ref": "#/components/parameters/security-center_auditLogCursor"}, {"$ref": "#/components/parameters/security-center_auditLogFieldChanged"}, {"$ref": "#/components/parameters/security-center_auditLogChangedBy"}, {"$ref": "#/components/parameters/security-center_auditLogSince"}, {"$ref": "#/components/parameters/security-center_auditLogBefore"}, {"$ref": "#/components/parameters/security-center_auditLogOrder"}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_auditLogResponse"}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Audit Log"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
