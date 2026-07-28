---
title: Create a bulk action job
page_id: operation-post-accounts-account-id-email-security-investigate-bulk-f1580669
path: operations/email-security
description: Creates a new bulk action job to move or release messages that match the provided search parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/bulk
operation_ids:
    - email_security_create_bulk_job
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a bulk action job

`POST /accounts/{account_id}/email-security/investigate/bulk`

Operation ID: `email_security_create_bulk_job`

Creates a new bulk action job to move or release messages that match the provided search parameters.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_bulk_job", "summary": "Create a bulk action job", "description": "Creates a new bulk action job to move or release messages that match the provided search parameters.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_BulkActionRequest"}}}}, "responses": {"201": {"description": "Bulk job created.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BulkJobDetail"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
