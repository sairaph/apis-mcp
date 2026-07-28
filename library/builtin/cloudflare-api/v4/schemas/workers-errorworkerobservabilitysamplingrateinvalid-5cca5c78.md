---
title: workers_ErrorWorkerObservabilitySamplingRateInvalid
page_id: schema-workers-errorworkerobservabilitysamplingrateinvalid-5cca5c78
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerObservabilitySamplingRateInvalid

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that an observability sampling rate is invalid.", "type": "integer", "enum": [100308]}, "message": {"description": "Message explaining that sampling rates must be between 0 and 1 inclusive.", "type": "string"}}, "required": ["code", "message"]}
```
