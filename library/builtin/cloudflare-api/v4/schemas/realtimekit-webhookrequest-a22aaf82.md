---
title: realtimekit_WebhookRequest
page_id: schema-realtimekit-webhookrequest-a22aaf82
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_WebhookRequest

```yaml
{"type": "object", "properties": {"enabled": {"description": "Set whether or not the webhook should be active when created", "type": "boolean", "default": true}, "events": {"description": "Events that this webhook will get triggered by", "type": "array", "items": {"enum": ["meeting.started", "meeting.ended", "meeting.participantJoined", "meeting.participantLeft", "meeting.chatSynced", "recording.statusUpdate", "livestreaming.statusUpdate", "meeting.transcript", "meeting.summary"], "type": "string"}, "example": ["meeting.started", "meeting.ended", "meeting.participantJoined", "meeting.participantLeft", "meeting.chatSynced", "recording.statusUpdate", "livestreaming.statusUpdate", "meeting.transcript", "meeting.summary"]}, "name": {"description": "Name of the webhook", "type": "string", "example": "All events webhook"}, "url": {"description": "URL this webhook will send events to", "type": "string", "format": "uri", "example": "https://webhook.site/b23a5bbd-c7b0-4ced-a9e2-78ae7889897e"}}, "required": ["name", "url", "events"]}
```
