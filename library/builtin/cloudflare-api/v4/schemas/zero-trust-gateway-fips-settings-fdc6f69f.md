---
title: zero-trust-gateway_fips-settings
page_id: schema-zero-trust-gateway-fips-settings-fdc6f69f
path: schemas
description: Specify FIPS settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_fips-settings

Specify FIPS settings.

```yaml
{"description": "Specify FIPS settings.", "type": "object", "properties": {"tls": {"description": "Enforce cipher suites and TLS versions compliant with FIPS 140-2.", "type": "boolean", "example": true, "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
