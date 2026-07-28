---
title: dns-records_DSRecord
page_id: schema-dns-records-dsrecord-805d7796
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_DSRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted DS content. See 'data' to set DS properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a DS record.", "type": "object", "properties": {"algorithm": {"description": "Algorithm.", "type": "number", "example": 3, "maximum": 255, "minimum": 0, "x-auditable": true}, "digest": {"description": "Digest.", "type": "string", "x-auditable": true}, "digest_type": {"description": "Digest Type.", "type": "number", "example": 1, "maximum": 255, "minimum": 0, "x-auditable": true}, "key_tag": {"description": "Key Tag.", "type": "number", "example": 1, "maximum": 65535, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "DS", "enum": ["DS"], "x-auditable": true}}, "type": "object"}], "title": "DS Record"}
```
