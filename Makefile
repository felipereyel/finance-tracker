all: statics templ

statics:
	curl https://cdn.tailwindcss.com/3.4.16 --output internal/embeded/assets/tailwind.js
	curl https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js --output internal/embeded/assets/htmx.js
	curl https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js --output internal/embeded/assets/echarts.min.js
	curl https://go-echarts.github.io/go-echarts-assets/assets/themes/wonderland.js --output internal/embeded/assets/wonderland.js

templ:
	go tool templ generate
