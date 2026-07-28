---
title: Layout Parsing
page_id: operation-post-paas-v4-layout-parsing-9b6c8b83
path: operations/untagged
description: Use the [GLM-OCR](/guides/vlm/glm-ocr) model to parse the layout of documents and images and extract text content. Support OCR recognition of images and PDF documents, returning detailed layout information and visualization results.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/layout_parsing
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# Layout Parsing

`POST /paas/v4/layout_parsing`

Use the [GLM-OCR](/guides/vlm/glm-ocr) model to parse the layout of documents and images and extract text content. Support OCR recognition of images and PDF documents, returning detailed layout information and visualization results.

## Definition

```yaml
{"summary": "Layout Parsing", "description": "Use the [GLM-OCR](/guides/vlm/glm-ocr) model to parse the layout of documents and images and extract text content. Support OCR recognition of images and PDF documents, returning detailed layout information and visualization results.", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/LayoutParsingRequest"}, "examples": {"Layout Parsing Example": {"value": {"model": "glm-ocr", "file": "https://cdn.bigmodel.cn/static/logo/introduction.png"}}}}}, "required": true}, "responses": {"200": {"description": "Business processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/LayoutParsingResponse"}}}}, "default": {"description": "Request failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
