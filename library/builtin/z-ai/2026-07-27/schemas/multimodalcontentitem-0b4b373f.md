---
title: MultimodalContentItem
page_id: schema-multimodalcontentitem-0b4b373f
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# MultimodalContentItem

```yaml
{"oneOf": [{"title": "Text", "type": "object", "properties": {"type": {"type": "string", "enum": ["text"], "description": "Content type is text", "default": "text"}, "text": {"type": "string", "description": "Text content"}}, "required": ["type", "text"], "additionalProperties": false}, {"title": "Image", "type": "object", "properties": {"type": {"type": "string", "enum": ["image_url"], "description": "Content type is image URL", "default": "image_url"}, "image_url": {"type": "object", "description": "Image information", "properties": {"url": {"type": "string", "description": "Image URL or Base64 encoding. Image size limit is under 5M per image, with pixels not exceeding 6000*6000. Supports jpg, png, jpeg formats."}}, "required": ["url"], "additionalProperties": false}}, "required": ["type", "image_url"], "additionalProperties": false}, {"title": "Audio", "type": "object", "properties": {"type": {"type": "string", "enum": ["input_audio"], "description": "Content type is audio input", "default": "input_audio"}, "input_audio": {"type": "object", "description": "Audio information", "properties": {"data": {"type": "string", "description": "Base64 encoding of audio file. Audio duration should not exceed 10 minutes. 1s audio = 12.5 Tokens, rounded up."}, "format": {"type": "string", "description": "Audio file format, supports wav and mp3", "enum": ["wav", "mp3"]}}, "required": ["data", "format"], "additionalProperties": false}}, "required": ["type", "input_audio"], "additionalProperties": false}, {"title": "Video", "type": "object", "properties": {"type": {"type": "string", "enum": ["video_url"], "description": "Content type is video URL", "default": "video_url"}, "video_url": {"type": "object", "description": "Video information.", "properties": {"url": {"type": "string", "description": "Video URL address."}}, "required": ["url"], "additionalProperties": false}}, "required": ["type", "video_url"], "additionalProperties": false}, {"title": "File", "type": "object", "properties": {"type": {"type": "string", "enum": ["file_url"], "description": "Content type is file URL", "default": "file_url"}, "file_url": {"type": "object", "description": "File information.", "properties": {"url": {"type": "string", "description": "File URL address. Supports formats such as PDF and Word, with a maximum of 50 pages for parsing."}}, "required": ["url"], "additionalProperties": false}}, "required": ["type", "file_url"], "additionalProperties": false}]}
```
