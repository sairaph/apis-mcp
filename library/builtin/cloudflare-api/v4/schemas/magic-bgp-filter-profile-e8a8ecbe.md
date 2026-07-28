---
title: magic_bgp_filter_profile
page_id: schema-magic-bgp-filter-profile-e8a8ecbe
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_bgp_filter_profile

```yaml
{"type": "object", "properties": {"created_on": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "Description of the filter profile", "type": "string", "example": "Allowed corporate subnets from on-premises", "default": "", "maxLength": 1024}, "id": {"$ref": "#/components/schemas/magic_schemas-identifier"}, "match_action": {"$ref": "#/components/schemas/magic_bgp_filter_match_action"}, "modified_on": {"type": "string", "format": "date-time", "readOnly": true}, "name": {"description": "Friendly name for the filter profile", "type": "string", "example": "Allowed On-Prem Imports", "maxLength": 255, "minLength": 1}, "targets": {"description": "List of CIDR prefixes. Each entry may carry an optional suffix that specifies which prefix lengths to match relative to the prefix length N: '{X,Y}' matches prefix lengths in the inclusive range [X, Y] where N <= X <= Y <= max (max is 32 for IPv4, 128 for IPv6), '{X}' matches exactly length X (equivalent to {X,X}), '+' is shorthand for {N, max} (the prefix and all more-specific subnets, including at length N itself; valid even when N is the maximum length). Omit the suffix to match the prefix exactly at length N.", "type": "array", "items": {"description": "CIDR notation prefix with optional + or {X,Y} suffix", "example": "10.0.0.0/8{8,32}", "type": "string"}, "maxItems": 1000}}, "required": ["id", "name", "description", "match_action", "targets"]}
```
