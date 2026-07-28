---
title: zones_settings
page_id: schema-zones-settings-fbe6a331
path: schemas
description: Settings available for the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_settings

Settings available for the zone.

```yaml
{"description": "Settings available for the zone.", "type": "array", "items": {"type": "object"}, "example": [{"id": "browser_check", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "browser_cache_ttl", "properties": [{"type": "range", "max": 31536000, "min": 1800, "name": "value", "suggested_values": [1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000]}]}, {"id": "browser_check", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "cache_key_fields", "properties": [{"type": "object", "name": "value", "properties": [{"type": "select", "allowEmpty": true, "choices": ["include", "exclude"], "multiple": false, "name": "query_string"}, {"type": "select", "allowEmpty": true, "choices": ["include", "exclude", "check_presence"], "multiple": true, "name": "header"}, {"type": "select", "allowEmpty": false, "choices": ["resolved"], "multiple": true, "name": "host"}, {"type": "select", "allowEmpty": true, "choices": ["include", "check_presence"], "multiple": true, "name": "cookie"}, {"type": "select", "allowEmpty": false, "choices": ["device_type", "geo", "lang"], "multiple": true, "name": "user"}]}]}, {"id": "cache_deception_armor", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "cache_level", "properties": [{"type": "select", "choices": ["bypass", "basic", "simplified", "aggressive", "cache_everything"], "multiple": false, "name": "value"}]}, {"id": "cache_ttl_by_status", "properties": [{"type": "object", "allowEmpty": false, "name": "value"}]}, {"id": "disable_apps", "properties": []}, {"id": "disable_performance", "properties": []}, {"id": "disable_security", "properties": []}, {"id": "edge_cache_ttl", "properties": [{"type": "range", "max": 2419200, "min": 7200, "name": "value", "suggested_values": [7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 518400, 604800, 1209600, 2419200]}]}, {"id": "email_obfuscation", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "forwarding_url", "properties": [{"type": "choice", "choices": [301, 302], "multiple": false, "name": "status_code"}, {"type": "forwardingUrl", "name": "url"}]}, {"id": "ip_geolocation", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "explicit_cache_control", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "rocket_loader", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "security_level", "properties": [{"type": "select", "choices": ["essentially_off", "low", "medium", "high", "under_attack"], "multiple": false, "name": "value"}]}, {"id": "server_side_exclude", "properties": [{"type": "toggle", "name": "value"}]}, {"id": "ssl", "properties": [{"type": "choice", "choices": ["off", "flexible", "full", "strict"], "multiple": false, "name": "value"}]}]}
```
