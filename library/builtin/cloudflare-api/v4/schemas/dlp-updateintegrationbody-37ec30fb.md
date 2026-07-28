---
title: dlp_UpdateIntegrationBody
page_id: schema-dlp-updateintegrationbody-37ec30fb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_UpdateIntegrationBody

```yaml
{"type": "object", "properties": {"active": {"description": "Whether this integration is enabled. If disabled, no risk changes will be exported to the third-party.", "type": "boolean"}, "reference_id": {"description": "A reference id that can be supplied by the client. Currently this should be set to the Access-Okta IDP ID (a UUIDv4).\nhttps://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider", "type": "string", "nullable": true}, "tenant_url": {"description": "The base url of the tenant, e.g. \"https://tenant.okta.com\".", "type": "string", "format": "uri"}}, "required": ["tenant_url", "active"]}
```
