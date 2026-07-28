---
title: Prediction
page_id: schema-prediction-6f2ddd1d
path: schemas
description: Static predicted output content. Supported models can use this to reduce latency when much of the response is known in advance.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Prediction

Static predicted output content. Supported models can use this to reduce latency when much of the response is known in advance.

```yaml
{"description": "Static predicted output content. Supported models can use this to reduce latency when much of the response is known in advance.", "example": {"content": "Expected response", "type": "content"}, "properties": {"content": {"anyOf": [{"type": "string"}, {"items": {"$ref": "#/components/schemas/PredictionContentText"}, "type": "array"}]}, "type": {"enum": ["content"], "type": "string"}}, "required": ["type", "content"], "type": ["object", "null"]}
```
