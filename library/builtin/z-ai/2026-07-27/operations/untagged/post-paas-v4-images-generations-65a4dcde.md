---
title: Generate Image
page_id: operation-post-paas-v4-images-generations-d0b1cb7b
path: operations/untagged
description: Use [GLM-Image](/guides/image/glm-image) series models to generate high-quality images from text prompts. Through quick and accurate understanding of user text descriptions, `AI` image expression becomes more precise and personalized.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/images/generations
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# Generate Image

`POST /paas/v4/images/generations`

Use [GLM-Image](/guides/image/glm-image) series models to generate high-quality images from text prompts. Through quick and accurate understanding of user text descriptions, `AI` image expression becomes more precise and personalized.

## Definition

```yaml
{"summary": "Generate Image", "description": "Use [GLM-Image](/guides/image/glm-image) series models to generate high-quality images from text prompts. Through quick and accurate understanding of user text descriptions, `AI` image expression becomes more precise and personalized.", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/CreateImageRequest"}, "examples": {"Generate Image Example": {"value": {"model": "glm-image", "prompt": "A cute little kitten sitting on a sunny windowsill, with the background of blue sky and white clouds.", "size": "1280x1280"}}}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ImageGenerationResponse"}}}}, "default": {"description": "Request Failed", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
