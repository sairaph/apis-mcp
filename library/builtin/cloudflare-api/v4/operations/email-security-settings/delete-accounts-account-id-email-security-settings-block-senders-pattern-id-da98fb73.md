---
title: Delete a blocked email sender
page_id: operation-delete-accounts-account-id-email-security-settings-block-senders-pattern-a09a56cb
path: operations/email-security-settings
description: Removes a blocked sender pattern. After deletion, emails from this sender will no longer be automatically blocked based on this rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}
operation_ids:
    - email_security_delete_blocked_sender
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a blocked email sender

`DELETE /accounts/{account_id}/email-security/settings/block_senders/{pattern_id}`

Operation ID: `email_security_delete_blocked_sender`

Removes a blocked sender pattern. After deletion, emails from this sender will no longer be automatically blocked based on this rule.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "pattern_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_BlockedSenderId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_blocked_sender", "summary": "Delete a blocked email sender", "description": "Removes a blocked sender pattern. After deletion, emails from this sender will no longer be automatically blocked based on this rule.", "responses": {"200": {"description": "Deleted blocked sender.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedBlockedSender"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.block-senders", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
