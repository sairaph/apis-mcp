---
title: Preview for non-detection messages
page_id: operation-post-accounts-account-id-email-security-investigate-preview-90126bf1
path: operations/email-security
description: Generates a preview image for a message that was not flagged as a detection. Useful for investigating benign messages. Returns a base64-encoded PNG screenshot of the email body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/preview
operation_ids:
    - email_security_post_preview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview for non-detection messages

`POST /accounts/{account_id}/email-security/investigate/preview`

Operation ID: `email_security_post_preview`

Generates a preview image for a message that was not flagged as a detection. Useful for investigating benign messages. Returns a base64-encoded PNG screenshot of the email body.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_post_preview", "summary": "Preview for non-detection messages", "description": "Generates a preview image for a message that was not flagged as a detection. Useful for investigating benign messages. Returns a base64-encoded PNG screenshot of the email body.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"postfix_id": {"$ref": "#/components/schemas/email-security_PostfixId"}}, "required": ["postfix_id"]}}}}, "responses": {"200": {"description": "Email preview.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_MessagePreview"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
