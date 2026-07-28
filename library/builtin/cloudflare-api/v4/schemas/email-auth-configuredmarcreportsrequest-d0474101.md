---
title: email-auth_ConfigureDmarcReportsRequest
page_id: schema-email-auth-configuredmarcreportsrequest-d0474101
path: schemas
description: Request body for PATCH /dmarc-reports
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_ConfigureDmarcReportsRequest

Request body for PATCH /dmarc-reports

```yaml
{"description": "Request body for PATCH /dmarc-reports", "type": "object", "properties": {"enabled": {"description": "Enable or disable DMARC reports for this zone", "type": "boolean", "example": true, "nullable": true}, "skip_wizard": {"description": "Skip the DMARC setup wizard", "type": "boolean", "example": false, "nullable": true}}}
```
