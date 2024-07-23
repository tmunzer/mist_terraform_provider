#/bin/sh


git status

read -p "Continue (y/N)? " resp

if echo "x$resp" | grep -vqi "xy"
then
    echo "Exiting..."
    exit 0
fi

echo ">> Importing GPG Sign Key"
