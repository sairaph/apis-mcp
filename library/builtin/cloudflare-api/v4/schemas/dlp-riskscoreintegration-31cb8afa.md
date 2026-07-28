---
title: dlp_RiskScoreIntegration
page_id: schema-dlp-riskscoreintegration-31cb8afa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_RiskScoreIntegration

```yaml
{"type": "object", "properties": {"account_tag": {"description": "The Cloudflare account tag.", "type": "string"}, "active": {"description": "Whether this integration is enabled and should export changes in risk score.", "type": "boolean"}, "created_at": {"description": "When the integration was created in RFC3339 format.", "type": "string", "format": "date-time"}, "id": {"description": "The id of the integration, a UUIDv4.", "type": "string", "format": "uuid"}, "integration_type": {"$ref": "#/components/schemas/dlp_RiskScoreIntegrationType"}, "reference_id": {"description": "A reference ID defined by the client.\nShould be set to the Access-Okta IDP integration ID.\nUseful when the risk-score integration needs to be associated with a secondary asset and recalled using that ID.", "type": "string"}, "tenant_url": {"description": "The base URL for the tenant. E.g. \"https://tenant.okta.com\".", "type": "string"}, "well_known_url": {"description": "The URL for the Shared Signals Framework configuration, e.g. \"/.well-known/sse-configuration/{integration_uuid}/\". https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1.", "type": "string"}}, "required": ["id", "account_tag", "integration_type", "reference_id", "tenant_url", "well_known_url", "active", "created_at"]}
```
