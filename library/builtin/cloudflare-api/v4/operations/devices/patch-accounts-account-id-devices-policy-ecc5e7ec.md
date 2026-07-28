---
title: Update the default device settings profile
page_id: operation-patch-accounts-account-id-devices-policy-f0b7b28b
path: operations/devices
description: Updates the default device settings profile for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/devices/policy
operation_ids:
    - devices-update-default-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the default device settings profile

`PATCH /accounts/{account_id}/devices/policy`

Operation ID: `devices-update-default-device-settings-policy`

Updates the default device settings profile for an account.

## Definition

```yaml
{"operationId": "devices-update-default-device-settings-policy", "summary": "Update the default device settings profile", "description": "Updates the default device settings profile for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"allow_mode_switch": {"$ref": "#/components/schemas/teams-devices_allow_mode_switch"}, "allow_updates": {"$ref": "#/components/schemas/teams-devices_allow_updates"}, "allowed_to_leave": {"$ref": "#/components/schemas/teams-devices_allowed_to_leave"}, "auto_connect": {"$ref": "#/components/schemas/teams-devices_auto_connect"}, "captive_portal": {"$ref": "#/components/schemas/teams-devices_captive_portal"}, "disable_auto_fallback": {"$ref": "#/components/schemas/teams-devices_disable_auto_fallback"}, "dns_search_suffixes": {"$ref": "#/components/schemas/teams-devices_dns_search_suffixes"}, "exclude": {"$ref": "#/components/schemas/teams-devices_exclude_request"}, "exclude_office_ips": {"$ref": "#/components/schemas/teams-devices_exclude_office_ips"}, "global_acceleration": {"$ref": "#/components/schemas/teams-devices_global_acceleration"}, "include": {"$ref": "#/components/schemas/teams-devices_include_request"}, "lan_allow_minutes": {"$ref": "#/components/schemas/teams-devices_lan_allow_minutes"}, "lan_allow_subnet_size": {"$ref": "#/components/schemas/teams-devices_lan_allow_subnet_size"}, "register_interface_ip_with_dns": {"$ref": "#/components/schemas/teams-devices_register_interface_ip_with_dns"}, "sccm_vpn_boundary_support": {"$ref": "#/components/schemas/teams-devices_sccm_vpn_boundary_support"}, "service_mode_v2": {"$ref": "#/components/schemas/teams-devices_service_mode_v2"}, "support_url": {"$ref": "#/components/schemas/teams-devices_support_url"}, "switch_locked": {"$ref": "#/components/schemas/teams-devices_switch_locked"}, "tunnel_protocol": {"$ref": "#/components/schemas/teams-devices_tunnel_protocol"}, "virtual_networks": {"$ref": "#/components/schemas/teams-devices_virtual_networks"}}}}}}, "responses": {"200": {"description": "Update the default device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_default_device_settings_response"}}}}, "4XX": {"description": "Update the default device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_default_device_settings_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
