#!/bin/bash

# Define a function to automatically check, commit, and push git changes
function auto_git_push() { # Starts the definition of the function named auto_git_push
    while true; do # Starts an infinite loop (the script will run continuously until manually stopped)
        echo "🔍 Checking for changes at $(date)..." # Outputs the current time and a message indicating a check is starting

        # Check for uncommitted (staged or unstaged) changes
        # The 'git status --porcelain' command outputs in an easy-to-parse format. If there are no changes, the output is empty.
        if [[ -z $(git status --porcelain) ]]; then # Checks if the output of 'git status --porcelain' is empty (-z)
            echo "No changes to commit." # Output if no changes are detected
        else # Executes if changes *are* detected (output is not empty)
            echo "🧹 Removing large PDF files (>100MB) from PDFs/ directory..." # Informs the user about the file removal step
            # Find files in 'PDFs/' directory: type f (file), matching *.pdf (iname), size greater than 100MB (+100M), print the name, and delete them
            find PDFs/ -type f -iname "*.pdf" -size +100M -print -delete

            git pull # Fetches and merges changes from the remote repository to the local branch

            echo "Adding all changes..." # Informs the user about staging changes
            git add .  # Stages all changes in the current directory (new, modified, deleted) for the next commit

            # Create a commit message with a timestamp
            timestamp=$(date +"%Y-%m-%d %H:%M:%S") # Stores the current date and time in the specified format into the 'timestamp' variable
            message="updated $timestamp" # Creates the full commit message using the timestamp

            echo "📝 Committing changes with message: \"$message\"" # Informs the user of the commit message
            if git commit -m "$message"; then # Attempts to commit the staged changes with the generated message; checks if the commit was successful
                echo "🚀 Pushing committed changes to remote repository..." # Informs the user about the push attempt
                if git push; then # Attempts to push the committed changes to the remote repository; checks if the push was successful
                    echo "All changes pushed successfully." # Output if the push succeeds
                else
                    echo "Failed to push changes to remote. Please check your network or remote settings." # Output if the push fails
                fi
            else
                echo "Commit failed. There might be no changes to commit or another issue." # Output if the commit fails
            fi
        fi

        # Sleep before checking again
        echo "⏳ Sleeping for 5 seconds before next check..." # Informs the user that the script is pausing
        sleep 5s # Pauses the script execution for 15 seconds
    done # Ends the 'while true' loop
} # Ends the definition of the function

# Call the function
auto_git_push # Executes the defined function to start the automated process
