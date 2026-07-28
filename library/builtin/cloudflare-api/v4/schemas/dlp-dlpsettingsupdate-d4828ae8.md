---
title: dlp_DlpSettingsUpdate
page_id: schema-dlp-dlpsettingsupdate-d4828ae8
path: schemas
description: |-
    Request model for updating DLP account-level settings.
    All fields are optional. Missing fields behave differently for PUT vs PATCH:
    - PUT: missing fields reset to initial (unconfigured) values.
    - PATCH: missing fields keep existing values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DlpSettingsUpdate

Request model for updating DLP account-level settings.
All fields are optional. Missing fields behave differently for PUT vs PATCH:
- PUT: missing fields reset to initial (unconfigured) values.
- PATCH: missing fields keep existing values.

```yaml
{"description": "Request model for updating DLP account-level settings.\nAll fields are optional. Missing fields behave differently for PUT vs PATCH:\n- PUT: missing fields reset to initial (unconfigured) values.\n- PATCH: missing fields keep existing values.", "type": "object", "properties": {"ai_context_analysis": {"description": "Whether AI context analysis is enabled at the account level.", "type": "boolean", "default": false, "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}, "ocr": {"description": "Whether OCR is enabled at the account level.", "type": "boolean", "default": false, "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}, "payload_logging": {"allOf": [{"$ref": "#/components/schemas/dlp_PayloadLogSettingUpdate"}], "x-stainless-terraform-configurability": "computed_optional"}}}
```
