---
title: List remediation jobs
page_id: operation-get-accounts-account-id-data-security-posture-remediations-jobs-d9936590
path: operations/remediations
description: List all remediation jobs tied to a specific Cloudflare Account. Note that `cursor` and `page` are mutually exclusive.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/remediations/jobs
operation_ids:
    - ListRemediationJobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List remediation jobs

`GET /accounts/{account_id}/data-security/posture/remediations/jobs`

Operation ID: `ListRemediationJobs`

List all remediation jobs tied to a specific Cloudflare Account. Note that `cursor` and `page` are mutually exclusive.

## Definition

```yaml
{"operationId": "ListRemediationJobs", "summary": "List remediation jobs", "description": "List all remediation jobs tied to a specific Cloudflare Account. Note that `cursor` and `page` are mutually exclusive.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"name": "cursor", "in": "query", "description": "A cursor for pagination.", "schema": {"type": "string"}}, {"$ref": "#/components/parameters/posture-api_Page"}, {"$ref": "#/components/parameters/posture-api_PerPage"}, {"$ref": "#/components/parameters/posture-api_Search"}, {"name": "min_updated_at", "in": "query", "description": "Filter to view remediations updated on or after the min updated datetime. Can be a date-time in ISO 8601 format or an epoch timestamp.", "schema": {"type": "string", "format": "date-time"}}, {"name": "max_updated_at", "in": "query", "description": "Filter to view remediations updated on or before the max updated datetime. Can be a date-time in ISO 8601 format or an epoch timestamp.", "schema": {"type": "string", "format": "date-time"}}, {"name": "status", "in": "query", "description": "Filter to view remediations with the given status.", "schema": {"$ref": "#/components/schemas/posture-api_RemediationJobStatusEnum"}}, {"name": "triggered_by_actor", "in": "query", "description": "Filter remediations by what kind of actor triggered them. Supports multiple comma-separated values.", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJobActorTypeEnum"}}, "example": ["user", "account_token"], "explode": false, "style": "form"}, {"$ref": "#/components/parameters/posture-api_IntegrationId"}, {"name": "order", "in": "query", "description": "An optional param to sort the results by the given field.", "schema": {"type": "string", "enum": ["created_at", "affliction_date", "integration_name", "status", "last_updated_at", "asset_name", "finding_type_name"]}}, {"$ref": "#/components/parameters/posture-api_Direction"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_PaginatedRemediationJobList"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}}, "security": [{"api_token": []}], "tags": ["remediations"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
