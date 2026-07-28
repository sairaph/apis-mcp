---
title: Get DMARC Report Status
page_id: operation-get-zones-zone-id-email-auth-dmarc-reports-b9dd4de5
path: operations/email-auth
description: |-
    Retrieves the current DMARC report configuration and status for a zone.
    Returns the RUA prefix, enabled status, approved sources, and DNS records.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/auth/dmarc-reports
operation_ids:
    - get_dmarc_reports_status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DMARC Report Status

`GET /zones/{zone_id}/email/auth/dmarc-reports`

Operation ID: `get_dmarc_reports_status`

Retrieves the current DMARC report configuration and status for a zone.
Returns the RUA prefix, enabled status, approved sources, and DNS records.

## Definition

```yaml
{"operationId": "get_dmarc_reports_status", "summary": "Get DMARC Report Status", "description": "Retrieves the current DMARC report configuration and status for a zone.\nReturns the RUA prefix, enabled status, approved sources, and DNS records.\n", "parameters": [{"$ref": "#/components/parameters/email-auth_zone_id"}], "responses": {"200": {"description": "DMARC report status retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-auth_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-auth_DmarcReportResponse"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-auth_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Auth"]}
```
