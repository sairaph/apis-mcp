---
title: smartshield_http_config
page_id: schema-smartshield-http-config-ec2844b5
path: schemas
description: Parameters specific to an HTTP or HTTPS health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_http_config

Parameters specific to an HTTP or HTTPS health check.

```yaml
{"description": "Parameters specific to an HTTP or HTTPS health check.", "type": "object", "properties": {"allow_insecure": {"description": "Do not validate the certificate when the health check uses HTTPS.", "type": "boolean", "default": false, "x-auditable": true}, "expected_body": {"description": "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.", "type": "string", "example": "success", "x-auditable": true}, "expected_codes": {"description": "The expected HTTP response codes (e.g. \"200\") or code ranges (e.g. \"2xx\" for all codes starting with 2) of the health check.", "type": "array", "items": {"type": "string"}, "example": ["2xx", "302"], "default": "200", "nullable": true, "x-auditable": true}, "follow_redirects": {"description": "Follow redirects if the origin returns a 3xx status code.", "type": "boolean", "default": false, "x-auditable": true}, "header": {"description": "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.", "type": "object", "example": {"Host": ["example.com"], "X-App-ID": ["abc123"]}, "additionalProperties": {"items": {"type": "string", "x-auditable": true}, "type": "array"}, "nullable": true}, "method": {"description": "The HTTP method to use for the health check.", "type": "string", "default": "GET", "enum": ["GET", "HEAD"], "x-auditable": true}, "path": {"description": "The endpoint path to health check against.", "type": "string", "example": "/health", "default": "/", "x-auditable": true}, "port": {"description": "Port number to connect to for the health check. Defaults to 80 if type is HTTP or 443 if type is HTTPS.", "type": "integer", "default": 80, "x-auditable": true}}, "nullable": true}
```
