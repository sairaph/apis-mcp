---
title: dlp_PayloadLogSettingUpdate
page_id: schema-dlp-payloadlogsettingupdate-376bbd75
path: schemas
description: |-
    Request model for payload log settings within the DLP settings endpoint.
    Unlike the legacy endpoint, null and missing are treated identically here
    (both mean "not provided" for PATCH, "reset to default" for PUT).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PayloadLogSettingUpdate

Request model for payload log settings within the DLP settings endpoint.
Unlike the legacy endpoint, null and missing are treated identically here
(both mean "not provided" for PATCH, "reset to default" for PUT).

```yaml
{"description": "Request model for payload log settings within the DLP settings endpoint.\nUnlike the legacy endpoint, null and missing are treated identically here\n(both mean \"not provided\" for PATCH, \"reset to default\" for PUT).", "type": "object", "properties": {"masking_level": {"default": "default", "allOf": [{"$ref": "#/components/schemas/dlp_PayloadLogMaskingLevel"}]}, "public_key": {"description": "Base64-encoded public key for encrypting payload logs.\n\n- Set to a non-empty base64 string to enable payload logging with the given key.\n- Set to an empty string to disable payload logging.\n- Omit or set to null to leave unchanged (PATCH) or reset to disabled (PUT).", "type": "string", "nullable": true}}}
```
