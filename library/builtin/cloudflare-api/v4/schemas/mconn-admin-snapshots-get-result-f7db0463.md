---
title: mconn_admin_snapshots_get_result
page_id: schema-mconn-admin-snapshots-get-result-f7db0463
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_admin_snapshots_get_result

```yaml
{"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "items": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_snapshot_metadata"}}}, "required": ["items", "count"]}
```
