---
title: InputImage
page_id: schema-inputimage-0fec6478
path: schemas
description: Image input content item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InputImage

Image input content item

```yaml
{"description": "Image input content item", "example": {"detail": "auto", "image_url": "https://example.com/image.jpg", "type": "input_image"}, "properties": {"detail": {"enum": ["auto", "high", "low", "original"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "image_url": {"type": ["string", "null"]}, "type": {"enum": ["input_image"], "type": "string"}}, "required": ["type", "detail"], "type": "object"}
```
