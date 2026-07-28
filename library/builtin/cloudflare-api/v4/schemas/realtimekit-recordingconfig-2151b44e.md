---
title: realtimekit_RecordingConfig
page_id: schema-realtimekit-recordingconfig-2151b44e
path: schemas
description: Recording Configurations to be used for this meeting. This level of configs takes higher preference over App level configs on the RealtimeKit developer portal.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_RecordingConfig

Recording Configurations to be used for this meeting. This level of configs takes higher preference over App level configs on the RealtimeKit developer portal.

```yaml
{"description": "Recording Configurations to be used for this meeting. This level of configs takes higher preference over App level configs on the RealtimeKit developer portal.\n", "type": "object", "properties": {"audio_config": {"$ref": "#/components/schemas/realtimekit_AudioConfig"}, "file_name_prefix": {"description": "Adds a prefix to the beginning of the file name of the recording.", "type": "string"}, "live_streaming_config": {"$ref": "#/components/schemas/realtimekit_LivestreamingConfig"}, "max_seconds": {"description": "Specifies the maximum duration for recording in seconds, ranging from a minimum of 60 seconds to a maximum of 24 hours.", "type": "number", "maximum": 86400, "minimum": 60}, "realtimekit_bucket_config": {"$ref": "#/components/schemas/realtimekit_realtimekitBucketConfig"}, "storage_config": {"$ref": "#/components/schemas/realtimekit_StorageConfig"}, "video_config": {"$ref": "#/components/schemas/realtimekit_VideoConfig"}}, "title": "RecordingConfig"}
```
