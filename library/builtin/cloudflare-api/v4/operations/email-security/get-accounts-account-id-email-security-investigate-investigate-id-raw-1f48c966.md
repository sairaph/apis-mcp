---
title: Get raw email content
page_id: operation-get-accounts-account-id-email-security-investigate-investigate-id-raw-5fc82022
path: operations/email-security
description: Returns the raw eml of any non-benign message.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}/raw
operation_ids:
    - email_security_get_message_raw
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get raw email content

`GET /accounts/{account_id}/email-security/investigate/{investigate_id}/raw`

Operation ID: `email_security_get_message_raw`

Returns the raw eml of any non-benign message.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_message_raw", "summary": "Get raw email content", "description": "Returns the raw eml of any non-benign message.", "responses": {"200": {"description": "Raw email content.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_MessageRaw"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate.raw", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
