---
title: posture-api_Webhook
page_id: schema-posture-api-webhook-fca5cbd7
path: schemas
description: Webhook configuration for sending finding notifications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_Webhook

Webhook configuration for sending finding notifications.

```yaml
{"description": "Webhook configuration for sending finding notifications.", "type": "object", "properties": {"authentication_type": {"description": "Type of authentication used for the webhook.", "type": "string", "example": "Bearer Auth", "enum": ["Basic Auth", "None", "Bearer Auth", "Static Headers", "HMAC-Signing"]}, "created_at": {"description": "Timestamp when the webhook configuration was created.", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00Z"}, "destination_url": {"description": "Target URL for the webhook configuration. Where resulting data will be sent.", "type": "string", "format": "uri", "example": "https://example.com/webhook"}, "headers": {"description": "List of header keys configured for this webhook. Values are not included for security reasons.", "type": "array", "items": {"properties": {"key": {"description": "Header key name (lowercase).", "type": "string", "example": "authorization"}, "value": {"description": "Header value. This field is never returned in API responses for security reasons.", "type": "string", "example": "my-authorization-value", "writeOnly": true}}, "type": "object"}}, "id": {"description": "Unique identifier for the specific webhook configuration.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}, "label": {"description": "Account-specified display label for the webhook configuration.", "type": "string", "example": "Send to Gmail"}, "status": {"description": "Current status of the webhook configuration. If disabled, data cannot be sent through this configuration.", "type": "string", "example": "enabled", "enum": ["enabled", "disabled"]}, "updated_at": {"description": "Timestamp when the webhook configuration was last updated.", "type": "string", "format": "date-time", "example": "2024-01-20T14:45:00Z"}, "version": {"description": "Version number of the configuration.", "type": "integer", "format": "uint32", "example": 1}}, "required": ["id", "label", "destination_url", "status", "authentication_type", "created_at", "updated_at", "version"]}
```
