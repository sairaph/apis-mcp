---
title: load-balancing_Host
page_id: schema-load-balancing-host-a689f6c2
path: schemas
description: The 'Host' header allows to override the hostname set in the HTTP request. Current support is 1 'Host' header override per origin.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_Host

The 'Host' header allows to override the hostname set in the HTTP request. Current support is 1 'Host' header override per origin.

```yaml
{"description": "The 'Host' header allows to override the hostname set in the HTTP request. Current support is 1 'Host' header override per origin.", "type": "array", "items": {"example": "example.com", "type": "string", "x-auditable": true}}
```
