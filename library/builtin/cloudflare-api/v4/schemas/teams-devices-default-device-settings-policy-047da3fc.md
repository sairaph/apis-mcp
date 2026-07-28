---
title: teams-devices_default_device_settings_policy
page_id: schema-teams-devices-default-device-settings-policy-047da3fc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_default_device_settings_policy

```yaml
{"type": "object", "properties": {"allow_mode_switch": {"$ref": "#/components/schemas/teams-devices_allow_mode_switch"}, "allow_updates": {"$ref": "#/components/schemas/teams-devices_allow_updates"}, "allowed_to_leave": {"$ref": "#/components/schemas/teams-devices_allowed_to_leave"}, "auto_connect": {"$ref": "#/components/schemas/teams-devices_auto_connect"}, "captive_portal": {"$ref": "#/components/schemas/teams-devices_captive_portal"}, "default": {"description": "Whether the policy will be applied to matching devices.", "type": "boolean", "example": true}, "disable_auto_fallback": {"$ref": "#/components/schemas/teams-devices_disable_auto_fallback"}, "dns_search_suffixes": {"$ref": "#/components/schemas/teams-devices_dns_search_suffixes"}, "enabled": {"description": "Whether the policy will be applied to matching devices.", "type": "boolean", "example": true, "default": true}, "exclude": {"$ref": "#/components/schemas/teams-devices_exclude"}, "exclude_office_ips": {"$ref": "#/components/schemas/teams-devices_exclude_office_ips"}, "fallback_domains": {"$ref": "#/components/schemas/teams-devices_fallback_domains"}, "gateway_unique_id": {"$ref": "#/components/schemas/teams-devices_gateway_unique_id"}, "global_acceleration": {"$ref": "#/components/schemas/teams-devices_global_acceleration"}, "include": {"$ref": "#/components/schemas/teams-devices_include"}, "policy_id": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}, "register_interface_ip_with_dns": {"$ref": "#/components/schemas/teams-devices_register_interface_ip_with_dns"}, "sccm_vpn_boundary_support": {"$ref": "#/components/schemas/teams-devices_sccm_vpn_boundary_support"}, "service_mode_v2": {"$ref": "#/components/schemas/teams-devices_service_mode_v2"}, "support_url": {"$ref": "#/components/schemas/teams-devices_support_url"}, "switch_locked": {"$ref": "#/components/schemas/teams-devices_switch_locked"}, "tunnel_protocol": {"$ref": "#/components/schemas/teams-devices_tunnel_protocol"}, "virtual_networks": {"$ref": "#/components/schemas/teams-devices_virtual_networks"}}}
```
