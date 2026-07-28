---
title: Top integrations by content findings
page_id: operation-post-accounts-account-id-analytics-query-data-security-content-findings-cae03159
path: operations/data-security
description: Returns the top N integrations ranked by total content findings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics/query/data-security/content-findings/top-n
operation_ids:
    - data-security-content-findings-top-n
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Top integrations by content findings

`POST /accounts/{account_id}/analytics/query/data-security/content-findings/top-n`

Operation ID: `data-security-content-findings-top-n`

Returns the top N integrations ranked by total content findings.

## Definition

```yaml
{"operationId": "data-security-content-findings-top-n", "summary": "Top integrations by content findings", "description": "Returns the top N integrations ranked by total content findings.\n", "parameters": [{"$ref": "#/components/parameters/art_account_id"}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"top_integrations": {"summary": "Top 10 integrations by content findings", "value": {"filters": [], "from": "2024-11-01T00:00:00Z", "n": 10, "to": "2024-11-08T00:00:00Z"}}}, "schema": {"$ref": "#/components/schemas/art_DataSecurityContentFindingsTopNQuery"}}}}, "responses": {"200": {"description": "Top integrations result.", "content": {"application/json": {"examples": {"success": {"summary": "Successful top integrations result", "value": {"errors": [], "messages": [{"code": 1000, "message": "API in beta: expect breaking changes."}], "result": [{"integrationId": "123e4567-e89b-12d3-a456-426614174000", "integrationName": "Google Workspace", "total": 42}, {"integrationId": "223e4567-e89b-12d3-a456-426614174001", "integrationName": "Microsoft 365", "total": 17}], "success": true}}}, "schema": {"$ref": "#/components/schemas/art_TopNResponse"}}}}, "400": {"$ref": "#/components/responses/art_BadRequest"}, "401": {"$ref": "#/components/responses/art_Unauthorized"}, "403": {"$ref": "#/components/responses/art_Forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Data Security"], "x-api-token-group": ["Zero Trust Read"]}
```
