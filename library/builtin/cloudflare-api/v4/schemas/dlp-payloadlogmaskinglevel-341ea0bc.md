---
title: dlp_PayloadLogMaskingLevel
page_id: schema-dlp-payloadlogmaskinglevel-341ea0bc
path: schemas
description: |-
    Masking level for payload logs.

    - `full`: The entire payload is masked.
    - `partial`: Only partial payload content is masked.
    - `clear`: No masking is applied to the payload content.
    - `default`: DLP uses its default masking behavior.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PayloadLogMaskingLevel

Masking level for payload logs.

- `full`: The entire payload is masked.
- `partial`: Only partial payload content is masked.
- `clear`: No masking is applied to the payload content.
- `default`: DLP uses its default masking behavior.

```yaml
{"description": "Masking level for payload logs.\n\n- `full`: The entire payload is masked.\n- `partial`: Only partial payload content is masked.\n- `clear`: No masking is applied to the payload content.\n- `default`: DLP uses its default masking behavior.", "type": "string", "enum": ["full", "partial", "clear", "default"]}
```
