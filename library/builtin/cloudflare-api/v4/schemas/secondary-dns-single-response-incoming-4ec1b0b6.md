---
title: secondary-dns_single_response_incoming
page_id: schema-secondary-dns-single-response-incoming-4ec1b0b6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_single_response_incoming

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/secondary-dns_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"auto_refresh_seconds": {"$ref": "#/components/schemas/secondary-dns_auto_refresh_seconds"}, "checked_time": {"$ref": "#/components/schemas/secondary-dns_time"}, "created_time": {"$ref": "#/components/schemas/secondary-dns_time"}, "id": {"$ref": "#/components/schemas/secondary-dns_identifier"}, "modified_time": {"$ref": "#/components/schemas/secondary-dns_time"}, "name": {"$ref": "#/components/schemas/secondary-dns_name"}, "peers": {"$ref": "#/components/schemas/secondary-dns_peers"}, "soa_serial": {"$ref": "#/components/schemas/secondary-dns_soa_serial"}}}}, "type": "object"}]}
```
