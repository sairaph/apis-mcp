---
title: access_google-apps-2
page_id: schema-access-google-apps-2-944fdafd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_google-apps-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config-2"}, {"properties": {"apps_domain": {"description": "Your companies TLD", "type": "string", "example": "mycompany.com"}}, "type": "object"}]}}, "type": "object"}], "title": "Google Workspace"}
```
