---
title: dns-records_SRVRecord
page_id: schema-dns-records-srvrecord-425d327b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_SRVRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Priority, weight, port, and SRV target. See 'data' for setting the individual component values.", "example": "10 IN SRV 5 8806 example.com.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a SRV record.", "type": "object", "properties": {"port": {"description": "The port of the service.", "type": "number", "example": 8806, "maximum": 65535, "minimum": 0, "x-auditable": true}, "priority": {"$ref": "#/components/schemas/dns-records_priority"}, "target": {"description": "A valid hostname.", "type": "string", "format": "hostname", "example": "example.com", "x-auditable": true}, "weight": {"description": "The record weight.", "type": "number", "example": 5, "maximum": 65535, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "SRV", "enum": ["SRV"], "x-auditable": true}}, "type": "object"}], "title": "SRV Record"}
```
