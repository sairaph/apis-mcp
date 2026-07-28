---
title: teams-devices_uptycs_config_request
page_id: schema-teams-devices-uptycs-config-request-0cb72ea4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_uptycs_config_request

```yaml
{"type": "object", "properties": {"api_url": {"description": "The Uptycs API URL.", "type": "string", "example": "rnd.uptycs.io", "x-auditable": true}, "client_key": {"description": "The Uptycs client secret.", "type": "string", "example": "example client key", "x-sensitive": true}, "client_secret": {"description": "The Uptycs client secret.", "type": "string", "example": "example client secret", "x-sensitive": true}, "customer_id": {"description": "The Uptycs customer ID.", "type": "string", "example": "example customer id"}}, "required": ["api_url", "client_key", "customer_id", "client_secret"], "title": "Uptycs Config"}
```
