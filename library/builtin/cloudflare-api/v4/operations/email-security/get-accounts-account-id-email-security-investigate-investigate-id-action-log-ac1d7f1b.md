---
title: Get action log for a message
page_id: operation-get-accounts-account-id-email-security-investigate-investigate-id-action-9e426f09
path: operations/email-security
description: Returns the list of post-delivery actions (moves, quarantine releases, previews, etc.) that have been applied to a specific email message.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}/action_log
operation_ids:
    - email_security_get_message_action_log
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get action log for a message

`GET /accounts/{account_id}/email-security/investigate/{investigate_id}/action_log`

Operation ID: `email_security_get_message_action_log`

Returns the list of post-delivery actions (moves, quarantine releases, previews, etc.) that have been applied to a specific email message.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_message_action_log", "summary": "Get action log for a message", "description": "Returns the list of post-delivery actions (moves, quarantine releases, previews, etc.) that have been applied to a specific email message.", "responses": {"200": {"description": "Action log for the email.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_ActionEntry"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate.get-action", "x-fern-sdk-method-name": "log", "x-forge-hidden": true, "x-stability": "beta"}
```
