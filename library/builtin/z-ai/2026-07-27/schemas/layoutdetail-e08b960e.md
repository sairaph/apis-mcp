---
title: LayoutDetail
page_id: schema-layoutdetail-e08b960e
path: schemas
description: Layout detail element
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# LayoutDetail

Layout detail element

```yaml
{"type": "object", "description": "Layout detail element", "properties": {"index": {"type": "integer", "description": "Element index", "example": 1}, "label": {"type": "string", "description": "Element type: image for images, text for text content, formula for inline formulas, table for tables", "enum": ["image", "text", "formula", "table"], "example": "text"}, "bbox_2d": {"type": "array", "description": "Normalized element coordinates [x1,y1,x2,y2]", "items": {"type": "number", "minimum": 0, "maximum": 1}, "minItems": 4, "maxItems": 4, "example": [0.1, 0.1, 0.5, 0.3]}, "content": {"type": "string", "description": "Element content (text / image URL / table HTML)", "example": "This is the content of the element"}, "height": {"type": "integer", "description": "Page height", "example": 800}, "width": {"type": "integer", "description": "Page width", "example": 600}}, "required": ["index", "label"]}
```
