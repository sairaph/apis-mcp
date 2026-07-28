---
title: csam-config-service_csam_scanner_third_party_update_request
page_id: schema-csam-config-service-csam-scanner-third-party-update-request-8f4628ae
path: schemas
description: Request body for updating CSAM Scanner configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# csam-config-service_csam_scanner_third_party_update_request

Request body for updating CSAM Scanner configuration.

```yaml
{"description": "Request body for updating CSAM Scanner configuration.\n", "type": "object", "properties": {"id": {"description": "The feature identifier.", "type": "string", "example": "csam_scanner", "enum": ["csam_scanner"]}, "value": {"$ref": "#/components/schemas/csam-config-service_csam_scanner_third_party_update_value"}}}
```
