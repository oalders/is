#!/bin/bash

go tool dist list | cut -d/ -f1 | uniq | sort
