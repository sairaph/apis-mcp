---
title: Retrieves Security Center Insight Context
page_id: operation-get-accounts-account-id-security-center-insights-issue-id-context-928bf0ec
path: operations/security-center-insights
description: Returns the full context payload for an insight. This endpoint is used for insights with large payloads that are not included inline in the list response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/security-center/insights/{issue_id}/context
operation_ids:
    - get-security-center-insight-context
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Security Center Insight Context

`GET /accounts/{account_id}/security-center/insights/{issue_id}/context`

Operation ID: `get-security-center-insight-context`

Returns the full context payload for an insight. This endpoint is used for insights with large payloads that are not included inline in the list response.

## Definition

```yaml
{"operationId": "get-security-center-insight-context", "summary": "Retrieves Security Center Insight Context", "description": "Returns the full context payload for an insight. This endpoint is used for insights with large payloads that are not included inline in the list response.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "issue_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"type": "object", "additionalProperties": true}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"]}
```
