---
title: Delete a bulk action job
page_id: operation-delete-accounts-account-id-email-security-investigate-bulk-job-id-0cc6af5f
path: operations/email-security
description: Deletes the job, removing it from all list and detail endpoints. Only jobs in a terminal state (`COMPLETED`, `CANCELLED`, `FAILED`, or `SKIPPED`) can be deleted. To stop an in-progress job without removing it, use the cancel endpoint instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/bulk/{job_id}
operation_ids:
    - email_security_delete_bulk_job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a bulk action job

`DELETE /accounts/{account_id}/email-security/investigate/bulk/{job_id}`

Operation ID: `email_security_delete_bulk_job`

Deletes the job, removing it from all list and detail endpoints. Only jobs in a terminal state (`COMPLETED`, `CANCELLED`, `FAILED`, or `SKIPPED`) can be deleted. To stop an in-progress job without removing it, use the cancel endpoint instead.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_bulk_job", "summary": "Delete a bulk action job", "description": "Deletes the job, removing it from all list and detail endpoints. Only jobs in a terminal state (`COMPLETED`, `CANCELLED`, `FAILED`, or `SKIPPED`) can be deleted. To stop an in-progress job without removing it, use the cancel endpoint instead.", "responses": {"200": {"description": "Job deleted.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"id": {"type": "string", "format": "uuid"}}, "required": ["id"]}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-stability": "beta"}
```
