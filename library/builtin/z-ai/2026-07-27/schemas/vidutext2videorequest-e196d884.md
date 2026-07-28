---
title: ViduText2VideoRequest
page_id: schema-vidutext2videorequest-e196d884
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ViduText2VideoRequest

```yaml
{"allOf": [{"type": "object", "properties": {"model": {"type": "string", "description": "The model code to be called.", "enum": ["viduq1-text"]}, "prompt": {"type": "string", "description": "Text description of the video, maximum input length of 512 characters."}, "style": {"type": "string", "description": "Style\nDefault: `general`\nOptional values: `general` , `anime`\n- `general`: General style, can be controlled using prompts to define the style.\n- `anime`: Anime style, optimized for anime-specific visuals. The style can be controlled using different anime-themed prompts.", "enum": ["general", "anime"]}, "duration": {"type": "integer", "description": "Video duration parameter.\nDefault: `5` , Optional: `5`.", "example": 5, "enum": [5]}, "aspect_ratio": {"type": "string", "description": "Aspect ratio\nDefault: `16:9`, Optional values: `16:9`, `9:16`, `1:1`", "example": "16:9", "enum": ["16:9", "9:16", "1:1"]}, "size": {"type": "string", "description": "Resolution parameter\nDefault: `1920x1080`, Optional: `1920x1080`", "example": "1920x1080", "enum": ["1920x1080"]}, "movement_amplitude": {"type": "string", "description": "Motion amplitude\nDefault: `auto` , Optional values:  `auto` ,`small` ,`medium` ,`large`", "example": "auto", "enum": ["auto", "small", "medium", "large"]}}, "required": ["model", "prompt"]}, {"$ref": "#/components/schemas/VideoCommonRequest"}]}
```
