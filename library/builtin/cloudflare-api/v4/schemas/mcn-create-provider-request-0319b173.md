---
title: mcn_create_provider_request
page_id: schema-mcn-create-provider-request-0319b173
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_create_provider_request

```yaml
{"type": "object", "properties": {"cloud_type": {"$ref": "#/components/schemas/mcn_cloud_type"}, "description": {"type": "string", "x-auditable": true}, "friendly_name": {"type": "string", "x-auditable": true}}, "required": ["friendly_name", "cloud_type"]}
```
