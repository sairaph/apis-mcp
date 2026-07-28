---
title: csam-config-service_csam_scanner_third_party_update_value
page_id: schema-csam-config-service-csam-scanner-third-party-update-value-299a406e
path: schemas
description: Writable CSAM Scanner feature configuration values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# csam-config-service_csam_scanner_third_party_update_value

Writable CSAM Scanner feature configuration values.

```yaml
{"description": "Writable CSAM Scanner feature configuration values.\n", "type": "object", "properties": {"email": {"description": "Notification email address for CSAM scan results. When changed,\nemail verification is triggered automatically.\n", "type": "string", "example": "user@example.com", "maxLength": 254}, "enabled": {"description": "Whether CSAM scanning is enabled for this zone.", "type": "boolean", "example": true}, "resend_email": {"description": "Set to true to trigger re-sending the email verification.\nWrite-only; never appears in responses (omitted when false).\n", "type": "boolean", "example": true}, "sources": {"description": "Map of scanning sources and their enabled state.", "type": "object", "example": {"source1": true}, "additionalProperties": {"type": "boolean"}}}}
```
