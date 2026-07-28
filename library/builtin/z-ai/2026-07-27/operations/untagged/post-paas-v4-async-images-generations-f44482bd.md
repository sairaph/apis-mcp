---
title: Generate Image (Async)
page_id: operation-post-paas-v4-async-images-generations-f959cee4
path: operations/untagged
description: Use the [GLM-Image](/guides/image/glm-image) series models to generate high-quality images from text prompts. Through quick and accurate understanding of user text descriptions, `AI` image expression becomes more precise and personalized. Only supports `GLM-Image` model.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/async/images/generations
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# Generate Image (Async)

`POST /paas/v4/async/images/generations`

Use the [GLM-Image](/guides/image/glm-image) series models to generate high-quality images from text prompts. Through quick and accurate understanding of user text descriptions, `AI` image expression becomes more precise and personalized. Only supports `GLM-Image` model.

## Definition

```yaml
{"summary": "Generate Image (Async)", "description": "Use the [GLM-Image](/guides/image/glm-image) series models to generate high-quality images from text prompts. Through quick and accurate understanding of user text descriptions, `AI` image expression becomes more precise and personalized. Only supports `GLM-Image` model.", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/AsyncCreateImageRequest"}, "examples": {"Image Generation Example": {"value": {"model": "glm-image", "prompt": "A cute little kitten sitting on a sunny windowsill, with the background of blue sky and white clouds.", "size": "1280x1280"}}}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/AsyncResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
