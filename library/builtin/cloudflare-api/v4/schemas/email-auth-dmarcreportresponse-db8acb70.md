---
title: email-auth_DmarcReportResponse
page_id: schema-email-auth-dmarcreportresponse-db8acb70
path: schemas
description: Response for GET/PATCH /dmarc-reports
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_DmarcReportResponse

Response for GET/PATCH /dmarc-reports

```yaml
{"description": "Response for GET/PATCH /dmarc-reports", "type": "object", "properties": {"approved_sources": {"description": "List of approved sending sources (omitted when empty)", "type": "array", "items": {"$ref": "#/components/schemas/email-auth_ApprovedSourceResponse"}}, "created": {"description": "Deprecated, use created_at", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00.12345Z", "deprecated": true, "x-stainless-deprecation-message": "Use `created_at` instead."}, "created_at": {"description": "Creation timestamp", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00.12345Z"}, "enabled": {"description": "Whether DMARC reports are enabled", "type": "boolean", "example": true}, "modified": {"description": "Deprecated, use modified_at", "type": "string", "format": "date-time", "example": "2024-01-15T11:45:00.12345Z", "deprecated": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"description": "Last modification timestamp", "type": "string", "format": "date-time", "example": "2024-01-15T11:45:00.12345Z"}, "records": {"$ref": "#/components/schemas/email-auth_ZoneDnsRecords"}, "rua_prefix": {"description": "Prefix for DMARC RUA addresses (32-char hex string)", "type": "string", "example": "9233c80fc89f43e3a7b749605f651868"}, "skip_wizard": {"description": "Whether to skip the setup wizard", "type": "boolean", "example": false}, "status": {"$ref": "#/components/schemas/email-auth_DmarcStatus"}, "tag": {"description": "Use `zone_id` instead", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "deprecated": true, "x-stainless-deprecation-message": "Use `zone_id` instead."}, "zone_id": {"description": "Zone identifier", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}}
```
