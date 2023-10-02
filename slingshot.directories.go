#!/bin/bash

SERVICE_NAME=$1

### AngularPay Service Directories ##
BASE_DIR="/angularpay/services"
SERVICE_DIR="${BASE_DIR}/${SERVICE_NAME}"
SERVICE_LOGS_DIR="${SERVICE_DIR}/logs"


if [ ! -d "$BASE_DIR" ];
then 
	echo "$BASE_DIR does not exist! Creating it..."
	mkdir -p "$BASE_DIR"
else
	echo "$BASE_DIR already exist. Moving on..."
fi
if [ ! -d "$SERVICE_DIR" ];
then 
	echo "$SERVICE_DIR does not exist! Creating it..."
	mkdir -p "$SERVICE_DIR"
else
	echo "$SERVICE_DIR already exist. Moving on..."
fi

if [ ! -d "$SERVICE_LOGS_DIR" ];
then 
	echo "$SERVICE_LOGS_DIR does not exist! Creating it..."
	mkdir -p "$SERVICE_LOGS_DIR"
else
	echo "$SERVICE_LOGS_DIR already exist. Moving on..."
fi


