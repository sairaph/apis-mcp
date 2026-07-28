---
title: stream_webhook_request
page_id: schema-stream-webhook-request-33e52b28
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_webhook_request

```yaml
{"type": "object", "properties": {"notificationUrl": {"$ref": "#/components/schemas/stream_notificationUrl"}, "notification_url": {"description": "The URL where webhooks will be sent.", "type": "string", "format": "uri", "example": "https://example.com", "x-auditable": true}}, "anyOf": [{"required": ["notificationUrl"]}, {"required": ["notification_url"]}]}
```
