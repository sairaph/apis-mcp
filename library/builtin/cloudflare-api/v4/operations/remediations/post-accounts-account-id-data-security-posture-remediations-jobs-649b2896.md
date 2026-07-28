---
title: Creates remediation jobs
page_id: operation-post-accounts-account-id-data-security-posture-remediations-jobs-d1723abc
path: operations/remediations
description: Create one or more remediation jobs tied to a specific Cloudflare Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/remediations/jobs
operation_ids:
    - CreateRemediationJobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates remediation jobs

`POST /accounts/{account_id}/data-security/posture/remediations/jobs`

Operation ID: `CreateRemediationJobs`

Create one or more remediation jobs tied to a specific Cloudflare Account.

## Definition

```yaml
{"operationId": "CreateRemediationJobs", "summary": "Creates remediation jobs", "description": "Create one or more remediation jobs tied to a specific Cloudflare Account.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_RemediationJobsCreateRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_CreateRemediationJobResponse"}}}}, "400": {"description": "Bad Request: Invalid request parameters", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_CreateRemediationJobResponse"}}}}}, "security": [{"api_token": []}], "tags": ["remediations"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
