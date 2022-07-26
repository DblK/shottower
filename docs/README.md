# Shottower official documentation

## Shotstack documentation
- [Main Website](https://shotstack.io/)
- [Doc & Guide](https://shotstack.io/docs/guide/)
- [API Reference](https://shotstack.io/docs/api/#shotstack)
- [Templates](https://shotstack.io/templates/) ([Private copy](./templates.md))
- [Demos](https://shotstack.io/templates/)
  - [Real Estate](https://shotstack.io/demo/domain-real-estate-api/) - [Source Code](https://github.com/shotstack/domain-real-estate-demo)
  - [Compress Video](https://shotstack.io/demo/compress-video/)
  - [Motivational Quote Image Maker](https://shotstack.io/demo/motivational-quote-maker/) - [Source Code](https://github.com/shotstack/pexels-quote-image-demo)
  - [Trim Video](https://shotstack.io/demo/trim-video/) - [Source Code](https://github.com/shotstack/trim-video-demo)
  - [Pexels Slideshow Video Maker](https://shotstack.io/demo/pexels-slideshow/)
  - [Christmas Personalized Video](https://shotstack.io/demo/christmas-video/) - [Source Code](https://github.com/shotstack/christmas-demo-2020)
  - [Watermark Video](https://shotstack.io/demo/watermarker/) - [Source Code](https://github.com/shotstack/watermark-demo)
  - [Picture in Picture Video Editor](https://shotstack.io/demo/picture-in-picture/) - [Source Code](https://github.com/shotstack/picture-in-picture-demo)
  - [Pexels Video Maker](https://shotstack.io/demo/pexels/)

## cURL Picture in Picture example
```sh
curl -d '{
    "timeline": {
        "background": "#000000",
        "tracks": [
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/examples/picture-in-picture/commentary.mp4",
                            "volume": 1
                        },
                        "start": 0,
                        "length": 4.96
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/examples/picture-in-picture/commentary.mp4",
                            "volume": 1,
                            "trim": 5
                        },
                        "start": 5,
                        "length": 4.96,
                        "scale": 0.35,
                        "position": "bottomRight"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/examples/picture-in-picture/commentary.mp4",
                            "volume": 1,
                            "trim": 10
                        },
                        "start": 10,
                        "length": 2.46,
                        "scale": 0.35,
                        "position": "topRight"
                    },
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/examples/picture-in-picture/commentary.mp4",
                            "volume": 1,
                            "trim": 12.5
                        },
                        "start": 12.5,
                        "length": 2.5,
                        "scale": 0.25,
                        "position": "topRight"
                    }
                ]
            },
            {
                "clips": [
                    {
                        "asset": {
                            "type": "video",
                            "src": "https://shotstack-assets.s3.ap-southeast-2.amazonaws.com/examples/picture-in-picture/code.mp4"
                        },
                        "start": 5,
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
}'  \
-H "Content-Type: application/json" \
-X POST http://0.0.0.0:4000/stage/render
```

## cURL command for the status
From the id given by the previous command:

```sh
curl -X GET http://0.0.0.0:4000/stage/render/d2b46ed6-998a-4d6b-9d91-b8cf0193a655 \
-H 'Accept: application/json' \
  -H 'x-api-key: freeKey'
```
Then you should have the `url` property to know where to download the render.
## cURL command for downloading renders

```sh
curl -X GET \
     -H "Content-Type: application/json" \
     -H "x-api-key: freeKey" \
     https://0.0.0.0:4000/dl/stage/renders/d2b46ed6-998a-4d6b-9d91-b8cf0193a655
```

## New features

### Burning subtitle

```json
"clips": [
    {
    "asset": {
        "type": "video",
        "src": "https://example.com/my_asset.mkv",
        "volume": 1,
        // Add a section subtitle if you want to burn it
        "subtitle": {
            // Specify the index of the subtitle stream (default: 0)
            "index": 0
        }
    },
    "start": 0,
    "length": 4.96
    },
    ...
]
```

### Use local resources

```json
"clips": [
    {
    "asset": {
        "type": "video",
        // Every source starting with `file` won't be fetched and used as is
        "src": "file:///Users/dblk/clips/my_asset.mp4",
        "volume": 1,
    },
    "start": 0,
    "length": 4.96
    },
    ...
]
```