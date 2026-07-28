---
title: d1_primary-location-hint
page_id: schema-d1-primary-location-hint-5a8033af
path: schemas
description: Specify the region to create the D1 primary, if available. If this option is omitted, the D1 will be created as close as possible to the current user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_primary-location-hint

Specify the region to create the D1 primary, if available. If this option is omitted, the D1 will be created as close as possible to the current user.

```yaml
{"description": "Specify the region to create the D1 primary, if available. If this option is omitted, the D1 will be created as close as possible to the current user.", "type": "string", "example": "wnam", "enum": ["wnam", "enam", "weur", "eeur", "apac", "oc"], "x-auditable": true}
```
