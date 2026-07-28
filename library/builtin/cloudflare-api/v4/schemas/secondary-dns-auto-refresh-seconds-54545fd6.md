---
title: secondary-dns_auto_refresh_seconds
page_id: schema-secondary-dns-auto-refresh-seconds-54545fd6
path: schemas
description: |-
    How often should a secondary zone auto refresh regardless of DNS NOTIFY.
    Not applicable for primary zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_auto_refresh_seconds

How often should a secondary zone auto refresh regardless of DNS NOTIFY.
Not applicable for primary zones.

```yaml
{"description": "How often should a secondary zone auto refresh regardless of DNS NOTIFY.\nNot applicable for primary zones.", "type": "number", "example": 86400, "default": 86400, "minimum": 300, "x-auditable": true}
```
