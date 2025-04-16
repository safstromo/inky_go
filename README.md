## Running

1. Clone project:

```
git clone https://github.com/safstromo/inky_go
```

2. Move to project directory

```
cd inky_go
```

3. Create a symlink to image directory:

```
ln -s /mnt/directory_with_images static/images
```

4. Build project

```
go build
```

5. Run project

```
./inky_go
```

## Development

### Hotreload commmands

Tailwind:
npx @tailwindcss/cli -i ./static/css/style.css -o ./static/css/tailwind.css --watch

Templ:
templ generate --watch

Air:
air
