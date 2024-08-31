#!/usr/bin/env bash
set -e
export GIT_COMMIT=$(git describe --tags)
if [ -z "$GIT_COMMIT" ]
then
    echo "'git describe --tags' did not provide any output."
    echo "This probably means that you either have not defined any tags "
    echo "or you are running this code from a non-git directory."
    exit 1
fi
echo "Git commit $GIT_COMMIT"
go build -ldflags "-X 'github.com/datacharmer/cliptool/cmd.CliptoolVersion=$GIT_COMMIT'"
#go build -ldflags "-X 'main.CliptoolVersion=$GIT_COMMIT'"
set -x
./cliptool --version
