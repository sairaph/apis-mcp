---
title: tls-certificates-and-hostnames_value
page_id: schema-tls-certificates-and-hostnames-value-91543be1
path: schemas
description: |-
    The TLS setting value.
    The type depends on the `setting_id` used in the request path:
    - `ciphers`: an array of allowed cipher suite strings in BoringSSL format (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`).
    - `min_tls_version`: a string indicating the minimum TLS version — one of `"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"` (e.g., `"1.2"`).
    - `http2`: a string indicating whether HTTP/2 is enabled — `"on"` or `"off"` (e.g., `"on"`).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_value

The TLS setting value.
The type depends on the `setting_id` used in the request path:
- `ciphers`: an array of allowed cipher suite strings in BoringSSL format (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`).
- `min_tls_version`: a string indicating the minimum TLS version — one of `"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"` (e.g., `"1.2"`).
- `http2`: a string indicating whether HTTP/2 is enabled — `"on"` or `"off"` (e.g., `"on"`).

```yaml
{"description": "The TLS setting value.\nThe type depends on the `setting_id` used in the request path:\n- `ciphers`: an array of allowed cipher suite strings in BoringSSL format (e.g., `[\"ECDHE-RSA-AES128-GCM-SHA256\", \"AES128-GCM-SHA256\"]`).\n- `min_tls_version`: a string indicating the minimum TLS version — one of `\"1.0\"`, `\"1.1\"`, `\"1.2\"`, or `\"1.3\"` (e.g., `\"1.2\"`).\n- `http2`: a string indicating whether HTTP/2 is enabled — `\"on\"` or `\"off\"` (e.g., `\"on\"`).", "oneOf": [{"description": "Used when `setting_id` is `ciphers`. An array of allowed cipher suite strings.", "example": ["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"], "items": {"type": "string"}, "type": "array", "x-auditable": true}, {"description": "Used when `setting_id` is `min_tls_version`. The minimum TLS version.", "enum": ["1.0", "1.1", "1.2", "1.3"], "example": "1.2", "type": "string", "x-auditable": true}, {"description": "Used when `setting_id` is `http2`. Whether HTTP/2 is enabled.", "enum": ["on", "off"], "example": "on", "type": "string", "x-auditable": true}]}
```
