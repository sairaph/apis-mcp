---
title: dnssec_dnssec_key_state
page_id: schema-dnssec-dnssec-key-state-cbdf48b9
path: schemas
description: Lifecycle state tag attached to the DNSSEC key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_dnssec_key_state

Lifecycle state tag attached to the DNSSEC key.

```yaml
{"description": "Lifecycle state tag attached to the DNSSEC key.", "type": "string", "example": "active", "enum": ["active", "publish", "external", "retired", "revoked", "removed"], "x-auditable": true}
```
