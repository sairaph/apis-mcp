---
title: zones_type
page_id: schema-zones-type-4ad8ae12
path: schemas
description: |-
    A full zone implies that DNS is hosted with Cloudflare. A partial zone is
    typically a partner-hosted zone or a CNAME setup.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_type

A full zone implies that DNS is hosted with Cloudflare. A partial zone is
typically a partner-hosted zone or a CNAME setup.

```yaml
{"description": "A full zone implies that DNS is hosted with Cloudflare. A partial zone is\ntypically a partner-hosted zone or a CNAME setup.\n", "type": "string", "example": "full", "default": "full", "enum": ["full", "partial", "secondary", "internal"]}
```
