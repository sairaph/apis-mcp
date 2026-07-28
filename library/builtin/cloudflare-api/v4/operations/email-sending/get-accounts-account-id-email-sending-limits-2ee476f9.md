---
title: Get sending limits
page_id: operation-get-accounts-account-id-email-sending-limits-5810e0d6
path: operations/email-sending
description: Returns the current daily sending quota for the account and, when a quota is resolved, the account's current usage against it. Quota is null when not yet available; usage is null when there is no resolved quota or usage is temporarily unavailable.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/sending/limits
operation_ids:
    - email-sending-get-sending-limits
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get sending limits

`GET /accounts/{account_id}/email/sending/limits`

Operation ID: `email-sending-get-sending-limits`

Returns the current daily sending quota for the account and, when a quota is resolved, the account's current usage against it. Quota is null when not yet available; usage is null when there is no resolved quota or usage is temporarily unavailable.

## Definition

```yaml
{"operationId": "email-sending-get-sending-limits", "summary": "Get sending limits", "description": "Returns the current daily sending quota for the account and, when a quota is resolved, the account's current usage against it. Quota is null when not yet available; usage is null when there is no resolved quota or usage is temporarily unavailable.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get sending limits response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_sending_limits_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Sending"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-sending.limits", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
