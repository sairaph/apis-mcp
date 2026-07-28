---
title: access_mcp_portal_props
page_id: schema-access-mcp-portal-props-8639b9a3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_mcp_portal_props

```yaml
{"type": "object", "properties": {"allow_authenticate_via_warp": {"$ref": "#/components/schemas/access_allow_authenticate_via_warp-2"}, "allowed_idps": {"$ref": "#/components/schemas/access_allowed_idps"}, "auto_redirect_to_identity": {"$ref": "#/components/schemas/access_auto_redirect_to_identity-2"}, "custom_deny_message": {"$ref": "#/components/schemas/access_custom_deny_message"}, "custom_deny_url": {"$ref": "#/components/schemas/access_custom_deny_url"}, "custom_non_identity_deny_url": {"$ref": "#/components/schemas/access_custom_non_identity_deny_url"}, "custom_pages": {"$ref": "#/components/schemas/access_custom_pages-2"}, "destinations": {"$ref": "#/components/schemas/access_destinations"}, "domain": {"$ref": "#/components/schemas/access_domain"}, "http_only_cookie_attribute": {"$ref": "#/components/schemas/access_http_only_cookie_attribute"}, "logo_url": {"$ref": "#/components/schemas/access_logo_url"}, "name": {"$ref": "#/components/schemas/access_name-8"}, "oauth_configuration": {"$ref": "#/components/schemas/access_oauth_configuration"}, "options_preflight_bypass": {"$ref": "#/components/schemas/access_options_preflight_bypass"}, "same_site_cookie_attribute": {"$ref": "#/components/schemas/access_same_site_cookie_attribute"}, "scim_config": {"$ref": "#/components/schemas/access_scim_config"}, "session_duration": {"$ref": "#/components/schemas/access_session_duration-2"}, "tags": {"$ref": "#/components/schemas/access_tags"}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "mcp_portal"}]}}, "required": ["type"], "title": "MCP Server Portal Application"}
```
