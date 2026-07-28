---
title: email-auth_SpfTree
page_id: schema-email-auth-spftree-9cc7f81d
path: schemas
description: Recursive SPF inspection tree
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_SpfTree

Recursive SPF inspection tree

```yaml
{"description": "Recursive SPF inspection tree\n", "type": "object", "properties": {"components": {"description": "Parsed SPF components (mechanisms)", "items": {"$ref": "#/components/schemas/email-auth_SpfComponent"}, "type": "array"}, "domain": {"description": "Domain being inspected", "type": "string", "example": "example.com"}, "errors": {"description": "All errors encountered during inspection, collected from the entire tree.\nThis includes errors from nested includes at any depth, providing a quick\noverview of all issues without needing to traverse the nested structure.\nEach error includes a `domain` field to identify where it occurred.\nEmpty array if no errors (omitted from JSON when empty).\n", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_InspectError"}, "example": []}, "record": {"description": "Raw SPF record content", "type": "string", "example": "v=spf1 ip4:203.0.113.1 include:spf.example.com -all"}, "total_lookups": {"description": "Total number of DNS lookups performed across all includes", "type": "integer", "example": 2}}, "required": ["domain", "record", "total_lookups", "components"]}
```
