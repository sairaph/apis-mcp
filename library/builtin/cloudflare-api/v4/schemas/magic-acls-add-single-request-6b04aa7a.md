---
title: magic_acls_add_single_request
page_id: schema-magic-acls-add-single-request-6b04aa7a
path: schemas
description: Bidirectional ACL policy for local network traffic within a site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_acls_add_single_request

Bidirectional ACL policy for local network traffic within a site.

```yaml
{"description": "Bidirectional ACL policy for local network traffic within a site.", "type": "object", "properties": {"description": {"description": "Description for the ACL.", "type": "string", "example": "Allows local traffic between PIN pads and cash register."}, "forward_locally": {"$ref": "#/components/schemas/magic_forward_locally"}, "lan_1": {"$ref": "#/components/schemas/magic_lan-acl-configuration"}, "lan_2": {"$ref": "#/components/schemas/magic_lan-acl-configuration"}, "name": {"description": "The name of the ACL.", "type": "string", "example": "PIN Pad - Cash Register"}, "protocols": {"type": "array", "items": {"description": "Array of allowed communication protocols between configured LANs. If no protocols are provided, all protocols are allowed.", "enum": ["tcp", "udp", "icmp"], "type": "string"}}, "unidirectional": {"$ref": "#/components/schemas/magic_unidirectional"}}, "required": ["name", "lan_1", "lan_2"]}
```
