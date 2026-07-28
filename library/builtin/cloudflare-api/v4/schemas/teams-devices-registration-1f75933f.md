---
title: teams-devices_registration
page_id: schema-teams-devices-registration-1f75933f
path: schemas
description: A WARP configuration tied to a single user. Multiple registrations can be created from a single WARP device.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_registration

A WARP configuration tied to a single user. Multiple registrations can be created from a single WARP device.

```yaml
{"description": "A WARP configuration tied to a single user. Multiple registrations can be created from a single WARP device.", "type": "object", "properties": {"created_at": {"description": "The RFC3339 timestamp when the registration was created.", "type": "string", "example": "2025-02-14T13:17:00Z", "x-auditable": true}, "deleted_at": {"description": "The RFC3339 timestamp when the registration was deleted.", "type": "string", "example": "2025-02-14T13:17:00Z", "nullable": true, "x-auditable": true}, "device": {"$ref": "#/components/schemas/teams-devices_registration_device_details"}, "id": {"description": "The ID of the registration.", "type": "string", "example": "11ffb86f-3f0c-4306-b4a2-e62f872b166a", "x-auditable": true}, "key": {"description": "The public key used to connect to the Cloudflare network.", "type": "string", "example": "U+QTP50RsWfeLGHF4tlGDnmGeuwtsz46KCHr5OyhWq00Rsdfl45mgnQAuEJ6CO0YrkyTl9FUf5iB0bwYR3g4EEFEHhtu6jFaqfMrBMBSz6itv9HQXkaR9OieKQ==", "x-auditable": true}, "key_type": {"description": "The type of encryption key used by the WARP client for the active key. Currently 'curve25519' for WireGuard and 'secp256r1' for MASQUE.", "type": "string", "example": "secp256r1", "nullable": true, "x-auditable": true}, "last_seen_at": {"description": "The RFC3339 timestamp when the registration was last seen.", "type": "string", "example": "2025-02-14T13:17:00Z", "x-auditable": true}, "policy": {"$ref": "#/components/schemas/teams-devices_policy_summary"}, "revoked_at": {"description": "The RFC3339 timestamp when the registration was revoked.", "type": "string", "example": "2025-02-14T13:17:00Z", "nullable": true, "x-auditable": true}, "tunnel_type": {"description": "Type of the tunnel - wireguard or masque.", "type": "string", "example": "masque", "nullable": true, "x-auditable": true}, "updated_at": {"description": "The RFC3339 timestamp when the registration was last updated.", "type": "string", "example": "2025-02-14T13:17:00Z", "x-auditable": true}, "user": {"$ref": "#/components/schemas/teams-devices_user"}, "virtual_ipv4": {"description": "The virtual IPv4 address assigned to the network interface of the tunnel for this registration.", "type": "string", "example": "100.96.0.1", "nullable": true, "x-auditable": true}, "virtual_ipv6": {"description": "The virtual IPv6 address assigned to the network interface of the tunnel for this registration.", "type": "string", "example": "2606:4700:0cf1:1000::1", "nullable": true, "x-auditable": true}}, "required": ["id", "key", "device", "created_at", "updated_at", "last_seen_at"]}
```
