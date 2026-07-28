---
title: teams-devices_tls_config_response
page_id: schema-teams-devices-tls-config-response-b1e55fec
path: schemas
description: The Managed Network TLS Config Response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_tls_config_response

The Managed Network TLS Config Response.

```yaml
{"description": "The Managed Network TLS Config Response.", "type": "object", "properties": {"sha256": {"description": "The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.", "type": "string", "example": "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"}, "tls_sockaddr": {"description": "A network address of the form \"host:port\" that the WARP client will use to detect the presence of a TLS host.", "type": "string", "example": "foobar:1234", "x-auditable": true}}, "required": ["tls_sockaddr"]}
```
