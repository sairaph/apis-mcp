---
title: turnstile_invalidate_immediately
page_id: schema-turnstile-invalidate-immediately-2ff7a792
path: schemas
description: |-
    If `invalidate_immediately` is set to `false`, the previous secret will
    remain valid for two hours. Otherwise, the secret is immediately
    invalidated, and requests using it will be rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# turnstile_invalidate_immediately

If `invalidate_immediately` is set to `false`, the previous secret will
remain valid for two hours. Otherwise, the secret is immediately
invalidated, and requests using it will be rejected.

```yaml
{"description": "If `invalidate_immediately` is set to `false`, the previous secret will\nremain valid for two hours. Otherwise, the secret is immediately\ninvalidated, and requests using it will be rejected.\n", "type": "boolean", "default": false, "x-auditable": true}
```
