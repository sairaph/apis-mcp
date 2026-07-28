---
title: stream_webhook_response_single
page_id: schema-stream-webhook-response-single-f91f9163
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_webhook_response_single

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"modified": {"description": "The date and time the webhook was last modified.", "type": "string", "format": "date-time", "example": "2014-01-02T02:20:00Z"}, "notificationUrl": {"description": "The URL where webhooks will be sent.", "type": "string", "format": "uri", "example": "https://example.com"}, "notification_url": {"description": "The URL where webhooks will be sent.", "type": "string", "format": "uri", "example": "https://example.com"}, "secret": {"description": "The secret used to verify webhook signatures.", "type": "string", "x-sensitive": true}}}}, "type": "object"}]}
```
