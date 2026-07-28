---
title: Create blocked email sender
page_id: operation-post-accounts-account-id-email-security-settings-block-senders-b9e073dc
path: operations/email-security-settings
description: Creates a new blocked sender pattern. Emails matching this pattern will be blocked from delivery. Patterns can be email addresses, domains, or IP addresses, and support regular expressions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/block_senders
operation_ids:
    - email_security_create_blocked_sender
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create blocked email sender

`POST /accounts/{account_id}/email-security/settings/block_senders`

Operation ID: `email_security_create_blocked_sender`

Creates a new blocked sender pattern. Emails matching this pattern will be blocked from delivery. Patterns can be email addresses, domains, or IP addresses, and support regular expressions.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_blocked_sender", "summary": "Create blocked email sender", "description": "Creates a new blocked sender pattern. Emails matching this pattern will be blocked from delivery. Patterns can be email addresses, domains, or IP addresses, and support regular expressions.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateBlockedSender"}}}}, "responses": {"201": {"description": "Created blocked sender.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BlockedSender"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
