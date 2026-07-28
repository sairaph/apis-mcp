---
title: InputReference
page_id: schema-inputreference-5f652288
path: schemas
description: A reference asset used to guide video generation. Image references are supported by all providers; audio and video references are only honored by providers that support them (currently BytePlus Seedance 2.0).
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InputReference

A reference asset used to guide video generation. Image references are supported by all providers; audio and video references are only honored by providers that support them (currently BytePlus Seedance 2.0).

```yaml
{"description": "A reference asset used to guide video generation. Image references are supported by all providers; audio and video references are only honored by providers that support them (currently BytePlus Seedance 2.0).", "discriminator": {"mapping": {"audio_url": "#/components/schemas/ContentPartAudio", "image_url": "#/components/schemas/ContentPartImage", "video_url": "#/components/schemas/ContentPartVideo"}, "propertyName": "type"}, "example": {"image_url": {"url": "https://example.com/image.png"}, "type": "image_url"}, "oneOf": [{"$ref": "#/components/schemas/ContentPartImage"}, {"$ref": "#/components/schemas/ContentPartAudio"}, {"$ref": "#/components/schemas/ContentPartVideo"}]}
```
