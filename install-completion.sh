#!/bin/bash

# Set the target directory where the file will be extracted
TARGET_DIR="$HOME/.gomni/bin"

# Check if there is a gomni-v*.tar.gz file in the current directory
TAR_FILE=$(ls gomni-v*.tar.gz 2> /dev/null)

# If the file is not found, return an error
if [ -z "$TAR_FILE" ]; then
    echo "Error: No gomni-v*.tar.gz file found in the current directory."
    exit 1
fi

# Create the target directory if it doesn't exist
mkdir -p "$TARGET_DIR"

# Extract the contents of the tar.gz file to the target directory
echo "Extracting $TAR_FILE to $TARGET_DIR..."
tar -xf "$TAR_FILE" -C "$TARGET_DIR"

# Check if extraction was successful
if [ $? -eq 0 ]; then
    echo "Extraction completed successfully."
else
    echo "Error: Failed to extract $TAR_FILE."
    exit 1
fi

export PATH=$PATH:"$HOME/.gomni/bin"

# Detect the shell type
if [ -n "$ZSH_VERSION" ]; then
    SHELL_TYPE="zsh"
elif [ -n "$BASH_VERSION" ]; then
    SHELL_TYPE="bash"
elif [ -n "$FISH_VERSION" ]; then
    SHELL_TYPE="fish"
else
    echo "Unsupported shell. Only Bash, Zsh, and Fish are supported."
    exit 1
fi

# Install auto-completion based on shell type
echo "Installing gomni completion for $SHELL_TYPE..."

case $SHELL_TYPE in
    bash)
		echo -e "\n\nexport PATH=\$PATH:\$HOME/.gomni/bin" | tee -a ~/.bashrc
        gomni completion bash > ~/.gomni/completion
        source ~/.gomni/completion
        echo "Bash completion installed successfully."
        ;;
    zsh)
		echo -e "\n\nexport PATH=\$PATH:\$HOME/.gomni/bin" | tee -a ~/.zshrc
        gomni completion zsh > "${fpath[1]}/_gomni"
        echo "Zsh completion installed successfully. Restart your shell to apply changes."
        ;;
    fish)
		echo "add PATH=$PATH:\"$HOME/.gomni/bin\" to your PATH variable."
        gomni completion fish > ~/.config/fish/completions/gomni.fish
        echo "Fish completion installed successfully."
        ;;
esac

