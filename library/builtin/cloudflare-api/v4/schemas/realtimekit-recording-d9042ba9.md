---
title: realtimekit_Recording
page_id: schema-realtimekit-recording-d9042ba9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_Recording

```yaml
{"type": "object", "properties": {"audio_download_url": {"description": "If the audio_config is passed, the URL for downloading the audio recording is returned.", "type": "string", "format": "uri", "nullable": true, "readOnly": true}, "download_url": {"description": "URL where the recording can be downloaded.", "type": "string", "format": "uri", "nullable": true, "readOnly": true}, "download_url_expiry": {"description": "Timestamp when the download URL expires.", "type": "string", "format": "date-time", "nullable": true, "readOnly": true}, "file_size": {"description": "File size of the recording, in bytes.", "type": "number", "nullable": true, "readOnly": true}, "id": {"description": "ID of the recording", "type": "string", "format": "uuid", "readOnly": true}, "invoked_time": {"description": "Timestamp when this recording was invoked.", "type": "string", "format": "date-time"}, "output_file_name": {"description": "File name of the recording.", "type": "string"}, "recording_duration": {"description": "Total recording time in seconds.", "type": "integer"}, "session_id": {"description": "ID of the meeting session this recording is for.", "type": "string", "format": "uuid", "nullable": true, "readOnly": true}, "started_time": {"description": "Timestamp when this recording actually started after being invoked. Usually a few seconds after `invoked_time`.", "type": "string", "format": "date-time", "nullable": true}, "status": {"description": "Current status of the recording.", "type": "string", "enum": ["INVOKED", "RECORDING", "UPLOADING", "UPLOADED", "ERRORED", "PAUSED"]}, "stopped_time": {"description": "Timestamp when this recording was stopped. Optional; is present only when the recording has actually been stopped.", "type": "string", "format": "date-time", "nullable": true}}, "required": ["id", "download_url", "download_url_expiry", "audio_download_url", "file_size", "session_id", "output_file_name", "status", "invoked_time", "started_time", "stopped_time"], "title": "Recording"}
```
