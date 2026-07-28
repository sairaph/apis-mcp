---
title: realtimekit_PatchWebhookRequest
page_id: schema-realtimekit-patchwebhookrequest-cc6a7883
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PatchWebhookRequest

```yaml
{"type": "object", "properties": {"enabled": {"type": "boolean", "default": true}, "events": {"description": "Events that the webhook will get triggered by", "type": "array", "items": {"enum": ["meeting.started", "meeting.ended", "meeting.participantJoined", "meeting.participantLeft", "recording.statusUpdate", "livestreaming.statusUpdate", "meeting.chatSynced", "meeting.transcript", "meeting.summary"], "type": "string"}}, "name": {"description": "Name of the webhook", "type": "string"}, "url": {"description": "URL the webhook will send events to", "type": "string", "format": "uri", "example": "https://webhook.site/b23a5bbd-c7b0-4ced-a9e2-78ae7889897e"}}}
```
