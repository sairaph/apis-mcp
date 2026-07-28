---
title: teams-devices_domain_joined_input_request
page_id: schema-teams-devices-domain-joined-input-request-4b79c8b8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_domain_joined_input_request

```yaml
{"type": "object", "properties": {"domain": {"description": "Domain.", "type": "string", "example": "example.com", "x-auditable": true}, "operating_system": {"description": "Operating System.", "type": "string", "example": "windows", "enum": ["windows"], "x-auditable": true}}, "required": ["operating_system"], "title": "Domain Joined"}
```
