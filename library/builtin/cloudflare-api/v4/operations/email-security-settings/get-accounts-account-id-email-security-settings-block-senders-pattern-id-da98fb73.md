---
title: Get a blocked email sender
page_id: operation-get-accounts-account-id-email-security-settings-block-senders-pattern-id-cec97e75
path: operations/email-security-settings
description: Retrieves details for a specific blocked sender pattern including its pattern type, value, and metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}
operation_ids:
    - email_security_get_blocked_sender
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a blocked email sender

`GET /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}`

Operation ID: `email_security_get_blocked_sender`

Retrieves details for a specific blocked sender pattern including its pattern type, value, and metadata.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "pattern_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_BlockedSenderId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_blocked_sender", "summary": "Get a blocked email sender", "description": "Retrieves details for a specific blocked sender pattern including its pattern type, value, and metadata.", "responses": {"200": {"description": "Blocked sender details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BlockedSender"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
