---
title: Get email trace
page_id: operation-get-accounts-account-id-email-security-investigate-investigate-id-trace-43d42948
path: operations/email-security
description: Retrieves delivery and processing trace information for an email message. Shows the delivery path, retraction history, and move operations performed on the message. Useful for debugging delivery issues.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}/trace
operation_ids:
    - email_security_get_message_trace
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get email trace

`GET /accounts/{account_id}/email-security/investigate/{investigate_id}/trace`

Operation ID: `email_security_get_message_trace`

Retrieves delivery and processing trace information for an email message. Shows the delivery path, retraction history, and move operations performed on the message. Useful for debugging delivery issues.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_message_trace", "summary": "Get email trace", "description": "Retrieves delivery and processing trace information for an email message. Shows the delivery path, retraction history, and move operations performed on the message. Useful for debugging delivery issues.", "responses": {"200": {"description": "Email trace.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_MessageTrace"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate.trace", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
