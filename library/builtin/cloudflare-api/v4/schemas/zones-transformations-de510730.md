---
title: zones_transformations
page_id: schema-zones-transformations-de510730
path: schemas
description: Media Transformations provides on-demand resizing, conversion and optimization for images and video served through Cloudflare's network. Refer to the [Image Transformations](https://developers.cloudflare.com/images/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_transformations

Media Transformations provides on-demand resizing, conversion and optimization for images and video served through Cloudflare's network. Refer to the [Image Transformations](https://developers.cloudflare.com/images/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.

```yaml
{"description": "Media Transformations provides on-demand resizing, conversion and optimization for images and video served through Cloudflare's network. Refer to the [Image Transformations](https://developers.cloudflare.com/images/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting. Shared between Image Transformations and Video Transformations.", "example": "transformations", "enum": ["transformations"]}, "value": {"$ref": "#/components/schemas/zones_image_resizing_value"}}}], "title": "Media Transformations"}
```
