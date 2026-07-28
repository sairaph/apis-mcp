---
title: realtimekit_WebhookEventsSuccessResponse
page_id: schema-realtimekit-webhookeventssuccessresponse-7582a67e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_WebhookEventsSuccessResponse

```yaml
{"type": "object", "properties": {"data": {"type": "array", "items": {"enum": ["meeting.started", "meeting.ended", "meeting.participantJoined", "meeting.participantLeft", "meeting.chatSynced", "recording.statusUpdate", "livestreaming.statusUpdate", "meeting.transcript", "meeting.summary"], "type": "string"}}, "success": {"type": "boolean", "example": true}}, "required": ["success", "data"]}
```
