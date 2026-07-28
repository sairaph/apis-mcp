---
title: access_app_launcher_props
page_id: schema-access-app-launcher-props-52d2fa43
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_app_launcher_props

```yaml
{"allOf": [{"$ref": "#/components/schemas/access_feature_app_props"}, {"properties": {"app_launcher_logo_url": {"$ref": "#/components/schemas/access_app_launcher_logo_url"}, "bg_color": {"$ref": "#/components/schemas/access_bg_color"}, "domain": {"example": "authdomain.cloudflareaccess.com", "readOnly": true}, "footer_links": {"$ref": "#/components/schemas/access_footer_links"}, "header_bg_color": {"$ref": "#/components/schemas/access_header_bg_color"}, "landing_page_design": {"$ref": "#/components/schemas/access_landing_page_design"}, "name": {"example": "App Launcher", "default": "App Launcher", "readOnly": true}, "skip_app_launcher_login_page": {"$ref": "#/components/schemas/access_skip_app_launcher_login_page"}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "app_launcher"}]}}}]}
```
