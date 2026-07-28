---
title: List bulk action jobs
page_id: operation-get-accounts-account-id-email-security-investigate-bulk-79abfdb1
path: operations/email-security
description: Returns a paginated list of bulk action jobs for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/bulk
operation_ids:
    - email_security_get_bulk_jobs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List bulk action jobs

`GET /accounts/{account_id}/email-security/investigate/bulk`

Operation ID: `email_security_get_bulk_jobs`

Returns a paginated list of bulk action jobs for the account.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_get_bulk_jobs", "summary": "List bulk action jobs", "description": "Returns a paginated list of bulk action jobs for the account.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"name": "action_type", "in": "query", "schema": {"type": "string", "enum": ["MOVE", "RELEASE"]}}, {"name": "status", "in": "query", "schema": {"type": "string", "enum": ["PENDING", "DISCOVERING", "PROCESSING", "COMPLETED", "FAILED", "CANCELLED", "SKIPPED"]}}], "responses": {"200": {"description": "Jobs retrieved successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_BulkJobDetail"}}, "result_info": {"$ref": "#/components/schemas/email-security_CursorResultInfo"}}, "required": ["result", "result_info"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
