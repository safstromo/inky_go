#!/bin/bash

# --- Configuration ---
VENV_DIR="./venv"  # Adjust this if your virtual environment is in a different location
PYTHON_SCRIPT="./update_screen.py" # Path to your Python script
REQUIREMENTS_FILE="requirements.txt"

# --- Create virtual environment if it doesn't exist ---
if [ ! -d "$VENV_DIR" ]; then
    echo "Creating virtual environment in $VENV_DIR..."
    python3 -m venv "$VENV_DIR"
    echo "Virtual environment created."
fi

# --- Activate the virtual environment ---
source "$VENV_DIR/bin/activate"

# --- Install dependencies if a requirements file exists ---
if [ -f "$REQUIREMENTS_FILE" ]; then
    echo "Installing dependencies from $REQUIREMENTS_FILE..."
    pip install -r "$REQUIREMENTS_FILE"
    echo "Dependencies installed."
fi

# --- Run the Python script with arguments ---
if [ -z "$1" ]; then
    echo "Usage: $0 <image_file> [saturation]"
    deactivate
    exit 1
fi

IMAGE_FILE="$1"
SATURATION="${2:-0.8}" # Default saturation if not provided

echo "Running $PYTHON_SCRIPT with image: $IMAGE_FILE and saturation: $SATURATION"
python "$PYTHON_SCRIPT" "$IMAGE_FILE" "$SATURATION"

# --- Deactivate the virtual environment ---
deactivate
