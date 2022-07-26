# Dissolves

- [Dissolve (fade) tracks](#dissolve-fade-tracks)
- [Slow dissolve (fade) tracks](#slow-dissolve-fade-tracks)
- [Fast dissolve (fade) tracks](#fast-dissolve-fade-tracks)
- [Fade to/from black](#fade-tofrom-black)

## Dissolve (fade) tracks

[Official example](https://shotstack.io/templates/dissolve-fade-video-clips/)

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
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 0,
                        "length": 5,
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 4,
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 8,
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

## Slow dissolve (fade) tracks

[Official example](https://shotstack.io/templates/dissolve-fade-video-clips-slow/)

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
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 0,
                        "length": 6,
                        "transition": {
                            "in": "fadeSlow",
                            "out": "fadeSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 4,
                        "length": 6,
                        "transition": {
                            "out": "fadeSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 8,
                        "length": 6,
                        "transition": {
                            "out": "fadeSlow"
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

## Fast dissolve (fade) tracks

[Official example](https://shotstack.io/templates/dissolve-fade-video-clips-fast/)

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
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 0,
                        "length": 5,
                        "transition": {
                            "in": "fadeFast",
                            "out": "fadeFast"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 4.5,
                        "length": 5,
                        "transition": {
                            "out": "fadeFast"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 9,
                        "length": 5,
                        "transition": {
                            "out": "fadeFast"
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

## Fade to/from black

[Official example](https://shotstack.io/templates/dissolve-fade-to-from-black/)

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
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach-cliffs.mp4"
                        },
                        "start": 0,
                        "length": 5,
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/road.mp4"
                        },
                        "start": 5,
                        "length": 5,
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/table-mountain.mp4"
                        },
                        "start": 10,
                        "length": 5,
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

