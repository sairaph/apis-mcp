---
title: firewall_asn_configuration
page_id: schema-firewall-asn-configuration-d45a9464
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_asn_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `asn` when specifying an Autonomous System Number (ASN) in the rule.", "type": "string", "example": "asn", "enum": ["asn"]}, "value": {"description": "The AS number to match.", "type": "string", "example": "AS12345"}}, "title": "An ASN configuration."}
```
