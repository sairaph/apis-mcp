---
title: Move multiple messages
page_id: operation-post-accounts-account-id-email-security-investigate-move-66a3c729
path: operations/email-security
description: Moves multiple messages to a specified mailbox folder (Inbox, JunkEmail, DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires active integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/move
operation_ids:
    - email_security_post_bulk_move
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Move multiple messages

`POST /accounts/{account_id}/email-security/investigate/move`

Operation ID: `email_security_post_bulk_move`

Moves multiple messages to a specified mailbox folder (Inbox, JunkEmail, DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires active integration.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_post_bulk_move", "summary": "Move multiple messages", "description": "Moves multiple messages to a specified mailbox folder (Inbox, JunkEmail, DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires active integration.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination": {"$ref": "#/components/schemas/email-security_MailboxDestination"}, "expected_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "ids": {"description": "List of message IDs to move.", "type": "array", "items": {"$ref": "#/components/schemas/email-security_InvestigateId"}, "x-auditable": true}, "postfix_ids": {"description": "Deprecated, use `ids` instead. End of life: November 1, 2026.", "type": "array", "items": {"$ref": "#/components/schemas/email-security_PostfixId"}, "deprecated": true, "x-auditable": true, "x-stainless-deprecation-message": "Use `ids` instead."}}, "required": ["destination"]}}}}, "responses": {"200": {"description": "Bulk move operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_MoveResponseItem"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
