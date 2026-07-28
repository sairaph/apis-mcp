---
title: allowedNodeAttributes
page_id: schema-allowednodeattributes-40791ec4
path: schemas
description: The list of custom node attributes that the OAuth app is allowed to set.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# allowedNodeAttributes

The list of custom node attributes that the OAuth app is allowed to set.

```yaml
type: array
items:
    type: string
example:
    - custom:myattribute
description: |
    The list of custom node attributes that the OAuth app is allowed to set.
```
