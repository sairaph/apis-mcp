---
title: VisionMultimodalContentItem
page_id: schema-visionmultimodalcontentitem-8fe0f908
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# VisionMultimodalContentItem

```yaml
{"oneOf": [{"title": "Text", "type": "object", "properties": {"type": {"type": "string", "enum": ["text"], "description": "Content type is text", "default": "text"}, "text": {"type": "string", "description": "Text content"}}, "required": ["type", "text"], "additionalProperties": false}, {"title": "Image", "type": "object", "properties": {"type": {"type": "string", "enum": ["image_url"], "description": "Content type is image URL", "default": "image_url"}, "image_url": {"type": "object", "description": "Image information", "properties": {"url": {"type": "string", "description": "Image URL or Base64 encoding. Image size limit is under 5M per image, with pixels not exceeding 6000*6000. GLM-5V GLM4.6V series are limited to 150 sheets, GLM4.5V limit 50 sheets. Supports jpg, png, jpeg formats."}}, "required": ["url"], "additionalProperties": false}}, "required": ["type", "image_url"], "additionalProperties": false}, {"title": "Video", "type": "object", "properties": {"type": {"type": "string", "enum": ["video_url"], "description": "Content type is video URL", "default": "video_url"}, "video_url": {"type": "object", "description": "Video information.", "properties": {"url": {"type": "string", "description": "Video URL address.The video size is limited to within 200 MB, GLM-5V GLM4.6V series are limited to 2 videos, GLM4.5V limit 1 video, and the format supports `mp4`，`mkv`，`mov`."}}, "required": ["url"], "additionalProperties": false}}, "required": ["type", "video_url"], "additionalProperties": false}, {"title": "File", "type": "object", "properties": {"type": {"type": "string", "enum": ["file_url"], "description": "Content type is file URL, not support passing both the `file_url` and `image_url` or `video_url` parameters at the same time.", "default": "file_url"}, "file_url": {"type": "object", "description": "File information.", "properties": {"url": {"type": "string", "description": "File URL address. Only GLM-5V-Turbo, GLM-4.6V, GLM-4.5V supported. Supports formats such as pdf、txt、word、jsonl、xlsx、pptx, with a maximum of 50."}}, "required": ["url"], "additionalProperties": false}}, "required": ["type", "file_url"], "additionalProperties": false}]}
```
