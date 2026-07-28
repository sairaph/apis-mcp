---
title: dlp_CreateIntegrationBody
page_id: schema-dlp-createintegrationbody-e11fd256
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CreateIntegrationBody

```yaml
{"type": "object", "properties": {"integration_type": {"$ref": "#/components/schemas/dlp_RiskScoreIntegrationType"}, "reference_id": {"description": "A reference id that can be supplied by the client. Currently this should be set to the Access-Okta IDP ID (a UUIDv4).\nhttps://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider", "type": "string", "nullable": true}, "tenant_url": {"description": "The base url of the tenant, e.g. \"https://tenant.okta.com\".", "type": "string", "format": "uri"}}, "required": ["integration_type", "tenant_url"]}
```
