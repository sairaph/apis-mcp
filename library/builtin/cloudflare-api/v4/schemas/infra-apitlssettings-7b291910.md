---
title: infra_ApiTlsSettings
page_id: schema-infra-apitlssettings-7b291910
path: schemas
description: |-
    TLS settings for a connectivity service.

    If omitted, the default mode (`verify_full`) is used.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_ApiTlsSettings

TLS settings for a connectivity service.

If omitted, the default mode (`verify_full`) is used.

```yaml
{"description": "TLS settings for a connectivity service.\n\nIf omitted, the default mode (`verify_full`) is used.", "type": "object", "properties": {"cert_verification_mode": {"description": "TLS certificate verification mode for the connection to the origin.\n\n- `\"verify_full\"` — verify certificate chain and hostname (default)\n- `\"verify_ca\"` — verify certificate chain only, skip hostname check\n- `\"disabled\"` — do not verify the server certificate at all", "type": "string", "example": "verify_full"}}, "required": ["cert_verification_mode"]}
```
