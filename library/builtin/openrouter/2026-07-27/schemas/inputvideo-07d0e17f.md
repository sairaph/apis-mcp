---
title: InputVideo
page_id: schema-inputvideo-07d0e17f
path: schemas
description: Video input content item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InputVideo

Video input content item

```yaml
{"description": "Video input content item", "example": {"type": "input_video", "video_url": "https://example.com/video.mp4"}, "properties": {"type": {"enum": ["input_video"], "type": "string"}, "video_url": {"description": "A base64 data URL or remote URL that resolves to a video file", "type": "string"}}, "required": ["type", "video_url"], "type": "object"}
```
