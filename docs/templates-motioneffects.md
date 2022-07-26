# Motion Effects

- [Slide text effect](#slide-text-effect)
- [Slide text effect (slow)](#slide-text-effect-slow)
- [Slide alternate left/right text affect](#slide-alternate-leftright-text-affect)
- [Slide up/down text effect](#slide-updown-text-effect)
- [Slide up/down logo text reveal](#slide-updown-logo-text-reveal)

## Slide text effect

[Official example](https://shotstack.io/templates/sliding-text-right-to-left/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>HELLO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 38px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.2,
                        "length": 5,
                        "transition": {
                            "in": "slideLeft",
                            "out": "slideLeft"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0.05
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>SHOTSTACK</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 38px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.3,
                        "length": 5,
                        "transition": {
                            "in": "slideLeft",
                            "out": "slideLeft"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>VIDEO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #25d3d0; font-size: 38px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.4,
                        "length": 5,
                        "transition": {
                            "in": "slideLeft",
                            "out": "slideLeft"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": -0.05
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

## Slide text effect (slow)

[Official example](https://shotstack.io/templates/slow-sliding-text-right-to-left/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>HELLO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 38px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.2,
                        "length": 5,
                        "transition": {
                            "in": "slideLeftSlow",
                            "out": "slideLeftSlow"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0.05
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>SHOTSTACK</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 38px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.4,
                        "length": 5,
                        "transition": {
                            "in": "slideLeftSlow",
                            "out": "slideLeftSlow"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>VIDEO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 38px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.6,
                        "length": 5,
                        "transition": {
                            "in": "slideLeftSlow",
                            "out": "slideLeftSlow"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": -0.05
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

## Slide alternate left/right text affect

[Official example](https://shotstack.io/templates/alternative-sliding-text-right-to-left/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>HELLO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 34px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.2,
                        "length": 5,
                        "transition": {
                            "in": "slideLeft",
                            "out": "slideRight"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0.05
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>SHOTSTACK</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 34px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.3,
                        "length": 5,
                        "transition": {
                            "in": "slideRight",
                            "out": "slideLeft"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>VIDEO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 34px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.4,
                        "length": 5,
                        "transition": {
                            "in": "slideLeft",
                            "out": "slideRight"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": -0.05
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

## Slide up/down text effect

[Official example](https://shotstack.io/templates/sliding-text-up-and-down/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>HELLO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 34px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.2,
                        "length": 5,
                        "transition": {
                            "in": "slideUp",
                            "out": "slideDown"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0.05
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>SHOTSTACK</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 34px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.3,
                        "length": 5,
                        "transition": {
                            "in": "slideUp",
                            "out": "slideDown"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>VIDEO</p>",
                            "css": "p { font-family: 'Montserrat ExtraBold'; color: #ffffff; font-size: 34px; text-align: left; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 1.4,
                        "length": 5,
                        "transition": {
                            "in": "slideUp",
                            "out": "slideDown"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": -0.05
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

## Slide up/down logo text reveal

[Official example](https://shotstack.io/templates/logo-subheading-slide-up-down/)

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
                            "type": "image",
                            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/branding/logo-reverse.png"
                        },
                        "start": 0.2,
                        "length": 4.8,
                        "transition": {
                            "in": "slideUp",
                            "out": "slideUp"
                        },
                        "fit": "none",
                        "offset": {
                            "y": 0.06
                        },
                        "scale": 1
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>THE CLOUD VIDEO EDITING API</p>",
                            "css": "p { font-family: 'Clear Sans'; color: #cccccc; font-size: 18px; }",
                            "width": 600,
                            "height": 60
                        },
                        "start": 0.4,
                        "length": 4.6,
                        "transition": {
                            "in": "slideDown",
                            "out": "slideDown"
                        },
                        "offset": {
                            "x": 0.036,
                            "y": 0
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