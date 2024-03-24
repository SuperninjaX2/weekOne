#!/bin/bash

copy_from_outside() {
    # Source directory
    srcDir="/storage/code/weekOne"

    # Destination directory
    destDir="/storage/internal/weekOne"

    # Call the copy function
    copy_dir "$srcDir" "$destDir"

    echo "Directory copied from outside to inside successfully!"
}

copy_from_inside() {
    # Source directory
    srcDir="/storage/internal/weekOne"

    # Destination directory
    destDir="/storage/code/weekOne"

    # Call the copy function
    copy_dir "$srcDir" "$destDir"

    echo "Directory copied from inside to outside successfully!"
}

copy_dir() {
    src="$1"
    dest="$2"

    # Create destination directory if it does not exist
    mkdir -p "$dest"

    # Get the content of the source directory
    for entry in "$src"/*; do
        # Ignore .git directory
        if [ "$(basename "$entry")" == ".git" ]; then
            continue
        fi

        if [ -d "$entry" ]; then
            # If the entry is a directory, recursively copy it
            copy_dir "$entry" "$dest/$(basename "$entry")"
        else
            # If the entry is a file, copy it
            cp "$entry" "$dest"
        fi
    done
}

# Prompt user for choice
read -p "Do you want to copy from outside to inside or inside to outside? (Type 'out' or 'in'): " choice

case "$choice" in
    out) copy_from_outside ;;
    in) copy_from_inside ;;
    *) echo "Invalid choice. Please type 'out' or 'in'." ;;
esac
