---
title: r2-slurper_JobLogResponse
page_id: schema-r2-slurper-joblogresponse-3b8d7672
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_JobLogResponse

```yaml
{"type": "object", "properties": {"createdAt": {"type": "string"}, "job": {"type": "string"}, "logType": {"type": "string", "enum": ["migrationStart", "migrationComplete", "migrationAbort", "migrationError", "migrationPause", "migrationResume", "migrationErrorFailedContinuation", "importErrorRetryExhaustion", "importSkippedStorageClass", "importSkippedOversized", "importSkippedEmptyObject", "importSkippedUnsupportedContentType", "importSkippedExcludedContentType", "importSkippedInvalidMedia", "importSkippedRequiresRetrieval"]}, "message": {"type": "string", "nullable": true}, "objectKey": {"type": "string", "nullable": true}}}
```
