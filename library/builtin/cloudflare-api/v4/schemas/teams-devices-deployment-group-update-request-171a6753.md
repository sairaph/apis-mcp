---
title: teams-devices_deployment_group_update_request
page_id: schema-teams-devices-deployment-group-update-request-171a6753
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_deployment_group_update_request

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/teams-devices_deployment_group_name"}, "policy_ids": {"description": "Replaces the entire list of policy IDs.", "type": "array", "items": {"type": "string"}}, "version_config": {"description": "Replaces the entire version_config array.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_version_config"}}}}
```
