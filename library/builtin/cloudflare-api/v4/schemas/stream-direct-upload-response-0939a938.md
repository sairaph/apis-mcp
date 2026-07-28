---
title: stream_direct_upload_response
page_id: schema-stream-direct-upload-response-0939a938
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_direct_upload_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"scheduledDeletion": {"$ref": "#/components/schemas/stream_scheduledDeletion"}, "uid": {"$ref": "#/components/schemas/stream_identifier"}, "uploadURL": {"description": "The URL an unauthenticated upload can use for a single `HTTP POST multipart/form-data` request.", "type": "string", "example": "www.example.com/samplepath"}, "watermark": {"$ref": "#/components/schemas/stream_watermarks"}}}}, "type": "object"}]}
```
