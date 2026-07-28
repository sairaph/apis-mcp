---
title: posture-api_CreateWebhookRequest
page_id: schema-posture-api-createwebhookrequest-e3c2660b
path: schemas
description: Request body for creating a new webhook configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_CreateWebhookRequest

Request body for creating a new webhook configuration.

```yaml
{"description": "Request body for creating a new webhook configuration.", "type": "object", "properties": {"authentication_type": {"description": "Type of authentication used for the webhook.", "type": "string", "example": "Bearer Auth", "enum": ["Basic Auth", "None", "Bearer Auth", "Static Headers", "HMAC-Signing"]}, "destination_url": {"description": "Target URL for the webhook configuration. Where resulting data will be sent.", "type": "string", "format": "uri", "example": "https://example.com/webhook"}, "headers": {"description": "List of custom headers to include in webhook requests.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_WebhookHeaderInput"}, "example": [{"key": "Authorization", "value": "Bearer token123"}, {"key": "X-Custom-Header", "value": "value"}]}, "label": {"description": "Account-specified display label for the webhook configuration.", "type": "string", "example": "Send to Slack"}, "signing_secret": {"description": "Secret key used for HMAC signing when authentication_type is \"HMAC-Signing\".", "type": "string", "example": "my-secret-key"}}, "required": ["label", "destination_url", "authentication_type"]}
```
