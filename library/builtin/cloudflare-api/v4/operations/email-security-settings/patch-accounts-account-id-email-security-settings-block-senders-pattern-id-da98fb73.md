---
title: Update a blocked email sender
page_id: operation-patch-accounts-account-id-email-security-settings-block-senders-pattern-1b0a72cd
path: operations/email-security-settings
description: Updates an existing blocked sender pattern. Only provided fields will be modified. The pattern will continue blocking emails until deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}
operation_ids:
    - email_security_update_blocked_sender
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a blocked email sender

`PATCH /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}`

Operation ID: `email_security_update_blocked_sender`

Updates an existing blocked sender pattern. Only provided fields will be modified. The pattern will continue blocking emails until deleted.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "pattern_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_BlockedSenderId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_blocked_sender", "summary": "Update a blocked email sender", "description": "Updates an existing blocked sender pattern. Only provided fields will be modified. The pattern will continue blocking emails until deleted.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateBlockedSender"}}}}, "responses": {"200": {"description": "Updated blocked sender.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_BlockedSender"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.block-senders", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-stability": "beta"}
```
