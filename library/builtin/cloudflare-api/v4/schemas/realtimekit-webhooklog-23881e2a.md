---
title: realtimekit_WebhookLog
page_id: schema-realtimekit-webhooklog-23881e2a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_WebhookLog

```yaml
{"type": "object", "properties": {"event": {"description": "Webhook event name.", "type": "string"}, "headers": {"description": "Headers sent with the webhook delivery attempt.", "type": "object", "additionalProperties": {"type": "string"}}, "id": {"description": "ID of the webhook delivery log.", "type": "string"}, "payload": {"description": "Webhook delivery payload.", "type": "object", "additionalProperties": true}, "status_code": {"description": "HTTP status code returned by the webhook endpoint.", "type": "integer"}, "timestamp": {"description": "Timestamp for the webhook delivery attempt.", "type": "string", "format": "date-time"}}}
```
