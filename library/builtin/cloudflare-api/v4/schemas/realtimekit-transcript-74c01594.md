---
title: realtimekit_Transcript
page_id: schema-realtimekit-transcript-74c01594
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_Transcript

```yaml
{"type": "object", "properties": {"sessionId": {"type": "string"}, "transcript_download_url": {"description": "URL where the transcript can be downloaded", "type": "string"}, "transcript_download_url_expiry": {"description": "Time when the download URL will expire", "type": "string"}}, "required": ["sessionId", "transcript_download_url", "transcript_download_url_expiry"]}
```
