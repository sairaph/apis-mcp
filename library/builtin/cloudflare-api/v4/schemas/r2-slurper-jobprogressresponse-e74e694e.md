---
title: r2-slurper_JobProgressResponse
page_id: schema-r2-slurper-jobprogressresponse-e74e694e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_JobProgressResponse

```yaml
{"type": "object", "properties": {"createdAt": {"type": "string"}, "failedObjects": {"type": "integer"}, "id": {"type": "string"}, "objects": {"type": "integer"}, "skippedObjects": {"type": "integer"}, "status": {"$ref": "#/components/schemas/r2-slurper_JobStatus"}, "transferredObjects": {"type": "integer"}}}
```
