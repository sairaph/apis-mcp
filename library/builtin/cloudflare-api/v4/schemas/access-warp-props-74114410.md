---
title: access_warp_props
page_id: schema-access-warp-props-74114410
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_warp_props

```yaml
{"allOf": [{"$ref": "#/components/schemas/access_feature_app_props"}, {"properties": {"domain": {"example": "authdomain.cloudflareaccess.com/warp", "readOnly": true}, "name": {"example": "Warp Login App", "default": "Warp Login App", "readOnly": true}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "warp"}]}}}]}
```
