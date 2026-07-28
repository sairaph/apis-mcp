---
title: realtimekit_TranscriptSummary
page_id: schema-realtimekit-transcriptsummary-5a8add82
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_TranscriptSummary

```yaml
{"type": "object", "properties": {"sessionId": {"type": "string"}, "summaryDownloadUrl": {"description": "URL where the summary of transcripts can be downloaded", "type": "string"}, "summaryDownloadUrlExpiry": {"description": "Time of Expiry before when you need to download the csv file.", "type": "string"}}, "required": ["sessionId", "summaryDownloadUrl", "summaryDownloadUrlExpiry"], "title": "TranscriptSummary"}
```
