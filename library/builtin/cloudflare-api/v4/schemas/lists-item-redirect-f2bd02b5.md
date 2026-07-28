---
title: lists_item_redirect
page_id: schema-lists-item-redirect-f2bd02b5
path: schemas
description: The definition of the redirect.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_item_redirect

The definition of the redirect.

```yaml
{"description": "The definition of the redirect.", "properties": {"include_subdomains": {"type": "boolean", "default": false, "x-auditable": true}, "preserve_path_suffix": {"type": "boolean", "default": false, "x-auditable": true}, "preserve_query_string": {"type": "boolean", "default": false, "x-auditable": true}, "source_url": {"type": "string", "example": "example.com/arch", "x-auditable": true}, "status_code": {"type": "integer", "default": 301, "enum": [301, 302, 307, 308], "x-auditable": true}, "subpath_matching": {"type": "boolean", "default": false, "x-auditable": true}, "target_url": {"type": "string", "example": "https://archlinux.org/", "x-auditable": true}}, "required": ["source_url", "target_url"]}
```
