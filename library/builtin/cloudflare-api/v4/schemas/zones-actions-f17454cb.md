---
title: zones_actions
page_id: schema-zones-actions-f17454cb
path: schemas
description: |-
    The set of actions to perform if the targets of this rule match the
    request. Actions can redirect to another URL or override settings, but
    not both.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_actions

The set of actions to perform if the targets of this rule match the
request. Actions can redirect to another URL or override settings, but
not both.

```yaml
{"description": "The set of actions to perform if the targets of this rule match the\nrequest. Actions can redirect to another URL or override settings, but\nnot both.\n", "type": "array", "items": {"discriminator": {"propertyName": "id"}, "oneOf": [{"$ref": "#/components/schemas/zones_always_use_https"}, {"$ref": "#/components/schemas/zones_automatic_https_rewrites"}, {"$ref": "#/components/schemas/zones_browser_cache_ttl"}, {"$ref": "#/components/schemas/zones_browser_check"}, {"$ref": "#/components/schemas/zones_bypass_cache_on_cookie"}, {"$ref": "#/components/schemas/zones_cache_by_device_type"}, {"$ref": "#/components/schemas/zones_cache_deception_armor"}, {"$ref": "#/components/schemas/zones_cache_key_fields"}, {"$ref": "#/components/schemas/zones_cache_level"}, {"$ref": "#/components/schemas/zones_cache_on_cookie"}, {"$ref": "#/components/schemas/zones_cache_ttl_by_status"}, {"$ref": "#/components/schemas/zones_disable_apps"}, {"$ref": "#/components/schemas/zones_disable_performance"}, {"$ref": "#/components/schemas/zones_disable_security"}, {"$ref": "#/components/schemas/zones_disable_zaraz"}, {"$ref": "#/components/schemas/zones_edge_cache_ttl"}, {"$ref": "#/components/schemas/zones_email_obfuscation"}, {"$ref": "#/components/schemas/zones_explicit_cache_control"}, {"$ref": "#/components/schemas/zones_forwarding_url"}, {"$ref": "#/components/schemas/zones_host_header_override"}, {"$ref": "#/components/schemas/zones_ip_geolocation"}, {"$ref": "#/components/schemas/zones_mirage"}, {"$ref": "#/components/schemas/zones_opportunistic_encryption"}, {"$ref": "#/components/schemas/zones_origin_error_page_pass_thru"}, {"$ref": "#/components/schemas/zones_polish"}, {"$ref": "#/components/schemas/zones_resolve_override"}, {"$ref": "#/components/schemas/zones_respect_strong_etag"}, {"$ref": "#/components/schemas/zones_response_buffering"}, {"$ref": "#/components/schemas/zones_rocket_loader"}, {"$ref": "#/components/schemas/zones_security_level"}, {"$ref": "#/components/schemas/zones_sort_query_string_for_cache"}, {"$ref": "#/components/schemas/zones_ssl"}, {"$ref": "#/components/schemas/zones_true_client_ip_header"}, {"$ref": "#/components/schemas/zones_waf"}]}, "example": [{"id": "browser_check", "value": "on"}], "x-stainless-skip": ["terraform"]}
```
