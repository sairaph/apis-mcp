---
title: cache-rules_origin_tls_compliance_modes
page_id: schema-cache-rules-origin-tls-compliance-modes-0d9b6e2a
path: schemas
description: Origin TLS Compliance Modes constrains the set of TLS key-exchange algorithms Cloudflare may use when establishing the TLS connection to the zone's origin. The value is a list of named compliance modes (currently `fips` and `pqh`). Multiple modes are combined as the intersection of their permitted algorithm lists.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_tls_compliance_modes

Origin TLS Compliance Modes constrains the set of TLS key-exchange algorithms Cloudflare may use when establishing the TLS connection to the zone's origin. The value is a list of named compliance modes (currently `fips` and `pqh`). Multiple modes are combined as the intersection of their permitted algorithm lists.

```yaml
{"description": "Origin TLS Compliance Modes constrains the set of TLS key-exchange algorithms Cloudflare may use when establishing the TLS connection to the zone's origin. The value is a list of named compliance modes (currently `fips` and `pqh`). Multiple modes are combined as the intersection of their permitted algorithm lists.", "type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "type": "string", "example": "origin_tls_compliance_modes", "enum": ["origin_tls_compliance_modes"], "x-auditable": true}}, "type": "object"}], "title": "Origin TLS Compliance Modes"}
```
