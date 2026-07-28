---
title: Create new findings export request
page_id: operation-post-accounts-account-id-data-security-posture-findings-export-e634be53
path: operations/exports
description: Creates a CSV export for findings and accepts optional filters in the payload.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/export
operation_ids:
    - CreateFindingExportCSV
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new findings export request

`POST /accounts/{account_id}/data-security/posture/findings/export`

Operation ID: `CreateFindingExportCSV`

Creates a CSV export for findings and accepts optional filters in the payload.

## Definition

```yaml
{"operationId": "CreateFindingExportCSV", "summary": "Create new findings export request", "description": "Creates a CSV export for findings and accepts optional filters in the payload.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_FindingExportFilterRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_export-job-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["exports"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
