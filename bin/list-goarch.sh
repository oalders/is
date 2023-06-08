#!/bin/bash

go tool dist list | cut -d/ -f2 | sort | uniq
