---
title: workers_exports_reconciliation_scenario
page_id: schema-workers-exports-reconciliation-scenario-055fd08e
path: schemas
description: |-
    Stable, machine-readable tag identifying which reconciliation
    scenario produced an error, warning, or info entry. Clients may
    branch on this value instead of parsing `message`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_scenario

Stable, machine-readable tag identifying which reconciliation
scenario produced an error, warning, or info entry. Clients may
branch on this value instead of parsing `message`.

```yaml
{"description": "Stable, machine-readable tag identifying which reconciliation\nscenario produced an error, warning, or info entry. Clients may\nbranch on this value instead of parsing `message`.\n", "type": "string", "enum": ["code_class_not_in_exports", "provisioned_class_missing_from_config", "config_export_not_in_code", "config_references_nonexistent_class", "orphaned_provisioned_namespace", "storage_type_mismatch", "free_tier_requires_sqlite", "invalid_export", "tombstone_delete_class_still_in_code", "tombstone_delete_blocked_by_external_bindings", "tombstone_renamed_to_occupied", "transferred_pending_not_found", "transferred_target_missing", "transferred_target_mismatch", "phase_one_transfer_source_missing", "phase_one_transfer_source_namespace_missing", "phase_one_transfer_target_class_provisioned", "phase_one_transfer_after_commit_mismatch", "phase_one_transfer_duplicate", "phase_one_transfer_target_in_dispatch_namespace", "phase_one_transfer_source_in_dispatch_namespace", "transferred_source_in_dispatch_namespace", "transferred_target_in_dispatch_namespace", "container_undeclared_reference", "container_class_not_durable_object", "container_wiring_inconsistent", "container_multiple_durable_objects", "transfer_container_parity_mismatch", "transfer_container_parity_mismatch_on_commit", "tombstone_class_still_in_code", "stale_tombstone", "transfer_receive_already_applied", "transfer_receive_cleanup_complete"]}
```
