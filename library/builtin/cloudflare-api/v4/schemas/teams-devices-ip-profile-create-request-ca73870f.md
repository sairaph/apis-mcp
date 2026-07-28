---
title: teams-devices_ip_profile_create_request
page_id: schema-teams-devices-ip-profile-create-request-ca73870f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_ip_profile_create_request

```yaml
{"type": "object", "properties": {"description": {"description": "An optional description of the Device IP profile.", "type": "string", "example": "example comment", "nullable": true, "x-auditable": true}, "enabled": {"description": "Whether the Device IP profile will be applied to matching devices.", "type": "boolean", "example": true, "default": true, "x-auditable": true}, "match": {"$ref": "#/components/schemas/teams-devices_ip_profile_match"}, "name": {"$ref": "#/components/schemas/teams-devices_ip_profile_name"}, "precedence": {"$ref": "#/components/schemas/teams-devices_ip_profile_precedence"}, "subnet_id": {"$ref": "#/components/schemas/teams-devices_ip_profile_subnet_id"}}, "required": ["name", "subnet_id", "precedence", "match"]}
```
