---
title: zones_ssl
page_id: schema-zones-ssl-945c30ab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_ssl

```yaml
{"type": "object", "properties": {"id": {"description": "Control options for the SSL feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.\n", "type": "string", "enum": ["ssl"], "x-auditable": true}, "value": {"description": "The encryption mode that Cloudflare uses to connect to your origin server.\n", "type": "string", "example": "full", "enum": ["off", "flexible", "full", "strict", "origin_pull"], "x-auditable": true}}, "title": "SSL", "x-stainless-skip": ["terraform"]}
```
