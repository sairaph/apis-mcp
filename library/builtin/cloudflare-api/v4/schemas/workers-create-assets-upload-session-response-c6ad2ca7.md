---
title: workers_create-assets-upload-session-response
page_id: schema-workers-create-assets-upload-session-response-c6ad2ca7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_create-assets-upload-session-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"buckets": {"description": "The requests to make to upload assets.", "type": "array", "items": {"description": "The set of assets to include in each request while uploading.", "items": {"description": "The file hash to include in this bucket.", "type": "string"}, "type": "array", "x-stainless-collection-type": "set"}, "x-stainless-collection-type": "set"}, "jwt": {"description": "A JWT to use as authentication for uploading assets.", "type": "string", "x-sensitive": true}}}}, "type": "object"}]}
```
