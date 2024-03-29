# Temporal Worker

Temporal Worker is a role for Temporal service used for hosting any
components responsible for performing background processing on the Temporal
cluster.

## Replicator

Replicator is a background worker responsible for consuming replication tasks
generated by remote Temporal clusters and pass it down to processor, so they
can be applied to local Temporal cluster.

### Quickstart for localhost development

1. Start Temporal development server for active zone:
    ```bash
    make start-cdc-active
    ```

2. Start Temporal development server for standby(passive) zone:
    ```bash
    make start-cdc-standby
    ```
   
3. Connect two Temporal clusters:
   ```bash
   tctl --ad 127.0.0.1:7233 adm cl upsert-remote-cluster --frontend_address "localhost:8233"
   tctl --ad 127.0.0.1:8233 adm cl upsert-remote-cluster --frontend_address "localhost:7233"
   ```

4. Create global namespaces
    ```bash
    tctl --ns sample namespace register --gd true --ac active --cl active standby
    ```

5. Failover between zones:

    Failover to standby:
    ```bash
    tctl --ns sample namespace update --ac standby
    ```
    Failback to active:
    ```bash
    tctl --ns sample namespace update --ac active
    ```
