---
title: nsc_Interconnect_GcpPartner_Body
page_id: schema-nsc-interconnect-gcppartner-body-51aeec33
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_Interconnect_GcpPartner_Body

```yaml
{"type": "object", "allOf": [{"properties": {"account": {"type": "string"}, "name": {"type": "string"}, "owner": {"type": "string"}, "type": {"type": "string"}}, "required": ["type", "name", "account"], "type": "object"}, {"properties": {"region": {"type": "string"}, "speed": {"$ref": "#/components/schemas/nsc_ApiBandwidth"}}, "required": ["region"], "type": "object"}], "title": "GcpPartner"}
```
