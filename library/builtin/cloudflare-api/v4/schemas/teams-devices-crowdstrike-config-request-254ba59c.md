---
title: teams-devices_crowdstrike_config_request
page_id: schema-teams-devices-crowdstrike-config-request-254ba59c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_crowdstrike_config_request

```yaml
{"type": "object", "properties": {"api_url": {"description": "The Crowdstrike API URL.", "type": "string", "example": "https://api.us-2.crowdstrike.com", "x-auditable": true}, "client_id": {"description": "The Crowdstrike client ID.", "type": "string", "example": "example client id"}, "client_secret": {"description": "The Crowdstrike client secret.", "type": "string", "example": "example client secret", "x-sensitive": true}, "customer_id": {"description": "The Crowdstrike customer ID.", "type": "string", "example": "example customer id", "x-auditable": true}}, "required": ["api_url", "customer_id", "client_id", "client_secret"], "title": "Crowdstrike Config"}
```
