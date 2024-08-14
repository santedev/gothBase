
# GothBase

GothBase is a Go template optimized for Linux development, using Air and a Makefile, along with the Goth and Gotth stack.

## Requirements: 
Ensure you have Golang and Linux. WSL may work if it supports curl, go commands, and chmod. 

Makefile and Air work with Bash commands.

## Installation steps
- clone the project with the next command: 
```bash
git clone https://github.com/santedev/gothBase
```

- Install Golang [here](https://go.dev/doc/install)

- Install Templ [here](https://templ.guide/quick-start/installation/)

Or run next command to install templ:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
## Run CLI
- Build the CLI:
```bash
go build -o ./GothBase
```
```bash
./GothBase
```
Or run next command:

```bash
go run .
```
**Note**: The output directory will be one level up from the cli cloned folder.
## Command Line Navigation
Command Line Navigation when input is required
- **Next Step**: Key **`Enter`**.
- **Previous Step**: Key **`ctrl+u`**.
- **Confirm Creation**: Key **`Enter`**. 

Command Line Navigation

- **Move Up/Down**: Use arrow keys or **`k`** and **`j`**.
- **Next Step**: Arrow **`right`** or **`n`**.
- **Previous Step**: Arrow **`left`** or **`b`**.
- **Select Options**: **`Enter`** or **`Space`**.
- **Confirm Creation**: Press **`y`** to confirm or **`n`** 
    to cancel.
## Post-Setup
with templ installed run next command 
```bash
templ generate
```
if errors encountered try
```bash
~/go/bin/templ generate
```
## For windows users

Here's how you can update the manual installation instructions for Windows, specifically for PowerShell:
Manual Installation for Windows (PowerShell)

- Install TailwindCSS:

```powershell
Invoke-WebRequest -Uri "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-win-x64.zip" -OutFile "tailwindcss.zip"
Expand-Archive -Path "tailwindcss.zip" -DestinationPath "."
Remove-Item -Path "tailwindcss.zip"
```
- Install JS Min Files:

```powershell
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/htmx.org/dist/htmx.min.js" -OutFile "public/scripts/htmx.min.js"
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/alpinejs/dist/cdn.min.js" -OutFile "public/scripts/alpine.js"
Invoke-WebRequest -Uri "https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js" -OutFile "public/scripts/jquery.min.js"
```