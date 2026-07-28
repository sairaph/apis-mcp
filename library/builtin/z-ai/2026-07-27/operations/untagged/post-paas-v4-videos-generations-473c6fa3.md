---
title: POST /paas/v4/videos/generations
page_id: operation-post-paas-v4-videos-generations-79fdaecb
path: operations/untagged
description: |-
    CogVideoX is a video generation large model developed by Z.AI, equipped with powerful video generation capabilities. Simply inputting text or images allows for effortless video creation.

    Vidu: A high-performance video large model that combines high consistency and high dynamism, with precise semantic understanding and exceptional reasoning speed.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/videos/generations
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# POST /paas/v4/videos/generations

`POST /paas/v4/videos/generations`

CogVideoX is a video generation large model developed by Z.AI, equipped with powerful video generation capabilities. Simply inputting text or images allows for effortless video creation.

Vidu: A high-performance video large model that combines high consistency and high dynamism, with precise semantic understanding and exceptional reasoning speed.

## Definition

```yaml
{"description": "CogVideoX is a video generation large model developed by Z.AI, equipped with powerful video generation capabilities. Simply inputting text or images allows for effortless video creation.\n\nVidu: A high-performance video large model that combines high consistency and high dynamism, with precise semantic understanding and exceptional reasoning speed.", "parameters": [{"$ref": "#/components/parameters/AcceptLanguage"}], "requestBody": {"content": {"application/json": {"schema": {"oneOf": [{"title": "CogVideoX-3", "$ref": "#/components/schemas/CogVideoX3Request"}, {"title": "Vidu: Text to Video", "$ref": "#/components/schemas/ViduText2VideoRequest"}, {"title": "Vidu: Image to Video", "$ref": "#/components/schemas/ViduImage2VideoRequest"}, {"title": "Vidu: First & Last Frame to Video", "$ref": "#/components/schemas/ViduFrames2VideoRequest"}, {"title": "Vidu: Ref to Video", "$ref": "#/components/schemas/ViduReference2VideoRequest"}]}, "examples": {"Text to Video Example": {"value": {"model": "cogvideox-3", "prompt": "A cat is playing with a ball.", "quality": "quality", "with_audio": true, "size": "1920x1080", "fps": 30}}, "Image to Video Example": {"value": {"model": "cogvideox-3", "image_url": "https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo.jpg", "prompt": "Make the picture move", "quality": "quality", "with_audio": true, "size": "1920x1080", "fps": 30}}, "First Last Frame to Video": {"value": {"model": "cogvideox-3", "image_url": ["https://gd-hbimg.huaban.com/ccee58d77afe8f5e17a572246b1994f7e027657fe9e6-qD66In_fw1200webp", "https://gd-hbimg.huaban.com/cc2601d568a72d18d90b2cc7f1065b16b2d693f7fa3f7-hDAwNq_fw1200webp"], "prompt": "Make the picture move", "quality": "quality", "with_audio": true, "size": "1920x1080", "fps": 30}}}}}, "required": true}, "responses": {"200": {"description": "Processing successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/VideoResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
