---
title: mconn_customer_device
page_id: schema-mconn-customer-device-c2950843
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_customer_device

```yaml
{"type": "object", "properties": {"id": {"$ref": "#/components/schemas/mconn_uuid"}, "serial_number": {"type": "string", "x-auditable": true}, "type": {"type": "string", "enum": ["MANAGED", "LICENSED"], "x-auditable": true}}, "required": ["id"]}
```
