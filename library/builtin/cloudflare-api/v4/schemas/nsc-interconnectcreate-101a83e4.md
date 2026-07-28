---
title: nsc_InterconnectCreate
page_id: schema-nsc-interconnectcreate-101a83e4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_InterconnectCreate

```yaml
{"type": "object", "discriminator": {"mapping": {"direct": "#/components/schemas/nsc_InterconnectCreate_Physical_Body", "gcp_partner": "#/components/schemas/nsc_InterconnectCreate_GcpPartner_Body"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/nsc_InterconnectCreate_Physical_Body"}, {"$ref": "#/components/schemas/nsc_InterconnectCreate_GcpPartner_Body"}]}
```
