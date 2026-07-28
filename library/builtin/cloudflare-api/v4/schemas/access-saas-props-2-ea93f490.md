---
title: access_saas_props-2
page_id: schema-access-saas-props-2-ea93f490
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saas_props-2

```yaml
{"type": "object", "properties": {"allowed_idps": {"$ref": "#/components/schemas/access_allowed_idps"}, "app_launcher_visible": {"$ref": "#/components/schemas/access_app_launcher_visible"}, "auto_redirect_to_identity": {"$ref": "#/components/schemas/access_auto_redirect_to_identity-2"}, "logo_url": {"$ref": "#/components/schemas/access_logo_url"}, "name": {"$ref": "#/components/schemas/access_name-8"}, "saas_app": {"type": "object", "oneOf": [{"$ref": "#/components/schemas/access_saml_saas_app-2"}, {"$ref": "#/components/schemas/access_oidc_saas_app-2"}]}, "type": {"description": "The application type.", "type": "string", "example": "saas"}}, "title": "SaaS Application"}
```
