---
title: teams-devices_deployment_group
page_id: schema-teams-devices-deployment-group-ffe6416f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_deployment_group

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/teams-devices_deployment_group_created_at"}, "id": {"$ref": "#/components/schemas/teams-devices_deployment_group_id"}, "name": {"$ref": "#/components/schemas/teams-devices_deployment_group_name"}, "policy_ids": {"$ref": "#/components/schemas/teams-devices_deployment_group_policy_ids"}, "updated_at": {"$ref": "#/components/schemas/teams-devices_deployment_group_updated_at"}, "version_config": {"description": "Contains version configurations for different target environments.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_version_config"}}}, "required": ["id", "name", "version_config", "created_at", "updated_at"]}
```
