---
title: Configure DMARC Reports
page_id: operation-patch-zones-zone-id-email-auth-dmarc-reports-172fc6a8
path: operations/email-auth
description: |-
    Updates the DMARC report configuration for a zone.
    At least one of `enabled` or `skip_wizard` must be provided.
    When enabling, the handler will ensure the DMARC RUA record exists in DNS.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/email/auth/dmarc-reports
operation_ids:
    - configure_dmarc_reports
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Configure DMARC Reports

`PATCH /zones/{zone_id}/email/auth/dmarc-reports`

Operation ID: `configure_dmarc_reports`

Updates the DMARC report configuration for a zone.
At least one of `enabled` or `skip_wizard` must be provided.
When enabling, the handler will ensure the DMARC RUA record exists in DNS.

## Definition

```yaml
{"operationId": "configure_dmarc_reports", "summary": "Configure DMARC Reports", "description": "Updates the DMARC report configuration for a zone.\nAt least one of `enabled` or `skip_wizard` must be provided.\nWhen enabling, the handler will ensure the DMARC RUA record exists in DNS.\n", "parameters": [{"$ref": "#/components/parameters/email-auth_zone_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-auth_ConfigureDmarcReportsRequest"}}}}, "responses": {"200": {"description": "DMARC report configuration updated successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-auth_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-auth_DmarcReportResponse"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-auth_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Auth"]}
```
