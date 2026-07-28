---
title: teams-devices_zero-trust-account-device-settings
page_id: schema-teams-devices-zero-trust-account-device-settings-d396505c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_zero-trust-account-device-settings

```yaml
{"type": "object", "properties": {"disable_for_time": {"description": "Sets the time limit, in seconds, that a user can use an override code to bypass WARP.", "type": "number"}, "external_emergency_signal_enabled": {"description": "Controls whether the external emergency disconnect feature is enabled.", "type": "boolean", "example": true}, "external_emergency_signal_fingerprint": {"description": "The SHA256 fingerprint (64 hexadecimal characters) of the HTTPS server certificate for the external_emergency_signal_url. If provided, the WARP client will use this value to verify the server's identity. The device will ignore any response if the server's certificate fingerprint does not exactly match this value.", "type": "string", "example": "abcd1234567890abcd1234567890abcd1234567890abcd1234567890abcd1234"}, "external_emergency_signal_interval": {"description": "The interval at which the WARP client fetches the emergency disconnect signal, formatted as a duration string (e.g., \"5m\", \"2m30s\", \"1h\"). Minimum 30 seconds.", "type": "string", "example": "5m"}, "external_emergency_signal_url": {"description": "The HTTPS URL from which to fetch the emergency disconnect signal. Must use HTTPS and have an IPv4 or IPv6 address as the host.", "type": "string", "example": "https://192.0.2.1/signal"}, "gateway_proxy_enabled": {"description": "Enable gateway proxy filtering on TCP.", "type": "boolean", "example": true}, "gateway_udp_proxy_enabled": {"description": "Enable gateway proxy filtering on UDP.", "type": "boolean", "example": true}, "root_certificate_installation_enabled": {"description": "Enable installation of cloudflare managed root certificate.", "type": "boolean", "example": true}, "use_zt_virtual_ip": {"description": "Enable using CGNAT virtual IPv4.", "type": "boolean", "example": true}}}
```
