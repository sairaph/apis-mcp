---
title: Retrieves Zone Security Center Insights
page_id: operation-get-zones-zone-id-security-center-insights-3ffe9fbd
path: operations/security-center-insights
description: Lists Security Center insights for a specific zone, showing zone-level security findings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/security-center/insights
operation_ids:
    - get-zone-security-center-insights
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves Zone Security Center Insights

`GET /zones/{zone_id}/security-center/insights`

Operation ID: `get-zone-security-center-insights`

Lists Security Center insights for a specific zone, showing zone-level security findings.

## Definition

```yaml
{"operationId": "get-zone-security-center-insights", "summary": "Retrieves Zone Security Center Insights", "description": "Lists Security Center insights for a specific zone, showing zone-level security findings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}, {"name": "dismissed", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_dismissed"}}, {"name": "issue_class", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueClasses"}}, {"name": "issue_type", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueTypes"}}, {"name": "product", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_products"}}, {"name": "severity", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_severityQueryParam"}}, {"name": "subject", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_subjects"}}, {"name": "issue_class~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueClasses"}}, {"name": "issue_type~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_issueTypes"}}, {"name": "product~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_products"}}, {"name": "severity~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_severityQueryParam"}}, {"name": "subject~neq", "in": "query", "schema": {"$ref": "#/components/schemas/security-center_subjects"}}, {"name": "page", "in": "query", "schema": {"default": 1, "allOf": [{"$ref": "#/components/schemas/security-center_page"}]}}, {"name": "per_page", "in": "query", "schema": {"default": 25, "allOf": [{"$ref": "#/components/schemas/security-center_perPage"}]}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"anyOf": [{"properties": {"count": {"$ref": "#/components/schemas/security-center_count"}, "issues": {"type": "array", "items": {"$ref": "#/components/schemas/security-center_issue"}}, "page": {"$ref": "#/components/schemas/security-center_page"}, "per_page": {"$ref": "#/components/schemas/security-center_perPage"}}, "type": "object"}]}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Insights"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"]}
```
