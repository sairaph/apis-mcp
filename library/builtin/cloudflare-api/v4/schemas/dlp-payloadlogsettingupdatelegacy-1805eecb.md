---
title: dlp_PayloadLogSettingUpdateLegacy
page_id: schema-dlp-payloadlogsettingupdatelegacy-1805eecb
path: schemas
description: Request model for the legacy payload log settings endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PayloadLogSettingUpdateLegacy

Request model for the legacy payload log settings endpoint.

```yaml
{"description": "Request model for the legacy payload log settings endpoint.", "type": "object", "properties": {"masking_level": {"allOf": [{"$ref": "#/components/schemas/dlp_PayloadLogMaskingLevel"}]}, "public_key": {"description": "Base64-encoded public key for encrypting payload logs.\n\n- Set to null or empty string to disable payload logging.\n- Set to a non-empty base64 string to enable payload logging with the given key.\n\nFor customers with configurable payload masking feature rolled out:\n- If the field is missing, the existing setting will be kept. Note that this is different from setting to null or empty string.\n\nFor all other customers:\n- If the field is missing, the existing setting will be cleared.", "type": "string", "nullable": true}}}
```
