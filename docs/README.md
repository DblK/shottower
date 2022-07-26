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