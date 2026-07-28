---
title: zero-trust-gateway_browser-isolation-settings
page_id: schema-zero-trust-gateway-browser-isolation-settings-51586ed2
path: schemas
description: Specify Clientless Browser Isolation settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_browser-isolation-settings

Specify Clientless Browser Isolation settings.

```yaml
{"description": "Specify Clientless Browser Isolation settings.", "type": "object", "properties": {"non_identity_enabled": {"description": "Specify whether to enable non-identity onramp support for Browser Isolation.", "type": "boolean", "example": true, "x-auditable": true}, "url_browser_isolation_enabled": {"description": "Specify whether to enable Clientless Browser Isolation.", "type": "boolean", "example": true, "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
