---
title: teams-devices_ip_profile
page_id: schema-teams-devices-ip-profile-0bae573a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_ip_profile

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/teams-devices_ip_profile_created_at"}, "description": {"$ref": "#/components/schemas/teams-devices_ip_profile_description"}, "enabled": {"$ref": "#/components/schemas/teams-devices_ip_profile_enabled"}, "id": {"$ref": "#/components/schemas/teams-devices_ip_profile_id"}, "match": {"$ref": "#/components/schemas/teams-devices_ip_profile_match"}, "name": {"$ref": "#/components/schemas/teams-devices_ip_profile_name"}, "precedence": {"$ref": "#/components/schemas/teams-devices_ip_profile_precedence"}, "subnet_id": {"$ref": "#/components/schemas/teams-devices_ip_profile_subnet_id"}, "updated_at": {"$ref": "#/components/schemas/teams-devices_ip_profile_updated_at"}}, "required": ["id", "name", "match", "description", "precedence", "subnet_id", "created_at", "updated_at", "enabled"]}
```
