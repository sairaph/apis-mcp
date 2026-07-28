---
title: teams-devices_ip_profile_update_request
page_id: schema-teams-devices-ip-profile-update-request-87b702d5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_ip_profile_update_request

```yaml
{"type": "object", "properties": {"description": {"description": "An optional description of the Device IP profile.", "type": "string", "example": "example comment", "x-auditable": true}, "enabled": {"$ref": "#/components/schemas/teams-devices_ip_profile_enabled"}, "match": {"$ref": "#/components/schemas/teams-devices_ip_profile_match"}, "name": {"$ref": "#/components/schemas/teams-devices_ip_profile_name"}, "precedence": {"$ref": "#/components/schemas/teams-devices_ip_profile_precedence"}, "subnet_id": {"$ref": "#/components/schemas/teams-devices_ip_profile_subnet_id"}}}
```
