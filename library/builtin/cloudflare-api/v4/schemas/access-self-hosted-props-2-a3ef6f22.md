---
title: access_self_hosted_props-2
page_id: schema-access-self-hosted-props-2-a3ef6f22
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_self_hosted_props-2

```yaml
{"type": "object", "properties": {"allow_iframe": {"$ref": "#/components/schemas/access_allow_iframe"}, "allowed_idps": {"$ref": "#/components/schemas/access_allowed_idps"}, "app_launcher_visible": {"$ref": "#/components/schemas/access_app_launcher_visible"}, "auto_redirect_to_identity": {"$ref": "#/components/schemas/access_auto_redirect_to_identity-2"}, "cors_headers": {"$ref": "#/components/schemas/access_cors_headers-2"}, "custom_deny_message": {"$ref": "#/components/schemas/access_custom_deny_message"}, "custom_deny_url": {"$ref": "#/components/schemas/access_custom_deny_url-2"}, "domain": {"$ref": "#/components/schemas/access_domain-3"}, "eager_redirect_cookie_setting": {"$ref": "#/components/schemas/access_eager_redirect_cookie_setting"}, "enable_binding_cookie": {"$ref": "#/components/schemas/access_enable_binding_cookie"}, "http_only_cookie_attribute": {"$ref": "#/components/schemas/access_http_only_cookie_attribute"}, "logo_url": {"$ref": "#/components/schemas/access_logo_url"}, "name": {"$ref": "#/components/schemas/access_name-8"}, "options_preflight_bypass": {"$ref": "#/components/schemas/access_options_preflight_bypass-2"}, "same_site_cookie_attribute": {"$ref": "#/components/schemas/access_same_site_cookie_attribute"}, "service_auth_401_redirect": {"$ref": "#/components/schemas/access_service_auth_401_redirect"}, "session_duration": {"$ref": "#/components/schemas/access_session_duration-4"}, "skip_interstitial": {"$ref": "#/components/schemas/access_skip_interstitial"}, "type": {"description": "The application type.", "type": "string", "example": "self_hosted"}, "use_clientless_isolation_app_launcher_url": {"$ref": "#/components/schemas/access_use_clientless_isolation_app_launcher_url"}}, "required": ["type", "domain"], "title": "Self Hosted Application"}
```
