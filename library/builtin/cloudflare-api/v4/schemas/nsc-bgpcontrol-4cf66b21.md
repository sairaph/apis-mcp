---
title: nsc_BgpControl
page_id: schema-nsc-bgpcontrol-4cf66b21
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_BgpControl

```yaml
{"type": "object", "properties": {"customer_asn": {"description": "ASN used on the customer end of the BGP session", "type": "integer", "format": "int32", "minimum": 0}, "extra_prefixes": {"description": "Extra set of static prefixes to advertise to the customer's end of the session", "type": "array", "items": {"format": "A.B.C.D/N", "type": "string"}, "example": "192.168.3.4/31"}, "md5_key": {"description": "MD5 key to use for session authentication.\n\nNote that *this is not a security measure*. MD5 is not a valid security mechanism, and the\nkey is not treated as a secret value. This is *only* supported for preventing\nmisconfiguration, not for defending against malicious attacks.\n\nThe MD5 key, if set, must be of non-zero length and consist only of the following types of\ncharacter:\n\n* ASCII alphanumerics: `[a-zA-Z0-9]`\n* Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \\|`\n\nIn other words, MD5 keys may contain any printable ASCII character aside from newline (0x0A),\nquotation mark (`\"`), vertical tab (0x0B), carriage return (0x0D), tab (0x09), form feed\n(0x0C), and the question mark (`?`). Requests specifying an MD5 key with one or more of\nthese disallowed characters will be rejected.", "type": "string", "nullable": true}}, "required": ["customer_asn", "extra_prefixes"]}
```
