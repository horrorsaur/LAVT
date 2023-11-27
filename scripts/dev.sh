#!/usr/bin/env bash

set -e

if [[ -e app.log ]]; then
  echo "cleaning up log file"
  rm app.log 
fi

echo "starting application using wails CLI"
wails dev -loglevel=info
