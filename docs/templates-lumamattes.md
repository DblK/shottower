# Luma Mattes

- [Simple luma matte transitions](#simple-luma-matte-transitions)
- [Multiple luma matte transitions](#multiple-luma-matte-transitions)
- [Arrow luma matte transition](#arrow-luma-matte-transition)
- [Arrow (45deg) luma matte transition](#arrow-45deg-luma-matte-transition)
- [Wipe transition](#wipe-transition)
- [Paint transition](#paint-transition)
- [Snow effect using luma matte](#snow-effect-using-luma-matte)
- [Circle mask](#circle-mask)

## Simple luma matte transitions

[Official example](https://shotstack.io/templates/simple-luma-matte-transitions/)

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
                            "text": "Luma Matte Demo",
                            "style": "marker",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 3.5,
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
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/circles/center-double.mp4"
                        },
                        "start": 2.24,
                        "length": 1.76
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 4,
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/waves/double-vertical.mp4"
                        },
                        "start": 4.92,
                        "length": 1.32
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 2.24,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/grids/pattern-3-invert.mp4"
                        },
                        "start": 8.08,
                        "length": 1.84
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 4.92,
                        "length": 5,
                        "transition": {
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 8.08,
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

## Multiple luma matte transitions

[Official example](https://shotstack.io/templates/multiple-luma-matte-transitions/)

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
                            "text": "Luma Matte Demo",
                            "style": "marker",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 3.5,
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
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/examples/luma-mattes/radial.mp4"
                        },
                        "start": 2.24,
                        "length": 1.76
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 4,
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/examples/luma-mattes/blocks-in.mp4"
                        },
                        "start": 4.92,
                        "length": 1.32
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 2.24,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/examples/luma-mattes/lines.mp4"
                        },
                        "start": 8.08,
                        "length": 1.84
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 4.92,
                        "length": 5,
                        "transition": {
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 8.08,
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

## Arrow luma matte transition

[Official example](https://shotstack.io/templates/arrow-luma-matte-transitions/)

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
                            "text": "Luma Matte Arrow Transition",
                            "style": "future",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 3.5,
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
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-down.mp4"
                        },
                        "start": 3,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 4,
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-up.mp4"
                        },
                        "start": 6,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 3,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-right.mp4"
                        },
                        "start": 9,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 6,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-left.mp4"
                        },
                        "start": 12,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 9,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/night-sky.mp4"
                        },
                        "start": 12,
                        "length": 4,
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

## Arrow (45deg) luma matte transition

[Official example](https://shotstack.io/templates/arrow-45deg-luma-matte-transitions/)

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
                            "text": "Luma Matte Demo",
                            "style": "future",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 3.5,
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
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-down-45.mp4"
                        },
                        "start": 3,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 4,
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-up-45.mp4"
                        },
                        "start": 6,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 3,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-right-45.mp4"
                        },
                        "start": 9,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 6,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://templates.shotstack.io/basic/asset/video/luma/double-arrow/double-arrow-left-45.mp4"
                        },
                        "start": 12,
                        "length": 2
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 9,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/night-sky.mp4"
                        },
                        "start": 12,
                        "length": 4,
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

## Wipe transition

[Official example](https://shotstack.io/templates/wipe-luma-matte-transitions/)

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
                            "text": "Luma Matte Wipe Transition",
                            "style": "marker",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 3.5,
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
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/soft-wipes/left.mp4"
                        },
                        "start": 2.24,
                        "length": 1.76
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 4,
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/soft-wipes/right.mp4"
                        },
                        "start": 4.72,
                        "length": 1.52
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 2.24,
                        "length": 4
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/soft-wipes/left.mp4"
                        },
                        "start": 8.08,
                        "length": 1.84
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 4.92,
                        "length": 5,
                        "transition": {
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 8.08,
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

## Paint transition

[Official example](https://shotstack.io/templates/brush-luma-matte-transitions/)

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
                            "text": "Luma Matte Paint Demo",
                            "style": "marker",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 4.5,
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
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/examples/luma-mattes/paint-left.mp4"
                        },
                        "start": 3.6,
                        "length": 1.4
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 5,
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/examples/luma-mattes/paint-right.mp4"
                        },
                        "start": 7.2,
                        "length": 1.4
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 3.6,
                        "length": 5
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/lake.mp4"
                        },
                        "start": 7.2,
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

## Snow effect using luma matte

[Official example](https://shotstack.io/templates/snow-luma-matte-effect/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3",
            "effect": "fadeOut"
        },
        "background": "#ffffff",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "Merry Christmas",
                            "style": "marker",
                            "size": "small"
                        },
                        "start": 0,
                        "length": 4.5,
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
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/snowflakes.mov"
                        },
                        "start": 0,
                        "length": 10,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 0,
                        "length": 10
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

## Circle mask

[Official example](https://shotstack.io/templates/cicle-mask-luma-matte-effect/)

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
                            "type": "luma",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/luma-mattes/static/circle-sd.jpg"
                        },
                        "start": 0,
                        "length": 10
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4"
                        },
                        "start": 0,
                        "length": 10
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/night-sky.mp4"
                        },
                        "start": 0,
                        "length": 10
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