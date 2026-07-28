---
title: realtimekit_Meeting
page_id: schema-realtimekit-meeting-97fe9c18
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_Meeting

```yaml
{"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string", "format": "date-time", "readOnly": true}, "id": {"description": "ID of the meeting.", "type": "string", "format": "uuid", "readOnly": true}, "live_stream_on_start": {"description": "Specifies if the meeting should start getting livestreamed on start.", "type": "boolean"}, "persist_chat": {"description": "Specifies if Chat within a meeting should persist for a week.", "type": "boolean"}, "record_on_start": {"description": "Specifies if the meeting should start getting recorded as soon as someone joins the meeting.", "type": "boolean"}, "recording_config": {"$ref": "#/components/schemas/realtimekit_RecordingConfig"}, "session_keep_alive_time_in_secs": {"description": "Time in seconds, for which a session remains active, after the last participant has left the meeting.", "type": "number", "default": 60, "maximum": 600, "minimum": 60}, "status": {"description": "Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an `INACTIVE` meeting.", "type": "string", "enum": ["ACTIVE", "INACTIVE"]}, "summarize_on_end": {"description": "Automatically generate summary of meetings using transcripts. Requires Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.", "type": "boolean"}, "title": {"description": "Title of the meeting.", "type": "string"}, "transcribe_on_end": {"description": "Automatically generate transcripts when the meeting ends.", "type": "boolean"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string", "format": "date-time", "readOnly": true}}, "required": ["id", "created_at", "updated_at"]}
```
