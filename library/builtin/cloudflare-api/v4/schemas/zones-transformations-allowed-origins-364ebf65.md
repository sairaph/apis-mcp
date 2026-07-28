---
title: zones_transformations_allowed_origins
page_id: schema-zones-transformations-allowed-origins-364ebf65
path: schemas
description: Media Transformations Allowed Origins restricts transformations for images and video served through Cloudflare's network. Refer to the [Image Transformations](https://developers.cloudflare.com/images/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_transformations_allowed_origins

Media Transformations Allowed Origins restricts transformations for images and video served through Cloudflare's network. Refer to the [Image Transformations](https://developers.cloudflare.com/images/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.

```yaml
{"description": "Media Transformations Allowed Origins restricts transformations for images and video served through Cloudflare's network. Refer to the [Image Transformations](https://developers.cloudflare.com/images/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting. Shared between Image Transformations and Video Transformations.", "example": "transformations_allowed_origins", "enum": ["transformations_allowed_origins"]}, "value": {"$ref": "#/components/schemas/zones_transformations_allowed_origins_value"}}}], "title": "Media Transformations Allowed Origins"}
```
