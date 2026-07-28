---
title: mconn_event
page_id: schema-mconn-event-737e9c01
path: schemas
description: |-
    Event kind plus event-specific payload fields.

    Event kinds:
    - `Init`: Initialized process
    - `Leave`: Stopped process
    - `StartAttestation`: Started attestation
    - `FinishAttestationSuccess`: Finished attestation
    - `FinishAttestationFailure`: Failed attestation
    - `StartRotateCryptKey`: Started crypt key rotation
    - `FinishRotateCryptKeySuccess`: Finished crypt key rotation
    - `FinishRotateCryptKeyFailure`: Failed crypt key rotation
    - `StartRotatePki`: Started PKI rotation
    - `FinishRotatePkiSuccess`: Finished PKI rotation
    - `FinishRotatePkiFailure`: Failed PKI rotation
    - `StartUpgrade`: Started upgrade
    - `FinishUpgradeSuccess`: Finished upgrade
    - `FinishUpgradeFailure`: Failed upgrade
    - `Reconcile`: Reconciled
    - `ConfigureCloudflaredTunnel`: Configured Cloudflared tunnel
    - `RekeyInstallBoth`: Installed initial inbound and outbound keys
    - `RekeyStart`: Installed new inbound key, kept old outbound
    - `RekeyRestart`: Restarted in-progress rekey with newer key material
    - `RekeyAdvance`: Confirmed traffic on new inbound key, swapped outbound to new
    - `RekeyComplete`: Deleted old keys
    - `RekeyReset`: Deleted all keys after receiving an unexpected key
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_event

Event kind plus event-specific payload fields.

Event kinds:
- `Init`: Initialized process
- `Leave`: Stopped process
- `StartAttestation`: Started attestation
- `FinishAttestationSuccess`: Finished attestation
- `FinishAttestationFailure`: Failed attestation
- `StartRotateCryptKey`: Started crypt key rotation
- `FinishRotateCryptKeySuccess`: Finished crypt key rotation
- `FinishRotateCryptKeyFailure`: Failed crypt key rotation
- `StartRotatePki`: Started PKI rotation
- `FinishRotatePkiSuccess`: Finished PKI rotation
- `FinishRotatePkiFailure`: Failed PKI rotation
- `StartUpgrade`: Started upgrade
- `FinishUpgradeSuccess`: Finished upgrade
- `FinishUpgradeFailure`: Failed upgrade
- `Reconcile`: Reconciled
- `ConfigureCloudflaredTunnel`: Configured Cloudflared tunnel
- `RekeyInstallBoth`: Installed initial inbound and outbound keys
- `RekeyStart`: Installed new inbound key, kept old outbound
- `RekeyRestart`: Restarted in-progress rekey with newer key material
- `RekeyAdvance`: Confirmed traffic on new inbound key, swapped outbound to new
- `RekeyComplete`: Deleted old keys
- `RekeyReset`: Deleted all keys after receiving an unexpected key

```yaml
{"description": "Event kind plus event-specific payload fields.\n\nEvent kinds:\n- `Init`: Initialized process\n- `Leave`: Stopped process\n- `StartAttestation`: Started attestation\n- `FinishAttestationSuccess`: Finished attestation\n- `FinishAttestationFailure`: Failed attestation\n- `StartRotateCryptKey`: Started crypt key rotation\n- `FinishRotateCryptKeySuccess`: Finished crypt key rotation\n- `FinishRotateCryptKeyFailure`: Failed crypt key rotation\n- `StartRotatePki`: Started PKI rotation\n- `FinishRotatePkiSuccess`: Finished PKI rotation\n- `FinishRotatePkiFailure`: Failed PKI rotation\n- `StartUpgrade`: Started upgrade\n- `FinishUpgradeSuccess`: Finished upgrade\n- `FinishUpgradeFailure`: Failed upgrade\n- `Reconcile`: Reconciled\n- `ConfigureCloudflaredTunnel`: Configured Cloudflared tunnel\n- `RekeyInstallBoth`: Installed initial inbound and outbound keys\n- `RekeyStart`: Installed new inbound key, kept old outbound\n- `RekeyRestart`: Restarted in-progress rekey with newer key material\n- `RekeyAdvance`: Confirmed traffic on new inbound key, swapped outbound to new\n- `RekeyComplete`: Deleted old keys\n- `RekeyReset`: Deleted all keys after receiving an unexpected key", "type": "object", "properties": {"k": {"description": "Event kind", "type": "string", "enum": ["Init", "Leave", "StartAttestation", "FinishAttestationSuccess", "FinishAttestationFailure", "StartRotateCryptKey", "FinishRotateCryptKeySuccess", "FinishRotateCryptKeyFailure", "StartRotatePki", "FinishRotatePkiSuccess", "FinishRotatePkiFailure", "StartUpgrade", "FinishUpgradeSuccess", "FinishUpgradeFailure", "Reconcile", "ConfigureCloudflaredTunnel", "RekeyInstallBoth", "RekeyStart", "RekeyRestart", "RekeyAdvance", "RekeyComplete", "RekeyReset"]}}, "additionalProperties": true, "required": ["k"]}
```
