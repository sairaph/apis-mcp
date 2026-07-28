---
title: posture-api_UpdateWebhookRequest
page_id: schema-posture-api-updatewebhookrequest-b9900cfe
path: schemas
description: Request body for updating an existing webhook configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_UpdateWebhookRequest

Request body for updating an existing webhook configuration.

```yaml
{"description": "Request body for updating an existing webhook configuration.", "type": "object", "properties": {"authentication_type": {"description": "Type of authentication used for the webhook.", "type": "string", "example": "Bearer Auth", "enum": ["Basic Auth", "None", "Bearer Auth", "Static Headers", "HMAC-Signing"]}, "destination_url": {"description": "Target URL for the webhook configuration. Where resulting data will be sent.", "type": "string", "format": "uri", "example": "https://example.com/webhook"}, "headers": {"description": "List of custom headers to include in webhook requests.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_WebhookHeaderInput"}, "example": [{"key": "Authorization", "value": "Bearer token123"}, {"key": "X-Custom-Header", "value": "value"}]}, "label": {"description": "Account-specified display label for the webhook configuration.", "type": "string", "example": "Send to Slack"}, "signing_secret": {"description": "Secret key used for HMAC signing when authentication_type is \"HMAC-Signing\".", "type": "string", "example": "my-secret-key"}, "status": {"description": "Status of the webhook configuration.", "type": "string", "example": "enabled", "enum": ["enabled", "disabled"]}}, "required": ["label", "destination_url", "authentication_type", "status"]}
```
