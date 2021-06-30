#!/usr/bin/env bash

BASE=${1:-datasets}

if [ "$1" = "-h" -o "$1" = "--help" ]; then
  cat <<HELP
Usage $0:
  You can specify the path base as first argument
HELP
  exit 1
fi

# http://ftp.lacnic.net/pub/stats/lacnic/RIR-Statistics-Exchange-Format.txt

URLS=(
ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest
ftp.ripe.net/ripe/stats/delegated-ripencc-latest
ftp.afrinic.net/pub/stats/afrinic/delegated-afrinic-latest
ftp.apnic.net/pub/stats/apnic/delegated-apnic-latest
ftp.lacnic.net/pub/stats/lacnic/delegated-lacnic-latest
)

mkdir -p datasets

for URL in ${URLS[@]}; do
  ZONE=`echo $URL | cut -d. -f2`
  echo "Quering to $ZONE..."
  curl -s $URL > $BASE/$ZONE
  echo -e "Finish $ZONE at $BASE/$ZONE"
done

