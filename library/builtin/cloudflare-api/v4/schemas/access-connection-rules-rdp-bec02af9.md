---
title: access_connection_rules_rdp
page_id: schema-access-connection-rules-rdp-bec02af9
path: schemas
description: The RDP-specific rules that define clipboard behavior for RDP connections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_connection_rules_rdp

The RDP-specific rules that define clipboard behavior for RDP connections.

```yaml
{"description": "The RDP-specific rules that define clipboard behavior for RDP connections.", "type": "object", "properties": {"allowed_clipboard_local_to_remote_formats": {"$ref": "#/components/schemas/access_allowed_clipboard_local_to_remote_formats"}, "allowed_clipboard_remote_to_local_formats": {"$ref": "#/components/schemas/access_allowed_clipboard_remote_to_local_formats"}}, "title": "RDP Connection Rules"}
```
