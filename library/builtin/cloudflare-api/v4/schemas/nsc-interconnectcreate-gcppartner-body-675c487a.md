---
title: nsc_InterconnectCreate_GcpPartner_Body
page_id: schema-nsc-interconnectcreate-gcppartner-body-675c487a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_InterconnectCreate_GcpPartner_Body

```yaml
{"type": "object", "allOf": [{"properties": {"account": {"type": "string"}, "type": {"type": "string"}}, "required": ["type", "account"], "type": "object"}, {"properties": {"bandwidth": {"$ref": "#/components/schemas/nsc_ApiBandwidth"}, "pairing_key": {"description": "Pairing key provided by GCP", "type": "string"}}, "required": ["pairing_key", "bandwidth"], "type": "object"}], "title": "GcpPartner"}
```
