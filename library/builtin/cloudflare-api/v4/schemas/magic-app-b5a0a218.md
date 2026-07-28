---
title: magic_app
page_id: schema-magic-app-b5a0a218
path: schemas
description: Collection of Hostnames and/or IP Subnets to associate with traffic decisions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_app

Collection of Hostnames and/or IP Subnets to associate with traffic decisions.

```yaml
{"description": "Collection of Hostnames and/or IP Subnets to associate with traffic decisions.", "type": "object", "oneOf": [{"$ref": "#/components/schemas/magic_account_app"}, {"$ref": "#/components/schemas/magic_managed_app"}]}
```
