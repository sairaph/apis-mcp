---
title: smartshield_query_healthcheck
page_id: schema-smartshield-query-healthcheck-6259f50d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_query_healthcheck

```yaml
{"type": "object", "properties": {"address": {"$ref": "#/components/schemas/smartshield_address"}, "check_regions": {"$ref": "#/components/schemas/smartshield_check_regions"}, "consecutive_fails": {"$ref": "#/components/schemas/smartshield_consecutive_fails"}, "consecutive_successes": {"$ref": "#/components/schemas/smartshield_consecutive_successes"}, "description": {"$ref": "#/components/schemas/smartshield_description"}, "http_config": {"$ref": "#/components/schemas/smartshield_http_config"}, "interval": {"$ref": "#/components/schemas/smartshield_interval"}, "name": {"$ref": "#/components/schemas/smartshield_name"}, "retries": {"$ref": "#/components/schemas/smartshield_retries"}, "suspended": {"$ref": "#/components/schemas/smartshield_suspended"}, "tcp_config": {"$ref": "#/components/schemas/smartshield_tcp_config"}, "timeout": {"$ref": "#/components/schemas/smartshield_timeout"}, "type": {"$ref": "#/components/schemas/smartshield_type"}}, "required": ["name", "address"]}
```
