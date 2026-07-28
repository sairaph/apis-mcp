---
title: ChatContentVideo
page_id: schema-chatcontentvideo-4405a130
path: schemas
description: Video input content part
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentVideo

Video input content part

```yaml
{"description": "Video input content part", "example": {"type": "video_url", "video_url": {"url": "https://example.com/video.mp4"}}, "properties": {"type": {"enum": ["video_url"], "type": "string"}, "video_url": {"$ref": "#/components/schemas/ChatContentVideoInput"}}, "required": ["type", "video_url"], "type": "object"}
```
