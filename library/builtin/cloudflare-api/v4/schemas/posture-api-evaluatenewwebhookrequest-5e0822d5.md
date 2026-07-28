---
title: posture-api_EvaluateNewWebhookRequest
page_id: schema-posture-api-evaluatenewwebhookrequest-5e0822d5
path: schemas
description: Request body for testing a webhook configuration before creating it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_EvaluateNewWebhookRequest

Request body for testing a webhook configuration before creating it.

```yaml
{"description": "Request body for testing a webhook configuration before creating it.", "type": "object", "properties": {"authentication_type": {"description": "Type of authentication to use for the test webhook request.", "type": "string", "example": "Bearer Auth", "enum": ["Basic Auth", "None", "Bearer Auth", "Static Headers", "HMAC-Signing"]}, "destination_url": {"description": "Target URL to send the test webhook event to.", "type": "string", "format": "uri", "example": "https://example.com/webhook"}, "headers": {"description": "List of custom headers to include in the test webhook request.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_WebhookHeaderInput"}, "example": [{"key": "Authorization", "value": "Bearer token123"}, {"key": "X-Custom-Header", "value": "value"}]}, "signing_secret": {"description": "Secret key used for HMAC signing when authentication_type is \"HMAC-Signing\".", "type": "string", "example": "my-secret-key"}}, "required": ["destination_url", "authentication_type"]}
```
