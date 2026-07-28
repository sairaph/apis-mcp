---
title: images_image_variant_options
page_id: schema-images-image-variant-options-66f9ee2b
path: schemas
description: Allows you to define image resizing sizes for different use cases.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_image_variant_options

Allows you to define image resizing sizes for different use cases.

```yaml
{"description": "Allows you to define image resizing sizes for different use cases.", "type": "object", "properties": {"fit": {"$ref": "#/components/schemas/images_image_variant_fit"}, "height": {"$ref": "#/components/schemas/images_image_variant_height"}, "metadata": {"$ref": "#/components/schemas/images_image_variant_schemas_metadata"}, "width": {"$ref": "#/components/schemas/images_image_variant_width"}}, "required": ["fit", "metadata", "width", "height"]}
```
