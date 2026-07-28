---
title: zero-trust-gateway_anti-virus-settings
page_id: schema-zero-trust-gateway-anti-virus-settings-a925b0af
path: schemas
description: Specify anti-virus settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_anti-virus-settings

Specify anti-virus settings.

```yaml
{"description": "Specify anti-virus settings.", "type": "object", "properties": {"enabled_download_phase": {"$ref": "#/components/schemas/zero-trust-gateway_enabled_download_phase"}, "enabled_upload_phase": {"$ref": "#/components/schemas/zero-trust-gateway_enabled_upload_phase"}, "fail_closed": {"$ref": "#/components/schemas/zero-trust-gateway_fail_closed"}, "notification_settings": {"$ref": "#/components/schemas/zero-trust-gateway_notification_settings"}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
