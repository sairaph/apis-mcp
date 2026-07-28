---
title: nsc_Interconnect
page_id: schema-nsc-interconnect-c7c4b785
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_Interconnect

```yaml
{"type": "object", "discriminator": {"mapping": {"direct": "#/components/schemas/nsc_Interconnect_Physical_Body", "gcp_partner": "#/components/schemas/nsc_Interconnect_GcpPartner_Body"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/nsc_Interconnect_Physical_Body"}, {"$ref": "#/components/schemas/nsc_Interconnect_GcpPartner_Body"}]}
```
