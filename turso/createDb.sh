#!/usr/bin/bash
read group name
turso db create $name --group $group 
