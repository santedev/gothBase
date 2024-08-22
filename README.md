
# GothBase

GothBase is a Go templating system optimized for Linux development, using Air and a Makefile, along with the Goth and Gotth stack.

- you can see the structure of the template at [example](https://github.com/santedev/gothBase/tree/example)
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
cd cli
```
```bash
go build -o ../GothBase
```
```bash
../GothBase
```
Or run next command:

```bash
cd cli
```
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

## Script and Style Customization

In the **public/scripts** directory, you'll find files like htmx.min.js, jquery.min.js, and alpine.js, depending on the options you selected during setup. For example, if you only need HTMX, you can opt for that by selecting all options except JavaScript and then choosing HTMX.

With the default setup, which includes JavaScript minified files, you're equipped to use jQuery's ecosystem if needed or leverage Alpine UI libraries with Tailwind UI. Explore [Pines JS](https://devdojo.com/pines), [Penguin UI](https://www.penguinui.com/), and [Alpine UI](https://alpinejs.dev/components) Components for more.

To add your own JavaScript scripts, place them in the **public/scripts** folder for consistency. Then, include them in **views/layout/** if you want to add the script to the base layout, either within the HTML `<head>` or `<body>` tags using the Templ component.

Similarly, the public directory is intended for your CSS files. You can create a css folder and add your stylesheets there, following a similar process as with JavaScript.

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
## CLI Screenshots
![Screenshot from 2024-08-14 12-47-21](https://github.com/user-attachments/assets/b83e15c6-a6dc-4a31-91f2-a4ab67a2c5d3)

![Screenshot from 2024-08-14 12-48-24](https://github.com/user-attachments/assets/c07a7ecd-6b7c-4b3e-ba87-c46326bff46c)
