Tailwind:
npx @tailwindcss/cli -i ./static/css/style.css -o ./static/css/tailwind.css --watch

Templ:
templ generate --watch

Air:
air

chromium-browser http://localhost:3000 --headless --screenshot=test.png --window-size="480,890" --disable-gpu --no-sandbox
