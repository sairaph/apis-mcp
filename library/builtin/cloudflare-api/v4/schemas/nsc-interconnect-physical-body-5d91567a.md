---
title: nsc_Interconnect_Physical_Body
page_id: schema-nsc-interconnect-physical-body-5d91567a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_Interconnect_Physical_Body

```yaml
{"type": "object", "allOf": [{"properties": {"account": {"type": "string"}, "name": {"type": "string"}, "owner": {"type": "string"}, "type": {"type": "string"}}, "required": ["type", "name", "account"], "type": "object"}, {"properties": {"facility": {"$ref": "#/components/schemas/nsc_FacilityInfo"}, "site": {"$ref": "#/components/schemas/nsc_CloudflareSite"}, "slot_id": {"type": "string", "format": "uuid"}, "speed": {"type": "string"}}, "required": ["slot_id", "site", "speed", "facility"], "type": "object"}], "title": "Physical"}
```
