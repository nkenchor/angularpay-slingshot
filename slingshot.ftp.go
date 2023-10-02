
#!/bin/bash

SERVICE_NAME=$1
SOURCE_FILE=$2
SOURCE_CONFIG=$3

BASE_DIR="/angularpay/services"
SERVICE_DIR="${BASE_DIR}/${SERVICE_NAME}"
ARTIFACT="${SERVICE_DIR}/${SERVICE_NAME}"
CONFIG="${SERVICE_DIR}/${SERVICE_NAME}.env"

if [ ! -d "$SERVICE_DIR" ];
then 
	echo "$SERVICE_DIR does not exist! Creating it..."
	mkdir -p "$SERVICE_DIR"
fi

if [ -f "$ARTIFACT" ];
then
	echo "Deleting old $SERVICE_NAME artifact from $SERVICE_DIR"
	sudo rm -f "${ARTIFACT}"
	sudo rm -f "${CONFIG}"
fi

echo "Copying new $SERVICE_NAME artifact to $SERVICE_DIR"
sudo cp "$SOURCE_FILE" "$SERVICE_DIR"
sudo cp "${SOURCE_CONFIG}" "$SERVICE_DIR"
