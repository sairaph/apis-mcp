---
title: teams-devices_workspace_one_config_request
page_id: schema-teams-devices-workspace-one-config-request-a968b8f4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_workspace_one_config_request

```yaml
{"type": "object", "properties": {"api_url": {"description": "The Workspace One API URL provided in the Workspace One Admin Dashboard.", "type": "string", "example": "https://as123.awmdm.com/API"}, "auth_url": {"description": "The Workspace One Authorization URL depending on your region.", "type": "string", "example": "https://na.uemauth.workspaceone.com/connect/token"}, "client_id": {"description": "The Workspace One client ID provided in the Workspace One Admin Dashboard.", "type": "string", "example": "example client id"}, "client_secret": {"description": "The Workspace One client secret provided in the Workspace One Admin Dashboard.", "type": "string", "example": "example client secret", "x-sensitive": true}}, "required": ["api_url", "auth_url", "client_id", "client_secret"], "title": "Workspace One Config"}
```
