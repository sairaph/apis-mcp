---
title: Delete Scan
page_id: operation-delete-accounts-account-id-vuln-scanner-scans-scan-id-38c3d023
path: operations/scans
description: |-
    Deletes a scan and all associated data.

    Only scans in a terminal state (`finished`, `failed`) may be deleted.
    Attempting to delete a scan that is still being created or executed
    (`created`, `scheduled`, `planning`, `running`) returns `400`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/scans/{scan_id}
operation_ids:
    - delete-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Scan

`DELETE /accounts/{account_id}/vuln_scanner/scans/{scan_id}`

Operation ID: `delete-scan`

Deletes a scan and all associated data.

Only scans in a terminal state (`finished`, `failed`) may be deleted.
Attempting to delete a scan that is still being created or executed
(`created`, `scheduled`, `planning`, `running`) returns `400`.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_scan_id"}]
```

## Definition

```yaml
{"operationId": "delete-scan", "summary": "Delete Scan", "description": "Deletes a scan and all associated data.\n\nOnly scans in a terminal state (`finished`, `failed`) may be deleted.\nAttempting to delete a scan that is still being created or executed\n(`created`, `scheduled`, `planning`, `running`) returns `400`.\n", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_delete-scan-response"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"]}
```
