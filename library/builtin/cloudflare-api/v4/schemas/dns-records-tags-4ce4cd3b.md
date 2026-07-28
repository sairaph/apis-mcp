---
title: dns-records_tags
page_id: schema-dns-records-tags-4ce4cd3b
path: schemas
description: Custom tags for the DNS record. This field has no effect on DNS responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_tags

Custom tags for the DNS record. This field has no effect on DNS responses.

```yaml
{"description": "Custom tags for the DNS record. This field has no effect on DNS responses.", "type": "array", "items": {"description": "Individual tag of the form name:value (the name must consist of only letters, numbers, underscores and hyphens)", "example": "owner:dns-team", "type": "string", "x-auditable": true}, "default": [], "x-stainless-collection-type": "set"}
```
