# Hello World

- [Hello World title](#hello-world-title)
- [Hello World title & video](#hello-world-title--video)
- [Hello World animated GIF](#hello-world-animated-gif)

## Hello World title

[Official example](https://shotstack.io/templates/hello-world-title/)

```json
{
    "timeline": {
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "Hello World",
                            "style": "future"
                        },
                        "start": 0,
                        "length": 5
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "mp4",
        "resolution": "hd"
    }
}
```

## Hello World title & video

[Official example](https://shotstack.io/templates/hello-world-title-video/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/motions.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "Hello World",
                            "style": "future",
                            "position": "left",
                            "size": "x-small"
                        },
                        "start": 4,
                        "length": 11,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        },
                        "effect": "zoomIn"
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/earth.mp4",
                            "trim": 5
                        },
                        "start": 0,
                        "length": 15,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "mp4",
        "resolution": "hd"
    }
}
```

## Hello World animated GIF

[Official example](https://shotstack.io/templates/hello-world-title-video-gif/)

```json
{
    "timeline": {
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "Hello World",
                            "style": "future",
                            "position": "left",
                            "size": "x-small"
                        },
                        "start": 4,
                        "length": 11,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        },
                        "effect": "zoomIn"
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/earth.mp4",
                            "trim": 5
                        },
                        "start": 0,
                        "length": 15,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "gif",
        "resolution": "mobile",
        "quality": "low",
        "fps": 12
    }
}
```
