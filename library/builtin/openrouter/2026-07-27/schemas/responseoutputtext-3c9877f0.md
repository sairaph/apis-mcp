---
title: ResponseOutputText
page_id: schema-responseoutputtext-3c9877f0
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ResponseOutputText

```yaml
{"example": {"annotations": [{"end_index": 42, "start_index": 0, "title": "Paris - Wikipedia", "type": "url_citation", "url": "https://en.wikipedia.org/wiki/Paris"}], "text": "The capital of France is Paris.", "type": "output_text"}, "properties": {"annotations": {"items": {"$ref": "#/components/schemas/OpenAIResponsesAnnotation"}, "type": "array"}, "logprobs": {"items": {"properties": {"bytes": {"items": {"type": "integer"}, "type": "array"}, "logprob": {"format": "double", "type": "number"}, "token": {"type": "string"}, "top_logprobs": {"items": {"properties": {"bytes": {"items": {"type": "integer"}, "type": "array"}, "logprob": {"format": "double", "type": "number"}, "token": {"type": "string"}}, "required": ["token", "bytes", "logprob"], "type": "object"}, "type": "array"}}, "required": ["token", "bytes", "logprob", "top_logprobs"], "type": "object"}, "type": "array"}, "text": {"type": "string"}, "type": {"enum": ["output_text"], "type": "string"}}, "required": ["type", "text"], "type": "object"}
```
