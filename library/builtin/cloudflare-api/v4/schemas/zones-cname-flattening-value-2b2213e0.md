---
title: zones_cname_flattening_value
page_id: schema-zones-cname-flattening-value-2b2213e0
path: schemas
description: Value of the cname flattening setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cname_flattening_value

Value of the cname flattening setting.

```yaml
{"description": "Value of the cname flattening setting.", "type": "string", "default": "flatten_at_root", "deprecated": true, "enum": ["flatten_at_root", "flatten_all"], "x-stainless-deprecation-message": "This zone setting is deprecated; please use the DNS Settings route instead. More information at https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-03-21"}
```
