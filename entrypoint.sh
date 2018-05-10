#!/bin/sh

trap 'exit' TERM

ISDAEMON='false'

while getopts b:c:k:g:d option
do
    case "${option}"
        in
        b)  echo "Copy CNI Binary to "${OPTARG}
            cp -f /go/bin/multus ${OPTARG}
            ;;
        c)
            echo "Copty CNI Conf to "${OPTARG}
            cp -f /tmp/multus-cni.conf ${OPTARG}
            ;;
        k) 
            echo "Copty CRD resource to "${OPTARG}
            cp -f /tmp/crdnetwork.yaml ${OPTARG}
            ;;
        g)
            echo "Try to embed the CNI ${OPTARG} to mulsu CNI and save in ${DEST_CNI}"
            echo "/go/bin/docker-multus-cni -srcDir "${OPTARG}" -output "${DEST_CNI}""
            echo "The current file of "${OPTARG}""
            ls ${OPTARG}
            /go/bin/docker-multus-cni -srcDir "${OPTARG}" -output "${DEST_CNI}"
            ;;
        d)
            ISDAEMON='true'
            ;;
    esac
done

if [ "$ISDAEMON" == "true" ] ;then
while true; do sleep 1; done
fi
