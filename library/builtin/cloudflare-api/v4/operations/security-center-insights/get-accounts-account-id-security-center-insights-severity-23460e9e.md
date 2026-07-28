---
title: Retrieves Security Center Insight Counts by Severity
page_id: operation-get-accounts-account-id-security-center-insights-severity-c70b82c7
path: operations/security-center-insights
description: Retrieves Security Center insight counts aggregated by severity level (critical, high, medium, low).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/security-center/insights/severity
operation_ids:
    - get-security-center-insight-counts-by-severity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Security Center Insight Counts by Severity

`GET /accounts/{account_id}/security-center/insights/severity`

Operation ID: `get-security-center-insight-counts-by-severity`

Retrieves Security Center insight counts aggregated by severity level (critical, high, medium, low).

## Definition

```yaml
{"operationId": "get-security-center-insight-counts-by-severity", "summary": "Retrieves Security Center Insight Counts by Severity", "description": "Retrieves Security Center insight counts aggregated by severity level (critical, high, medium, low).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "dismissed", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_dismissed"}}, {"name": "issue_class", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueClasses"}}, {"name": "issue_type", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueTypes"}}, {"name": "product", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_products"}}, {"name": "severity", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_severityQueryParam"}}, {"name": "subject", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_subjects"}}, {"name": "issue_class~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueClasses"}}, {"name": "issue_type~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueTypes"}}, {"name": "product~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_products"}}, {"name": "severity~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_severityQueryParam"}}, {"name": "subject~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_subjects"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_valueCountsResponse"}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"]}
```
