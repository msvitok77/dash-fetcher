# DASH-fetcher
It's a simple command line utility to fetch .mpd files. It supports parsing files from local storage and from remote server.
It produces output in JSON format similar to this:
```json
{
  "periods": [
    {
      "audios": [
        {
          "codec": "mp4a.40.2",
          "bitrate": "34189",
          "channels": 2,
          "language": "eng"
        }
      ],
      "videos": [
        {
          "codec": "avc1.640028",
          "bitrate": "999120",
          "resolution": "1920x1080"
        },
        {
          "codec": "avc1.640028",
          "bitrate": "2003095",
          "resolution": "1920x1080"
        },
        {
          "codec": "hev1.1.6.L120.90",
          "bitrate": "1029626",
          "resolution": "1920x1080"
        },
        {
          "codec": "hev1.1.6.L120.90",
          "bitrate": "2067007",
          "resolution": "1920x1080"
        }
      ]
    }
  ]
}
```

# Usage
`dash-fetcher -p [URL]`

# Testing scenario
1. `make build`
2. `make start-file-server`
3. Test http protocol (files {1..6}.mpd)
  * `./dash-fetcher -p "http://localhost:8080/1.mpd"`
4. Test https protocol (files {1..6}.mpd)
  * `./dash-fetcher -p "https://localhost/1.mpd"`
6. Test files protocol (files {1..6}.mpd)
  * `./dash-fetcher -p "file://${HOME}/{cloned-folder}/cmd/file-server/mpds/1.mpd"`
