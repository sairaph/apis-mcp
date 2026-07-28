---
title: load-balancing_resource_reference
page_id: schema-load-balancing-resource-reference-3d78f106
path: schemas
description: A reference to a load balancer resource.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_resource_reference

A reference to a load balancer resource.

```yaml
{"description": "A reference to a load balancer resource.", "type": "object", "properties": {"reference_type": {"description": "When listed as a reference, the type (direction) of the reference.", "type": "string", "enum": ["referral", "referrer"], "x-auditable": true}, "references": {"description": "A list of references to (referrer) or from (referral) this resource.", "type": "array", "items": {"description": "A reference to a load balancer resource.", "type": "object"}, "example": [{"reference_type": "referrer", "resource_id": "699d98642c564d2e855e9661899b7252", "resource_name": "www.example.com", "resource_type": "load_balancer"}, {"reference_type": "referral", "resource_id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "resource_name": "Login page monitor", "resource_type": "monitor"}]}, "resource_id": {"type": "string", "example": "17b5962d775c646f3f9725cbc7a53df4"}, "resource_name": {"description": "The human-identifiable name of the resource.", "type": "string", "example": "primary-dc-1", "x-auditable": true}, "resource_type": {"description": "The type of the resource.", "type": "string", "example": "pool", "enum": ["load_balancer", "monitor", "pool"], "x-auditable": true}}}
```
