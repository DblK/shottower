# Basic edits

- [Title, image & video](#title-image--video)
- [1 minute video](#1-minute-video)

## Title, image & video

[Official example](https://shotstack.io/templates/basic-edits-title-image-video/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/berlin.mp3"
        },
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "Hello World",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 0,
                        "length": 5,
                        "transition": {
                            "in": "slideRight",
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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 5,
                        "effect": "zoomIn"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4",
                            "trim": 5
                        },
                        "start": 5,
                        "length": 5
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "mp4",
        "resolution": "sd"
    }
}
```

## 1 minute video

[Official example](https://shotstack.io/templates/one-minute-video-edit/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "1 MINUTE SAMPLE VIDEO",
                            "size": "large",
                            "style": "chunk"
                        },
                        "start": 0,
                        "length": 10,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/cars-tunnel.mp4"
                        },
                        "start": 0,
                        "length": 5,
                        "transition": {
                            "in": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 5,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 10,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/city-timelapse.mp4"
                        },
                        "start": 15,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/drone.mp4"
                        },
                        "start": 20,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/headland-houses.mp4"
                        },
                        "start": 25,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 30,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/night-sky.mp4"
                        },
                        "start": 35,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/rain-on-roof.mp4"
                        },
                        "start": 40,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 45,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/skate-park.mp4"
                        },
                        "start": 50,
                        "length": 5
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/skateboarding.mp4"
                        },
                        "start": 55,
                        "length": 5,
                        "transition": {
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
