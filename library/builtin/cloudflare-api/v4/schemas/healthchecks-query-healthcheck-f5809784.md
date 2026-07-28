---
title: healthchecks_query_healthcheck
page_id: schema-healthchecks-query-healthcheck-f5809784
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# healthchecks_query_healthcheck

```yaml
{"type": "object", "properties": {"address": {"$ref": "#/components/schemas/healthchecks_address"}, "check_regions": {"$ref": "#/components/schemas/healthchecks_check_regions"}, "consecutive_fails": {"$ref": "#/components/schemas/healthchecks_consecutive_fails"}, "consecutive_successes": {"$ref": "#/components/schemas/healthchecks_consecutive_successes"}, "description": {"$ref": "#/components/schemas/healthchecks_description"}, "http_config": {"$ref": "#/components/schemas/healthchecks_http_config"}, "interval": {"$ref": "#/components/schemas/healthchecks_interval"}, "name": {"$ref": "#/components/schemas/healthchecks_name"}, "retries": {"$ref": "#/components/schemas/healthchecks_retries"}, "suspended": {"$ref": "#/components/schemas/healthchecks_suspended"}, "tcp_config": {"$ref": "#/components/schemas/healthchecks_tcp_config"}, "timeout": {"$ref": "#/components/schemas/healthchecks_timeout"}, "type": {"$ref": "#/components/schemas/healthchecks_type"}}, "required": ["name", "address"]}
```
