---
title: teams-devices_kolide_config_request
page_id: schema-teams-devices-kolide-config-request-ca5801ac
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_kolide_config_request

```yaml
{"type": "object", "properties": {"client_id": {"description": "The Kolide client ID.", "type": "string", "example": "example client id"}, "client_secret": {"description": "The Kolide client secret.", "type": "string", "example": "example client secret", "x-sensitive": true}}, "required": ["client_id", "client_secret"], "title": "Kolide Config"}
```
