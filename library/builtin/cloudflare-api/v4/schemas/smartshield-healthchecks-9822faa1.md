---
title: smartshield_healthchecks
page_id: schema-smartshield-healthchecks-9822faa1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_healthchecks

```yaml
{"type": "object", "properties": {"address": {"$ref": "#/components/schemas/smartshield_address"}, "check_regions": {"$ref": "#/components/schemas/smartshield_check_regions"}, "consecutive_fails": {"$ref": "#/components/schemas/smartshield_consecutive_fails"}, "consecutive_successes": {"$ref": "#/components/schemas/smartshield_consecutive_successes"}, "created_on": {"$ref": "#/components/schemas/smartshield_timestamp"}, "description": {"$ref": "#/components/schemas/smartshield_description"}, "failure_reason": {"$ref": "#/components/schemas/smartshield_failure_reason"}, "http_config": {"$ref": "#/components/schemas/smartshield_http_config"}, "id": {"$ref": "#/components/schemas/smartshield_identifier"}, "interval": {"$ref": "#/components/schemas/smartshield_interval"}, "modified_on": {"$ref": "#/components/schemas/smartshield_timestamp"}, "name": {"$ref": "#/components/schemas/smartshield_name"}, "retries": {"$ref": "#/components/schemas/smartshield_retries"}, "status": {"$ref": "#/components/schemas/smartshield_status"}, "suspended": {"$ref": "#/components/schemas/smartshield_suspended"}, "tcp_config": {"$ref": "#/components/schemas/smartshield_tcp_config"}, "timeout": {"$ref": "#/components/schemas/smartshield_timeout"}, "type": {"$ref": "#/components/schemas/smartshield_type"}}}
```
