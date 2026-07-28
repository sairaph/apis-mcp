---
title: intel_schemas-ip
page_id: schema-intel-schemas-ip-c907989d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_schemas-ip

```yaml
{"type": "object", "properties": {"belongs_to_ref": {"description": "Specifies a reference to the autonomous systems (AS) that the IP address belongs to.", "type": "object", "properties": {"country": {"type": "string", "example": "US", "x-auditable": true}, "description": {"type": "string", "example": "CLOUDFLARENET", "x-auditable": true}, "id": {"type": "string", "example": "autonomous-system--2fa28d71-3549-5a38-af05-770b79ad6ea8", "x-auditable": true}, "type": {"description": "Infrastructure type of this ASN.", "type": "string", "example": "hosting_provider", "enum": ["hosting_provider", "isp", "organization"], "x-auditable": true}, "value": {"type": "string", "x-auditable": true}}}, "ip": {"$ref": "#/components/schemas/intel_ip"}, "risk_types": {"type": "array", "items": {"properties": {"id": {"type": "number"}, "name": {"type": "string", "x-auditable": true}, "super_category_id": {"type": "number"}}, "type": "object"}, "example": [{"id": 131, "name": "Phishing", "super_category_id": 21}]}}}
```
