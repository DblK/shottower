# Ken burns

- [Ken burns slideshow effect](#ken-burns-slideshow-effect)
- [Ken burns slideshow effect (slow)](#ken-burns-slideshow-effect-slow)
- [Ken burns slideshow effect (fast)](#ken-burns-slideshow-effect-fast)

## Ken burns slideshow effect

[Official example](https://shotstack.io/templates/ken-burns-effect-slideshow/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/advertising.mp3",
            "effect": "fadeInFadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://images.pexels.com/photos/189349/pexels-photo-189349.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 0,
                        "length": 5,
                        "effect": "zoomIn",
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
                            "src": "https://images.pexels.com/photos/1680140/pexels-photo-1680140.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 4,
                        "length": 5,
                        "effect": "slideUp",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/847393/pexels-photo-847393.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 8,
                        "length": 5,
                        "effect": "slideLeft",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1295138/pexels-photo-1295138.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 12,
                        "length": 5,
                        "effect": "zoomOut",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1452701/pexels-photo-1452701.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 16,
                        "length": 5,
                        "effect": "slideDown",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/756856/pexels-photo-756856.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 20,
                        "length": 5,
                        "effect": "slideRight",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1533720/pexels-photo-1533720.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 24,
                        "length": 5,
                        "effect": "zoomIn",
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

## Ken burns slideshow effect (slow)

[Official example](https://shotstack.io/templates/ken-burns-effect-slideshow-slow/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/advertising.mp3",
            "effect": "fadeInFadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://images.pexels.com/photos/189349/pexels-photo-189349.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 0,
                        "length": 6,
                        "effect": "zoomInSlow",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1680140/pexels-photo-1680140.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 4,
                        "length": 6,
                        "effect": "slideUpSlow",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/847393/pexels-photo-847393.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 8,
                        "length": 6,
                        "effect": "slideLeftSlow",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1295138/pexels-photo-1295138.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 12,
                        "length": 6,
                        "effect": "zoomOutSlow",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1452701/pexels-photo-1452701.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 16,
                        "length": 6,
                        "effect": "slideDownSlow",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/756856/pexels-photo-756856.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 20,
                        "length": 6,
                        "effect": "slideRightSlow",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1533720/pexels-photo-1533720.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 24,
                        "length": 6,
                        "effect": "zoomInSlow",
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

## Ken burns slideshow effect (fast)

[Official example](https://shotstack.io/templates/ken-burns-effect-slideshow-fast/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/advertising.mp3",
            "effect": "fadeInFadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://images.pexels.com/photos/189349/pexels-photo-189349.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 0,
                        "length": 5,
                        "effect": "zoomInFast",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1680140/pexels-photo-1680140.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 4.5,
                        "length": 5,
                        "effect": "slideUpFast",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/847393/pexels-photo-847393.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 9,
                        "length": 5,
                        "effect": "slideLeftFast",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1295138/pexels-photo-1295138.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 13.5,
                        "length": 5,
                        "effect": "zoomOutFast",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1452701/pexels-photo-1452701.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 18,
                        "length": 5,
                        "effect": "slideDownFast",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/756856/pexels-photo-756856.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 22.5,
                        "length": 5,
                        "effect": "slideRightFast",
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
                            "type": "image",
                            "src": "https://images.pexels.com/photos/1533720/pexels-photo-1533720.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=720&w=1280"
                        },
                        "start": 27,
                        "length": 5,
                        "effect": "zoomInFast",
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