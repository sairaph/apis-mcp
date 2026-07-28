---
title: zero-trust-gateway_rules
page_id: schema-zero-trust-gateway-rules-0ecc6242
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_rules

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/zero-trust-gateway_action"}, "created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "deleted_at": {"$ref": "#/components/schemas/zero-trust-gateway_deleted_at"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description-2"}, "device_posture": {"$ref": "#/components/schemas/zero-trust-gateway_device_posture"}, "enabled": {"$ref": "#/components/schemas/zero-trust-gateway_enabled"}, "expiration": {"$ref": "#/components/schemas/zero-trust-gateway_expiration"}, "filters": {"$ref": "#/components/schemas/zero-trust-gateway_filters"}, "id": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}, "identity": {"$ref": "#/components/schemas/zero-trust-gateway_identity"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-3"}, "precedence": {"$ref": "#/components/schemas/zero-trust-gateway_precedence"}, "read_only": {"$ref": "#/components/schemas/zero-trust-gateway_read_only"}, "rule_settings": {"$ref": "#/components/schemas/zero-trust-gateway_rule-settings"}, "schedule": {"$ref": "#/components/schemas/zero-trust-gateway_schedule"}, "sharable": {"$ref": "#/components/schemas/zero-trust-gateway_sharable"}, "source_account": {"$ref": "#/components/schemas/zero-trust-gateway_source_account"}, "traffic": {"$ref": "#/components/schemas/zero-trust-gateway_traffic"}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "version": {"$ref": "#/components/schemas/zero-trust-gateway_version"}, "warning_status": {"$ref": "#/components/schemas/zero-trust-gateway_warning_status"}}, "required": ["name", "precedence", "enabled", "action", "traffic", "filters"]}
```
