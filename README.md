# inky_go

This an application for the Inky Imperssion 7.3" with a Raspberry Pi Zero.

Its a webserver that serves an HTMX page with images from a folder and in intervalls takes a screenshot of the page then runs a script to update the frame with the screenshot.

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

Set envs to build for raspberrypi zero if you compile on another system then the pi.

```
export GOOS=linux
export GOARCH=arm
export GOARM=7
```

Build

```
go build
```

5. Run project

```
./inky_go
```

The binary needs thes following folders: static, scripts

## Misc

In the script folder there are a couple of scripts to update the frame without the webserver.

You can run update_from_folder.sh in a set intervall using systemd by adding the service and timer files to /etc/systemd/system then enable the timer service.

## Development

### Hotreload commmands

Tailwind:
npx @tailwindcss/cli -i ./static/css/style.css -o ./static/css/tailwind.css --watch

Templ:
templ generate --watch

Air:
air
