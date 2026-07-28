---
title: Check Post-Quantum TLS support
page_id: operation-get-radar-post-quantum-tls-support-a5576779
path: operations/radar-post-quantum
description: Tests whether a hostname or IP address supports Post-Quantum (PQ) TLS key exchange. Returns information about the negotiated key exchange algorithm, whether it uses PQ cryptography, and any detected TLS implementation bugs (Split ClientHello, HRR failure, etc.).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/post_quantum/tls/support
operation_ids:
    - radar-get-post-quantum-tls-support
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check Post-Quantum TLS support

`GET /radar/post_quantum/tls/support`

Operation ID: `radar-get-post-quantum-tls-support`

Tests whether a hostname or IP address supports Post-Quantum (PQ) TLS key exchange. Returns information about the negotiated key exchange algorithm, whether it uses PQ cryptography, and any detected TLS implementation bugs (Split ClientHello, HRR failure, etc.).

## Definition

```yaml
{"operationId": "radar-get-post-quantum-tls-support", "summary": "Check Post-Quantum TLS support", "description": "Tests whether a hostname or IP address supports Post-Quantum (PQ) TLS key exchange. Returns information about the negotiated key exchange algorithm, whether it uses PQ cryptography, and any detected TLS implementation bugs (Split ClientHello, HRR failure, etc.).", "parameters": [{"name": "host", "in": "query", "description": "Hostname or IP address to test for Post-Quantum TLS support, optionally with port (defaults to 443).", "required": true, "schema": {"description": "Hostname or IP address to test for Post-Quantum TLS support, optionally with port (defaults to 443).", "type": "string", "example": "cloudflare.com", "minLength": 1}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"bugs": {"type": "object", "properties": {"hrrFailure": {"description": "Server sends a HelloRetryRequest but fails to complete the handshake after the client sends the second ClientHello. Often caused by non-compliant TLS 1.3 implementations on shared hosting providers.", "type": "boolean"}, "splitClientHello": {"description": "Server rejects fragmented ClientHello caused by large PQ keyshare, but accepts classical (non-PQ) handshakes. Typically caused by middleboxes or firewalls that cannot reassemble split TLS ClientHello messages.", "type": "boolean"}, "unknownKeyshare": {"description": "Server cannot handle an unknown key exchange algorithm in the ClientHello keyshare extension. Compliant servers should respond with HelloRetryRequest for a supported algorithm.", "type": "boolean"}}, "required": ["splitClientHello", "hrrFailure", "unknownKeyshare"]}, "host": {"description": "The host that was tested", "type": "string"}, "kex": {"description": "TLS CurveID of the negotiated key exchange", "type": "number"}, "kexName": {"description": "Human-readable name of the key exchange algorithm", "type": "string"}, "pq": {"description": "Whether the negotiated key exchange uses Post-Quantum cryptography (specifically X25519MLKEM768)", "type": "boolean"}}, "required": ["kex", "kexName", "pq", "host", "bugs"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Post-Quantum"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.post-quantum.tls", "x-fern-sdk-method-name": "support", "x-forge-hidden": true}
```
