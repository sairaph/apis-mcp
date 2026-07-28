---
title: Create a remediation jobs export
page_id: operation-post-accounts-account-id-data-security-posture-remediations-jobs-export-7d0d88d4
path: operations/exports
description: Creates a CSV export for remediation jobs and accepts optional filters in the payload.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/remediations/jobs/export
operation_ids:
    - CreateRemediationJobsExportCSV
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a remediation jobs export

`POST /accounts/{account_id}/data-security/posture/remediations/jobs/export`

Operation ID: `CreateRemediationJobsExportCSV`

Creates a CSV export for remediation jobs and accepts optional filters in the payload.

## Definition

```yaml
{"operationId": "CreateRemediationJobsExportCSV", "summary": "Create a remediation jobs export", "description": "Creates a CSV export for remediation jobs and accepts optional filters in the payload.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_RemediationJobsExportFilterRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_export-job-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["exports"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
