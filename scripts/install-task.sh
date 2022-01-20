#!/usr/bin/env bash

is_user_root() { [ ${EUID:-$(id -u)} -eq 0 ]; }

TASK_PATH=$(which task)
if [ -z "$TASK_PATH" ]; then
  if is_user_root; then
    curl -sL https://taskfile.dev/install.sh | sh
    sudo mv bin/task /usr/local/bin
    echo "task binary added to \$PATH"
    rm -R bin
  else
    echo "you need to be a sudo to add the binary to \$PATH"
  fi
else
  echo "Task has already been installed"
fi
