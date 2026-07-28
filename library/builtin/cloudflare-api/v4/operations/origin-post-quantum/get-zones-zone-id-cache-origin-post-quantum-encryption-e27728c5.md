---
title: Get Origin Post-Quantum Encryption setting
page_id: operation-get-zones-zone-id-cache-origin-post-quantum-encryption-3a512f7f
path: operations/origin-post-quantum
description: Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when connecting to your origin. Preferred instructs Cloudflare to opportunistically send a Post-Quantum keyshare in the first message to the origin (for fastest connections when the origin supports and prefers PQ), supported means that PQ algorithms are advertised but only used when requested by the origin, and off means that PQ algorithms are not advertised.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/cache/origin_post_quantum_encryption
operation_ids:
    - zone-cache-settings-get-origin-post-quantum-encryption-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Origin Post-Quantum Encryption setting

`GET /zones/{zone_id}/cache/origin_post_quantum_encryption`

Operation ID: `zone-cache-settings-get-origin-post-quantum-encryption-setting`

Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when connecting to your origin. Preferred instructs Cloudflare to opportunistically send a Post-Quantum keyshare in the first message to the origin (for fastest connections when the origin supports and prefers PQ), supported means that PQ algorithms are advertised but only used when requested by the origin, and off means that PQ algorithms are not advertised.

## Definition

```yaml
{"operationId": "zone-cache-settings-get-origin-post-quantum-encryption-setting", "summary": "Get Origin Post-Quantum Encryption setting", "description": "Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when connecting to your origin. Preferred instructs Cloudflare to opportunistically send a Post-Quantum keyshare in the first message to the origin (for fastest connections when the origin supports and prefers PQ), supported means that PQ algorithms are advertised but only used when requested by the origin, and off means that PQ algorithms are not advertised.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "Get Origin Post-Quantum Encryption setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_origin_post_quantum_encryption_response_value"}]}}}}, "4XX": {"description": "Get Origin Post-Quantum Encryption setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Post-Quantum"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "Zone Read", "Zone Write"], "x-cfDeprecation": {"description": "Origin post-quantum encryption selection is now handled automatically by Auto-Origin TLS KEX (`/zones/{zone_id}/settings/auto_origin_tls_kex`) combined with Origin TLS Compliance Modes (`/zones/{zone_id}/settings/origin_tls_compliance_modes`). As that rollout proceeds, the manual `origin_post_quantum_encryption` setting becomes ineffective and the endpoint will be removed on 2026-07-19.", "display": true, "eol": "2026-07-19", "id": "origin_pqe_deprecation"}, "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
