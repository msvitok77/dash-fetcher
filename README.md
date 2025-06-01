# DASH-fetcher
It's a simple command line utility to fetch .mpd files. It supports parsing files from local storage and from remote server.
It produces output in JSON format similar to this:
```json
{
  "audios": [
    {
      "codec": "mp4a.40.2",
      "bitrate": "64008",
      "channels": 2,
      "language": "en"
    },
    {
      "codec": "mp4a.40.2",
      "bitrate": "128008",
      "channels": 2,
      "language": "en"
    }
  ],
  "videos": [
    {
      "codec": "avc1.42C00D",
      "bitrate": "401000",
      "resolution": "224x100"
    },
    {
      "codec": "avc1.42C016",
      "bitrate": "751000",
      "resolution": "448x200"
    },
    {
      "codec": "avc1.4D401F",
      "bitrate": "1001000",
      "resolution": "784x350"
    },
    {
      "codec": "avc1.640028",
      "bitrate": "1501000",
      "resolution": "1680x750"
    },
    {
      "codec": "avc1.640028",
      "bitrate": "2200000",
      "resolution": "1680x750"
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
