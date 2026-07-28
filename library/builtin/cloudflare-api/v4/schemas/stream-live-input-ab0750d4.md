---
title: stream_live_input
page_id: schema-stream-live-input-ab0750d4
path: schemas
description: Details about a live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_live_input

Details about a live input.

```yaml
{"description": "Details about a live input.", "type": "object", "properties": {"created": {"$ref": "#/components/schemas/stream_live_input_created"}, "deleteRecordingAfterDays": {"$ref": "#/components/schemas/stream_live_input_recording_deletion"}, "enabled": {"$ref": "#/components/schemas/stream_live_input_enabled"}, "keysRotatedAt": {"$ref": "#/components/schemas/stream_live_input_keys_rotated_at"}, "meta": {"$ref": "#/components/schemas/stream_live_input_metadata"}, "modified": {"$ref": "#/components/schemas/stream_live_input_modified"}, "preferLowLatency": {"$ref": "#/components/schemas/stream_live_input_prefer_low_latency"}, "recording": {"$ref": "#/components/schemas/stream_live_input_recording_settings"}, "rtmps": {"$ref": "#/components/schemas/stream_input_rtmps"}, "rtmpsPlayback": {"$ref": "#/components/schemas/stream_playback_rtmps"}, "srt": {"$ref": "#/components/schemas/stream_input_srt"}, "srtPlayback": {"$ref": "#/components/schemas/stream_playback_srt"}, "status": {"$ref": "#/components/schemas/stream_live_input_status"}, "uid": {"$ref": "#/components/schemas/stream_live_input_identifier"}, "webRTC": {"$ref": "#/components/schemas/stream_input_webrtc"}, "webRTCPlayback": {"$ref": "#/components/schemas/stream_playback_webrtc"}}}
```
