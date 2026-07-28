---
title: Create a device settings profile
page_id: operation-post-accounts-account-id-devices-policy-a40b09a5
path: operations/devices
description: Creates a device settings profile to be applied to certain devices matching the criteria.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/policy
operation_ids:
    - devices-create-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a device settings profile

`POST /accounts/{account_id}/devices/policy`

Operation ID: `devices-create-device-settings-policy`

Creates a device settings profile to be applied to certain devices matching the criteria.

## Definition

```yaml
{"operationId": "devices-create-device-settings-policy", "summary": "Create a device settings profile", "description": "Creates a device settings profile to be applied to certain devices matching the criteria.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"allow_mode_switch": {"$ref": "#/components/schemas/teams-devices_allow_mode_switch"}, "allow_updates": {"$ref": "#/components/schemas/teams-devices_allow_updates"}, "allowed_to_leave": {"$ref": "#/components/schemas/teams-devices_allowed_to_leave"}, "auto_connect": {"$ref": "#/components/schemas/teams-devices_auto_connect"}, "captive_portal": {"$ref": "#/components/schemas/teams-devices_captive_portal"}, "description": {"allOf": [{"$ref": "#/components/schemas/teams-devices_schemas-description"}, {"default": ""}]}, "disable_auto_fallback": {"$ref": "#/components/schemas/teams-devices_disable_auto_fallback"}, "dns_search_suffixes": {"$ref": "#/components/schemas/teams-devices_dns_search_suffixes"}, "enabled": {"description": "Whether the policy will be applied to matching devices.", "type": "boolean", "example": true, "default": true}, "exclude": {"$ref": "#/components/schemas/teams-devices_exclude_request"}, "exclude_office_ips": {"$ref": "#/components/schemas/teams-devices_exclude_office_ips"}, "global_acceleration": {"$ref": "#/components/schemas/teams-devices_global_acceleration"}, "include": {"$ref": "#/components/schemas/teams-devices_include_request"}, "lan_allow_minutes": {"$ref": "#/components/schemas/teams-devices_lan_allow_minutes"}, "lan_allow_subnet_size": {"$ref": "#/components/schemas/teams-devices_lan_allow_subnet_size"}, "match": {"$ref": "#/components/schemas/teams-devices_schemas-match"}, "name": {"description": "The name of the device settings profile.", "type": "string", "example": "Allow Developers", "maxLength": 100}, "precedence": {"$ref": "#/components/schemas/teams-devices_precedence"}, "register_interface_ip_with_dns": {"$ref": "#/components/schemas/teams-devices_register_interface_ip_with_dns"}, "sccm_vpn_boundary_support": {"$ref": "#/components/schemas/teams-devices_sccm_vpn_boundary_support"}, "service_mode_v2": {"$ref": "#/components/schemas/teams-devices_service_mode_v2"}, "support_url": {"$ref": "#/components/schemas/teams-devices_support_url"}, "switch_locked": {"$ref": "#/components/schemas/teams-devices_switch_locked"}, "tunnel_protocol": {"$ref": "#/components/schemas/teams-devices_tunnel_protocol"}, "virtual_networks": {"$ref": "#/components/schemas/teams-devices_virtual_networks"}}, "required": ["name", "precedence", "match"]}}}}, "responses": {"200": {"description": "Create a device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_device_settings_response"}}}}, "4XX": {"description": "Create a device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_device_settings_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
