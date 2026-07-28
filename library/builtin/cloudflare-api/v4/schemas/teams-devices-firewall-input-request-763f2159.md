---
title: teams-devices_firewall_input_request
page_id: schema-teams-devices-firewall-input-request-763f2159
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_firewall_input_request

```yaml
{"type": "object", "properties": {"enabled": {"description": "Enabled.", "type": "boolean", "example": true, "x-auditable": true}, "operating_system": {"description": "Operating System.", "type": "string", "example": "windows", "enum": ["windows", "mac"], "x-auditable": true}}, "required": ["operating_system", "enabled"], "title": "Firewall"}
```
