---
title: Archives Security Center Insight
page_id: operation-put-accounts-account-id-security-center-insights-issue-id-dismiss-962b993f
path: operations/security-center-insights
description: Archives a Security Center insight for an account, removing it from the active insights list while preserving historical data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/security-center/insights/{issue_id}/dismiss
operation_ids:
    - archive-security-center-insight
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Archives Security Center Insight

`PUT /accounts/{account_id}/security-center/insights/{issue_id}/dismiss`

Operation ID: `archive-security-center-insight`

Archives a Security Center insight for an account, removing it from the active insights list while preserving historical data.

## Definition

```yaml
{"operationId": "archive-security-center-insight", "summary": "Archives Security Center Insight", "description": "Archives a Security Center insight for an account, removing it from the active insights list while preserving historical data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "issue_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"dismiss": {"type": "boolean", "default": true, "x-auditable": true}}}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-single"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"]}
```
