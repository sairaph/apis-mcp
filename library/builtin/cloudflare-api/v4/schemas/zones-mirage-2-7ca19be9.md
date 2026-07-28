---
title: zones_mirage-2
page_id: schema-zones-mirage-2-7ca19be9
path: schemas
description: |-
    Automatically optimize image loading for website visitors on mobile
    devices. Refer to [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed)
    for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_mirage-2

Automatically optimize image loading for website visitors on mobile
devices. Refer to [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed)
for more information.

```yaml
{"description": "Automatically optimize image loading for website visitors on mobile\ndevices. Refer to [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed)\nfor more information.\n", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "mirage", "enum": ["mirage"]}, "value": {"$ref": "#/components/schemas/zones_mirage_value"}}}], "deprecated": true, "title": "Mirage Image Optimization", "x-stainless-deprecation-message": "Mirage is being deprecated. More information at https://developers.cloudflare.com/speed/optimization/images/mirage/"}
```
