---
title: stream_downloads_response
page_id: schema-stream-downloads-response-56c99879
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_downloads_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-single"}, {"properties": {"result": {"description": "An object with download type keys. Each key is optional and only present if that download type has been created.", "type": "object", "properties": {"audio": {"description": "The audio-only download. Only present if this download type has been created.", "allOf": [{"$ref": "#/components/schemas/stream_downloads"}]}, "default": {"description": "The default video download. Only present if this download type has been created.", "allOf": [{"$ref": "#/components/schemas/stream_downloads"}]}}}}, "type": "object"}]}
```
