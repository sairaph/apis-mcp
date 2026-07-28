---
title: ImageGenerationResponse
page_id: schema-imagegenerationresponse-a25e6aa7
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ImageGenerationResponse

```yaml
{"type": "object", "properties": {"created": {"type": "integer", "example": 1760335349, "description": "Request creation time, in `Unix` timestamp format, unit is seconds."}, "data": {"type": "array", "description": "Array, containing the generated image `URL`. Currently, the array only contains one image.", "items": {"type": "object", "properties": {"url": {"type": "string", "description": "Image link. The temporary link expires after `30` days, please store it promptly."}}, "required": ["url"]}}, "content_filter": {"type": "array", "description": "Array, containing content safety related information.", "items": {"type": "object", "properties": {"role": {"type": "string", "description": "Safety enforcement stage, including `role = assistant` model inference, `role = user` user input, `role = history` historical context.", "enum": ["assistant", "user", "history"]}, "level": {"type": "integer", "description": "Severity level `level 0-3`, `level 0` is most severe, `3` is least severe.", "minimum": 0, "maximum": 3}}}}}}
```
