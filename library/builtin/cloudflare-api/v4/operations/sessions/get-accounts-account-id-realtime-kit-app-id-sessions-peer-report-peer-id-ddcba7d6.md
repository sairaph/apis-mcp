---
title: Fetch details of peer
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-peer-report-peer-id-3a2409de
path: operations/sessions
description: Returns participant details for the given peer ID along with call statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/peer-report/{peer_id}
operation_ids:
    - GetParticipantDataFromPeerId
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch details of peer

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/peer-report/{peer_id}`

Operation ID: `GetParticipantDataFromPeerId`

Returns participant details for the given peer ID along with call statistics.

## Definition

```yaml
{"operationId": "GetParticipantDataFromPeerId", "summary": "Fetch details of peer", "description": "Returns participant details for the given peer ID along with call statistics.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "filters", "in": "query", "description": "Filter to apply to the peer report.", "schema": {"type": "string", "example": "device_info", "enum": ["device_info", "ip_information", "precall_network_information", "events", "quality_stats"]}}, {"name": "include_peer_events", "in": "query", "description": "if true, response includes all the peer events of participant.", "schema": {"type": "boolean", "default": false}}, {"name": "peer_id", "in": "path", "description": "ID of the peer", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetParticipantDataFromPeerId"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
