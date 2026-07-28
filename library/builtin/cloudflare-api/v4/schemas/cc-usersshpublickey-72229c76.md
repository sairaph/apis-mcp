---
title: cc_UserSSHPublicKey
page_id: schema-cc-usersshpublickey-72229c76
path: schemas
description: SSH public key provided by the user
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_UserSSHPublicKey

SSH public key provided by the user

```yaml
{"description": "SSH public key provided by the user", "type": "object", "properties": {"name": {"description": "Optional human readable name for this key", "type": "string"}, "public_key": {"$ref": "#/components/schemas/cc_SSHPublicKey"}}, "required": ["public_key"]}
```
