# Cropping

- [Top & left crop](#top--left-crop)
- [Splitscreen crop effects](#splitscreen-crop-effects)
- [Splitscreen with zoom](#splitscreen-with-zoom)


## Top & left crop

[Official example](https://shotstack.io/templates/simple-crop-top-left/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/berlin.mp3"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "Crop Example",
                            "style": "minimal",
                            "size": "small",
                            "position": "top"
                        },
                        "start": 0,
                        "length": 4,
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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/wave-barrel.jpg",
                            "crop": {
                                "top": 0.5,
                                "bottom": 0,
                                "left": 0,
                                "right": 0
                            }
                        },
                        "start": 0,
                        "length": 4,
                        "effect": "zoomIn"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4",
                            "crop": {
                                "top": 0,
                                "bottom": 0,
                                "left": 0.5,
                                "right": 0
                            }
                        },
                        "start": 4,
                        "length": 4
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

## Splitscreen crop effects

[Official example](https://shotstack.io/templates/splitscreen-crop-effects/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://feeds.soundcloud.com/stream/824693758-unminus-majesty.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "CROP DEMO",
                            "style": "future",
                            "size": "x-small",
                            "position": "bottomLeft"
                        },
                        "start": 0,
                        "length": 4,
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/surfers.mp4",
                            "trim": 5,
                            "crop": {
                                "left": 0.25,
                                "right": 0.25
                            }
                        },
                        "start": 0,
                        "length": 4,
                        "transition": {
                            "in": "carouselDown",
                            "out": "carouselUp"
                        },
                        "offset": {
                            "x": -0.25
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach.mp4",
                            "crop": {
                                "top": 0.248,
                                "bottom": 0.25
                            }
                        },
                        "start": 4,
                        "length": 4,
                        "offset": {
                            "y": 0.25
                        },
                        "transition": {
                            "in": "carouselLeft",
                            "out": "carouselRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/city-timelapse.mp4",
                            "crop": {
                                "top": 0.25,
                                "bottom": 0.25,
                                "right": 0.25
                            }
                        },
                        "start": 8,
                        "length": 4,
                        "offset": {
                            "x": 0.25
                        },
                        "scale": 0.8,
                        "transition": {
                            "in": "carouselLeft",
                            "out": "carouselRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/beach.mp4",
                            "crop": {
                                "left": 0.375,
                                "right": 0.375
                            }
                        },
                        "start": 12,
                        "length": 4,
                        "offset": {
                            "x": -0.375
                        },
                        "transition": {
                            "in": "carouselUp",
                            "out": "carouselDown"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/surfers.mp4",
                            "trim": 10,
                            "crop": {
                                "left": 0.375,
                                "right": 0.375
                            }
                        },
                        "start": 12.3,
                        "length": 4,
                        "offset": {
                            "x": 0.375
                        },
                        "transition": {
                            "in": "carouselUp",
                            "out": "carouselDown"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/skier.mp4",
                            "crop": {
                                "left": 0.25,
                                "right": 0.25
                            }
                        },
                        "start": 0,
                        "length": 4,
                        "offset": {
                            "x": 0.25
                        },
                        "transition": {
                            "in": "carouselUp",
                            "out": "carouselDown"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/snow-forest.mp4",
                            "crop": {
                                "top": 0.25,
                                "bottom": 0.248
                            }
                        },
                        "start": 4,
                        "length": 4,
                        "offset": {
                            "y": -0.25
                        },
                        "transition": {
                            "in": "carouselRight",
                            "out": "carouselLeft"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/mountains.mp4",
                            "crop": {
                                "left": 0.25,
                                "right": 0.25
                            }
                        },
                        "start": 8,
                        "length": 4,
                        "offset": {
                            "x": -0.25
                        },
                        "scale": 0.8,
                        "transition": {
                            "in": "carouselUp",
                            "out": "carouselDown"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/skier.mp4",
                            "crop": {
                                "left": 0.375,
                                "right": 0.375
                            }
                        },
                        "start": 12.1,
                        "length": 4,
                        "offset": {
                            "x": -0.125
                        },
                        "transition": {
                            "in": "carouselUp",
                            "out": "carouselDown"
                        }
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/mountains.mp4",
                            "crop": {
                                "left": 0.375,
                                "right": 0.375
                            }
                        },
                        "start": 12.2,
                        "length": 4,
                        "offset": {
                            "x": 0.125
                        },
                        "transition": {
                            "in": "carouselUp",
                            "out": "carouselDown"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp</p>",
                            "width": 1,
                            "height": 1,
                            "background": "#000000"
                        },
                        "start": 16.4,
                        "length": 0.04
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

## Splitscreen with zoom

[Official example](https://shotstack.io/templates/splitscreen-zoom-effect/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://feeds.soundcloud.com/stream/824693758-unminus-majesty.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "title",
                            "text": "CROP ZOOM DEMO",
                            "style": "future",
                            "size": "x-small",
                            "position": "bottomLeft"
                        },
                        "start": 0,
                        "length": 4,
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
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/surfers.mp4",
                            "crop": {
                                "left": 0.25,
                                "right": 0.25
                            }
                        },
                        "start": 0,
                        "length": 5,
                        "offset": {
                            "x": -0.25
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/footage/surfers.mp4",
                            "crop": {
                                "top": 0.248,
                                "bottom": 0.25
                            }
                        },
                        "start": 0,
                        "length": 5,
                        "offset": {
                            "x": 0.25
                        },
                        "scale": 2
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