---
title: stream_storage_use_response
page_id: schema-stream-storage-use-response-4ff06af5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_storage_use_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"creator": {"$ref": "#/components/schemas/stream_creator"}, "totalStorageMinutes": {"description": "The total minutes of video content stored in the account. May contain decimal values.", "type": "number", "format": "float"}, "totalStorageMinutesLimit": {"description": "The storage capacity alloted for the account.", "type": "integer"}, "videoCount": {"description": "The total count of videos associated with the account.", "type": "integer"}}}}, "type": "object"}]}
```
