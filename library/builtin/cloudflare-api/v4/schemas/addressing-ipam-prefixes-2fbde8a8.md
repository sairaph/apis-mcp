---
title: addressing_ipam-prefixes
page_id: schema-addressing-ipam-prefixes-2fbde8a8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_ipam-prefixes

```yaml
{"type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/addressing_account_identifier"}, "advertised": {"$ref": "#/components/schemas/addressing_advertised"}, "advertised_modified_at": {"$ref": "#/components/schemas/addressing_advertised_modified_at_nullable"}, "approved": {"$ref": "#/components/schemas/addressing_approved"}, "asn": {"$ref": "#/components/schemas/addressing_asn"}, "cidr": {"$ref": "#/components/schemas/addressing_cidr"}, "created_at": {"$ref": "#/components/schemas/addressing_timestamp"}, "delegate_loa_creation": {"$ref": "#/components/schemas/addressing_delegate_loa_creation"}, "description": {"$ref": "#/components/schemas/addressing_description"}, "id": {"$ref": "#/components/schemas/addressing_prefix_identifier"}, "irr_validation_state": {"$ref": "#/components/schemas/addressing_validation_state"}, "loa_document_id": {"$ref": "#/components/schemas/addressing_loa_document_identifier"}, "modified_at": {"$ref": "#/components/schemas/addressing_timestamp"}, "on_demand_enabled": {"$ref": "#/components/schemas/addressing_on_demand_enabled"}, "on_demand_locked": {"$ref": "#/components/schemas/addressing_on_demand_locked"}, "ownership_validation_state": {"$ref": "#/components/schemas/addressing_validation_state"}, "ownership_validation_token": {"$ref": "#/components/schemas/addressing_ownership_validation_token"}, "rpki_validation_state": {"$ref": "#/components/schemas/addressing_validation_state"}}}
```
