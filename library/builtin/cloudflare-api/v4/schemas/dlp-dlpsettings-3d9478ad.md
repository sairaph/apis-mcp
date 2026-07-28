---
title: dlp_DlpSettings
page_id: schema-dlp-dlpsettings-3d9478ad
path: schemas
description: DLP account-level settings response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_DlpSettings

DLP account-level settings response.

```yaml
{"description": "DLP account-level settings response.", "type": "object", "properties": {"ai_context_analysis": {"description": "Whether AI context analysis is enabled at the account level.", "type": "boolean"}, "ocr": {"description": "Whether OCR is enabled at the account level.", "type": "boolean"}, "payload_logging": {"$ref": "#/components/schemas/dlp_PayloadLogSetting"}}, "required": ["payload_logging", "ai_context_analysis", "ocr"]}
```
