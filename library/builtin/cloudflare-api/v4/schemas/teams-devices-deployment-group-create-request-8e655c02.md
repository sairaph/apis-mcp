---
title: teams-devices_deployment_group_create_request
page_id: schema-teams-devices-deployment-group-create-request-8e655c02
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_deployment_group_create_request

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/teams-devices_deployment_group_name"}, "policy_ids": {"description": "Contains an optional list of policy IDs assigned to a group.", "type": "array", "items": {"type": "string"}}, "version_config": {"description": "Contains at least one version configuration.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_version_config"}, "minItems": 1}}, "required": ["name", "version_config"]}
```
