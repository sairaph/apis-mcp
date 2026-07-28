---
title: dns-records_SSHFPRecord
page_id: schema-dns-records-sshfprecord-7874cbf7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_SSHFPRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted SSHFP content. See 'data' to set SSHFP properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a SSHFP record.", "type": "object", "properties": {"algorithm": {"description": "Algorithm.", "type": "number", "example": 2, "maximum": 255, "minimum": 0, "x-auditable": true}, "fingerprint": {"description": "Fingerprint.", "type": "string", "x-auditable": true}, "type": {"description": "Type.", "type": "number", "example": 1, "maximum": 255, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "SSHFP", "enum": ["SSHFP"], "x-auditable": true}}, "type": "object"}], "title": "SSHFP Record"}
```
