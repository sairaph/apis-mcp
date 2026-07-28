---
title: magic_cf1_site_ramp
page_id: schema-magic-cf1-site-ramp-89a7f064
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_cf1_site_ramp

```yaml
{"type": "object", "properties": {"created_on": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "A human-provided description of the ramp.", "type": "string", "example": "Primary CF GRE tunnel"}, "gre": {"$ref": "#/components/schemas/magic_cf1_site_managed_ramp"}, "gre_interconnect": {"$ref": "#/components/schemas/magic_cf1_site_managed_ramp"}, "id": {"allOf": [{"$ref": "#/components/schemas/magic_identifier"}], "readOnly": true}, "ipsec": {"$ref": "#/components/schemas/magic_cf1_site_managed_ramp"}, "mconn": {"$ref": "#/components/schemas/magic_cf1_site_managed_ramp"}, "modified_on": {"type": "string", "format": "date-time", "readOnly": true}, "mpls_interconnect": {"$ref": "#/components/schemas/magic_cf1_site_managed_ramp"}, "name": {"description": "A human-provided name describing the ramp that should be unique within the CF1 Site.", "type": "string", "example": "primary_gre_ramp"}, "type": {"allOf": [{"$ref": "#/components/schemas/magic_cf1_site_ramp_type"}], "readOnly": true}}, "oneOf": [{"required": ["gre"]}, {"required": ["ipsec"]}, {"required": ["gre_interconnect"]}, {"required": ["mpls_interconnect"]}, {"required": ["mconn"]}], "required": ["id", "name", "type", "created_on", "modified_on"]}
```
