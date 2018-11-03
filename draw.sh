#!/usr/bin/env bash

set -e

function help() {
  cat <<HELP
    Usage:
      $0 <ASN> <DEEP_LEVEL>

    ASN       = autonomous system number
    DEEP_LEVEL  = upstream deep level (See internet-topology help)
HELP
}

[ $# -eq 0 -o "$1" = "-h" -o "$1" = "--help" ] && ( help; exit 1 )

if [[ ! $1 =~ ^[0-9]+$ ]] || [[ $2 && ! $2 =~ ^[0-9]+$ ]]; then
  echo "<ASN> and <DEEP_LEVEL> are integers"
  exit 2
fi

if [ ! $2 ]; then
  echo "default to 1"
  DEEP_LEVEL=1
else
  DEEP_LEVEL=$2
fi

echo -e "Start with values\n - ASN=$1\n - DEEP_LEVEL=$DEEP_LEVEL\n"

for ((i = 1; i <= $DEEP_LEVEL; i++)); do
  dir=draw/level-$i/asn-$1
  prefix=$dir/level-$i

  if [ ! -f $prefix.data -o ! -f $prefix.dot ]; then
    echo -e "Quering ASN $1 in deep $i\n\n"
    mkdir -p $dir
    ./internet-topology --asn=$1 --deep=$i -pi --output $prefix
  fi
  [ ! -f $prefix-dot.svg ] && dot -Tsvg $prefix.dot > $prefix-dot.svg
  [ ! -f $prefix-neato.svg ] &&  neato -Tsvg $prefix.dot > $prefix-neato.svg
  [ ! -f $prefix-twopi.svg ] &&  twopi -Tsvg $prefix.dot > $prefix-twopi.svg
  [ ! -f $prefix-circo.svg ] &&  circo -Tsvg $prefix.dot > $prefix-circo.svg
  [ ! -f $prefix-fdp.svg ] &&  fdp -Tsvg $prefix.dot > $prefix-fdp.svg
  [ ! -f $prefix-sfdp.svg ] &&  sfdp -Tsvg $prefix.dot > $prefix-sfdp.svg
done

echo -e "\nDone!"