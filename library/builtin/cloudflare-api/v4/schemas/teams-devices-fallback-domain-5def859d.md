---
title: teams-devices_fallback_domain
page_id: schema-teams-devices-fallback-domain-5def859d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_fallback_domain

```yaml
{"type": "object", "properties": {"description": {"description": "A description of the fallback domain, displayed in the client UI.", "type": "string", "example": "Domain bypass for local development", "maxLength": 100, "x-auditable": true}, "dns_server": {"description": "A list of IP addresses to handle domain resolution.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_ip"}}, "suffix": {"description": "The domain suffix to match when resolving locally.", "type": "string", "example": "example.com", "x-auditable": true}}, "required": ["suffix"]}
```
