#!/bin/bash

SERVICE_NAME=$1
BINARY_NAME=$2
ENVIRONMENT_NAME=$3
DO_REINSTALL=$4

BASE_DIR="/angularpay/services"
SERVICE_DIR="${BASE_DIR}/${SERVICE_NAME}"
ARTIFACT="${SERVICE_DIR}/${BINARY_NAME}"
SERVICE_LOGS_DIR="${SERVICE_DIR}/logs"


SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"


echo "Configuring $SERVICE_NAME for $ENVIRONMENT_NAME environment"

### create service file for Go! applications
		
	echo "Stopping service $SERVICE_NAME"
	sudo systemctl stop "$SERVICE_NAME.service"

	if [ "$DO_REINSTALL" -eq 1 ]; 
	then
		echo "Deleting service $SERVICE_NAME ..."
		sudo rm "$SERVICE_FILE"
	fi

	if [ ! -f "$SERVICE_FILE" ];
		then
			echo "$SERVICE_FILE [go service file] does not exist. Creating it..."
### Write into the service file ##
echo '[Unit]
Description='${SERVICE_NAME}' Service
After=network.target
[Service]
WorkingDirectory='${SERVICE_DIR}'
#path to executable.
#executable is a bash script
ExecStart='${ARTIFACT}'
Restart=on-failure
Type=simple
User=root
RestartSec=10
startLimitIntervalSec=60
[Install]
WantedBy=multi-user.target' | sudo tee "$SERVICE_FILE"
	else 
		echo "$SERVICE_FILE [go service file] already exist. Moving on..."
	fi
	
	echo "Reloading daemon service ..."
	sudo systemctl daemon-reload

	echo "Enabling $SERVICE_NAME ..."
	sudo systemctl enable "$SERVICE_NAME.service"
	
	echo "Starting $SERVICE_NAME ..."
	sudo systemctl start "$SERVICE_NAME.service"




