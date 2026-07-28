---
title: magic_app_subnets
page_id: schema-magic-app-subnets-5d190af8
path: schemas
description: IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently unsupported)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_app_subnets

IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently unsupported)

```yaml
{"description": "IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently unsupported)", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/magic_cidr"}, {"example": "1.1.1.1/32"}]}}
```
