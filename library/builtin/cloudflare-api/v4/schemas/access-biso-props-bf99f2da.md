---
title: access_biso_props
page_id: schema-access-biso-props-bf99f2da
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_biso_props

```yaml
{"allOf": [{"$ref": "#/components/schemas/access_feature_app_props"}, {"properties": {"domain": {"example": "authdomain.cloudflareaccess.com/browser", "readOnly": true}, "name": {"example": "Clientless Web Isolation", "default": "Clientless Web Isolation", "readOnly": true}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "biso"}]}}}]}
```
