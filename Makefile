PORT ?= 8000
run: build
	@./bin/app
build:
	@~/go/bin/templ generate && \
	 npx tailwindcss -i tailwind/css/app.css -o public/styles.css && \
		 go build -o bin/app .
css:
	@npx tailwindcss -i tailwind/css/app.css -o public/styles.css --watch
templ:
	~/go/bin/templ generate --watch --proxy=http://localhost:$(PORT)
build-js:
	@curl -sLo public/scripts/htmx.min.js https://cdn.jsdelivr.net/npm/htmx.org/dist/htmx.min.js && \
	curl -sLo public/scripts/alpine.js https://cdn.jsdelivr.net/npm/alpinejs/dist/cdn.min.js && \
	curl -sLo public/scripts/jquery.min.js https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js
tailwind:
	@npm install -D tailwindcss
bundle-all: build-js build-css
	@echo "Bundling complete!"
build-js:
	@npx esbuild public/modules/main.js --minify --outfile=public/scripts/bundle.min.js
build-css:
	@npx esbuild public/styles/main.css --minify --outfile=public/styles/bundle.min.css
npm-pkg:
	@npm install -D tailwindcss esbuild