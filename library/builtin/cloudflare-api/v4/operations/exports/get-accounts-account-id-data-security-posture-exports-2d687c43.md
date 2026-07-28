---
title: List all export jobs
page_id: operation-get-accounts-account-id-data-security-posture-exports-c1504f7a
path: operations/exports
description: List all export jobs for a given requestor's organization
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/exports
operation_ids:
    - ListExportJobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all export jobs

`GET /accounts/{account_id}/data-security/posture/exports`

Operation ID: `ListExportJobs`

List all export jobs for a given requestor's organization

## Definition

```yaml
{"operationId": "ListExportJobs", "summary": "List all export jobs", "description": "List all export jobs for a given requestor's organization", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}, {"name": "status", "in": "query", "description": "Filter on export job's status", "schema": {"$ref": "#/components/schemas/posture-api_StatusEnum"}}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_PaginatedExportJobList"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["exports"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
