---
title: Archives Security Center Insight
page_id: operation-put-accounts-account-id-intel-attack-surface-report-issue-id-dismiss-afaa13b9
path: operations/security-center-insights
description: Deprecated endpoint for archiving Security Center insights. Use the newer archive-security-center-insight endpoint instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/intel/attack-surface-report/{issue_id}/dismiss
operation_ids:
    - archive-security-center-insight-deprecated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Archives Security Center Insight

`PUT /accounts/{account_id}/intel/attack-surface-report/{issue_id}/dismiss`

Operation ID: `archive-security-center-insight-deprecated`

Deprecated endpoint for archiving Security Center insights. Use the newer archive-security-center-insight endpoint instead.

## Definition

```yaml
{"operationId": "archive-security-center-insight-deprecated", "summary": "Archives Security Center Insight", "description": "Deprecated endpoint for archiving Security Center insights. Use the newer archive-security-center-insight endpoint instead.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "issue_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"dismiss": {"type": "boolean", "default": true, "x-auditable": true}}}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-single"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"]}
```
