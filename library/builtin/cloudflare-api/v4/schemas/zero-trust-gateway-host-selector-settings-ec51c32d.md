---
title: zero-trust-gateway_host-selector-settings
page_id: schema-zero-trust-gateway-host-selector-settings-ec51c32d
path: schemas
description: Enable host selection in egress policies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_host-selector-settings

Enable host selection in egress policies.

```yaml
{"description": "Enable host selection in egress policies.", "type": "object", "properties": {"enabled": {"description": "Specify whether to enable filtering via hosts for egress policies.", "type": "boolean", "example": false, "nullable": true, "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
