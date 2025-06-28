#!/bin/bash
echo "host all all 0.0.0.0/0 trust" >> "$PGDATA/pg_hba.conf"
echo "host replication all 0.0.0.0/0 trust" >> "$PGDATA/pg_hba.conf"
