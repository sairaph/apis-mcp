---
title: zero-trust-gateway_sandbox
page_id: schema-zero-trust-gateway-sandbox-cb679b6d
path: schemas
description: Specify whether to enable the sandbox.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_sandbox

Specify whether to enable the sandbox.

```yaml
{"description": "Specify whether to enable the sandbox.", "type": "object", "properties": {"enabled": {"description": "Specify whether to enable the sandbox.", "type": "boolean", "example": true, "nullable": true, "x-auditable": true}, "fallback_action": {"description": "Specify the action to take when the system cannot scan the file.", "type": "string", "enum": ["allow", "block"], "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
