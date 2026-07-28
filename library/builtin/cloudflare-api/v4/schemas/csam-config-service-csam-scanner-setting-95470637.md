---
title: csam-config-service_csam_scanner_setting
page_id: schema-csam-config-service-csam-scanner-setting-95470637
path: schemas
description: CSAM Scanner configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# csam-config-service_csam_scanner_setting

CSAM Scanner configuration for a zone.

```yaml
{"description": "CSAM Scanner configuration for a zone.", "type": "object", "properties": {"editable": {"description": "Whether the feature state can be changed. When false, the zone\nor account may be locked by Trust & Safety.\n", "type": "boolean", "example": true}, "id": {"description": "The feature identifier.", "type": "string", "example": "csam_scanner", "enum": ["csam_scanner"]}, "modified_on": {"description": "When the setting was last modified. Currently always null as the\nserver does not populate this field.\n", "type": "string", "format": "date-time", "nullable": true}, "value": {"$ref": "#/components/schemas/csam-config-service_csam_scanner_value"}}}
```
