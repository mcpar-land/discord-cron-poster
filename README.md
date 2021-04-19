# Discord Cron Poster

You, a discriminating Discord user, have an important schedule to keep to.

- Wednesday is [Wednesday, My Dudes](https://www.youtube.com/watch?v=du-TY1GUFGk)
- Thursday is [Out of Touch Thursday](https://knowyourmeme.com/memes/out-of-touch-thursday)
- Friday is [Flat Fuck Friday](https://knowyourmeme.com/memes/flat-fuck-friday)
- Christmas
- [September 21st](https://www.youtube.com/watch?v=CG7YHFT4hjw)

And probably many more.

Naturally, missing a single post would be a disaster, but it's a hassle to keep to.

This microservice will post Discord Webhook messages on a regular schedule, according to standard cron expressions.

## Use

download the executable, or compile it yourself.

Create a file called `config.json` in the same directory as the executable. Here's an example:

```json
{
  "url": "(Your webhook URL goes here)",
  "jobs": [
    {
      "cron": "0 0 * * FRI",
      "webhook": {
        "content": "This message will post every friday!"
      }
    },
    {
      "cron": "0 0 * * WED",
      "webhook": {
        "content": "This message will post every wednesday!"
      }
    },
    {
      "cron": "0 0 25 12 *",
      "webhook": {
        "content": "Merry christmas!"
      }
    }
  ]
}
```

Run the executable from the console.

```
./discord-cron-poster
```

When a specified cron occurs, it will post the webhook associated with it.

## Uploading media

To upload media with your webhook, refer to a file relative to the config file.
If your folder has these files:

```
- discord-cron-poster.exe
- config.json
- finally_friday.mp4
```

This config will upload `finally_friday.mp4` every friday.

```json
{
  "url": "(Your webhook URL goes here)",
  "jobs": [
    {
      "cron": "0 0 * * FRI",
      "webhook": {
        "content": "This message will post every friday!",
        "file": "./finally_friday.mp4"
      }
    }
  ]
}
```

You can also post embeds in accordance with Discord's [Embed Format](https://discord.com/developers/docs/resources/channel#embed-object)

```json
{
  "url": "(Your webhook URL goes here)",
  "jobs": [
    {
      "cron": "@every 1d",
      "webhook": {
        "content": "Check out my cool embeds today!",
        "embeds": [
          {
            "content": "This is a cool embed"
          },
          {
            "title": "This is a cooler embed",
            "description": "This is the embed description",
            "color": 16711935,
            "fields": [
              {
                "name": "Field 1 name",
                "value": "Field 1 value"
              },
              {
                "name": "Field 2 name",
                "value": "Field 2 value",
                "inline": true
              }
            ]
          }
        ]
      }
    }
  ]
}
```

## Flags and Options

These flags can go both in the top-level of your config file, and work as command line flags.

### `url`

Changes the Webhook URL.

```
./discord-cron-poster --url https://your-discord-webhook-url-here
```

```json
{
	"url": "https://your-discord-webhook-url-here",
	"jobs": [...]
}
```

### `media`

Changes the directory to search for media files, relative to the location of the config file.

```
- discord-cron-poster.exe
- config.json
- media_folder
	- meme.png
	- meme2.png
```

```
./discord-cron-poster --media ./media_folder
```

```json
{
  "url": "(url here)",
  "media": "./media_folder",
  "jobs": [
    {
      "cron": "@every 1d",
      "webhook": {
        "file": "./meme2.png"
      }
    }
  ]
}
```

### `tz`

Changes the cron timezone.

```
./discord-cron-poster --tz America/New_York
```

```json
{
	"url": "(url here)",
	"tz": "America/New_York",
	"jobs": [...]
}
```

### `config`

Changes the location of the config file. (This one can't go in the config file itself, obviously)

```
./discord-cron-poster --config ./config/file/path/other_config.json
```

## References

- [Cron Format Used](https://pkg.go.dev/github.com/robfig/cron/v3#hdr-CRON_Expression_Format)
- [Embed Builder](https://embedbuilder.nadekobot.me/)

---

(I mostly wrote this as a practice project while learning Go, but maybe somebody out there will find use for it.)
