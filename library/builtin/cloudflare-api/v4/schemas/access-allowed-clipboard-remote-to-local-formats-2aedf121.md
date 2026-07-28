---
title: access_allowed_clipboard_remote_to_local_formats
page_id: schema-access-allowed-clipboard-remote-to-local-formats-2aedf121
path: schemas
description: Clipboard formats allowed when copying from remote RDP session to local machine.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_allowed_clipboard_remote_to_local_formats

Clipboard formats allowed when copying from remote RDP session to local machine.

```yaml
{"description": "Clipboard formats allowed when copying from remote RDP session to local machine.", "type": "array", "items": {"$ref": "#/components/schemas/access_rdp_clipboard_format"}, "example": ["text", "file"]}
```
