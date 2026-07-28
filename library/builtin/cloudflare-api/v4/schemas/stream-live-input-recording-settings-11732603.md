---
title: stream_live_input_recording_settings
page_id: schema-stream-live-input-recording-settings-11732603
path: schemas
description: Records the input to a Cloudflare Stream video. Behavior depends on the mode. In most cases, the video will initially be viewable as a live video and transition to on-demand after a condition is satisfied.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_live_input_recording_settings

Records the input to a Cloudflare Stream video. Behavior depends on the mode. In most cases, the video will initially be viewable as a live video and transition to on-demand after a condition is satisfied.

```yaml
{"description": "Records the input to a Cloudflare Stream video. Behavior depends on the mode. In most cases, the video will initially be viewable as a live video and transition to on-demand after a condition is satisfied.", "type": "object", "properties": {"allowedOrigins": {"$ref": "#/components/schemas/stream_live_input_recording_allowedOrigins"}, "hideLiveViewerCount": {"$ref": "#/components/schemas/stream_live_input_recording_hideLiveViewerCount"}, "mode": {"$ref": "#/components/schemas/stream_live_input_recording_mode"}, "requireSignedURLs": {"$ref": "#/components/schemas/stream_live_input_recording_requireSignedURLs"}, "timeoutSeconds": {"$ref": "#/components/schemas/stream_live_input_recording_timeoutSeconds"}}, "example": {"hideLiveViewerCount": false, "mode": "off", "requireSignedURLs": false, "timeoutSeconds": 0}}
```
