---
title: teams-devices_service_mode_v2
page_id: schema-teams-devices-service-mode-v2-5f767afa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_service_mode_v2

```yaml
{"type": "object", "properties": {"mode": {"description": "The mode to run the WARP client under.", "type": "string", "example": "proxy", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "port": {"description": "The port number when used with proxy mode.", "type": "number", "example": 3000, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "x-stainless-terraform-configurability": "computed_optional"}
```
