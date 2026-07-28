---
title: stream_live_input_response_collection
page_id: schema-stream-live-input-response-collection-97e5b75c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_live_input_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"liveInputs": {"type": "array", "items": {"$ref": "#/components/schemas/stream_live_input_object_without_url"}}, "range": {"description": "The total number of remaining live inputs based on cursor position.", "type": "integer", "example": 1000}, "total": {"description": "The total number of live inputs that match the provided filters.", "type": "integer", "example": 35586}}}}, "type": "object"}]}
```
