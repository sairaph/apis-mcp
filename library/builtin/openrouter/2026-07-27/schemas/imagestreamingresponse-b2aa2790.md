---
title: ImageStreamingResponse
page_id: schema-imagestreamingresponse-b2aa2790
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageStreamingResponse

```yaml
{"example": {"data": {"b64_json": "<base64-encoded-partial-image>", "partial_image_index": 0, "type": "image_generation.partial_image"}}, "properties": {"data": {"anyOf": [{"$ref": "#/components/schemas/ImageGenPartialImageEvent"}, {"$ref": "#/components/schemas/ImageGenTextChunkEvent"}, {"$ref": "#/components/schemas/ImageGenCompletedEvent"}, {"$ref": "#/components/schemas/ImageGenStreamErrorEvent"}]}}, "required": ["data"], "type": "object"}
```
