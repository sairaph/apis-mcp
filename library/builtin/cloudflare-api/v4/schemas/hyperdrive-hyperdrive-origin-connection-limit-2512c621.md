---
title: hyperdrive_hyperdrive-origin-connection-limit
page_id: schema-hyperdrive-hyperdrive-origin-connection-limit-2512c621
path: schemas
description: |-
    The (soft) maximum number of connections the Hyperdrive is allowed to make to the origin database.

    Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts.
    If not specified, defaults to 20 for free tier and 60 for paid tier.
    Contact Cloudflare if you need a higher limit.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-origin-connection-limit

The (soft) maximum number of connections the Hyperdrive is allowed to make to the origin database.

Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts.
If not specified, defaults to 20 for free tier and 60 for paid tier.
Contact Cloudflare if you need a higher limit.

```yaml
{"description": "The (soft) maximum number of connections the Hyperdrive is allowed to make to the origin database.\n\nMaximum allowed: 20 for free tier accounts, 100 for paid tier accounts.\nIf not specified, defaults to 20 for free tier and 60 for paid tier.\nContact Cloudflare if you need a higher limit.\n", "type": "integer", "example": 60, "minimum": 5, "x-auditable": true}
```
