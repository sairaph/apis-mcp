---
title: stream_downloads
page_id: schema-stream-downloads-e858779d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_downloads

```yaml
{"type": "object", "properties": {"percentComplete": {"$ref": "#/components/schemas/stream_download_percent_complete"}, "status": {"$ref": "#/components/schemas/stream_download_status"}, "url": {"$ref": "#/components/schemas/stream_download_url"}}, "required": ["status", "percentComplete"]}
```
