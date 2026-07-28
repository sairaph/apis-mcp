---
title: Move a message
page_id: operation-post-accounts-account-id-email-security-investigate-investigate-id-move-b9fe6cb2
path: operations/email-security
description: Moves a single message to a specified mailbox folder (Inbox, JunkEmail, DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires active integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}/move
operation_ids:
    - email_security_post_message_move
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Move a message

`POST /accounts/{account_id}/email-security/investigate/{investigate_id}/move`

Operation ID: `email_security_post_message_move`

Moves a single message to a specified mailbox folder (Inbox, JunkEmail, DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires active integration.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}]
```

## Definition

```yaml
{"operationId": "email_security_post_message_move", "summary": "Move a message", "description": "Moves a single message to a specified mailbox folder (Inbox, JunkEmail, DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires active integration.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination": {"$ref": "#/components/schemas/email-security_MailboxDestination"}, "expected_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}}, "required": ["destination"]}}}}, "responses": {"200": {"description": "Move operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_MoveResponseItem"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
