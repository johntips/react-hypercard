#!/bin/bash - 
set -o nounset                              # Treat unset variables as an error
set -x

env | grep ON_WATCH >/dev/null
if [ $? -eq 0 ]; then
    godo server --watch
  else
      go install
        go_app
      fi
fi
