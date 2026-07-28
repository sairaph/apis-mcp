---
title: waitingroom_cookie_suffix
page_id: schema-waitingroom-cookie-suffix-d9430276
path: schemas
description: Appends a '_' + a custom suffix to the end of Cloudflare Waiting Room's cookie name(__cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be `__cf_waitingroom_abcd`. This field is required if using `additional_routes`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_cookie_suffix

Appends a '_' + a custom suffix to the end of Cloudflare Waiting Room's cookie name(__cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be `__cf_waitingroom_abcd`. This field is required if using `additional_routes`.

```yaml
{"description": "Appends a '_' + a custom suffix to the end of Cloudflare Waiting Room's cookie name(__cf_waitingroom). If `cookie_suffix` is \"abcd\", the cookie name will be `__cf_waitingroom_abcd`. This field is required if using `additional_routes`.", "type": "string", "example": "abcd", "default": "", "x-auditable": true}
```
