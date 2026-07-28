---
title: magic_bgp_config
page_id: schema-magic-bgp-config-6fdbf907
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_bgp_config

```yaml
{"type": "object", "properties": {"customer_asn": {"description": "ASN used on the customer end of the BGP session", "type": "integer", "format": "int32", "minimum": 0}, "export_filter_id": {"description": "ID of the BGP filter profile applied to routes advertised to the customer.", "type": "string", "example": "a1b2c3d4e5f647890a1b2c3d4e5f6789"}, "extra_prefixes": {"description": "Prefixes in this list will be advertised to the customer device, in addition to the routes in the Magic routing table.", "type": "array", "items": {"format": "cidr", "type": "string"}}, "import_filter_id": {"description": "ID of the BGP filter profile applied to routes received from the customer.", "type": "string", "example": "a1b2c3d4e5f647890a1b2c3d4e5f6789"}, "md5_key": {"description": "MD5 key to use for session authentication.\n\nNote that *this is not a security measure*. MD5 is not a valid security mechanism, and the\nkey is not treated as a secret value. This is *only* supported for preventing\nmisconfiguration, not for defending against malicious attacks.\n\nThe MD5 key, if set, must be of non-zero length and consist only of the following types of\ncharacter:\n\n* ASCII alphanumerics: `[a-zA-Z0-9]`\n* Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \\|`\n\nIn other words, MD5 keys may contain any printable ASCII character aside from newline (0x0A),\nquotation mark (`\"`), vertical tab (0x0B), carriage return (0x0D), tab (0x09), form feed\n(0x0C), and the question mark (`?`). Requests specifying an MD5 key with one or more of\nthese disallowed characters will be rejected.", "type": "string"}}, "required": ["customer_asn"]}
```
