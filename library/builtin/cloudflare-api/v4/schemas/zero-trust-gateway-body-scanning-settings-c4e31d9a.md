---
title: zero-trust-gateway_body-scanning-settings
page_id: schema-zero-trust-gateway-body-scanning-settings-c4e31d9a
path: schemas
description: Specify the DLP inspection mode.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_body-scanning-settings

Specify the DLP inspection mode.

```yaml
{"description": "Specify the DLP inspection mode.", "type": "object", "properties": {"inspection_mode": {"description": "Specify the inspection mode as either `deep` or `shallow`.", "type": "string", "example": "deep", "enum": ["deep", "shallow"], "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
