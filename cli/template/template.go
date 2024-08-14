package template

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type Selections struct {
	ProjectName  string
	Default      bool
	Minimalistic bool
	Auth         bool
	DatabaseSQL  bool
	Javascript   bool
	Dockerfile   bool
	Htmx         bool
	Tailwind     bool
}

func GenerateTemplate(selections Selections) error {
	selections.ProjectName = parseToPath(selections.ProjectName)
	var projectDir = filepath.Join("..", selections.ProjectName)

	err := createMainFiles(projectDir, selections)
	if err != nil {
		return err
	}
	err = createNonMinFiles(projectDir, selections)
	if err != nil {
		return err
	}
	err = createSQLFiles(projectDir, selections)
	if err != nil {
		return err
	}
	err = createAuthFiles(projectDir, selections)
	if err != nil {
		return err
	}
	err = createTailwindFiles(projectDir, selections)
	if err != nil {
		return err
	}
	if err := runCommand(projectDir, "go", "mod", "init", selections.ProjectName); err != nil {
		return fmt.Errorf("failed to initialize Go module: %w", err)
	}
	if err := runCommand(projectDir, "go", "mod", "tidy"); err != nil {
		return fmt.Errorf("failed to tidy Go module: %w", err)
	}
	
	pubDir := "public/scripts"
	if selections.Minimalistic {
		pubDir = "public"
	}
	if err := os.MkdirAll(path.Join(projectDir, pubDir), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", pubDir, err)
	}
	
	if (selections.Htmx || selections.Default) && !selections.Minimalistic {
		if err := runCommand(
			projectDir,
			"curl",
			"-sLo",
			"public/scripts/htmx.min.js",
			"https://cdn.jsdelivr.net/npm/htmx.org/dist/htmx.min.js"); err != nil {
			return fmt.Errorf("failed to download HTMX script: %w", err)
		}
	}
	
	if (selections.Javascript || selections.Default) && !selections.Minimalistic {
		if err := runCommand(
			projectDir,
			"curl",
			"-sLo",
			"public/scripts/alpine.js",
			"https://cdn.jsdelivr.net/npm/alpinejs/dist/cdn.min.js"); err != nil {
			return fmt.Errorf("failed to download Alpine.js script: %w", err)
		}
		if err := runCommand(
			projectDir,
			"curl",
			"-sLo",
			"public/scripts/jquery.min.js",
			"https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js"); err != nil {
			return fmt.Errorf("failed to download jQuery script: %w", err)
		}
	}
	
	if (selections.Tailwind || selections.Default) && !selections.Minimalistic {
		if err := runCommand(
			projectDir,
			"curl",
			"-sLo",
			"tailwindcss",
			"https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64"); err != nil {
			return fmt.Errorf("failed to download TailwindCSS binary: %w", err)
		}
		if err := runCommand(projectDir, "chmod", "+x", "tailwindcss"); err != nil {
			return fmt.Errorf("failed to make TailwindCSS binary executable: %w", err)
		}
		if err := runCommand(projectDir, "./tailwindcss", "-i", "tailwind/css/app.css", "-o", "public/styles.css"); err != nil {
			return fmt.Errorf("failed to generate TailwindCSS styles: %w", err)
		}
	}	
	return nil
}

func createFile(projectDir, fileName, ext, srcPath, outputPath string, selections Selections) error {
	t, err := template.New(fileName + ".txt").ParseFiles(srcPath)
	if err != nil {
		return err
	}
	var tpl bytes.Buffer

	err = t.Execute(&tpl, selections)
	if err != nil {
		return err
	}
	dirPath := filepath.Join(projectDir, outputPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(dirPath, fmt.Sprintf("%s.%s", fileName, ext))
	m, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer m.Close()

	_, err = io.Copy(m, &tpl)
	if err != nil {
		return err
	}
	return nil
}

func createFileNoExt(projectDir, fileName, srcPath, outputPath string, selections Selections) error {
	t, err := template.New(fileName + ".txt").ParseFiles(srcPath)
	if err != nil {
		return err
	}
	var tpl bytes.Buffer

	err = t.Execute(&tpl, selections)
	if err != nil {
		return err
	}
	dirPath := filepath.Join(projectDir, outputPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(dirPath, fileName)
	m, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer m.Close()

	_, err = io.Copy(m, &tpl)
	if err != nil {
		return err
	}
	return nil
}

func createMainFiles(projectDir string, selections Selections) error {
	err := createFile(projectDir, "main", "go", "sample/main_go/main.txt", "", selections)
	if err != nil {
		return err
	}
	err = createFile(projectDir, "config", "go", "sample/config_go/config.txt", "config", selections)
	if err != nil {
		return err
	}

	err = createFile(projectDir, "middleware", "go", "sample/middleware_go/middleware.txt", "handlers/middleware", selections)
	if err != nil {
		return err
	}

	err = createFile(projectDir, "render", "go", "sample/render_go/render.txt", "handlers/render", selections)
	if err != nil {
		return err
	}

	err = createFile(projectDir, "base", "templ", "sample/views_templ/layouts/base.txt", "views/layouts", selections)
	if err != nil {
		return err
	}
	err = createFile(projectDir, ".air", "toml", "sample/air/.air.txt", "", selections)
	if err != nil {
		return err
	}
	err = createFileNoExt(projectDir, "Makefile", "sample/makefile/Makefile.txt", "", selections)
	if err != nil {
		return err
	}
	if selections.Dockerfile || selections.Default {
		err = createFileNoExt(projectDir, "Dockerfile", "sample/dockerfile/Dockerfile.txt", "", selections)
		if err != nil {
			return err
		}
	}
	err = createFileNoExt(projectDir, ".gitignore", "sample/gitignore/.gitignore.txt", "", selections)
	if err != nil {
		return err
	}
	return nil
}

func createNonMinFiles(projectDir string, selections Selections) error {
	if !selections.Minimalistic {
		err := createFile(projectDir, "static_dev", "go", "sample/statics_go/static_dev.txt", "", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "static_prod", "go", "sample/statics_go/static_prod.txt", "", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "handleHome", "go", "sample/handlers_go/handleHome.txt", "handlers", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "components", "templ", "sample/views_templ/components/components.txt", "views/components", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "index", "templ", "sample/views_templ/home/index.txt", "views/home", selections)
		if err != nil {
			return err
		}
		err = createFileNoExt(projectDir, ".env", "sample/env/.env.txt", "", selections)
		if err != nil {
			return err
		}
	}
	return nil
}

func createSQLFiles(projectDir string, selections Selections) error {
	if (selections.DatabaseSQL || selections.Default) && !selections.Minimalistic {
		err := createFile(projectDir, "store", "go", "sample/store_go/store.txt", "services/store", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "storage", "go", "sample/store_go/storage.txt", "services/store", selections)
		if err != nil {
			return err
		}

	}
	return nil
}

func createAuthFiles(projectDir string, selections Selections) error {
	if (selections.Auth || selections.Default) && !selections.Minimalistic {
		err := createFile(projectDir, "auth", "go", "sample/auth_go/auth.txt", "services/auth", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "session", "go", "sample/auth_go/session.txt", "services/auth", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "handleAuth", "go", "sample/handlers_go/handleAuth.txt", "handlers", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "index", "templ", "sample/views_templ/login/index.txt", "views/login", selections)
		if err != nil {
			return err
		}
	}
	return nil
}

func createTailwindFiles(projectDir string, selections Selections) error {
	if (selections.Tailwind || selections.Default) && !selections.Minimalistic {
		err := createFile(projectDir, "app", "css", "sample/tailwind/css/app.txt", "tailwind/css", selections)
		if err != nil {
			return err
		}
		err = createFile(projectDir, "tailwind.config", "js", "sample/tailwind/tailwind.config.txt", "", selections)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseToPath(str string) string {
	words := strings.Split(str, " ")
	if len(words) == 0 {
		return ""
	}
	for i := range words {
		words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
	}
	words[0] = strings.ToLower(words[0][:1]) + words[0][1:]
	return strings.Join(words, "")
}

func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run()
}
