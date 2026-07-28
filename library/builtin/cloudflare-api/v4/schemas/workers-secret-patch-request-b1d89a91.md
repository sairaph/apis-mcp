---
title: workers_secret-patch-request
page_id: schema-workers-secret-patch-request-b1d89a91
path: schemas
description: JSON Merge Patch (RFC 7396) request body for bulk secret changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_secret-patch-request

JSON Merge Patch (RFC 7396) request body for bulk secret changes.

```yaml
{"description": "JSON Merge Patch (RFC 7396) request body for bulk secret changes.\n", "type": "object", "properties": {"secrets": {"description": "Map of secret names to secret values:\n- Set to a secret object to create or update.\n- Set to `null` to delete.\n- Omit to leave unchanged.\n", "type": "object", "additionalProperties": {"allOf": [{"$ref": "#/components/schemas/workers_secret"}], "nullable": true, "type": "object"}}, "version_tags": {"description": "Optional version tags to apply to the new script version.", "type": "object", "additionalProperties": true}}}
```
