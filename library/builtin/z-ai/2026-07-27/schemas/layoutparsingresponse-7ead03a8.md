---
title: LayoutParsingResponse
page_id: schema-layoutparsingresponse-7ead03a8
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# LayoutParsingResponse

```yaml
{"type": "object", "properties": {"id": {"type": "string", "description": "Task ID", "example": "task_123456789"}, "created": {"type": "integer", "format": "int64", "description": "Request creation time, Unix timestamp in seconds", "example": 1727156815}, "model": {"type": "string", "description": "Model name", "example": "GLM-OCR"}, "md_results": {"type": "string", "description": "Recognition result in Markdown format", "example": "# Doc title\nThis is the document content..."}, "layout_details": {"type": "array", "description": "Detailed layout information", "items": {"type": "array", "items": {"$ref": "#/components/schemas/LayoutDetail"}}}, "layout_visualization": {"type": "array", "description": "Recognition result image URLs", "items": {"type": "string"}}, "data_info": {"$ref": "#/components/schemas/DataInfo"}, "usage": {"type": "object", "description": "Token usage statistics returned when the model call ends.", "properties": {"prompt_tokens": {"type": "number", "description": "Number of tokens in user input"}, "completion_tokens": {"type": "number", "description": "Number of output tokens"}, "prompt_tokens_details": {"type": "object", "properties": {"cached_tokens": {"type": "number", "description": "Number of tokens served from cache"}}}, "total_tokens": {"type": "integer", "description": "Total number of tokens"}}}, "request_id": {"type": "string", "description": "Request ID", "example": "req_123456789"}}, "required": ["id", "created", "model"]}
```
