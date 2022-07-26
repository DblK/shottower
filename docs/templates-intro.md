# Intros

- [Curtain reveal effect intro](#curtain-reveal-effect-intro)
- [Panels and text reveal effect intro](#panels-and-text-reveal-effect-intro)
- [Horizontal panels animation intro](#horizontal-panels-animation-intro)
- [Diagonal sliding panel intro](#diagonal-sliding-panel-intro)
- [Falling bars animation effect intro](#falling-bars-animation-effect-intro)
- [Diagonal bars animation intro](#diagonal-bars-animation-intro)
- [Color fade transition intro](#color-fade-transition-intro)

## Curtain reveal effect intro

[Official example](https://shotstack.io/templates/curtain-reveal-effect-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 960,
                            "height": 1080,
                            "background": "#131463"
                        },
                        "start": 0,
                        "length": 1.4,
                        "position": "left",
                        "transition": {
                            "out": "carouselLeftSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 960,
                            "height": 1080,
                            "background": "#131463"
                        },
                        "start": 0,
                        "length": 1.4,
                        "position": "right",
                        "transition": {
                            "out": "carouselRightSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/corporate-colour.png"
                        },
                        "start": 0,
                        "length": 3,
                        "fit": "none",
                        "scale": 0.65,
                        "offset": {
                            "y": 0.1
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Main Title</p>",
                            "css": "p { font-family: \"Barlow ExtraBold\"; color: #555555; font-size: 60px; }",
                            "width": 1000,
                            "height": 70,
                            "position": "bottom"
                        },
                        "start": 0.8,
                        "length": 2.2,
                        "offset": {
                            "y": -0.12
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Sub heading and more information</p>",
                            "css": "p { font-family: \"Barlow\"; color: #777777; font-size: 50px; }",
                            "width": 1000,
                            "height": 120,
                            "position": "top"
                        },
                        "start": 1,
                        "length": 2,
                        "offset": {
                            "y": -0.2
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>email@example.com</p>",
                            "css": "p { font-family: \"Barlow\"; color: #CC000000; font-size: 30px; }",
                            "width": 800,
                            "height": 100
                        },
                        "start": 1.2,
                        "length": 1.8,
                        "position": "bottom",
                        "offset": {
                            "y": 0.03
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 160,
                            "background": "#f3f3f9"
                        },
                        "start": 0.6,
                        "length": 2.4,
                        "position": "bottom",
                        "transition": {
                            "in": "carouselUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 1080,
                            "background": "#ddddee"
                        },
                        "start": 0,
                        "length": 3
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "mp4",
        "resolution": "1080"
    }
}
```

## Panels and text reveal effect intro

[Official example](https://shotstack.io/templates/panels-text-reveal-effect-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/carmart-color-white.png"
                        },
                        "start": 0,
                        "length": 5,
                        "fit": "none",
                        "scale": 0.65,
                        "position": "topLeft",
                        "offset": {
                            "x": 0.04,
                            "y": -0.1
                        },
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
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 700,
                            "height": 80,
                            "background": "#fffcf2"
                        },
                        "start": 1.5,
                        "length": 1,
                        "position": "left",
                        "offset": {
                            "x": 0.035,
                            "y": 0.03
                        },
                        "transition": {
                            "in": "slideDown",
                            "out": "slideDown"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 770,
                            "height": 80,
                            "background": "#fffcf2"
                        },
                        "start": 1.7,
                        "length": 1,
                        "position": "left",
                        "offset": {
                            "x": 0.035,
                            "y": -0.05
                        },
                        "transition": {
                            "in": "slideDown",
                            "out": "slideDown"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Primary text on this line</p>",
                            "css": "p { font-family: \"Barlow ExtraBold\"; color: #ffffff; font-size: 60px; text-align: left; }",
                            "width": 800,
                            "height": 70
                        },
                        "start": 1.9,
                        "length": 3.1,
                        "position": "left",
                        "offset": {
                            "x": 0.04,
                            "y": 0.03
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Secondary text on this line</p>",
                            "css": "p { font-family: \"Barlow\"; color: #ffffff; font-size: 60px; text-align: left; }",
                            "width": 800,
                            "height": 70
                        },
                        "start": 2.1,
                        "length": 2.9,
                        "position": "left",
                        "offset": {
                            "x": 0.04,
                            "y": -0.05
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>email@example.com</p>",
                            "css": "p { font-family: \"Barlow\"; color: #000000; font-size: 30px; text-align: left; }",
                            "width": 400,
                            "height": 100
                        },
                        "start": 2.2,
                        "length": 2.8,
                        "position": "bottomRight",
                        "offset": {
                            "y": 0.1
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 700,
                            "height": 20,
                            "background": "#ffe499"
                        },
                        "start": 1.9,
                        "length": 3.1,
                        "position": "left",
                        "offset": {
                            "x": 0.035,
                            "y": 0.01
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 770,
                            "height": 20,
                            "background": "#ffe499"
                        },
                        "start": 2.1,
                        "length": 2.9,
                        "position": "left",
                        "offset": {
                            "x": 0.035,
                            "y": -0.067
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 450,
                            "height": 100,
                            "background": "#fff8e6"
                        },
                        "start": 1.5,
                        "length": 3.5,
                        "position": "bottomRight",
                        "offset": {
                            "y": 0.1
                        },
                        "transition": {
                            "in": "carouselLeft"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 460,
                            "height": 100,
                            "background": "#2B35FC"
                        },
                        "start": 1.4,
                        "length": 3.6,
                        "position": "bottomRight",
                        "offset": {
                            "y": 0.1
                        },
                        "transition": {
                            "in": "carouselLeft"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 960,
                            "height": 1080,
                            "background": "#ffbc00"
                        },
                        "start": 0.7,
                        "length": 4.3,
                        "position": "right",
                        "transition": {
                            "in": "carouselLeftSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1200,
                            "height": 1080,
                            "background": "#ffcd40"
                        },
                        "start": 0.5,
                        "length": 4.5,
                        "position": "left",
                        "transition": {
                            "in": "carouselRightSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 1080,
                            "background": "#FF8E00"
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
        "resolution": "1080"
    }
}
```

## Horizontal panels animation intro

[Official example](https://shotstack.io/templates/horizontal-panels-animation-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/news-black.png"
                        },
                        "start": 0.5,
                        "length": 4.5,
                        "fit": "none",
                        "offset": {
                            "y": 0.22
                        },
                        "transition": {
                            "in": "slideDown"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Main Headline</p>",
                            "css": "p { font-family: 'Barlow ExtraBold'; color: #ffffff; font-size: 80px; }",
                            "width": 1000,
                            "height": 100
                        },
                        "start": 1.9,
                        "length": 3.1,
                        "offset": {
                            "y": -0.14
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Subtitle text and more information</p>",
                            "css": "p { font-family: 'Barlow'; color: #ffffff; font-size: 60px; }",
                            "width": 1000,
                            "height": 200,
                            "position": "top"
                        },
                        "start": 2.1,
                        "length": 2.9,
                        "offset": {
                            "y": -0.28
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>example@mail.com</p>",
                            "css": "p { font-family: 'Barlow'; color: #ffffff; font-size: 40px; }",
                            "width": 1000,
                            "height": 50
                        },
                        "start": 2.3,
                        "length": 2.7,
                        "position": "bottom",
                        "offset": {
                            "y": -0.12
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 500,
                            "background": "#a51028"
                        },
                        "start": 0.7,
                        "length": 4.3,
                        "position": "bottom",
                        "transition": {
                            "in": "carouselUpSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 1080,
                            "background": "#ffffff"
                        },
                        "start": 0.6,
                        "length": 4.4,
                        "transition": {
                            "in": "carouselDownSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 1080,
                            "background": "#e9e9e9"
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
        "resolution": "1080"
    }
}
```

## Diagonal sliding panel intro

[Official example](https://shotstack.io/templates/diagonal-panel-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extralight.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/real-estate-black.png"
                        },
                        "start": 1.3,
                        "length": 3.7,
                        "fit": "none",
                        "scale": 0.65,
                        "position": "topLeft",
                        "offset": {
                            "x": 0.19,
                            "y": -0.1
                        },
                        "transition": {
                            "in": "slideRight"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Main Heading</p>",
                            "css": "p { font-family: 'Barlow ExtraBold'; color: #000000; font-size: 60px; }",
                            "width": 800,
                            "height": 80
                        },
                        "start": 1.9,
                        "length": 3.1,
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": 0.04
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Subheading and more info</p>",
                            "css": "p { font-family: 'Barlow'; color: #000000; font-size: 50px; }",
                            "width": 800,
                            "height": 150,
                            "position": "top"
                        },
                        "start": 2.1,
                        "length": 2.9,
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": -0.07
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>www.example.com</p>",
                            "css": "p { font-family: 'Barlow ExtraLight'; color: #000000; font-size: 40px; }",
                            "width": 800,
                            "height": 50
                        },
                        "start": 2.2,
                        "length": 2.8,
                        "position": "left",
                        "offset": {
                            "x": 0.05,
                            "y": -0.18
                        },
                        "transition": {
                            "in": "slideUp"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1200,
                            "height": 1500,
                            "background": "#e6f6e5"
                        },
                        "start": 0.7,
                        "length": 4.3,
                        "position": "left",
                        "offset": {
                            "x": -0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": -7
                            }
                        },
                        "transition": {
                            "in": "slideRightSlow"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 1080,
                            "background": "#0ea800"
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
        "resolution": "1080"
    }
}
```

## Falling bars animation effect intro

[Official example](https://shotstack.io/templates/falling-bars-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extralight.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#0036dd",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/news-white.png"
                        },
                        "start": 1.5,
                        "length": 3.5,
                        "fit": "none",
                        "offset": {
                            "y": 0.2
                        },
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
                            "type": "html",
                            "html": "<p>Main Headline</p>",
                            "css": "p { font-family: 'Barlow ExtraBold'; color: #f2f2f2; font-size: 80px; }",
                            "width": 1000,
                            "height": 100
                        },
                        "start": 1.9,
                        "length": 3.1,
                        "offset": {
                            "y": -0.02
                        },
                        "transition": {
                            "in": "slideRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Summary copy and more information</p>",
                            "css": "p { font-family: 'Barlow'; color: #f2f2f2; font-size: 60px; }",
                            "width": 1000,
                            "height": 150,
                            "position": "top"
                        },
                        "start": 2.1,
                        "length": 2.9,
                        "offset": {
                            "y": -0.14
                        },
                        "transition": {
                            "in": "slideRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p><b>Tel:</b> 999-999-9999</p>",
                            "css": "p { font-family: 'Barlow ExtraLight'; color: #f2f2f2; font-size: 40px; } ",
                            "width": 1000,
                            "height": 50
                        },
                        "start": 2.3,
                        "length": 2.7,
                        "position": "center",
                        "offset": {
                            "y": -0.25
                        },
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
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 384,
                            "height": 1080,
                            "background": "#dddddd"
                        },
                        "start": 0,
                        "length": 1.2,
                        "position": "left",
                        "offset": {
                            "x": 0
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 384,
                            "height": 1080,
                            "background": "#dddddd"
                        },
                        "start": 0,
                        "length": 1.4,
                        "position": "left",
                        "offset": {
                            "x": 0.2
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 384,
                            "height": 1080,
                            "background": "#dddddd"
                        },
                        "start": 0,
                        "length": 1.6,
                        "position": "left",
                        "offset": {
                            "x": 0.4
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 384,
                            "height": 1080,
                            "background": "#dddddd"
                        },
                        "start": 0,
                        "length": 1.8,
                        "position": "left",
                        "offset": {
                            "x": 0.6
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 384,
                            "height": 1080,
                            "background": "#dddddd"
                        },
                        "start": 0,
                        "length": 2,
                        "position": "left",
                        "offset": {
                            "x": 0.8
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "mp4",
        "resolution": "1080"
    }
}
```

## Diagonal bars animation intro

[Official example](https://shotstack.io/templates/diagonal-bars-animation-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extralight.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#ffffff",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "image",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/real-estate-black.png"
                        },
                        "start": 1.5,
                        "length": 3.5,
                        "fit": "none",
                        "scale": 0.8,
                        "offset": {
                            "y": 0.2
                        },
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
                            "type": "html",
                            "html": "<p>Main Heading</p>",
                            "css": "p { font-family: 'Barlow ExtraBold'; color: #ffe512; font-size: 80px; }",
                            "width": 1000,
                            "height": 90
                        },
                        "start": 1.9,
                        "length": 3.1,
                        "offset": {
                            "y": -0.05
                        },
                        "transition": {
                            "in": "slideRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Subheading and more info</p>",
                            "css": "p { font-family: 'Barlow'; color: #333333; font-size: 60px; }",
                            "width": 1000,
                            "height": 200,
                            "position": "top"
                        },
                        "start": 2.1,
                        "length": 2.9,
                        "offset": {
                            "y": -0.19
                        },
                        "transition": {
                            "in": "slideRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p><b>Tel:</b> 999-999-9999</p>",
                            "css": "p { font-family: 'Barlow ExtraLight'; color: #333333; font-size: 40px; } ",
                            "width": 1000,
                            "height": 50
                        },
                        "start": 2.3,
                        "length": 2.7,
                        "offset": {
                            "y": -0.36
                        },
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
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 380,
                            "height": 1280,
                            "background": "#ffe512"
                        },
                        "start": 0,
                        "length": 1.2,
                        "position": "left",
                        "offset": {
                            "x": -0.1,
                            "y": 0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": 10
                            }
                        },
                        "transition": {
                            "out": "carouselUpSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 380,
                            "height": 1280,
                            "background": "#ffe512"
                        },
                        "start": 0,
                        "length": 1.4,
                        "position": "left",
                        "offset": {
                            "x": 0.1,
                            "y": 0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": 10
                            }
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 380,
                            "height": 1280,
                            "background": "#ffe512"
                        },
                        "start": 0,
                        "length": 1.6,
                        "position": "left",
                        "offset": {
                            "x": 0.3,
                            "y": 0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": 10
                            }
                        },
                        "transition": {
                            "out": "carouselUpSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 380,
                            "height": 1280,
                            "background": "#ffe512"
                        },
                        "start": 0,
                        "length": 1.8,
                        "position": "left",
                        "offset": {
                            "x": 0.5,
                            "y": 0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": 10
                            }
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 380,
                            "height": 1280,
                            "background": "#ffe512"
                        },
                        "start": 0,
                        "length": 2,
                        "position": "left",
                        "offset": {
                            "x": 0.7,
                            "y": 0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": 10
                            }
                        },
                        "transition": {
                            "out": "carouselUpSlow"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 380,
                            "height": 1280,
                            "background": "#ffe512"
                        },
                        "start": 0,
                        "length": 2.2,
                        "position": "left",
                        "offset": {
                            "x": 0.9,
                            "y": 0.1
                        },
                        "transform": {
                            "rotate": {
                                "angle": 10
                            }
                        },
                        "transition": {
                            "out": "carouselDownSlow"
                        }
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "gif",
        "resolution": "1080"
    }
}
```

## Color fade transition intro

[Official example](https://shotstack.io/templates/color-fade-transition-intro-video/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extrabold.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-extralight.ttf"
            },
            {
                "src": "https://templates.shotstack.io/basic/asset/font/barlow-regular.ttf"
            }
        ],
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/fireworks.mp3",
            "effect": "fadeOut"
        },
        "background": "#00bdff",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>&nbsp;</p>",
                            "width": 1920,
                            "height": 1080,
                            "background": "#8ad424"
                        },
                        "start": 0,
                        "length": 1,
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
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/logos/heart-health-white.png"
                        },
                        "start": 1,
                        "length": 3,
                        "fit": "none",
                        "offset": {
                            "y": 0.25
                        },
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
                            "type": "html",
                            "html": "<p>Main Heading</p>",
                            "css": "p { font-family: 'Barlow ExtraBold'; color: #ffffff; font-size: 80px; }",
                            "width": 1000,
                            "height": 100
                        },
                        "start": 1.4,
                        "length": 2.6,
                        "transition": {
                            "in": "slideRight"
                        },
                        "offset": {
                            "y": -0.08
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Subheading text and more details</p>",
                            "css": "p { font-family: 'Barlow'; color: #ffffff; font-size: 60px; }",
                            "width": 1000,
                            "height": 150,
                            "position": "top"
                        },
                        "start": 1.6,
                        "length": 2.4,
                        "offset": {
                            "y": -0.2
                        },
                        "transition": {
                            "in": "slideRight"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p><b>Tel:</b> 999-999-9999</p>",
                            "css": "p { font-family: 'Barlow ExtraLight'; color: #ffffff; font-size: 40px; } ",
                            "width": 1000,
                            "height": 50
                        },
                        "start": 1.8,
                        "length": 2.2,
                        "offset": {
                            "y": -0.32
                        },
                        "transition": {
                            "in": "fade"
                        }
                    }
                ]
            }
        ]
    },
    "output": {
        "format": "gif",
        "resolution": "1080"
    }
}
```