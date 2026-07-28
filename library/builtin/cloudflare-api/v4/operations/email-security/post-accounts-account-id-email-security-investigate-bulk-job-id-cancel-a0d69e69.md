---
title: Cancel a bulk action job
page_id: operation-post-accounts-account-id-email-security-investigate-bulk-job-id-cancel-0ffe3fc9
path: operations/email-security
description: Cancels the job, marking it as cancelled and stopping any pending message processing. The job record remains visible in list and detail endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel
operation_ids:
    - email_security_cancel_bulk_job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Cancel a bulk action job

`POST /accounts/{account_id}/email-security/investigate/bulk/{job_id}/cancel`

Operation ID: `email_security_cancel_bulk_job`

Cancels the job, marking it as cancelled and stopping any pending message processing. The job record remains visible in list and detail endpoints.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}]
```

## Definition

```yaml
{"operationId": "email_security_cancel_bulk_job", "summary": "Cancel a bulk action job", "description": "Cancels the job, marking it as cancelled and stopping any pending message processing. The job record remains visible in list and detail endpoints.", "responses": {"200": {"description": "Job cancelled.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BulkJobDetail"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
