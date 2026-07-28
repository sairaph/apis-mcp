---
title: realtimekit_Webhook
page_id: schema-realtimekit-webhook-eddf81d8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_Webhook

```yaml
{"type": "object", "properties": {"created_at": {"description": "Timestamp when this webhook was created", "type": "string", "format": "date-time", "example": "2022-05-28T07:01:53.075Z"}, "enabled": {"description": "Set to true if the webhook is active", "type": "boolean"}, "events": {"description": "Events this webhook will send updates for", "type": "array", "items": {"enum": ["meeting.started", "meeting.ended", "meeting.participantJoined", "meeting.participantLeft", "meeting.chatSynced", "recording.statusUpdate", "livestreaming.statusUpdate", "meeting.transcript", "meeting.summary"], "type": "string"}, "example": ["meeting.started", "meeting.ended", "meeting.participantJoined", "meeting.participantLeft", "meeting.chatSynced", "recording.statusUpdate", "livestreaming.statusUpdate", "meeting.transcript", "meeting.summary"]}, "id": {"description": "ID of the webhook", "type": "string", "format": "uuid", "example": "0d1f069d-43bb-489a-ad8c-7eb95592ba8e", "readOnly": true}, "name": {"description": "Name of the webhook", "type": "string", "example": "All events webhook"}, "updated_at": {"description": "Timestamp when this webhook was updated", "type": "string", "format": "date-time", "example": "2022-05-28T07:01:53.075Z"}, "url": {"description": "URL the webhook will send events to", "type": "string", "format": "uri", "example": "https://webhook.site/b23a5bbd-c7b0-4ced-a9e2-78ae7889897e"}}, "required": ["id", "name", "url", "events", "created_at", "updated_at", "enabled"]}
```
