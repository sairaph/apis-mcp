---
title: dns-records_URIRecord
page_id: schema-dns-records-urirecord-1e59b250
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_URIRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted URI content. See 'data' to set URI properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a URI record.", "type": "object", "properties": {"target": {"description": "The record content.", "type": "string", "example": "http://example.com/example.html", "x-auditable": true}, "weight": {"description": "The record weight.", "type": "number", "example": 20, "maximum": 65535, "minimum": 0, "x-auditable": true}}}, "priority": {"$ref": "#/components/schemas/dns-records_priority"}, "type": {"description": "Record type.", "type": "string", "example": "URI", "enum": ["URI"], "x-auditable": true}}, "type": "object"}], "title": "URI Record"}
```
