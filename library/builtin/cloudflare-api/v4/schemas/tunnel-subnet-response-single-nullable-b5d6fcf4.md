---
title: tunnel_subnet_response_single_nullable
page_id: schema-tunnel-subnet-response-single-nullable-b5d6fcf4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_subnet_response_single_nullable

```yaml
{"allOf": [{"$ref": "#/components/schemas/tunnel_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true, "properties": {"capacity": {"$ref": "#/components/schemas/tunnel_subnet_capacity"}, "comment": {"$ref": "#/components/schemas/tunnel_subnet_comment"}, "created_at": {"$ref": "#/components/schemas/tunnel_created_at"}, "deleted_at": {"$ref": "#/components/schemas/tunnel_deleted_at"}, "id": {"$ref": "#/components/schemas/tunnel_subnet_id"}, "is_default_network": {"$ref": "#/components/schemas/tunnel_subnet_is_default_network"}, "name": {"$ref": "#/components/schemas/tunnel_subnet_name"}, "network": {"$ref": "#/components/schemas/tunnel_subnet_ip_network"}, "subnet_type": {"$ref": "#/components/schemas/tunnel_subnet_type"}}}}, "type": "object"}]}
```
