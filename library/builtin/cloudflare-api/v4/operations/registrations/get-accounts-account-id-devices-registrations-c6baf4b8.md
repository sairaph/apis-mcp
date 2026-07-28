---
title: List registrations
page_id: operation-get-accounts-account-id-devices-registrations-5788d751
path: operations/registrations
description: Lists WARP registrations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/registrations
operation_ids:
    - list-registrations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List registrations

`GET /accounts/{account_id}/devices/registrations`

Operation ID: `list-registrations`

Lists WARP registrations.

## Definition

```yaml
{"operationId": "list-registrations", "summary": "List registrations", "description": "Lists WARP registrations.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_user_id"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_seen_after"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_seen_before"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_status"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_per_page"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_search"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_sort_by"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_sort_order"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_cursor"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_id"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_device_id"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_policy_id"}, {"$ref": "#/components/parameters/teams-devices_devices_list_registrations_param_include"}], "responses": {"200": {"description": "List of registrations response.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": [{"created_at": "2025-02-14T13:17:00Z", "deleted_at": null, "device": {"client_version": "1.0.0", "id": "32aa0404-78f1-49a4-99e0-97f575081356", "name": "My Device"}, "id": "11ffb86f-3f0c-4306-b4a2-e62f872b166a", "key": "U+QTP50RsWfeLGHF4tlGDnmGeuwtsz46KCHr5OyhWq00Rsdfl45mgnQAuEJ6CO0YrkyTl9FUf5iB0bwYR3g4EEFEHhtu6jFaqfMrBMBSz6itv9HQXkaR9OieKQ==", "key_type": "secp256r1", "last_seen_at": "2025-02-14T13:17:00Z", "revoked_at": null, "tunnel_type": "masque", "updated_at": "2025-02-14T13:17:00Z", "user": {"email": "alice@example.org", "id": "30323c1f-318d-4ec9-92c7-5a8c4d25c4fc", "name": "Alice"}}, {"created_at": "2025-02-15T10:20:00Z", "deleted_at": null, "device": {"client_version": "1.0.1", "id": "43bb1515-8902-50b5-aa01-a88686192467", "name": "Bob's Laptop"}, "id": "22eedc7a-4a1d-5417-c5b3-f73a983c277b", "key": "V/RSP61StXgfmLHJG5umHEonHfvxtz57LDIs6PziXr11Stegm56nhrRBvFK7DP1ZsLzUm0GVg6jC1cxZS4h5FFGFJiju7kGbrgNsCNCT77juw0IRYlS0QpjgLR==", "key_type": "secp256r1", "last_seen_at": "2025-02-15T10:25:00Z", "revoked_at": null, "tunnel_type": "masque", "updated_at": "2025-02-15T10:25:00Z", "user": {"email": "bob@example.com", "id": "41434d2a-429e-5fd0-a3d8-6b9d5e36d5ad", "name": "Bob"}}], "result_info": {"count": 2, "cursor": "ais86dftf.asdf7ba8", "per_page": 10, "total_count": null}, "success": true}, "schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_registration"}}, "result_info": {"$ref": "#/components/schemas/teams-devices_cursor_result_info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Registrations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.registrations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
