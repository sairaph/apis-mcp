---
title: mcn_cloud_platform_client
page_id: schema-mcn-cloud-platform-client-6cac0b22
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_cloud_platform_client

```yaml
{"type": "object", "properties": {"client_type": {"type": "string", "enum": ["MAGIC_WAN_CLOUD_ONRAMP"], "x-auditable": true}, "id": {"$ref": "#/components/schemas/mcn_platform_client_id"}, "name": {"type": "string", "x-auditable": true}}, "required": ["client_type", "name", "id"]}
```
