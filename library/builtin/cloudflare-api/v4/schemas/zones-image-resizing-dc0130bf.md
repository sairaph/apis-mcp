---
title: zones_image_resizing
page_id: schema-zones-image-resizing-dc0130bf
path: schemas
description: Image Transformations provides on-demand resizing, conversion and optimization for images served through Cloudflare's network. Refer to the [Image Transformations documentation](https://developers.cloudflare.com/images/) for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_image_resizing

Image Transformations provides on-demand resizing, conversion and optimization for images served through Cloudflare's network. Refer to the [Image Transformations documentation](https://developers.cloudflare.com/images/) for more information.

```yaml
{"description": "Image Transformations provides on-demand resizing, conversion and optimization for images served through Cloudflare's network. Refer to the [Image Transformations documentation](https://developers.cloudflare.com/images/) for more information.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "image_resizing", "enum": ["image_resizing"]}, "value": {"$ref": "#/components/schemas/zones_image_resizing_value"}}}], "title": "Image Transformations"}
```
