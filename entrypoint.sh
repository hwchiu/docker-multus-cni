#!/bin/sh

while getopts b:c:k: option
do
    case "${option}"
        in
        b)  echo "Copy CNI Binary to "${OPTARG}
            cp -f /tmp/multus ${OPTARG}
            ;;
        c)
            echo "Copty CNI Conf to "${OPTARG}
            cp -f /tmp/multus-cni.conf ${OPTARG}
            ;;
        k) 
            echo "Copty CRD resource to "${OPTARG}
            cp -f /tmp/crdnetwork.yaml ${OPTARG}
            ;;
    esac
done
