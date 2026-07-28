---
title: access_saas_props
page_id: schema-access-saas-props-cdb9a8b7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saas_props

```yaml
{"type": "object", "properties": {"allowed_idps": {"$ref": "#/components/schemas/access_allowed_idps"}, "app_launcher_visible": {"$ref": "#/components/schemas/access_app_launcher_visible"}, "auto_redirect_to_identity": {"$ref": "#/components/schemas/access_auto_redirect_to_identity-2"}, "custom_pages": {"$ref": "#/components/schemas/access_custom_pages-2"}, "logo_url": {"$ref": "#/components/schemas/access_logo_url"}, "name": {"$ref": "#/components/schemas/access_name-8"}, "saas_app": {"type": "object", "oneOf": [{"$ref": "#/components/schemas/access_saml_saas_app"}, {"$ref": "#/components/schemas/access_oidc_saas_app"}]}, "scim_config": {"$ref": "#/components/schemas/access_scim_config"}, "tags": {"$ref": "#/components/schemas/access_tags"}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "saas"}]}}, "title": "SaaS Application"}
```
