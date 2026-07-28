---
title: access_connection_rules_ssh
page_id: schema-access-connection-rules-ssh-0cddd4aa
path: schemas
description: The SSH-specific rules that define how users may connect to the targets secured by your application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_connection_rules_ssh

The SSH-specific rules that define how users may connect to the targets secured by your application.

```yaml
{"description": "The SSH-specific rules that define how users may connect to the targets secured by your application.", "type": "object", "properties": {"allow_email_alias": {"$ref": "#/components/schemas/access_allow_email_alias"}, "usernames": {"$ref": "#/components/schemas/access_usernames"}}, "required": ["usernames"], "title": "SSH Connection Rules"}
```
