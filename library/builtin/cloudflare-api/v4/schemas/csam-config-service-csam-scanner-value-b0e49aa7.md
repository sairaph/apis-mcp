---
title: csam-config-service_csam_scanner_value
page_id: schema-csam-config-service-csam-scanner-value-b0e49aa7
path: schemas
description: |-
    The CSAM Scanner feature configuration values. Contains the
    notification email and scanning enablement settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# csam-config-service_csam_scanner_value

The CSAM Scanner feature configuration values. Contains the
notification email and scanning enablement settings.

```yaml
{"description": "The CSAM Scanner feature configuration values. Contains the\nnotification email and scanning enablement settings.\n", "type": "object", "properties": {"email": {"description": "Notification email address for CSAM scan results. Masked in\nresponses unless explicitly unmasked via admin endpoint.\n", "type": "string", "example": "**********", "maxLength": 254}, "email_state": {"description": "Current verification state of the notification email.\n", "type": "string", "example": "valid", "enum": ["valid", "pending", "unverified"]}, "enabled": {"description": "Whether CSAM scanning is enabled for this zone.", "type": "boolean", "example": true}, "sources": {"description": "Map of scanning sources and their enabled state.", "type": "object", "example": {"source1": true}, "additionalProperties": {"type": "boolean"}}, "zone_plan": {"description": "The zone's plan level.", "type": "string", "example": "ent", "readOnly": true}}}
```
