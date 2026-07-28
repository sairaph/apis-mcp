---
title: LayoutParsingRequest
page_id: schema-layoutparsingrequest-8af7935e
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# LayoutParsingRequest

```yaml
{"type": "object", "required": ["model", "file"], "properties": {"model": {"type": "string", "description": "Model code: `glm-ocr`", "example": "glm-ocr", "enum": ["glm-ocr"]}, "file": {"type": "string", "description": "Image or PDF document to be recognized, supports URL and base64. Supported image formats: PDF, JPG, PNG. Single image ≤10MB, PDF ≤50MB, maximum support 30 pages", "example": "https://cdn.bigmodel.cn/static/logo/introduction.png"}, "return_crop_images": {"type": "boolean", "description": "Whether to return screenshot information", "default": false}, "need_layout_visualization": {"type": "boolean", "description": "Whether to return detailed layout image result information", "default": false}, "start_page_id": {"type": "integer", "description": "Start page number for parsing when PDF is provided", "minimum": 1}, "end_page_id": {"type": "integer", "description": "End page number for parsing when PDF is provided", "minimum": 1}, "request_id": {"type": "string", "description": "Passed by the user side, needs to be unique; used to distinguish each request, 6–64 characters. If not provided by the user side, the platform will generate one by default.", "minLength": 6, "maxLength": 64}, "user_id": {"type": "string", "description": "Unique ID for the end user, 6–128 characters. Avoid using sensitive information.", "minLength": 6, "maxLength": 128}}}
```
