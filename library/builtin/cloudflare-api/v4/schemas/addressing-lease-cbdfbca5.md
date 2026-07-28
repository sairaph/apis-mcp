---
title: addressing_lease
page_id: schema-addressing-lease-cbdfbca5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_lease

```yaml
{"type": "object", "properties": {"active_from": {"description": "Timestamp of the moment the lease was created.\n", "type": "string", "format": "date-time", "example": "2020-01-01T00:00:00Z"}, "cidrs": {"description": "CIDRs attached to the lease", "type": "array", "items": {"$ref": "#/components/schemas/addressing_schemas-cidr"}, "example": ["192.0.2.100/32", "192.0.2.101/32"]}, "created_at": {"$ref": "#/components/schemas/addressing_created_at"}, "id": {"$ref": "#/components/schemas/addressing_lease_id"}, "modified_at": {"$ref": "#/components/schemas/addressing_modified_at"}, "owner_id": {"$ref": "#/components/schemas/addressing_lease_owner_id"}, "purpose": {"description": "Describes the purpose of the addresses.", "type": "string", "example": "Spectrum Static IPs"}}}
```
