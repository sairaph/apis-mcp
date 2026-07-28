---
title: Data security findings timeseries
page_id: operation-post-accounts-account-id-analytics-query-data-security-findings-timeseri-05a03e90
path: operations/data-security
description: Returns merged time-bucketed CASB findings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics/query/data-security/findings/timeseries
operation_ids:
    - data-security-findings-timeseries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Data security findings timeseries

`POST /accounts/{account_id}/analytics/query/data-security/findings/timeseries`

Operation ID: `data-security-findings-timeseries`

Returns merged time-bucketed CASB findings.

## Definition

```yaml
{"operationId": "data-security-findings-timeseries", "summary": "Data security findings timeseries", "description": "Returns merged time-bucketed CASB findings.\n", "parameters": [{"$ref": "#/components/parameters/art_account_id"}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"findings_timeseries": {"summary": "Weekly findings timeseries", "value": {"filters": [], "from": "2024-11-01T00:00:00Z", "to": "2024-11-08T00:00:00Z"}}}, "schema": {"$ref": "#/components/schemas/art_DataSecurityFindingsTimeseriesQuery"}}}}, "responses": {"200": {"description": "Findings timeseries result.", "content": {"application/json": {"examples": {"success": {"summary": "Successful findings timeseries", "value": {"errors": [], "messages": [{"code": 1000, "message": "API in beta: expect breaking changes."}], "result": {"slots": [{"content": {"cloud": 150, "saas": 23}, "posture": {"cloud": 0, "saas": 5}, "timestamp": "2024-11-05T00:00:00Z"}, {"content": {"cloud": 180, "saas": 30}, "posture": {"cloud": 0, "saas": 7}, "timestamp": "2024-11-06T00:00:00Z"}]}, "success": true}}}, "schema": {"$ref": "#/components/schemas/art_DataSecurityFindingsTimeseriesResponse"}}}}, "400": {"$ref": "#/components/responses/art_BadRequest"}, "401": {"$ref": "#/components/responses/art_Unauthorized"}, "403": {"$ref": "#/components/responses/art_Forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Data Security"], "x-api-token-group": ["Zero Trust Read"]}
```
