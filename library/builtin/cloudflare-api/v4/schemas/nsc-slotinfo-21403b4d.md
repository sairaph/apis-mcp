---
title: nsc_SlotInfo
page_id: schema-nsc-slotinfo-21403b4d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_SlotInfo

```yaml
{"type": "object", "properties": {"account": {"$ref": "#/components/schemas/nsc_AccountTag"}, "facility": {"$ref": "#/components/schemas/nsc_FacilityInfo"}, "id": {"description": "Slot ID", "type": "string", "format": "uuid"}, "occupied": {"description": "Whether the slot is occupied or not", "type": "boolean"}, "site": {"type": "string"}, "speed": {"type": "string"}}, "required": ["id", "occupied", "site", "speed", "facility"]}
```
