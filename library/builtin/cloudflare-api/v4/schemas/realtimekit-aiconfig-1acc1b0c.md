---
title: realtimekit_AIConfig
page_id: schema-realtimekit-aiconfig-1acc1b0c
path: schemas
description: The AI Config allows you to customize the behavior of meeting transcriptions and summaries
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_AIConfig

The AI Config allows you to customize the behavior of meeting transcriptions and summaries

```yaml
{"description": "The AI Config allows you to customize the behavior of meeting transcriptions and summaries", "type": "object", "properties": {"summarization": {"$ref": "#/components/schemas/realtimekit_SummarizationConfig"}, "transcription": {"$ref": "#/components/schemas/realtimekit_TranscriptionConfig"}}, "title": "AIConfig"}
```
