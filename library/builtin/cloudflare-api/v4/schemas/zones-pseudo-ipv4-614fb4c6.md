---
title: zones_pseudo_ipv4
page_id: schema-zones-pseudo-ipv4-614fb4c6
path: schemas
description: The value set for the Pseudo IPv4 setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_pseudo_ipv4

The value set for the Pseudo IPv4 setting.

```yaml
{"description": "The value set for the Pseudo IPv4 setting.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "Value of the Pseudo IPv4 setting.", "default": "pseudo_ipv4", "enum": ["pseudo_ipv4"]}, "value": {"$ref": "#/components/schemas/zones_pseudo_ipv4_value"}}}], "title": "Pseudo IPv4 Value"}
```
