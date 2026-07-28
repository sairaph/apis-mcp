---
title: ChatContentItems
page_id: schema-chatcontentitems-3af8a3f5
path: schemas
description: Content part for chat completion messages
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentItems

Content part for chat completion messages

```yaml
{"description": "Content part for chat completion messages", "discriminator": {"mapping": {"file": "#/components/schemas/ChatContentFile", "image_url": "#/components/schemas/ChatContentImage", "input_audio": "#/components/schemas/ChatContentAudio", "input_video": "#/components/schemas/Legacy_ChatContentVideo", "text": "#/components/schemas/ChatContentText", "video_url": "#/components/schemas/ChatContentVideo"}, "propertyName": "type"}, "example": {"text": "Hello, world!", "type": "text"}, "oneOf": [{"$ref": "#/components/schemas/ChatContentText"}, {"$ref": "#/components/schemas/ChatContentImage"}, {"$ref": "#/components/schemas/ChatContentAudio"}, {"$ref": "#/components/schemas/Legacy_ChatContentVideo"}, {"$ref": "#/components/schemas/ChatContentVideo"}, {"$ref": "#/components/schemas/ChatContentFile"}]}
```
