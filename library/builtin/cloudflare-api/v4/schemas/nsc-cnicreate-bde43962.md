---
title: nsc_CniCreate
page_id: schema-nsc-cnicreate-bde43962
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_CniCreate

```yaml
{"type": "object", "properties": {"account": {"$ref": "#/components/schemas/nsc_AccountTag"}, "bgp": {"$ref": "#/components/schemas/nsc_BgpControl"}, "interconnect": {"type": "string"}, "magic": {"$ref": "#/components/schemas/nsc_MagicSettings"}}, "required": ["interconnect", "account", "magic"]}
```
