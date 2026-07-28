---
title: teams-devices_unique_client_id_input_request
page_id: schema-teams-devices-unique-client-id-input-request-4a334e07
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_unique_client_id_input_request

```yaml
{"type": "object", "properties": {"id": {"description": "List ID.", "type": "string", "example": "da3de859-8f6e-47ea-a2b5-b2433858471f", "x-auditable": true}, "operating_system": {"description": "Operating System.", "type": "string", "example": "android", "enum": ["android", "ios", "chromeos"], "x-auditable": true}}, "required": ["operating_system", "id"], "title": "Unique Client ID"}
```
