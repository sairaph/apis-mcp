---
title: images_image_variants
page_id: schema-images-image-variants-993beb8c
path: schemas
description: Object specifying available variants for an image.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_image_variants

Object specifying available variants for an image.

```yaml
{"description": "Object specifying available variants for an image.", "type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/images_image_thumbnail_url"}, {"$ref": "#/components/schemas/images_image_hero_url"}, {"$ref": "#/components/schemas/images_image_original_url"}]}, "example": ["https://imagedelivery.net/MTt4OTd0b0w5aj/107b9558-dd06-4bbd-5fef-9c2c16bb7900/thumbnail", "https://imagedelivery.net/MTt4OTd0b0w5aj/107b9558-dd06-4bbd-5fef-9c2c16bb7900/hero", "https://imagedelivery.net/MTt4OTd0b0w5aj/107b9558-dd06-4bbd-5fef-9c2c16bb7900/original"], "readOnly": true}
```
