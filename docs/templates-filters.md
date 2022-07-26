# Filters

- [All filters](#all-filters)
- [Boost filter](#boost-filter)
- [Contrast filter](#contrast-filter)
- [Darken filter](#darken-filter)
- [Greyscale filter](#greyscale-filter)
- [Lighten filter](#lighten-filter)
- [Muted filter](#muted-filter)
- [Negative fIlter](#negative-filter)

## All filters

[Official example](https://shotstack.io/templates/all-filters-sequence/)

```json
{
    "timeline": {
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "original",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 0,
                        "length": 2,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "boost",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 2,
                        "length": 3,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "contrast",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 5,
                        "length": 3,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "muted",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 8,
                        "length": 3,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "darken",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 11,
                        "length": 3,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "lighten",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 14,
                        "length": 3,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "greyscale",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 17,
                        "length": 3,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "title",
                            "text": "negative",
                            "style": "minimal",
                            "size": "x-small"
                        },
                        "start": 20,
                        "length": 3,
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
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 0
                        },
                        "start": 0,
                        "length": 3
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 2
                        },
                        "start": 2,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "boost"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 5
                        },
                        "start": 5,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "contrast"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 8
                        },
                        "start": 8,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "muted"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 11
                        },
                        "start": 11,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "darken"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 14
                        },
                        "start": 14,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "lighten"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 17
                        },
                        "start": 17,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "greyscale"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4",
                            "trim": 20
                        },
                        "start": 20,
                        "length": 4,
                        "transition": {
                            "in": "reveal"
                        },
                        "filter": "negative"
                    }
                ]
            }
        ],
        "background": "#000000",
        "soundtrack": {
            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/music/freeflow.mp3",
            "effect": "fadeInFadeOut"
        }
    },
    "output": {
        "format": "mp4",
        "resolution": "hd"
    }
}
```

## Boost filter

[Official example](https://shotstack.io/templates/boost-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "boost"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "boost"
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

## Contrast filter

[Official example](https://shotstack.io/templates/contrast-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "contrast"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "contrast"
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

## Darken filter

[Official example](https://shotstack.io/templates/darken-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "darken"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "darken"
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

## Greyscale filter

[Official example](https://shotstack.io/templates/greyscale-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "greyscale"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "greyscale"
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

## Lighten filter

[Official example](https://shotstack.io/templates/lighten-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "lighten"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "lighten"
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

## Muted filter

[Official example](https://shotstack.io/templates/muted-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "muted"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "muted"
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

## Negative fIlter

[Official example](https://shotstack.io/templates/negative-filter/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg"
                        },
                        "start": 0,
                        "length": 3,
                        "filter": "negative"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 3,
                        "length": 3,
                        "filter": "negative"
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