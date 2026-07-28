---
title: Retrieves Zone Security Center Insight Counts by Severity
page_id: operation-get-zones-zone-id-security-center-insights-severity-9d1e93e3
path: operations/security-center-insights
description: Retrieves zone-specific Security Center insight counts aggregated by severity level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/security-center/insights/severity
operation_ids:
    - get-zone-security-center-insight-counts-by-severity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Zone Security Center Insight Counts by Severity

`GET /zones/{zone_id}/security-center/insights/severity`

Operation ID: `get-zone-security-center-insight-counts-by-severity`

Retrieves zone-specific Security Center insight counts aggregated by severity level.

## Definition

```yaml
{"operationId": "get-zone-security-center-insight-counts-by-severity", "summary": "Retrieves Zone Security Center Insight Counts by Severity", "description": "Retrieves zone-specific Security Center insight counts aggregated by severity level.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "dismissed", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_dismissed"}}, {"name": "issue_class", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueClasses"}}, {"name": "issue_type", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueTypes"}}, {"name": "product", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_products"}}, {"name": "severity", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_severityQueryParam"}}, {"name": "subject", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_subjects"}}, {"name": "issue_class~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueClasses"}}, {"name": "issue_type~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueTypes"}}, {"name": "product~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_products"}}, {"name": "severity~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_severityQueryParam"}}, {"name": "subject~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_subjects"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_valueCountsResponse"}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
