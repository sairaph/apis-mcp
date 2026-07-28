---
title: List messages for a bulk action job
page_id: operation-get-accounts-account-id-email-security-investigate-bulk-job-id-messages-32f8c68f
path: operations/email-security
description: Returns the individual messages associated with a bulk action job, including their processing status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/bulk/{job_id}/messages
operation_ids:
    - email_security_get_bulk_job_messages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List messages for a bulk action job

`GET /accounts/{account_id}/email-security/investigate/bulk/{job_id}/messages`

Operation ID: `email_security_get_bulk_job_messages`

Returns the individual messages associated with a bulk action job, including their processing status.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_bulk_job_messages", "summary": "List messages for a bulk action job", "description": "Returns the individual messages associated with a bulk action job, including their processing status.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"name": "status", "in": "query", "schema": {"type": "string", "enum": ["PENDING", "DISCOVERING", "PROCESSING", "COMPLETED", "FAILED", "CANCELLED", "SKIPPED"]}}], "responses": {"200": {"description": "Messages for the bulk job.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_BulkActionMessageDetail"}}, "result_info": {"$ref": "#/components/schemas/email-security_CursorResultInfo"}}, "required": ["result", "result_info"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
