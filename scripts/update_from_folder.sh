#!/bin/bash

# --- Configuration ---
VENV_DIR="./venv"  # Adjust this if your virtual environment is in a different location
PYTHON_SCRIPT="./update_screen.py" # Path to your Python script
REQUIREMENTS_FILE="requirements.txt"
IMAGE_DIR="./images/" # Directory containing the images

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

# --- Get a random image from the images directory ---
if [ ! -d "$IMAGE_DIR" ]; then
    echo "Error: Image directory '$IMAGE_DIR' does not exist."
    deactivate
    exit 1
fi

# Get a list of files in the image directory, filter for common image extensions, and pick a random one
shopt -s nullglob # Handle the case where no image files are found
image_files=($(find "$IMAGE_DIR" -maxdepth 1 -type f -regex '.*\.\(jpg\|jpeg\|png\|gif\)$'))
shopt -u nullglob

if [ ${#image_files[@]} -eq 0 ]; then
    echo "Error: No image files found in '$IMAGE_DIR'."
    deactivate
    exit 1
fi

random_index=$((RANDOM % ${#image_files[@]}))
IMAGE_FILE="${image_files[$random_index]}"
echo "Randomly selected image: $IMAGE_FILE"

# --- Run the Python script with arguments ---
SATURATION="${1:-0.8}" # Default saturation if not provided

echo "Running $PYTHON_SCRIPT with image: $IMAGE_FILE and saturation: $SATURATION"
python "$PYTHON_SCRIPT" "$IMAGE_FILE" "$SATURATION"

# --- Deactivate the virtual environment ---
deactivate
