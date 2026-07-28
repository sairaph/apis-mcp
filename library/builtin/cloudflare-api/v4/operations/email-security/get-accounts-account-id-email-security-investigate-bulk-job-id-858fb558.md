---
title: Get bulk action job details
page_id: operation-get-accounts-account-id-email-security-investigate-bulk-job-id-79de22ef
path: operations/email-security
description: Returns the status and details of a specific bulk action job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/bulk/{job_id}
operation_ids:
    - email_security_get_bulk_job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get bulk action job details

`GET /accounts/{account_id}/email-security/investigate/bulk/{job_id}`

Operation ID: `email_security_get_bulk_job`

Returns the status and details of a specific bulk action job.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "job_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_bulk_job", "summary": "Get bulk action job details", "description": "Returns the status and details of a specific bulk action job.", "responses": {"200": {"description": "Job details retrieved.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BulkJobDetail"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
