---
title: teams-devices_intune_config_request
page_id: schema-teams-devices-intune-config-request-acac7223
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_intune_config_request

```yaml
{"type": "object", "properties": {"client_id": {"description": "The Intune client ID.", "type": "string", "example": "example client id"}, "client_secret": {"description": "The Intune client secret.", "type": "string", "example": "example client secret", "x-sensitive": true}, "customer_id": {"description": "The Intune customer ID.", "type": "string", "example": "example customer id"}}, "required": ["customer_id", "client_id", "client_secret"], "title": "Intune Config"}
```
