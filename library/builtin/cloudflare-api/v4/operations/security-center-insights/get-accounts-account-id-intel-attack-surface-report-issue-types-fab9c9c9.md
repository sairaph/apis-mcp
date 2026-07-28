---
title: Retrieves Security Center Issues Types
page_id: operation-get-accounts-account-id-intel-attack-surface-report-issue-types-75bf8fbb
path: operations/security-center-insights
description: Lists all available issue types in Security Center, describing categories of security issues.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/attack-surface-report/issue-types
operation_ids:
    - get-security-center-issue-types
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Security Center Issues Types

`GET /accounts/{account_id}/intel/attack-surface-report/issue-types`

Operation ID: `get-security-center-issue-types`

Lists all available issue types in Security Center, describing categories of security issues.

## Definition

```yaml
{"operationId": "get-security-center-issue-types", "summary": "Retrieves Security Center Issues Types", "description": "Lists all available issue types in Security Center, describing categories of security issues.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"anyOf": [{"items": {"type": "string"}, "type": "array"}]}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
