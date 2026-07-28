---
title: access_self_hosted_domains
page_id: schema-access-self-hosted-domains-0c5df557
path: schemas
description: List of public domains that Access will secure. This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_self_hosted_domains

List of public domains that Access will secure. This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.

```yaml
{"description": "List of public domains that Access will secure. This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.\n", "type": "array", "items": {"description": "A domain that Access will secure.", "type": "string"}, "example": ["test.example.com/admin", "test.anotherexample.com/staff"], "default": [], "deprecated": true}
```
