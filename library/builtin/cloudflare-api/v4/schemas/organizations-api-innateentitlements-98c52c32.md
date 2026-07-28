---
title: organizations-api_InnateEntitlements
page_id: schema-organizations-api-innateentitlements-98c52c32
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_InnateEntitlements

```yaml
{"type": "object", "properties": {"allow_add_subdomain": {"$ref": "#/components/schemas/organizations-api_BoolAllocation"}, "allow_auto_accept_invites": {"$ref": "#/components/schemas/organizations-api_BoolAllocation"}, "cname_setup_allowed": {"$ref": "#/components/schemas/organizations-api_BoolAllocation"}, "custom_entitlements": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Entitlement"}, "nullable": true}, "mhs_certificate_count": {"$ref": "#/components/schemas/organizations-api_MaxCountAllocation"}, "partial_setup_allowed": {"$ref": "#/components/schemas/organizations-api_BoolAllocation"}}, "required": ["allow_add_subdomain", "cname_setup_allowed", "partial_setup_allowed", "allow_auto_accept_invites", "mhs_certificate_count", "custom_entitlements"]}
```
