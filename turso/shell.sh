#/usr/bin/bash
read nameOrUrl sql
turso db shell $nameOrUrl "$sql"
