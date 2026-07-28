---
title: Legacy_ChatContentVideo
page_id: schema-legacy-chatcontentvideo-5297abc5
path: schemas
description: Video input content part (legacy format - deprecated)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Legacy_ChatContentVideo

Video input content part (legacy format - deprecated)

```yaml
{"deprecated": true, "description": "Video input content part (legacy format - deprecated)", "example": {"type": "input_video", "video_url": {"url": "https://example.com/video.mp4"}}, "properties": {"type": {"enum": ["input_video"], "type": "string"}, "video_url": {"$ref": "#/components/schemas/Legacy_ChatContentVideoInput"}}, "required": ["type", "video_url"], "type": "object"}
```
