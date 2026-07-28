---
title: magic_dhcp_option
page_id: schema-magic-dhcp-option-4280b988
path: schemas
description: A custom DHCP option to include in DHCP responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_dhcp_option

A custom DHCP option to include in DHCP responses.

```yaml
{"description": "A custom DHCP option to include in DHCP responses.", "type": "object", "properties": {"code": {"description": "DHCP option number (1-254). Options 0 and 255 are reserved by RFC 2132. Options 3, 6, and 51 are not allowed because they conflict with connector-managed configuration.", "type": "integer", "maximum": 254, "minimum": 1}, "type": {"description": "The type of the option value. text: a string (max 255 bytes). hex: colon-separated hex bytes (e.g. \"01:04:aa:bb:cc\", max 255 bytes). ip: an IPv4 address (e.g. \"10.20.30.40\"). byte: an unsigned integer 0-255 (1 byte). short: an unsigned integer 0-65535 (2 bytes). integer: an unsigned integer 0-4294967295 (4 bytes).\n", "type": "string", "enum": ["text", "hex", "ip", "byte", "short", "integer"]}, "value": {"description": "The option value, interpreted according to the type field.", "type": "string"}}, "example": {"code": 66, "type": "ip", "value": "10.20.30.40"}, "required": ["code", "type", "value"]}
```
