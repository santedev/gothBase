GothBase

GothBase is a Go template optimized for Linux development, using Air and a Makefile, along with the Goth and Gotth stack.
Installation
On Linux

    Install Golang: Follow the instructions here.

    Install Templ:
        For installation, follow the guide here, or run:

        bash

    go install github.com/a-h/templ/cmd/templ@latest

Build the CLI:

    To get the binary:

    bash

go build -o ../GothBase

Or run:

bash

        cd cli && go run .

    Note: The output directory will be one level up from the cli folder.

    Requirements: Ensure you have Golang and Linux. WSL can work if it supports curl, go commands, and chmod. Makefile and Air work with Bash.

On Windows (PowerShell)

    Run the CLI:
        Navigate to the cli directory:

        powershell

cd cli
go run .

Or build the binary:

powershell

    cd cli
    go build -o ../GothBase
    ../GothBase

Setup Project:

    Set your project name in the first input.
    Use Enter to go to the next step, and Ctrl+U to go back.
    Select options using keyboard arrows (up/down) and Enter to select. Use k for moving up, j for moving down, n for next step, and b for the previous step. Use Space or Enter to select options. Press y or n to confirm or cancel.

Post-Setup:

    If successful, run:

    powershell

templ generate

If errors occur, try:

powershell

    ~/go/bin/templ generate

Manual Installation for PowerShell:

    Install TailwindCSS:

    powershell

Invoke-WebRequest -Uri "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-win-x64.zip" -OutFile "tailwindcss.zip"
Expand-Archive -Path "tailwindcss.zip" -DestinationPath "."
Remove-Item -Path "tailwindcss.zip"

Install JS Min Files:

powershell

        Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/htmx.org/dist/htmx.min.js" -OutFile "public/scripts/htmx.min.js"
        Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/alpinejs/dist/cdn.min.js" -OutFile "public/scripts/alpine.js"
        Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js" -OutFile "public/scripts/jquery.min.js"

For detailed command-line navigation, refer to the Command Line Navigation section.
Command Line Navigation

    Move Up/Down: Use arrow keys or k/j.
    Next Step: Arrow right or n.
    Previous Step: Arrow left or b.
    Select Options: Enter or Space.
    Confirm Creation: Press y to confirm or n to cancel.
