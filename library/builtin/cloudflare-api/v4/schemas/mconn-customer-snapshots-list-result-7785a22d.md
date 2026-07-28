---
title: mconn_customer_snapshots_list_result
page_id: schema-mconn-customer-snapshots-list-result-7785a22d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_customer_snapshots_list_result

```yaml
{"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "items": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_snapshot_metadata"}}}, "required": ["items", "count"]}
```
