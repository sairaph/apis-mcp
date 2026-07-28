---
title: Release messages from quarantine
page_id: operation-post-accounts-account-id-email-security-investigate-release-b5a0ca31
path: operations/email-security
description: Delivers one or more quarantined messages to their intended recipients, for cases where a message was incorrectly quarantined. The response includes delivery status for each recipient.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/release
operation_ids:
    - email_security_post_release
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Release messages from quarantine

`POST /accounts/{account_id}/email-security/investigate/release`

Operation ID: `email_security_post_release`

Delivers one or more quarantined messages to their intended recipients, for cases where a message was incorrectly quarantined. The response includes delivery status for each recipient.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_post_release", "summary": "Release messages from quarantine", "description": "Delivers one or more quarantined messages to their intended recipients, for cases where a message was incorrectly quarantined. The response includes delivery status for each recipient.", "requestBody": {"description": "A list of investigate IDs identifying the messages to release.", "required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_InvestigateId"}}}}}, "responses": {"200": {"description": "Release operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_ReleaseResponse"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate.release", "x-fern-sdk-method-name": "bulk", "x-forge-hidden": true, "x-stability": "beta"}
```
