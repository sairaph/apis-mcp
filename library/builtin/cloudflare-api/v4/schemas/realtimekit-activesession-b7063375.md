---
title: realtimekit_ActiveSession
page_id: schema-realtimekit-activesession-b7063375
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_ActiveSession

```yaml
{"type": "object", "properties": {"associated_id": {"description": "ID of the meeting this session is associated with. In the case of V2 meetings, it is always a UUID. In V1 meetings, it is a room name of the form `abcdef-ghijkl`", "type": "string"}, "breakout_rooms": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_ActiveSession"}, "readOnly": true}, "created_at": {"description": "timestamp when session created", "type": "string"}, "ended_at": {"description": "timestamp when session ended", "type": "string"}, "id": {"description": "ID of the session", "type": "string", "readOnly": true}, "live_participants": {"description": "number of participants currently in the session", "type": "number"}, "max_concurrent_participants": {"description": "number of maximum participants that were in the session", "type": "number"}, "meeting_display_name": {"description": "Title of the meeting this session belongs to", "type": "string"}, "minutes_consumed": {"description": "number of minutes consumed since the session started", "type": "number"}, "organization_id": {"description": "App id that hosted this session", "type": "string"}, "started_at": {"description": "timestamp when session started", "type": "string"}, "status": {"description": "current status of session", "type": "string", "enum": ["LIVE", "ENDED"]}, "type": {"description": "type of session", "type": "string", "enum": ["meeting", "livestream", "participant"]}, "updated_at": {"description": "timestamp when session was last updated", "type": "string"}}, "required": ["id", "associated_id", "meeting_display_name", "type", "status", "live_participants", "max_concurrent_participants", "minutes_consumed", "organization_id", "started_at", "created_at", "updated_at"]}
```
