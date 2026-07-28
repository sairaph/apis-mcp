---
title: access_allowed_clipboard_local_to_remote_formats
page_id: schema-access-allowed-clipboard-local-to-remote-formats-3d708a71
path: schemas
description: Clipboard formats allowed when copying from local machine to remote RDP session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_allowed_clipboard_local_to_remote_formats

Clipboard formats allowed when copying from local machine to remote RDP session.

```yaml
{"description": "Clipboard formats allowed when copying from local machine to remote RDP session.", "type": "array", "items": {"$ref": "#/components/schemas/access_rdp_clipboard_format"}, "example": ["text", "file"]}
```
