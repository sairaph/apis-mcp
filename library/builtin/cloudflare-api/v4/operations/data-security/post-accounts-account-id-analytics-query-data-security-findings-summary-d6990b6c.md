---
title: Data security findings summary
page_id: operation-post-accounts-account-id-analytics-query-data-security-findings-summary-699fb229
path: operations/data-security
description: Returns aggregate current-period and previous-period totals for CASB findings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics/query/data-security/findings/summary
operation_ids:
    - data-security-findings-summary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Data security findings summary

`POST /accounts/{account_id}/analytics/query/data-security/findings/summary`

Operation ID: `data-security-findings-summary`

Returns aggregate current-period and previous-period totals for CASB findings.

## Definition

```yaml
{"operationId": "data-security-findings-summary", "summary": "Data security findings summary", "description": "Returns aggregate current-period and previous-period totals for CASB findings.\n", "parameters": [{"$ref": "#/components/parameters/art_account_id"}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"findings_summary": {"summary": "Summary for a week", "value": {"filters": [], "from": "2024-11-01T00:00:00Z", "to": "2024-11-08T00:00:00Z"}}}, "schema": {"$ref": "#/components/schemas/art_DataSecurityFindingsSummaryQuery"}}}}, "responses": {"200": {"description": "Findings summary result.", "content": {"application/json": {"examples": {"success": {"summary": "Successful findings summary", "value": {"errors": [], "messages": [{"code": 1000, "message": "API in beta: expect breaking changes."}], "result": {"currentTotal": [{"findingProduct": "Cloud", "findingType": "Content", "findingsTotal": 48291}, {"findingProduct": "SaaS", "findingType": "Posture", "findingsTotal": 1205}], "previousTotal": [{"findingProduct": "Cloud", "findingType": "Content", "findingsTotal": 41033}, {"findingProduct": "SaaS", "findingType": "Posture", "findingsTotal": 982}]}, "success": true}}}, "schema": {"$ref": "#/components/schemas/art_SummaryResponse"}}}}, "400": {"$ref": "#/components/responses/art_BadRequest"}, "401": {"$ref": "#/components/responses/art_Unauthorized"}, "403": {"$ref": "#/components/responses/art_Forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Data Security"], "x-api-token-group": ["Zero Trust Read"]}
```
