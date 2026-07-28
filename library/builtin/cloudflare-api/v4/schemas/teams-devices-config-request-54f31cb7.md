---
title: teams-devices_config_request
page_id: schema-teams-devices-config-request-54f31cb7
path: schemas
description: The configuration object containing third-party integration information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_config_request

The configuration object containing third-party integration information.

```yaml
{"description": "The configuration object containing third-party integration information.", "type": "object", "example": {"api_url": "https://as123.awmdm.com/API", "auth_url": "https://na.uemauth.workspaceone.com/connect/token", "client_id": "example client id", "client_secret": "example client secret"}, "oneOf": [{"$ref": "#/components/schemas/teams-devices_workspace_one_config_request"}, {"$ref": "#/components/schemas/teams-devices_crowdstrike_config_request"}, {"$ref": "#/components/schemas/teams-devices_uptycs_config_request"}, {"$ref": "#/components/schemas/teams-devices_intune_config_request"}, {"$ref": "#/components/schemas/teams-devices_kolide_config_request"}, {"$ref": "#/components/schemas/teams-devices_tanium_config_request"}, {"$ref": "#/components/schemas/teams-devices_sentinelone_s2s_config_request"}, {"$ref": "#/components/schemas/teams-devices_custom_s2s_config_request"}]}
```
