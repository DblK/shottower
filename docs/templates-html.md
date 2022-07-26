# HTML

- [Simple HTML text styling](#simple-html-text-styling)
- [Multiple custom fonts using HTML](#multiple-custom-fonts-using-html)
- [Subtitles using HTML](#subtitles-using-html)
- [Text flat drop shadow](#text-flat-drop-shadow)
- [Text blur drop shadow](#text-blur-drop-shadow)
- [Flat box shadow](#flat-box-shadow)
- [HTML offset label](#html-offset-label)
- [Color mattes (background panels)](#color-mattes-background-panels)
- [Custom fonts using merge fields](#custom-fonts-using-merge-fields)
- [Expanding lower third](#expanding-lower-third)
- [Padded table](#padded-table)

## Simple HTML text styling

[Official example](https://shotstack.io/templates/basic-html-styling/)

```json
{
    "timeline": {
        "background": "#000000",
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf"
            }
        ],
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Style text in video using <b>HTML</b> and <u>CSS</u></p>",
                            "css": "p { font-family: 'Open Sans'; color: #ffffff; font-size: 42px; text-align: center; } b { color: #21bcb9; font-weight: normal; } u { color: #e784ff; text-decoration: none; }",
                            "width": 450,
                            "height": 200
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

## Multiple custom fonts using HTML

[Official example](https://shotstack.io/templates/multiple-fonts-html-styling/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/private/epic.mp3"
        },
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/IndieFlower-Regular.ttf"
            },
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/LilitaOne-Regular.ttf"
            }
        ],
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>In 2007 the Opera browser added <b>video</b> to <u>HTML</u>.</p>",
                            "css": "p { font-family: \"Lilita One\"; color: #ffffff; font-size: 34px; text-align: left; } b { color: #25d3d0 }",
                            "width": 600,
                            "height": 400,
                            "background": "transparent"
                        },
                        "start": 1.2,
                        "length": 5,
                        "position": "left",
                        "offset": {
                            "x": 0.1,
                            "y": 0.1
                        },
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>In 2019 the Shotstack API added <b>HTML</b> to <u>video</u>.</p>",
                            "css": "p { font-family: \"Lilita One\"; color: #ffffff; font-size: 34px; text-align: right; } b { color: #25d3d0; }",
                            "width": 600,
                            "height": 400
                        },
                        "start": 6.2,
                        "length": 5,
                        "position": "right",
                        "offset": {
                            "x": -0.1,
                            "y": -0.1
                        },
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

## Subtitles using HTML

[Official example](https://shotstack.io/templates/subtitles-html/)

```json
{
    "timeline": {
        "background": "#000000",
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf"
            }
        ],
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Hi, my name's Scott Ko, as an entrepreneur,</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 0.54,
                        "length": 2.48,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>I cannot overstate how important it is these days to use video as a tool to</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 3.18,
                        "length": 4.399,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>reach your audience, your community, and your customers.</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 7.681,
                        "length": 3.0789,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>People connect with stories and video allows us to be the most authentic we can</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 11.25,
                        "length": 4.43,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>be in order to tell those stories.</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 15.781,
                        "length": 1.879,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>And so if you can present in front of a camera and you have the right tools to</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 18.12,
                        "length": 3.559,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>support and amplify you, you can be unstoppable.</p>",
                            "css": "p { font-family: \"Open Sans\"; font-size: 24px; color: #ffffff; }",
                            "width": 650,
                            "height": 80,
                            "background": "#77000000"
                        },
                        "start": 21.781,
                        "length": 3.379,
                        "position": "bottom",
                        "offset": {
                            "x": 0,
                            "y": 0.1
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://github.com/shotstack/test-media/raw/main/captioning/scott-ko.mp4",
                            "volume": 1
                        },
                        "start": 0,
                        "length": 25.85
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

## Text flat drop shadow

[Official example](https://shotstack.io/templates/flat-text-shadow-html/)

```json
{
    "timeline": {
        "background": "#21bcb9",
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf"
            }
        ],
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Flat Drop Shadow</p>",
                            "css": "p { font-family: 'Open Sans'; color: #ffffff; font-size: 42px; text-align: center; }",
                            "width": 450,
                            "height": 200
                        },
                        "start": 0,
                        "length": 5
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Flat Drop Shadow</p>",
                            "css": "p { font-family: 'Open Sans'; color: #77000000; font-size: 42px; text-align: center; }",
                            "width": 450,
                            "height": 200
                        },
                        "start": 0,
                        "length": 5,
                        "offset": {
                            "x": 0.0015,
                            "y": -0.0015
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

## Text blur drop shadow

[Official example](https://shotstack.io/templates/html-text-drop-shadow-blur/)

```json
{
    "timeline": {
        "background": "#21bcb9",
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf"
            }
        ],
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Fake Drop Shadow</p>",
                            "css": "p { font-family: 'Open Sans'; color: #ffffff; font-size: 42px; text-align: center; }",
                            "width": 450,
                            "height": 200
                        },
                        "start": 0,
                        "length": 5
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Fake Drop Shadow</p>",
                            "css": "p { font-family: 'Open Sans'; color: #CC000000; font-size: 42px; text-align: center; }",
                            "width": 450,
                            "height": 200
                        },
                        "start": 0,
                        "length": 5,
                        "filter": "blur"
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

## Flat box shadow

[Official example](https://shotstack.io/templates/flat-box-shadow-html/)

```json
{
    "timeline": {
        "background": "#000000",
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/palmtrees.mp3",
            "effect": "fadeOut"
        },
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf"
            }
        ],
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<table border='0'><tr><td><p>BOX SHADOW</p></td></tr></table>",
                            "css": "table { background-color: #000000; } td { padding-top: 16px; padding-bottom: 16px; } p { color: #FFFFFF; font-size: 34px; font-family: 'Open Sans'; font-weight: bold; margin: 20px; text-align: center; }",
                            "width": 450,
                            "height": 200
                        },
                        "start": 1,
                        "length": 6,
                        "transition": {
                            "in": "slideUpFast",
                            "out": "slideDownFast"
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<table border='0'><tr><td><p>BOX SHADOW</p></td></tr></table>",
                            "css": "table { background-color: #ffd312; } td { padding-top: 16px; padding-bottom: 16px; } p { color: #ffd312; font-size: 34px; font-family: 'Open Sans'; font-weight: bold; margin: 20px; text-align: center; }",
                            "width": 450,
                            "height": 200
                        },
                        "start": 1,
                        "length": 6,
                        "offset": {
                            "x": 0.0075,
                            "y": -0.0065
                        },
                        "transition": {
                            "in": "slideUpFast",
                            "out": "slideDownFast"
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
                        "start": 0,
                        "length": 8,
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
        "resolution": "hd",
        "aspectRatio": "9:16"
    }
}
```

## HTML offset label

[Official example](https://shotstack.io/templates/offset-html-label/)

```json
{
    "timeline": {
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/LilitaOne-Regular.ttf"
            }
        ],
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>My Label Here</p>",
                            "css": "p { font-family: \"Lilita One\"; color: #ffffff; font-size: 30px; }",
                            "width": 300,
                            "height": 60,
                            "background": "#80ffffff"
                        },
                        "start": 0,
                        "length": 5,
                        "position": "bottomRight",
                        "offset": {
                            "x": -0.1093,
                            "y": 0.0833
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

## Color mattes (background panels)

[Official example](https://shotstack.io/templates/html-color-matte-panels/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/freepd/advertising.mp3"
        },
        "background": "#FFFFFF",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>HELLO CYAN</p>",
                            "css": "p { font-family: \"Montserrat ExtraBold\"; color: #ffffff; font-size: 64px; text-align: center; }",
                            "width": 1280,
                            "height": 720,
                            "background": "transparent"
                        },
                        "start": 1,
                        "length": 5,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p></p>",
                            "width": 1280,
                            "height": 720,
                            "background": "#45dae8"
                        },
                        "start": 0,
                        "length": 7
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>HELLO PINK</p>",
                            "css": "p { font-family: \"Montserrat ExtraBold\"; color: #ffffff; font-size: 64px; text-align: center; }",
                            "width": 1280,
                            "height": 720,
                            "background": "transparent"
                        },
                        "start": 8,
                        "length": 5,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p></p>",
                            "width": 1280,
                            "height": 720,
                            "background": "#da8afc"
                        },
                        "start": 7,
                        "length": 7
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

## Custom fonts using merge fields

[Official example](https://shotstack.io/templates/custom-html-fonts-merge-fields/)

```json
{
    "merge": [
        {
            "find": "FONT",
            "replace": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/IndieFlower-Regular.ttf"
        },
        {
            "find": "FONT_NAME",
            "replace": "Indie Flower"
        }
    ],
    "timeline": {
        "soundtrack": {
            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/private/epic.mp3"
        },
        "fonts": [
            {
                "src": "{{FONT}}"
            },
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/LilitaOne-Regular.ttf"
            }
        ],
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>Shotstack enables businesses to massively <b>scale</b> their video production capabilities using <b>code</b>.</p>",
                            "css": "p { font-family: \"Lilita One\"; color: #ffffff; font-size: 34px; text-align: left; } b { color: #25d3d0 }",
                            "width": 600,
                            "height": 400,
                            "background": "transparent"
                        },
                        "start": 1.2,
                        "length": 5,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.2
                        }
                    },
                    {
                        "asset": {
                            "type": "html",
                            "html": "<p>This new HTML asset type supports <b>highlighted text</b> and <b>wrapping</b> to create flexible and stylish layouts.</p>",
                            "css": "p { font-family: \"{{FONT_NAME}}\"; color: #ffffff; font-size: 34px; text-align: left; } b { color: #25d3d0 }",
                            "width": 600,
                            "height": 400
                        },
                        "start": 6.2,
                        "length": 5,
                        "transition": {
                            "in": "fade",
                            "out": "fade"
                        },
                        "position": "left",
                        "offset": {
                            "x": 0.04
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-pubic-files.s3-ap-southeast-2.amazonaws.com/tutorials/hello-world/earth.mp4",
                            "trim": 6
                        },
                        "start": 0,
                        "length": 6
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/cat.hd.mp4",
                            "trim": 5
                        },
                        "start": 6,
                        "length": 6
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

## Expanding lower third

[Official example](https://shotstack.io/templates/html-lower-third-expanding-highlite/)

```json
{
    "timeline": {
        "soundtrack": {
            "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/berlin.mp3"
        },
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Bold.ttf"
            }
        ],
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "html",
                            "html": "<table border='0' width='1000'><tr><td bgcolor='#FF0000' width='10'>&nbsp;</td><td width='990'><p>This is a lower third example that will expand as the amount of text increases.</p></td></tr></table>",
                            "css": "table { background-color: #77000000; } td { padding-top: 16px; padding-bottom: 16px; } p { color: #FFFFFF; font-size: 34px; font-family: 'Open Sans'; font-weight: bold; margin: 20px; }",
                            "height": 200,
                            "width": 1000,
                            "position": "bottom"
                        },
                        "start": 0,
                        "length": 8,
                        "position": "bottom",
                        "offset": {
                            "y": 0.05
                        },
                        "transition": {
                            "in": "slideUp",
                            "out": "slideDown"
                        }
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
                        "length": 8,
                        "effect": "zoomIn"
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

## Padded table

[Official example](https://shotstack.io/templates/expanding-padded-html-text-box/)

```json
{
    "output": {
        "format": "mp4",
        "resolution": "hd"
    },
    "timeline": {
        "tracks": [
            {
                "clips": [
                    {
                        "start": 0,
                        "length": 15,
                        "position": "right",
                        "asset": {
                            "width": 660,
                            "html": "<table cellpadding='16'><tr><td><p>The black-footed ferret is a very interesting animal. Some say it possesses the secret to the location of the Holy Grail.</p></td></tr></table>",
                            "css": "table { background-color: #33000000; } p { color: #FFFFFF; font-size: 28px; font-family: 'Open Sans' }",
                            "position": "left",
                            "type": "html",
                            "height": 660
                        },
                        "offset": {
                            "x": -0.025,
                            "y": 0
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "start": 0,
                        "length": 15,
                        "scale": 0.4,
                        "position": "left",
                        "asset": {
                            "type": "image",
                            "src": "https://cdm-citadel-core-2.s3.amazonaws.com/51ab442b-13b6-483d-aef3-7d71779bf3c3.png"
                        },
                        "offset": {
                            "x": 0.025,
                            "y": 0
                        }
                    }
                ]
            },
            {
                "clips": [
                    {
                        "start": 0,
                        "length": 15,
                        "scale": 1,
                        "position": "center",
                        "asset": {
                            "type": "video",
                            "src": "https://cdm-citadel-core-2.s3.amazonaws.com/534e4c13-f558-4e42-b95b-f735ef62de29.mp4"
                        },
                        "offset": {
                            "x": 0,
                            "y": 0
                        }
                    }
                ]
            }
        ],
        "background": "#000000",
        "fonts": [
            {
                "src": "https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf"
            }
        ]
    }
}
```